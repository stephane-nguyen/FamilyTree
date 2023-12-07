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
	Id          int    `db:"id" json:"id"`
    UserId      *int   `db:"user_id" json:"user_id"` 
    Firstname   string `db:"firstname" json:"firstname"`
    Middlename  string `db:"middlename" json:"middlename"`
    Lastname    string `db:"lastname" json:"lastname"`
    Birthdate   string `db:"birthdate" json:"birthdate"`
    Gender      string `db:"gender" json:"gender"`
    City        string `db:"city" json:"city"`
    Country     string `db:"country" json:"country"`
    Description *string `db:"description" json:"description"`
    Photo       *[]byte`db:"photo" json:"photo"`
}

type NewPersonInput struct {
	Firstname   string 
    Middlename  string 
    Lastname    string 
	Birthdate 	string 
    Gender      string 
    City        string 
    Country     string 
	Description string
    Photo     	string
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
	query := "insert into Person (firstname, middlename, lastname, birthdate, gender, city, country, description, photo) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	res, err := s.Conn.Exec(query, person.Firstname, person.Middlename, person.Lastname, person.Birthdate, person.Gender, person.City, person.Country, person.Description, person.Photo)
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
	_, err := s.Conn.Exec("UPDATE Person SET firstname=?, middlename=?, lastname=?, birthdate=?, gender=?, city=?, country=?, description=?, photo=? WHERE id=?", id)
	
	return err
}

func (s *PersonStorage) DeletePersonById(id int) error {
	_, err := s.Conn.Exec("DELETE FROM Person WHERE id=?", id)
	
	return err
}
