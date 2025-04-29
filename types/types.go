package types

type PersonStore interface {
	GetPersons() ([]Person, error)
	NewPerson(Person) (int, error)
	PersonChange(Person) error
	DeletePerson(id int) error
}

type PersonService interface {
	GetPersons() ([]Person, error)
	NewPerson(Person) (int, error)
	PersonChange(Person) error
	DeletePerson(id int) error
}

type Person struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type NewPerson struct {
	Name       string `json:"name" validate:"required"`
	Surname    string `json:"surname" validate:"required"`
	Patronymic string `json:"patronymic" validate:"required"`
}
