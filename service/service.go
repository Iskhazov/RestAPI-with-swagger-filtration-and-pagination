package service

import (
	"awesomeProject2/types"
	"awesomeProject2/utils"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type LayerService struct {
	store types.PersonStore
}

func NewLayerService(store types.PersonStore) *LayerService {
	return &LayerService{store: store}
}

// GetPersonById fetches a person by ID from the store.
func (l *LayerService) GetPersonById(id int) (*types.DBPerson, error) {
	return l.store.GetPersonById(id)
}

// GetPeople returns a paginated and filtered list of people.
func (l *LayerService) GetPeople(request types.PageToken, size int, filters []types.Filter) (*types.GetPeopleResponse, error) {
	people, err := l.store.GetPeople(request, size, filters)
	if err != nil {
		return nil, err
	}

	var token string
	if len(people) > size {
		token, err = utils.EncodeToken(&types.PageToken{
			Id: people[size-1].ID,
		})
		if err != nil {
			return nil, err
		}
		people = people[:size]
	} else {
		token = ""
	}

	return &types.GetPeopleResponse{
		People:        people,
		NextPageToken: token,
	}, nil
}

// CreatePerson enriches and stores a new person entity.
func (l *LayerService) CreatePerson(person types.NewPerson) error {
	log.Printf("Creating new person: %s %s", person.Name, person.Surname)
	return l.store.CreatePerson(NewDBPerson(person))
}

// PersonChange updates an existing person in the store.
func (l *LayerService) PersonChange(id int, person types.DBPerson) error {
	log.Printf("Updating person with ID: %d", id)
	return l.store.PersonChange(id, person)
}

// DeletePerson removes a person from the store by ID.
func (l *LayerService) DeletePerson(id int) error {
	log.Printf("Deleting person with ID: %d", id)
	return l.store.DeletePerson(id)
}

// NewDBPerson enriches data using external APIs and maps to DBPerson.
func NewDBPerson(person types.NewPerson) types.DBPerson {
	urlAge := "https://api.agify.io/?name=" + person.Name
	urlGender := "https://api.genderize.io/?name=" + person.Name
	urlCountry := "https://api.nationalize.io/?name=" + person.Name

	age := UrlAge(urlAge)
	gender := UrlGender(urlGender)
	country := UrlCountry(urlCountry)

	countryCode := GetRandomCountry(country.Country)

	log.Printf("Enriched data for %s: Age=%d, Gender=%s, Country=%s", person.Name, age.Age, gender.Gender, countryCode)

	return types.DBPerson{
		Name:       person.Name,
		Surname:    person.Surname,
		Patronymic: person.Patronymic,
		Age:        age.Age,
		Gender:     gender.Gender,
		Country:    countryCode,
	}
}

// UrlAge makes a request to agify.io to predict age.
func UrlAge(urlAge string) types.AgifyResponse {
	resp, err := http.Get(urlAge)
	RequestError(err)

	defer resp.Body.Close()
	ResponseError(resp)

	var age types.AgifyResponse
	err = json.NewDecoder(resp.Body).Decode(&age)
	if err != nil {
		log.Fatalf("Error decoding JSON from agify: %v", err)
	}
	return age
}

// UrlGender makes a request to genderize.io to predict gender.
func UrlGender(urlGender string) types.GenderizeResponse {
	resp, err := http.Get(urlGender)
	RequestError(err)

	defer resp.Body.Close()
	ResponseError(resp)

	var gender types.GenderizeResponse
	err = json.NewDecoder(resp.Body).Decode(&gender)
	if err != nil {
		log.Fatalf("Error decoding JSON from genderize: %v", err)
	}
	return gender
}

// UrlCountry makes a request to nationalize.io to predict country.
func UrlCountry(urlGender string) types.NationalizeResponse {
	resp, err := http.Get(urlGender)
	RequestError(err)

	defer resp.Body.Close()
	ResponseError(resp)

	var country types.NationalizeResponse
	err = json.NewDecoder(resp.Body).Decode(&country)
	if err != nil {
		log.Fatalf("Error decoding JSON from nationalize: %v", err)
	}
	return country
}

// GetRandomCountry selects a country randomly, weighted by probability.
func GetRandomCountry(countries []types.CountryElement) string {
	rand.NewSource(time.Now().UnixNano())

	total := 0.0
	for _, c := range countries {
		total += c.Probability
	}

	r := rand.Float64() * total

	current := 0.0
	for _, c := range countries {
		current += c.Probability
		if r <= current {
			return c.CountryID
		}
	}

	return ""
}

// RequestError logs and exits if an error occurred during request.
func RequestError(err error) {
	if err != nil {
		log.Fatalf("Request Error: %v", err)
	}
}

// ResponseError logs and exits if response is not 200 OK.
func ResponseError(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Response Error: %s", resp.Status)
	}
}
