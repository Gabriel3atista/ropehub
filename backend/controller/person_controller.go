package controller

import (
	"net/http"
	"server/model"
	"server/usecase"
	"strconv"

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

	person, err := p.personUseCase.GetPerson()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, person)
}

func (p *personController) CreatePerson(ctx *gin.Context) {

	var person model.Person

	err := ctx.BindJSON(&person)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertPerson, err := p.personUseCase.CreatePerson(person)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertPerson)
}

func (p *personController) GetPersonById(ctx *gin.Context) {

	id := ctx.Param("id_person")

	if id == "" {
		response := model.Response{
			Message: "id_person is missing",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	personId, err := strconv.Atoi(id)

	if err != nil {
		response := model.Response{
			Message: "id_person must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	person, err := p.personUseCase.GetPersonById(personId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if person == nil {
		response := model.Response{
			Message: "person not found in database",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, person)
}
