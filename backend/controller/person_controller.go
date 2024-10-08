package controller

import (
	"net/http"
	"server/model"
	"server/usecase"

	"github.com/gin-gonic/gin"
)

type personController struct {
	personUseCase usecase.PersonUsecase
}

func NewPersonController(usecase usecase.PersonUsecase) personController {
	return personController{
		personUseCase: usecase,
	}
}

func (p *personController) GetPerson(ctx *gin.Context) {
	person := model.Person{
		ID:    1,
		Name:  "Gabriel",
		Email: "gabriel3atista@gmail.com",
	}

	ctx.JSON(http.StatusOK, person)
}
