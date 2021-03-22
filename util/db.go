package util

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func InitDatabase(dbc DBConfig, logmode, singularTable bool, log *zap.SugaredLogger) *gorm.DB {

	dbConfig := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbc.Host, dbc.Username, dbc.DBName, dbc.Password)
	db, err := gorm.Open("postgres", dbConfig)

	db.LogMode(logmode)
	db.SingularTable(singularTable)

	if err != nil {
		log.Debug("Could not connect to the database:", err)
	}

	return db
}
