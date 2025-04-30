package service

import (
	"awesomeProject2/types"
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

func (l *LayerService) GetPeople() ([]types.DBPerson, error) {
	return l.store.GetPeople()
}

func (l *LayerService) CreatePerson(person types.NewPerson) error {

	return l.store.CreatePerson(NewDBPerson(person))
}

func (l *LayerService) PersonChange(person types.DBPerson) error {
	return l.store.PersonChange(person)
}

func (l *LayerService) DeletePerson(id int) error {
	return l.store.DeletePerson(id)
}

func NewDBPerson(person types.NewPerson) types.DBPerson {

	urlAge := "https://api.agify.io/?name=" + person.Name
	urlGender := "https://api.genderize.io/?name=" + person.Name
	urlCountry := "https://api.nationalize.io/?name=" + person.Name

	age := UrlAge(urlAge)
	gender := UrlGender(urlGender)
	country := UrlCountry(urlCountry)

	countryCode := GetRandomCountry(country.Country)

	return types.DBPerson{
		Name:       person.Name,
		Surname:    person.Surname,
		Patronymic: person.Patronymic,
		Age:        age.Age,
		Gender:     gender.Gender,
		Country:    countryCode,
	}
}

func UrlAge(urlAge string) types.AgifyResponse {
	resp, err := http.Get(urlAge)
	RequestError(err)

	defer resp.Body.Close()
	ResponseError(resp)

	var age types.AgifyResponse
	err = json.NewDecoder(resp.Body).Decode(&age)
	if err != nil {
		log.Fatalf("Error JSON: %v", err)
	}
	return age
}

func UrlGender(urlGender string) types.GenderizeResponse {
	resp, err := http.Get(urlGender)
	RequestError(err)

	defer resp.Body.Close()
	ResponseError(resp)

	var gender types.GenderizeResponse
	err = json.NewDecoder(resp.Body).Decode(&gender)
	if err != nil {
		log.Fatalf("Error JSON: %v", err)
	}
	return gender
}

func UrlCountry(urlGender string) types.NationalizeResponse {
	resp, err := http.Get(urlGender)
	RequestError(err)

	defer resp.Body.Close()
	ResponseError(resp)

	var country types.NationalizeResponse
	err = json.NewDecoder(resp.Body).Decode(&country)
	if err != nil {
		log.Fatalf("Error JSON: %v", err)
	}
	return country
}

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

func RequestError(err error) {
	if err != nil {
		log.Fatalf("Request Error: %v", err)
	}
}

func ResponseError(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Response Error: %s", resp.Status)
	}
}
