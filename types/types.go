package types

type PersonStore interface {
	GetPersonById(id int) (*DBPerson, error)
	GetPeople(request PageToken, size int, filters []Filter) ([]DBPerson, error)
	CreatePerson(DBPerson) error
	PersonChange(int, DBPerson) error
	DeletePerson(int) error
}

type PersonService interface {
	GetPersonById(id int) (*DBPerson, error)
	GetPeople(request PageToken, size int, filters []Filter) (*GetPeopleResponse, error)
	CreatePerson(person NewPerson) error
	PersonChange(int, DBPerson) error
	DeletePerson(int) error
}

type Filter struct {
	Field  string   `json:"field" example:"gender"`
	Values []string `json:"values" example:"male/female"`
}

type PageToken struct {
	Id int `json:"id"`
}

type GetPeopleRequest struct {
	PageToken string   `json:"page_token" example:""`
	Size      int      `json:"size" example:"20"`
	Filters   []Filter `json:"filters"`
}

type GetPeopleResponse struct {
	People        []DBPerson `json:"people"`
	NextPageToken string     `json:"next_page_token"`
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
	Name       string `json:"name" validate:"required" example:"Ivan"`
	Surname    string `json:"surname" validate:"required" example:"Ivanov"`
	Patronymic string `json:"patronymic" validate:"required" example:"Ivanovich"`
}

type AgifyResponse struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}

type GenderizeResponse struct {
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
	Count       int     `json:"count"`
}

type NationalizeResponse struct {
	Name        string           `json:"name"`
	Country     []CountryElement `json:"country"`
	Probability float32          `json:"probability"`
	Count       int              `json:"count"`
}

type CountryElement struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type SwagPerson struct {
	Name       string `json:"name" example:"John"`
	Surname    string `json:"surname" example:"Doe"`
	Patronymic string `json:"patronymic" example:"Ivanovich"`
	Age        int    `json:"age" example:"30"`
	Gender     string `json:"gender" example:"male"`
	Country    string `json:"country" example:"USA"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"something wrong"`
}
