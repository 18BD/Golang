package db

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
)

const (
	host = "127.0.0.1"
	port = "5432"
	user = "postgres"
	password = "12345"
	dbname = "Library"
)


func OpenConnection() (*pg.DB, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	options := &pg.Options{
		User:     user,
		Password: password,
		Addr:     address,
		Database: dbname,
		PoolSize: 50,
	}

	con := pg.Connect(options)
	if con == nil {
		return nil, errors.New("cannot connect to postgres")
	}

	return con, nil
}
