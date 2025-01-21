package repository

import (
	"database/sql"
)

var repo Repository

type Models struct {
	Company Company
	User    User
}

func New(conn *sql.DB) *Models {
	if conn != nil {
		repo = NewPsqlRepo(conn)
	}
	// else {
	// 	repo = newTestRepository(conn)
	// }

	return &Models{
		Company: Company{},
		User:    User{},
	}
}
