package services

import (
	"bank/api/persons/models"
	"bank/api/persons/repository"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

var (
	_service *PersonService
)

type PersonService struct {
	repository *repository.PersonRepository
}

func GetPersonService() *PersonService {
	if _service == nil {
		_service = &PersonService{
			repository: repository.GetPersonRepository(),
		}
	}
	return _service
}

func (service *PersonService) GetPerson(id uuid.UUID) (*models.Person, error) {
	var person, found = service.repository.GetPerson(id)
	if found {
		return person, nil
	}
	return nil, errors.New(fmt.Sprintf("Person not found with id %v", id))
}

func (service *PersonService) GetPersons(page int, size int) (*[]models.Person, *int64, *int64) {
	return service.repository.GetPersons(page, size)
}

func (service *PersonService) CreatePerson(person *models.Person) (*models.Person, error) {
	result, e := validatePersonModel(person)

	if result {
		service.repository.CreatePerson(person)

		return person, nil
	} else {
		return nil, e
	}
}

func (service *PersonService) UpdatePerson(id uuid.UUID, person *models.Person) (*models.Person, error) {
	result, e := validatePersonModel(person)

	if result {
		var _person, err = service.repository.UpdatePerson(id, person)

		if err != nil {
			return nil, err
		}
		return _person, nil
	} else {
		return nil, e
	}
}

func (service *PersonService) DeletePerson(id *uuid.UUID) error {
	return service.repository.DeletePerson(id)
}

func validatePersonModel(person *models.Person) (bool, error) {
	if strings.TrimSpace(person.Name) == "" {
		return false, errors.New("name not provided")
	}
	return true, nil
}
