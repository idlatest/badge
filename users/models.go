package users

import (
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/idlatest/badge/common"
)

type User struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:created_at`
}

type Response struct {
	Token string `json:"token"`
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

func (u User) Get(email string) (User, error) {
	user := User{}

	db := common.Db()

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(email))
		if err != nil {
			return err
		}

		value, err := item.Value()
		if err != nil {
			return err
		}

		err = common.GetInterfaceFromGob(value, &user)

		return err
	})

	// fmt.Println("Login successful")
	return user, err
}
