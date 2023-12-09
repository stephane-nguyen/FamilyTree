package storage

import (
	"github.com/jmoiron/sqlx"
)

type MemberStorage struct {
	Conn *sqlx.DB
}

func NewMemberStorage(conn *sqlx.DB) *MemberStorage {
	return &MemberStorage{Conn: conn}
}

type Member_DB struct {
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

type NewMemberInput struct {
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

func (s *MemberStorage) GetAllMembers() ([]Member_DB, error){
	members := []Member_DB{}

	err := s.Conn.Select(&members, "SELECT * FROM Member")
	if err != nil {
		
		return members, err
	}
	return members, nil
}

func (s *MemberStorage) GetMembersByCountry(country string) ([]Member_DB, error){
	members := []Member_DB{}

	err := s.Conn.Select(&members, "SELECT * FROM Member WHERE country = ?", country)
	if err != nil {
		return members, err
	}

	return members, nil
}

func (s *MemberStorage) GetMemberById(id int) (Member_DB, error){
	member := Member_DB{}

	err := s.Conn.Get(&member, "SELECT * FROM Member WHERE id = ?", id)
	if err != nil {
		return member, err
	}

	return member, nil
}

func (s *MemberStorage) CreateNewMember(Member NewMemberInput) (int, error) {
	query := "insert into Member (firstname, middlename, lastname, birthdate, gender, city, country, description, photo) values (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	res, err := s.Conn.Exec(query, Member.Firstname, Member.Middlename, Member.Lastname, Member.Birthdate, Member.Gender, Member.City, Member.Country, Member.Description, Member.Photo)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *MemberStorage) UpdateMemberById(id int) error {
	_, err := s.Conn.Exec("UPDATE Member SET firstname=?, middlename=?, lastname=?, birthdate=?, gender=?, city=?, country=?, description=?, photo=? WHERE id=?", id)
	
	return err
}

func (s *MemberStorage) DeleteMemberById(id int) error {
	_, err := s.Conn.Exec("DELETE FROM Member WHERE id=?", id)
	
	return err
}
