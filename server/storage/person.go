package storage

import (
	"time"
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
	Birthdate 	time.Time `json:"birthdate"`
    Gender      string `json:"gender"`
    City        string `json:"city"`
    Country     string `json:"country"`
	Photo		*[]byte`json:"photo"`
}
// type Person_DB struct {
// 	Id			int 	  `json:"id" db:"id"`
//     Firstname   string    `json:"firstname" db:"firstname"`
//     Middlename  string    `json:"middlename" db:"middlename"`
//     Lastname    string    `json:"lastname" db:"lastname"`
//     Birthdate   time.Time `json:"birthdate" db:"birthdate"`
//     Gender      string    `json:"gender" db:"gender"`
//     City        string    `json:"city" db:"city"`
//     Country     string    `json:"country" db:"country"`
//     Photo       *[]byte   `json:"photo" db:"photo"`
// }

func (s *PersonStorage) GetAllPersons() ([]Person_DB, error){
	persons := []Person_DB{}
	err := s.Conn.Select(&persons, "SELECT * FROM Person")
	if err != nil {
		return persons, err
	}
	return persons, nil
}

// func (s *PersonStorage) CreateNewPerson(person NewPerson) (int, error) {
// 	query := "insert into person (firstname, middlename, lastname, birthdate, gender, city, country, photo) values (?, ?, ?, ?, ?, ?, ?, ?)"
// 	res, err := s.Conn.Exec(query, person.firstname, person.middlename, person.lastname, person.birthdate, person.gender, person.city, person.country, person.photo)
// 	if err != nil {
// 		return 0, err
// 	}
// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return int(id), nil
// }