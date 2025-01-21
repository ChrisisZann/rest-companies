package repository

import "database/sql"

// ================================================================================
// Repository interface
// --------------------------------------------------------------------------------
type Repository interface {
	// Company
	CreateNewCompany(name, description, ct string, amountOfEmployees int, registered bool) error
	DeleteCompany(name string) error
	PatchCompanyName(name, input string) error
	PatchCompanyDescription(name, input string) error
	PatchCompanyAmtEmp(name string, input int) error
	PatchCompanyReg(name string, input bool) error
	PatchCompanyType(name, input string) error
	SelectSingleCompany(name string) (Company, error)

	// User
	CreateNewUser(username, password string) error
	ValidateUser(username, password string) error
}

// ================================================================================
// postgres repository
// --------------------------------------------------------------------------------
type psqlRepo struct {
	DB *sql.DB
}

// psqlRepo implements Repository
func NewPsqlRepo(conn *sql.DB) Repository {
	return &psqlRepo{
		DB: conn,
	}
}

// ================================================================================
// test repository
// --------------------------------------------------------------------------------
type testRepository struct {
	DB *sql.DB
}

func newTestRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}

// ================================================================================
