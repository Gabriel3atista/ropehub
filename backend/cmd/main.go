package main

import (
	"server/controller"
	"server/db"
	"server/repository"
	"server/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	PersonRepository := repository.NewPersonRepository(dbConnection)
	PersonUseCase := usecase.NewPersonUseCase(PersonRepository)
	PersonController := controller.NewPersonController(PersonUseCase)

	router.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.GET("/persons", PersonController.GetPerson)
	router.POST("/person", PersonController.CreatePerson)
	router.GET("/person/:id_person", PersonController.GetPersonById)

	router.Run(":8080")
}
