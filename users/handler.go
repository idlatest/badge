package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	data, err := user.Get(user.Email)
	if err != nil {
		fmt.Println(err)
		return
	}

	if data.Email != "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Println("Email already exist")
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	user.Password = string(passwordHash)

	err = user.Add()
	if err != nil {
		fmt.Println(err)
	}

	_, tokenString, err := tokenAuth.Encode(jwtauth.Claims{"email": user.Email})
	if err != nil {
		fmt.Println(err)
	}

	response := Response{Token: tokenString}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	fmt.Println(tokenString)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	data, err := user.Get(user.Email)
	if err != nil {
		fmt.Println(err)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password)) != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Println("password not match")
		return
	}

	_, tokenString, err := tokenAuth.Encode(jwtauth.Claims{"email": user.Email})
	if err != nil {
		fmt.Println(err)
	}

	response := Response{Token: tokenString}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	fmt.Println("data", data)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {

}
