package controller

import (
	"popstar/boilerplate/util"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

type Auth struct {
	Log    *zap.SugaredLogger
	Config *ini.File
	Db     *gorm.DB
	Echo   *echo.Echo
}

func (s *Auth) Start(env util.Environment, prefix string) error {

	return nil
}
