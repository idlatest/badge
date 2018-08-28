package common

import (
	"log"

	"github.com/dgraph-io/badger"
)

func Db() *badger.DB {
	opts := badger.DefaultOptions
	opts.Dir = "./database"
	opts.ValueDir = "./database"

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()

	return db
}
