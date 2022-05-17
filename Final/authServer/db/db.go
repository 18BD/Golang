package database

import (
	"context"
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


func OpenConn() (*pg.DB, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	options := &pg.Options{
		User:     user,
		Password: password,
		Addr:     address,
		Database: dbname,
	}

	con := pg.Connect(options)
	if con == nil {
		return nil, errors.New("cannot connect to postgres")
	}

	ctx := context.Background()

	if err := con.Ping(ctx); err != nil {
		return nil, err
	}

	return con, nil
}
