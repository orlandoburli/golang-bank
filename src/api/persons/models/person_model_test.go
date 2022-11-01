package models

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Should_Create_Person_With_Address(t *testing.T) {

	var person = Person{
		ID:        uuid.UUID{},
		Name:      "Orlando Burli",
		Cpf:       "12345678901",
		Birthday:  time.Date(1981, 02, 13, 0, 0, 0, 0, time.Local),
		Gender:    "Male",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: sql.NullTime{},
		Addresses: []Address{{
			ID:          uuid.UUID{},
			StreetName:  "Green Street",
			Number:      "65",
			City:        "Boston",
			State:       "Massachusetts",
			Country:     "EUA",
			PersonRefer: uuid.UUID{},
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   time.Time{},
		}},
	}

	assert.NotNil(t, person)
	assert.NotNil(t, person.ID)
	assert.Equal(t, person.Name, "Orlando Burli")
	assert.Equal(t, person.Cpf, "12345678901")
	assert.Equal(t, person.Birthday, time.Date(1981, 02, 13, 0, 0, 0, 0, time.Local))
	assert.NotNil(t, person.CreatedAt)
	assert.NotNil(t, person.UpdatedAt)
	assert.NotNil(t, person.DeletedAt)

	assert.NotNil(t, person.Addresses)
	assert.Len(t, person.Addresses, 1)
	assert.Equal(t, person.Addresses[0].StreetName, "Green Street")
	assert.Equal(t, person.Addresses[0].City, "Boston")
	assert.Equal(t, person.Addresses[0].State, "Massachusetts")
	assert.Equal(t, person.Addresses[0].Country, "EUA")

	assert.NotNil(t, person.Addresses[0].CreatedAt)
	assert.NotNil(t, person.Addresses[0].UpdatedAt)
	assert.NotNil(t, person.Addresses[0].DeletedAt)
}
