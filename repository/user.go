package repository

import (
	"errors"
	"log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:token`
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

	// token, err := GenerateJWT(username, jwtKey)
	// if err != nil {
	// 	return &User{}, err
	// }
	// thisRepo.tokens = append(thisRepo.tokens, token)
	// log.Printf("Saved token: %v for user: %v", token, username)

	// return &User{Username: username, Password: "******", Token: token}, nil
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
		return errors.New("wrong username" + username)
	}

	if user.Password != password {
		return errors.New("wrong password" + password)
	}
	// claims, err := ValidateJWT(token, jwtKey)
	if err != nil {
		return errors.New("Failed to validate user")
	}
	log.Println("Validated user:", username)

	return nil
}

// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.RegisteredClaims
// }

// func GenerateJWT(username string, signingKey []byte) (string, error) {
// 	claims := Claims{
// 		username,
// 		jwt.RegisteredClaims{
// 			// A usual scenario is to set the expiration time relative to the current time
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			NotBefore: jwt.NewNumericDate(time.Now()),
// 			Issuer:    username,
// 			Subject:   "companies",
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	ss, err := token.SignedString(signingKey)
// 	return ss, err

// }

// func ValidateJWT(str_token string, privateKey []byte) (*Claims, error) {
// 	claims := &Claims{}
// 	token, err := jwt.ParseWithClaims(str_token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return privateKey, nil
// 	})
// 	if err != nil {
// 		log.Println("error in ValidateJWT() : ParseWithClaims()", err)
// 		return nil, err
// 	} else if !token.Valid {
// 		return nil, errors.New("Invalid token")
// 	}
// 	return claims, nil
// }
