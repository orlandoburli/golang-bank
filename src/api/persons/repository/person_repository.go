package repository

import (
	"bank/api/persons/db"
	"bank/api/persons/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	_repository *PersonRepository
)

type PersonRepository struct {
	DB *gorm.DB
}

func GetPersonRepository() *PersonRepository {
	if _repository == nil {
		_repository = &PersonRepository{
			DB: db.OpenConnection(),
		}
	}
	return _repository
}

func (repository *PersonRepository) CreatePerson(person *models.Person) (*models.Person, error) {
	var result = repository.DB.Create(&person)

	if result.Error == nil {
		return person, nil
	} else {
		return nil, result.Error
	}
	return person, nil
}

func (repository *PersonRepository) UpdatePerson(id uuid.UUID, person *models.Person) (*models.Person, error) {
	var _person, found = repository.GetPerson(id)

	if found {
		copyAttributes(person, _person)

		var result = repository.DB.Save(&_person)

		if result.Error == nil {
			return _person, nil
		} else {
			return nil, result.Error
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Person not found with id %v", id))
	}
}

func copyAttributes(source *models.Person, target *models.Person) {
	target.Name = source.Name
	target.Cpf = source.Cpf
	target.Birthday = source.Birthday
	target.Gender = source.Gender

	for _, address := range source.Addresses {
		var targetAddress models.Address
		var _, ta = target.FindById(address.ID)

		if ta != nil {
			ta = &targetAddress
		} else {
			ta = &models.Address{}
			target.Addresses = append(target.Addresses, *ta)
		}

		targetAddress.StreetName = address.StreetName
		targetAddress.Number = address.Number
		targetAddress.City = address.City
		targetAddress.State = address.State
		targetAddress.Country = address.Country
		targetAddress.PersonRefer = source.ID
	}
}

func (repository *PersonRepository) DeletePerson(id *uuid.UUID) error {
	return repository.DB.Delete(&models.Person{}, id).Error
}

func (repository *PersonRepository) GetPersons(page int, size int) (*[]models.Person, *int64, *int64) {
	var persons = make([]models.Person, 0)
	var count int64
	var pages int64

	var result1 = repository.DB.Model(&models.Person{}).Count(&count)
	var result2 = repository.DB.Order("name ASC").Limit(size).Offset((page - 1) * size).Find(&persons)

	if result1.Error != nil {
		panic(result1.Error)
	}

	if result2.Error != nil {
		panic(result2.Error)
	}

	pages = int64(float64(count) / float64(size))
	var remain = (count) % int64(size)

	if remain > 0 {
		pages++
	}

	return &persons, &count, &pages
}

func (repository *PersonRepository) GetPerson(id uuid.UUID) (*models.Person, bool) {
	var person = models.Person{}

	result := repository.DB.First(&person, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, false
	}

	return &person, true
}
