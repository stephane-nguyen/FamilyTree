package storage

import (
	"github.com/jmoiron/sqlx"
)

type PersonStorage struct {
	Conn *sqlx.DB
}

func NewPersonStorage(conn *sqlx.DB) *PersonStorage {
	return &PersonStorage{Conn: conn}
}

type Person_DB struct {
	Id			int `json:"id"`
    Firstname   string `json:"firstname"`
    Middlename  string `json:"middlename"`
    Lastname    string `json:"lastname"`
	Birthdate 	string `json:"birthdate"`
    Gender      string `json:"gender"`
    City        string `json:"city"`
    Country     string `json:"country"`
	Photo		*[]byte`json:"photo"`
}

type NewPersonInput struct {
	Firstname   string 
    Middlename  string 
    Lastname    string 
	Birthdate 	string 
    Gender      string 
    City        string 
    Country     string 
    Photo     string
}

func (s *PersonStorage) GetAllPersons() ([]Person_DB, error){
	persons := []Person_DB{}

	err := s.Conn.Select(&persons, "SELECT * FROM Person")
	if err != nil {
		
		return persons, err
	}
	return persons, nil
}

func (s *PersonStorage) GetpersonsByCountry(country string) ([]Person_DB, error){
	persons := []Person_DB{}

	err := s.Conn.Select(&persons, "SELECT * FROM Person WHERE country = ?", country)
	if err != nil {
		return persons, err
	}

	return persons, nil
}

func (s *PersonStorage) GetPersonById(id int) (Person_DB, error){
	person := Person_DB{}

	err := s.Conn.Get(&person, "SELECT * FROM Person WHERE id = ?", id)
	if err != nil {
		return person, err
	}

	return person, nil
}

func (s *PersonStorage) CreateNewPerson(person NewPersonInput) (int, error) {
	query := "insert into Person (firstname, middlename, lastname, birthdate, gender, city, country, photo) values (?, ?, ?, ?, ?, ?, ?, ?)"

	res, err := s.Conn.Exec(query, person.Firstname, person.Middlename, person.Lastname, person.Birthdate, person.Gender, person.City, person.Country, person.Photo)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *PersonStorage) UpdatePersonById(id int) error {
	_, err := s.Conn.Exec("UPDATE Person SET firstname=?, middlename=?, lastname=?, birthdate=?, gender=?, city=?, country=?, photo=? WHERE id=?", id)
	
	return err
}

func (s *PersonStorage) DeletePersonById(id int) error {
	_, err := s.Conn.Exec("DELETE FROM Person WHERE id=?", id)
	
	return err
}
