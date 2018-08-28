package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
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

	fmt.Println("register")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
}
