package repository

func (s *User) CreateNewUser(username, password string) error {
	return repo.CreateNewUser(username, password)
}

func (s *User) ValidateUser(username, password string) error {
	return repo.ValidateUser(username, password)
}
