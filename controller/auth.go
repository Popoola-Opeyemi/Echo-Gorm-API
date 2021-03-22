package controller

import (
	"popstar/boilerplate/util"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

type Auth struct {
	Log    *zap.SugaredLogger
	Config *ini.File
	Db     *gorm.DB
}

func (s *Auth) Start(env util.Environment, prefix string) error {

	s.Log = env.Log
	s.Config = env.Config
	s.Db = env.Db

	e := env.Echo
	g := e.Group(prefix)
	g.GET("/auth", s.Test)

	return nil
}

func (s *Auth) Test(c echo.Context) error {

	c.JSON(200, "hello from test")

	return nil
}
