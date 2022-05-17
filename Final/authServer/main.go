package main

import (
	"github.com/labstack/echo/v4"
	"grpcAsan/authServer/ctrl"
	database "grpcAsan/authServer/db"
	"log"
)



func main() {
	dbConn, err := database.OpenConn()
	if err != nil {
		log.Fatalln(err)
	}

	controller := ctrl.New(dbConn)

	e := echo.New()
	e.POST("/", controller.AuthUser)
	e.POST("/adduser", controller.AddUser)
	e.Logger.Fatal(e.Start(":3000"))
}
