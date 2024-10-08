package main

import (
	"server/controller"
	"server/db"
	"server/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, error := db.ConnectDB()

	if error != nil {
		panic(error)
	}

	PersonUseCase := usecase.NewPersonUseCase()

	PersonController := controller.NewPersonController(PersonUseCase)

	server.GET("/person", PersonController.GetPerson)

	server.Run(":8080")
}
