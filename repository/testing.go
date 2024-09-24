package repository

func (thisRepo *testRepository) CreateNewCompany(name, description, ct string, amountOfEmployees int, registered bool) error {
	return nil
}

func (thisRepo *testRepository) DeleteCompany(name string) error {
	return nil
}

func (thisRepo *testRepository) PatchCompanyName(name, input string) error {
	return nil
}

func (thisRepo *testRepository) PatchCompanyDescription(name, input string) error {
	return nil
}

func (thisRepo *testRepository) PatchCompanyAmtEmp(name string, input int) error {
	return nil
}

func (thisRepo *testRepository) PatchCompanyReg(name string, input bool) error {
	return nil
}

func (thisRepo *testRepository) PatchCompanyType(name, input string) error {
	return nil
}

func (thisRepo *testRepository) SelectSingleCompany(name string) (Company, error) {

	return Company{}, nil
}

func (thisRepo *testRepository) CreateNewUser(username, password string) error {
	return nil
}

func (thisRepo *testRepository) ValidateUser(username, password string) error {
	return nil
}
