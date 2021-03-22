package util

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattes/migrate/database/postgres"

	"gopkg.in/ini.v1"
)

type DBConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
}

func InitDatabase(cfg *ini.File, logmode, singularTable bool) (*gorm.DB, error) {

	dbCfg := cfg.Section("database").KeysHash()

	fmt.Println(dbCfg["user"])

	dbConfig := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbCfg["dbhost"],
		dbCfg["user"],
		dbCfg["dbname"],
		dbCfg["password"])

	db, err := gorm.Open("postgres", dbConfig)

	if err != nil {
		fmt.Println("Could not connect to the database:", err)
		return db, err
	}

	db.LogMode(logmode)
	db.SingularTable(singularTable)

	return db, nil
}
