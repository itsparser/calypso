package db

import (
	"github.com/workfoxes/calypso/pkg/config"
	"github.com/workfoxes/calypso/pkg/constant"
	"github.com/workfoxes/calypso/pkg/utils"
)

func getDSN() string {
	dsn := utils.StringFormat(constant.DB_DSN,
		config.GetValue(constant.DbHost),
		config.GetValue(constant.DbUser),
		config.GetValue(constant.DbPassword),
		config.GetValue(constant.DbName),
		config.GetValue(constant.DbPort),
		config.GetDefault(constant.DbSsl, "disable"))
	return dsn
}
