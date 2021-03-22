package util

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
)

type DBConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
}

func InitDatabase(cfg *ini.File, logmode, singularTable bool) *gorm.DB {

	dbCfg := cfg.Section("db").KeysHash()

	dbConfig := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbCfg["user"],
		dbCfg["dbname"],
		dbCfg["password"])

	db, err := gorm.Open("postgres", dbConfig)

	db.LogMode(logmode)
	db.SingularTable(singularTable)

	if err != nil {
		AppEnv.Log.Debug("Could not connect to the database:", err)
	}

	return db
}
