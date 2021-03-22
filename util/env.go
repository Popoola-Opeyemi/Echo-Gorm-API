package util

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

type SharedEnvironment struct {
	Log    *zap.SugaredLogger
	Config *ini.File
	Db     *gorm.DB
}

type Environment struct {
	Log    *zap.SugaredLogger
	Config *ini.File
	Db     *gorm.DB
	Echo   *echo.Echo
}

var AppEnv SharedEnvironment

func init() {
	if AppEnv == nil {
		AppEnv = SharedEnvironment{}
	}
}
