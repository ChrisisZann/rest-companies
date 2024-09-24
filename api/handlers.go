package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"xm-companies/events"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tsawler/toolbox"
)

// type MiddlewareSrv struct{}

func (a *api) rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello There, this is the root response from xm-company-api!")
}

func (a *api) company(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		var t toolbox.Tools
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprintln(w, http.StatusBadRequest, "Need to provide a name")
			return
		}

		company, err := a.config.Models.Company.SelectSingleCompany(name)
		if err != nil {
			_ = t.ErrorJSON(w, err, http.StatusBadRequest) // http.NotFound(w, r)
		}
		if company.Uuid == "" {
			fmt.Fprintln(w, http.StatusNoContent, "No Content")
			return
		}
		_ = t.WriteJSON(w, http.StatusOK, company)

	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func (a *api) auth_company(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST":
		var t toolbox.Tools
		err := r.ParseForm()
		if err != nil {
			log.Panic("failed to parse form data", err)
			_ = t.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		description := r.FormValue("description")
		registered := r.FormValue("registered")
		companyType := r.FormValue("type")

		i_amountOfEmployees, err := strconv.Atoi(r.FormValue("amount_of_employees"))
		if err != nil {
			log.Println("ERROR:", err)
			_ = t.ErrorJSON(w, err, http.StatusBadRequest)
			// http.Error(w, "bad argument", http.StatusBadRequest)
			return
		}

		b_registered, err := strconv.ParseBool(registered)
		if err != nil {
			log.Println("ERROR:", err)
			_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		}

		err = a.config.Models.Company.CreateNewCompany(name, description, companyType, i_amountOfEmployees, b_registered)
		if err != nil {
			log.Println("ERROR:", err)
			_ = t.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		_ = t.WriteJSON(w, http.StatusCreated, nil)
		// w.WriteHeader(http.StatusCreated)

	// case "PUT":

	case "PATCH":
		var t toolbox.Tools
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprintln(w, http.StatusBadRequest, "Need to provide a name")
			return
		}

		field := r.URL.Query().Get("field")
		if name == "" {
			fmt.Fprintln(w, http.StatusBadRequest, "Need to provide a field to update")
			return
		}

		value := r.URL.Query().Get("value")
		if name == "" {
			fmt.Fprintln(w, http.StatusBadRequest, "Need to provide a new value")
			return
		}

		switch field {
		case "name":
			err := a.config.Models.Company.PatchCompanyName(name, value)
			if err != nil {
				_ = t.ErrorJSON(w, err, http.StatusBadRequest) // http.NotFound(w, r)
			}
			_ = t.WriteJSON(w, http.StatusOK, "patched with new name: "+value)
		case "description":
			err := a.config.Models.Company.PatchCompanyDescription(name, value)
			if err != nil {
				_ = t.ErrorJSON(w, err, http.StatusBadRequest) // http.NotFound(w, r)
			}
		case "amount_of_employees":
			i_value, err := strconv.Atoi(r.FormValue("amount_of_employees"))
			if err != nil {
				log.Println("ERROR:", err)
				_ = t.ErrorJSON(w, err, http.StatusBadRequest)
				// http.Error(w, "bad argument", http.StatusBadRequest)
				return
			}

			err = a.config.Models.Company.PatchCompanyAmtEmp(name, i_value)
			if err != nil {
				_ = t.ErrorJSON(w, err, http.StatusBadRequest) // http.NotFound(w, r)
			}
		case "registered":
			b_value, err := strconv.ParseBool(value)
			if err != nil {
				log.Println("ERROR:", err)
				_ = t.ErrorJSON(w, err, http.StatusBadRequest)
			}

			err = a.config.Models.Company.PatchCompanyReg(name, b_value)
			if err != nil {
				_ = t.ErrorJSON(w, err, http.StatusBadRequest) // http.NotFound(w, r)
			}
		case "type":
			err := a.config.Models.Company.PatchCompanyName(name, value)
			if err != nil {
				_ = t.ErrorJSON(w, err, http.StatusBadRequest) // http.NotFound(w, r)
			}
		default:
			_ = t.ErrorJSON(w, errors.New("invalid field:"+field), http.StatusBadRequest) // http.NotFound(w, r)
		}

	case "DELETE":
		var t toolbox.Tools
		name := r.URL.Query().Get("name")

		err := a.config.Models.Company.DeleteCompany(name)
		if err != nil {
			_ = t.ErrorJSON(w, err, http.StatusBadRequest) // http.NotFound(w, r)
		}

		e := events.Event{Type: "delete", Payload: []byte(name)}
		a.hub.PublishEvent(e)
		a.internalPublisher.WriteStream(e)

		_ = t.WriteJSON(w, http.StatusOK, "deleted "+name)

	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func (a *api) user(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST":
		var t toolbox.Tools
		err := r.ParseForm()
		if err != nil {
			log.Panic("failed to parse form data", err)
			_ = t.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")

		err = a.config.Models.Company.CreateNewUser(username, password)
		if err != nil {
			log.Println("ERROR:", err)
			_ = t.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		token, err := GenerateJWT(username, a.config.JwtKey)
		if err != nil {
			log.Println("Failed to generate token")
			return
		}
		log.Printf("Created token: %v for user: %v", token, username)
		// w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})

		// _ = t.WriteJSON(w, http.StatusCreated, username)

	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func (a *api) login(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST":
		var t toolbox.Tools
		err := r.ParseForm()
		if err != nil {
			log.Panic("failed to parse form data", err)
			_ = t.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")

		ctx := r.Context()
		if val, ok := ctx.Value("username").(string); ok {
			log.Printf("user %s already authenticated", val)
			_ = t.WriteJSON(w, http.StatusAccepted, "successful validation")
			return
		}

		err = a.config.Models.Company.ValidateUser(username, password)
		if err != nil {
			log.Println("ERROR:", err)
			_ = t.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}
		// e := events.Event{
		// 	Type:    "login",
		// 	Payload: []byte(username),
		// }
		e := events.Event{Type: "login", Payload: []byte(username)}
		a.hub.PublishEvent(e)
		a.internalPublisher.WriteStream(e)

		_ = t.WriteJSON(w, http.StatusAccepted, "successful validation")

	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func (a *api) ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	username := r.Context().Value("username").(string)
	_ = t.WriteJSON(w, http.StatusAccepted, "successful authentication: "+username)
}

func (a *api) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		str_token := r.Header.Get("Authorization")

		if str_token == "" {
			http.Error(w, "Missing Tokenb", http.StatusUnauthorized)
			return
		}
		// log.Println("Auth(): Authorization:", str_token)

		claims := &XmClaims{}

		// log.Println("Auth(): Parsing XmClaims:", str_token)

		claims, err := ValidateJWT(str_token, a.config.JwtKey)
		if err != nil {
			log.Println("error:", err)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type XmClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string, signingKey []byte) (string, error) {
	claims := XmClaims{
		username,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "yippee",
			Subject:   "companies",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println("signingKey:", signingKey)
	ss, err := token.SignedString(signingKey)
	return ss, err

}

func ValidateJWT(str_token string, privateKey []byte) (*XmClaims, error) {
	// log.Println("ValidatingJWT")
	// log.Println("str_token", str_token)
	// log.Println("privateKey", string(privateKey))

	var token_only string
	if strings.Contains(str_token, " ") {
		parts := strings.Split(str_token, " ")

		if len(parts) == 2 {
			token_only = parts[1]
			// log.Println("Key:", token_only)
			log.Println("Ignoring authendication scheme:", parts[0])
		} else {
			log.Println("Invalid authorization header format")
		}
	} else {
		token_only = str_token
	}

	token, err := jwt.ParseWithClaims(token_only, &XmClaims{}, func(token *jwt.Token) (interface{}, error) {
		// return privateKey, nil
		return privateKey, nil
	})
	if err != nil {
		log.Println("error in ValidateJWT() : ParseWithClaims()", err)
		return nil, err
	} else if !token.Valid {
		return nil, errors.New("Invalid token")
	}
	claims, ok := token.Claims.(*XmClaims)
	if ok {
		log.Println("Username:", claims.Username, "Issuer:", claims.RegisteredClaims.Issuer, "Subject:", claims.RegisteredClaims.Subject)
	}
	return claims, nil
}
