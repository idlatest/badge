package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/idlatest/badge/common"
)

type User struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:created_at`
}

func (u User) Add() error {

	db := common.Db()

	tnx := db.NewTransaction(true)

	defer tnx.Discard()

	u.CreatedAt = time.Now()
	responseBytes, err := common.GetGobFromInterface(u)
	if err != nil {
		fmt.Println(err)
	}

	err = tnx.Set([]byte(u.Email), responseBytes)
	if err != nil {
		log.Fatal(err)
	}

	if err = tnx.Commit(nil); err != nil {
		log.Fatal(err)
	}

	return err
}

func (u User) login(w http.ResponseWriter, r *http.Request) {
	var user User

	_ = json.NewDecoder(r.Body).Decode(&user)

	fmt.Println("Login successful")
	return
}
