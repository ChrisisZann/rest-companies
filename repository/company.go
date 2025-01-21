package repository

func (s *Company) CreateNewCompany(name, description, ct string, amountOfEmployees int, registered bool) error {
	return repo.CreateNewCompany(name, description, ct, amountOfEmployees, registered)
}

func (s *Company) SelectSingleCompany(name string) (Company, error) {
	return repo.SelectSingleCompany(name)
}

func (s *Company) PatchCompanyName(name, input string) error {
	return repo.PatchCompanyName(name, input)
}

func (s *Company) PatchCompanyDescription(name, input string) error {
	return repo.PatchCompanyDescription(name, input)
}

func (s *Company) PatchCompanyAmtEmp(name string, input int) error {
	return repo.PatchCompanyAmtEmp(name, input)
}

func (s *Company) PatchCompanyReg(name string, input bool) error {
	return repo.PatchCompanyReg(name, input)
}

func (s *Company) PatchCompanyType(name, input string) error {
	return repo.PatchCompanyType(name, input)
}

func (s *Company) DeleteCompany(name string) error {
	return repo.DeleteCompany(name)
}
