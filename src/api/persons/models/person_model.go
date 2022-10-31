package models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Person struct {
	ID        uuid.UUID    `json:"id" gorm:"primaryKey" gorm:"type:uuid"`
	Name      string       `json:"name" gorm:"type:varchar" gorm:"size:255"`
	Cpf       string       `json:"cpf" gorm:"type:varchar" gorm:"size:11"`
	Birthday  time.Time    `json:"birthday,string" gorm:"type:date"`
	Gender    string       `json:"gender" gorm:"type:varchar" gorm:"size:10"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt" gorm:"index"`
	Addresses []Address    `json:"addresses" gorm:"foreignKey:PersonRefer"`
}

type Address struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	StreetName  string    `json:"streetName"`
	Number      string    `json:"number"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	PersonRefer uuid.UUID `json:ignore`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt" gorm:"index"`
}

func (person *Person) BeforeCreate(tx *gorm.DB) error {
	person.ID = uuid.New()
	return nil
}

type PersonResult struct {
	Data         []Person `json:"data"`
	TotalRecords int64    `json:"totalRecords"`
	Pages        int64    `json:"pages"`
}

func (person *Person) FindById(id uuid.UUID) (int, *Address) {
	for i, address := range person.Addresses {
		if address.ID == id {
			return i, &address
		}
	}
	return -1, nil
}
