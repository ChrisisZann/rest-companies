package repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type companyType int

const (
	Corporations companyType = iota
	NonProfit
	Cooperative
	SoleProprietorship
	Undefined
)

// Value - Implementation of valuer for database/sql
func (ct companyType) Value() (driver.Value, error) {
	return ct.String(), nil
}

// Scan - Implement the database/sql scanner interface
func (ct *companyType) Scan(value interface{}) error {
	if value == nil {
		*ct = Undefined
		return nil
	}
	if str_type, err := driver.String.ConvertValue(value); err == nil {

		if v, ok := str_type.(string); ok {
			*ct, err = parseStringFromCT(v)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("failed to scan YesNoEnum")
}

func parseStringFromCT(input string) (companyType, error) {
	switch input {
	case "Corporations":
		return Corporations, nil
	case "NonProfit":
		return NonProfit, nil
	case "Cooperative":
		return Cooperative, nil
	case "Sole Proprietorship":
		return SoleProprietorship, nil
	default:
		return Undefined, errors.New("error, invalid type")
	}
}

func (ct *companyType) String() string {
	switch *ct {
	case 0:
		return "Corporations"
	case 1:
		return "NonProfit"
	case 2:
		return "Cooperative"
	case 3:
		return "Sole Proprietorship"
	default:
		return "error"
	}
}

// UnmarshalJSON custom implementation
func (t *companyType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	case "Corporations":
		*t = Corporations
	case "NonProfit":
		*t = NonProfit
	case "Cooperative":
		*t = Cooperative
	case "Sole Proprietorship":
		*t = SoleProprietorship
	default:
		return fmt.Errorf("invalid company type")
	}
	return nil
}

// MarshalJSON custom implementation
func (t companyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

type Company struct {
	Uuid              string      `json:"uuid"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	AmountOfEmployees int         `json:"amount_of_employees"`
	Registered        bool        `json:"registered"`
	Type              companyType `json:"type"`
	SysCreationDate   time.Time   `json:"sys_creation_date"`
	SysUpdateDate     time.Time   `json:"sys_update_date"`
}

func (s *Company) CreateNewCompany(name, description, ct string, amountOfEmployees int, registered bool) error {
	return repo.CreateNewCompany(name, description, ct, amountOfEmployees, registered)
}

func (thisRepo *psqlRepo) CreateNewCompany(name, description, ct string, amountOfEmployees int, registered bool) error {

	query := `INSERT INTO xm.companies(
	company_uuid,company_name, description, amount_of_employees, registered,company_type) 
	VALUES ($1,$2,$3,$4,$5,$6);`

	res, err := thisRepo.DB.Exec(query, uuid.New().String(), name, description, amountOfEmployees, registered, ct)
	if err != nil {
		log.Println("CreateNewCompany() :  PSQL : ", err)
		return errors.New("psql error")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("CreateNewCompany() : Failed to read result rows")
	}
	if rowsAffected == 0 {
		return errors.New("failed to create new sensor")
	}
	return nil
}

func (s *Company) SelectSingleCompany(name string) (Company, error) {
	return repo.SelectSingleCompany(name)
}

func (thisRepo *psqlRepo) SelectSingleCompany(name string) (Company, error) {

	query := `SELECT company_uuid,company_name, COALESCE(description,''),amount_of_employees, registered,company_type,sys_creation_date, sys_update_date
	FROM xm.companies WHERE company_name=$1;`

	rows, err := thisRepo.DB.Query(query, name)
	if err != nil {
		log.Println("Error in: SelectSingleCompany()", err)
	}
	// fmt.Println("rows: ", rows)
	defer rows.Close()

	rows.Next()
	var company Company
	err = rows.Scan(
		&company.Uuid,
		&company.Name,
		&company.Description,
		&company.AmountOfEmployees,
		&company.Registered,
		&company.Type,
		&company.SysCreationDate,
		&company.SysUpdateDate,
	)
	if err != nil {
		return company, err
	}
	return company, nil
}

func (s *Company) PatchCompanyName(name, input string) error {
	return repo.PatchCompanyName(name, input)
}
func (thisRepo *psqlRepo) PatchCompanyName(name, input string) error {
	query := `UPDATE xm.companies SET company_name=$2, sys_update_date=CURRENT_TIMESTAMP WHERE company_name=$1;`

	rows, err := thisRepo.DB.Query(query, name, input)
	defer rows.Close()
	if err != nil {
		log.Println("Error in: PatchCompanyName()", err)
		return err
	}
	log.Println("Updates company:", name, "to", input)
	return nil
}

func (s *Company) PatchCompanyDescription(name, input string) error {
	return repo.PatchCompanyDescription(name, input)
}
func (thisRepo *psqlRepo) PatchCompanyDescription(name, input string) error {
	query := `UPDATE xm.companies SET description=$2, sys_update_date=CURRENT_TIMESTAMP WHERE company_name=$1;`

	rows, err := thisRepo.DB.Query(query, name, input)
	defer rows.Close()
	if err != nil {
		log.Println("Error in: PatchCompanyName()", err)
		return err
	}
	log.Println("Updates company:", name, "to", input)
	return nil
}

func (s *Company) PatchCompanyAmtEmp(name string, input int) error {
	return repo.PatchCompanyAmtEmp(name, input)
}
func (thisRepo *psqlRepo) PatchCompanyAmtEmp(name string, input int) error {
	query := `UPDATE xm.companies SET amountOfEmployees=$2, sys_update_date=CURRENT_TIMESTAMP WHERE company_name=$1;`

	rows, err := thisRepo.DB.Query(query, name, input)
	defer rows.Close()
	if err != nil {
		log.Println("Error in: PatchCompanyName()", err)
		return err
	}
	log.Println("Updates company:", name, "to", input)
	return nil
}

func (s *Company) PatchCompanyReg(name string, input bool) error {
	return repo.PatchCompanyReg(name, input)
}
func (thisRepo *psqlRepo) PatchCompanyReg(name string, input bool) error {
	query := `UPDATE xm.companies SET registered=$2, sys_update_date=CURRENT_TIMESTAMP WHERE company_name=$1;`

	rows, err := thisRepo.DB.Query(query, name, input)
	defer rows.Close()
	if err != nil {
		log.Println("Error in: PatchCompanyName()", err)
		return err
	}
	log.Println("Updates company:", name, "to", input)
	return nil
}

func (s *Company) PatchCompanyType(name, input string) error {
	return repo.PatchCompanyType(name, input)
}
func (thisRepo *psqlRepo) PatchCompanyType(name, input string) error {

	_, err := parseStringFromCT(input)
	if err != nil {
		log.Println("received unsupported company type")
		return err
	}

	query := `UPDATE xm.companies SET company_type=$2, sys_update_date=CURRENT_TIMESTAMP WHERE company_name=$1;`

	rows, err := thisRepo.DB.Query(query, name, input)
	defer rows.Close()
	if err != nil {
		log.Println("Error in: PatchCompanyName()", err)
		return err
	}
	log.Println("Updates company:", name, "to", input)
	return nil
}

func (s *Company) DeleteCompany(name string) error {
	return repo.DeleteCompany(name)
}

func (thisRepo *psqlRepo) DeleteCompany(name string) error {

	query := `DELETE FROM xm.companies WHERE company_name=$1;`

	rows, err := thisRepo.DB.Query(query, name)
	defer rows.Close()
	if err != nil {
		log.Println("Error in: DeleteCompany()", err)
		return err
	}
	log.Println("Deleted company:", name)
	return nil
}
