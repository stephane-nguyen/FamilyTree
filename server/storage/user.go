package storage

import (
	"github.com/jmoiron/sqlx"
)

type UserStorage struct {
	Conn *sqlx.DB
}

func NewUserStorage(conn *sqlx.DB) *UserStorage {
	return &UserStorage{Conn: conn}
}

type User_DB struct {
	Id          int    `db:"id" json:"id"`
    Username	string `db:"username" json:"username"`
	Email      string `db:"email" json:"email"`
	Password   string `db:"password" json:"password"`
}

type NewUserInput struct {
	Username string
	Password string
	Email    string
}

func (s *UserStorage) GetAllUsers() ([]User_DB, error){
	Users := []User_DB{}

	err := s.Conn.Select(&Users, "SELECT * FROM User")
	if err != nil {
		
		return Users, err
	}
	return Users, nil
}

func (s *UserStorage) GetUserById(id int) (User_DB, error){
	User := User_DB{}

	err := s.Conn.Get(&User, "SELECT * FROM User WHERE id = ?", id)
	if err != nil {
		return User, err
	}

	return User, nil
}

func (s *UserStorage) CreateNewUser(User NewUserInput) (int, error) {
	query := "insert into User (username, email, password) values (?, ?, ?)"

	res, err := s.Conn.Exec(query, User.Username, User.Email, User.Password)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *UserStorage) UpdateUserById(id int) error {
	_, err := s.Conn.Exec("UPDATE User SET username=?, password=?, email=? WHERE id=?", id)
	
	return err
}

func (s *UserStorage) DeleteUserById(id int) error {
	_, err := s.Conn.Exec("DELETE FROM User WHERE id=?", id)
	
	return err
}
