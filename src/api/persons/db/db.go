package db

import (
	"bank/api/persons/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type ConnectionPoolInfo struct {
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxLifetime time.Duration
	ConnectionMaxIdleTime time.Duration
}

func OpenConnection() *gorm.DB {
	url := "host=localhost user=postgres password=b4nk4p1 dbname=persons-db port=5432 sslmode=disable application_name=bank-api-persons TimeZone=America/Sao_Paulo"

	db, e := gorm.Open(postgres.Open(url), &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           schema.NamingStrategy{},
		FullSaveAssociations:                     true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	handleError(e)

	ConfigureConnectionPool(db)

	return db
}

func ConfigureConnectionPool(db *gorm.DB) {
	sqlDb, e := db.DB()

	handleError(e)

	info := GetConnectionPoolInfo()

	sqlDb.SetMaxOpenConns(info.MaxOpenConnections)
	sqlDb.SetMaxIdleConns(info.MaxIdleConnections)
	sqlDb.SetConnMaxLifetime(info.ConnectionMaxLifetime)
	sqlDb.SetConnMaxIdleTime(info.ConnectionMaxIdleTime)
}

func GetConnectionPoolInfo() ConnectionPoolInfo {
	// TODO: Get this info externally
	return ConnectionPoolInfo{
		MaxOpenConnections:    10,
		MaxIdleConnections:    10,
		ConnectionMaxLifetime: time.Minute,
		ConnectionMaxIdleTime: 10 * time.Minute,
	}
}

func Migrate() {
	var conn = OpenConnection()
	migrateModels(conn)
	shutdown(conn)
}

func migrateModels(db *gorm.DB) *gorm.DB {
	e := db.AutoMigrate(&models.Person{}, &models.Address{})

	handleError(e)

	return db
}

func shutdown(db *gorm.DB) {
	var sqlDB, err = db.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.Close()
}

func handleError(e error) {
	if e != nil {
		panic(fmt.Sprintf("Error connecting to database: %v", e))
	}
}
