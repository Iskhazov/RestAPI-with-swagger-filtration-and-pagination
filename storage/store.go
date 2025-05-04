package storage

import (
	"awesomeProject2/types"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// Разобраться с внешним API

func (s *Store) GetPeople(request types.PageToken, size int) ([]types.DBPerson, error) {
	query := `
	SELECT * FROM people
	WHERE id > $1
	ORDER BY id
	LIMIT $2
	`
	rows, err := s.db.Query(query, request.Id, size+1)
	if err != nil {
		return nil, err
	}

	people := make([]types.DBPerson, 0)
	for rows.Next() {
		p, err := scanRowsIntoPeople(rows)
		if err != nil {
			return nil, err
		}
		people = append(people, *p)
	}

	return people, nil

}
func (s *Store) GetPersonById(id int) (*types.DBPerson, error) {
	rows, err := s.db.Query("SELECT * FROM people WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	person := new(types.DBPerson)
	for rows.Next() {
		person, err = scanRowsIntoPeople(rows)
		if err != nil {
			return nil, err
		}
	}
	if person.ID == 0 {
		return nil, fmt.Errorf("PERSON NOT FOUND")
	}
	return person, nil
}

func (s *Store) CreatePerson(person types.DBPerson) error {
	_, err := s.db.Exec("INSERT INTO people(name,surname,patronymic,age,gender,country) VALUES($1,$2,$3,$4,$5,$6)",
		person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Country)
	return err
}

func (s *Store) PersonChange(id int, person types.DBPerson) error {
	_, err := s.db.Exec("UPDATE people SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, country = $6 WHERE id=$7",
		person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Country, id)
	return err
}

func (s *Store) DeletePerson(id int) error {
	_, err := s.db.Exec("DELETE FROM people WHERE id = $1", id)
	return err
}

func scanRowsIntoPeople(rows *sql.Rows) (*types.DBPerson, error) {
	person := new(types.DBPerson)
	err := rows.Scan(
		&person.ID,
		&person.Name,
		&person.Surname,
		&person.Patronymic,
		&person.Age,
		&person.Gender,
		&person.Country,
	)
	if err != nil {
		return nil, err
	}
	return person, nil
}
