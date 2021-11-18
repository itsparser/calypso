package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/workfoxes/calypso/pkg/config"
	"github.com/workfoxes/calypso/pkg/constant"
	"github.com/workfoxes/calypso/pkg/log"
	"github.com/workfoxes/calypso/pkg/utils"
)

var DB *gorm.DB

func getDSN() string {
	dsn := utils.StringFormat(constant.DB_DSN,
		config.Get(constant.DB_HOST),
		config.Get(constant.DB_USER),
		config.Get(constant.DB_PASSWORD),
		config.Get(constant.DB_NAME),
		config.Get(constant.DB_PORT),
		config.GetDefault(constant.DB_SSL, "disable"))
	return dsn
}

func Init() {
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
}

func Begin() *gorm.DB {
	return DB.Begin()
}
