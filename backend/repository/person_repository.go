package repository

import (
	"database/sql"
	"fmt"
	"server/model"
)

type PersonRepository struct {
	connection *sql.DB
}

func NewPersonRepository(connection *sql.DB) PersonRepository {
	return PersonRepository{
		connection: connection,
	}
}

func (pr *PersonRepository) GetPerson() ([]model.Person, error) {
	query := "SELECT id, name, email FROM person"
	rows, error := pr.connection.Query(query)

	if error != nil {
		fmt.Println(error)

		return []model.Person{}, error
	}

	var personList []model.Person
	var personObj model.Person

	for rows.Next() {
		error = rows.Scan(
			&personObj.ID,
			&personObj.Name,
			&personObj.Email,
		)

		if error != nil {
			fmt.Println(error)

			return []model.Person{}, error
		}

		personList = append(personList, personObj)
	}

	rows.Close()

	return personList, nil
}
