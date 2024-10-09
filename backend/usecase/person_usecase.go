package usecase

import (
	"server/model"
	"server/repository"
)

type PersonUsecase struct {
	repository repository.PersonRepository
}

func NewPersonUseCase(repo repository.PersonRepository) PersonUsecase {
	return PersonUsecase{
		repository: repo,
	}
}

func (pu *PersonUsecase) GetPerson() ([]model.Person, error) {
	return pu.repository.GetPerson()
}

func (pu *PersonUsecase) CreatePerson(person model.Person) (model.Person, error) {
	personId, err := pu.repository.CreatePerson(person)

	if err != nil {
		return model.Person{}, err
	}

	person.ID = personId

	return person, nil
}

func (pu *PersonUsecase) GetPersonById(id_person int) (*model.Person, error) {
	person, err := pu.repository.GetPersonById(id_person)

	if err != nil {
		return nil, err
	}

	return person, nil
}
