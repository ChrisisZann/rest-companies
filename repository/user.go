package repository

func (s *Company) CreateNewUser(username, password string) error {
	return repo.CreateNewUser(username, password)
}

func (s *Company) ValidateUser(username, password string) error {
	return repo.ValidateUser(username, password)
}
