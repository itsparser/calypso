package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/workfoxes/calypso/pkg/log"
)

var DB *gorm.DB

type Database struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	log.Info(getDSN())
	DB, err := gorm.Open(postgres.Open(getDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connect to the database : ", err)
	}
	pgDB, err := DB.DB()
	if err != nil {
		log.Panic("Error connect to the DB : ", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	pgDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	pgDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	pgDB.SetConnMaxLifetime(time.Hour)
	return DB
}

func Begin() *gorm.DB {
	return DB.Begin()
}
