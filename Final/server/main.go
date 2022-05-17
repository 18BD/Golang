package main

import (
	"grpcAsan/server/ctrl"
	"grpcAsan/server/db"
	"log"
)

const (
	network = "tcp"
	address = "localhost:9000"
)

func main() {
	dbConn, err := db.OpenConnection()
	if err != nil {
		log.Fatalln(err)
	}

	server, err := ctrl.New(network, address, dbConn)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(server)
}
