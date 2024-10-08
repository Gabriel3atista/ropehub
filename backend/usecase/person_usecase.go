package usecase

import "server/model"

type PersonUsecase struct {
	// Repository
}

func NewPersonUseCase() PersonUsecase {
	return PersonUsecase{}
}

func (pu *PersonUsecase) GetPerson() ([]model.Person, error) {
	return []model.Person{}, nil
}
