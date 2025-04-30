package types

type PersonStore interface {
	GetPeople() ([]DBPerson, error)
	CreatePerson(DBPerson) error
	PersonChange(DBPerson) error
	DeletePerson(id int) error
}

type PersonService interface {
	GetPeople() ([]DBPerson, error)
	CreatePerson(person NewPerson) error
	PersonChange(DBPerson) error
	DeletePerson(id int) error
}

type Person struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type DBPerson struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Country    string `json:"country"`
}

type NewPerson struct {
	Name       string `json:"name" validate:"required"`
	Surname    string `json:"surname" validate:"required"`
	Patronymic string `json:"patronymic" validate:"required"`
}

type AgifyResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

type GenderizeResponse struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Propability float32 `json:"probability"`
	Count       int     `json:"count"`
}

type NationalizeResponse struct {
	Name        string           `json:"name"`
	Country     []CountryElement `json:"country"`
	Propability float32          `json:"probability"`
	Count       int              `json:"count"`
}

type CountryElement struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}
