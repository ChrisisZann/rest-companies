package repository

import (
	"errors"
	"log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (s *Company) CreateNewUser(username, password string) error {
	return repo.CreateNewUser(username, password)
}

func (thisRepo *psqlRepo) CreateNewUser(username, password string) error {
	query := `INSERT INTO xm.users(
		username, password) 
		VALUES ($1,$2);`

	res, err := thisRepo.DB.Exec(query, username, password)
	if err != nil {
		log.Println("CreateNewUser() :  PSQL : ", err)
		return errors.New("psql error")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("CreateNewCompany() : Failed to read result rows")
	}
	if rowsAffected == 0 {
		return errors.New("failed to create new user")
	}
	return nil
}

func (s *Company) ValidateUser(username, password string) error {
	return repo.ValidateUser(username, password)
}

func (thisRepo *psqlRepo) ValidateUser(username, password string) error {
	query := `select * FROM xm.users WHERE username=$1;`

	rows, err := thisRepo.DB.Query(query, username)
	if err != nil {
		log.Println("ValidateUser() :  PSQL : ", err)
		return errors.New("psql error")
	}
	defer rows.Close()

	rows.Next()
	var user User
	err = rows.Scan(
		&user.Username,
		&user.Password,
	)
	if err != nil {
		log.Println("ValidateUser() :  rows.Scan : ", err)
		return errors.New("Failed to validate user, wrong username" + username)
	}

	if user.Password != password {
		return errors.New("wrong password" + password)
	}
	log.Println("Validated user:", username)

	return nil
}
