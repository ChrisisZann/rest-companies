package repository

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

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
