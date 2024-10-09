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
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)

		return []model.Person{}, err
	}

	var personList []model.Person
	var personObj model.Person

	for rows.Next() {
		err = rows.Scan(
			&personObj.ID,
			&personObj.Name,
			&personObj.Email,
		)

		if err != nil {
			fmt.Println(err)

			return []model.Person{}, err
		}

		personList = append(personList, personObj)
	}

	rows.Close()

	return personList, nil
}

func (pr *PersonRepository) CreatePerson(person model.Person) (int, error) {

	var id int

	query, err := pr.connection.Prepare("INSERT INTO person" +
		"(name, email)" +
		" VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(person.Name, person.Email).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *PersonRepository) GetPersonById(id_person int) (*model.Person, error) {

	query, err := pr.connection.Prepare("SELECT * FROM person WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var person model.Person

	err = query.QueryRow(id_person).Scan(
		&person.ID,
		&person.Name,
		&person.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	query.Close()

	return &person, nil
}
