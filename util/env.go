package util

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"
)

type SharedEnvironment struct {
	Log    *zap.SugaredLogger
	Config *ini.File
	Db     *gorm.DB
}

var AppEnv SharedEnvironment

func init() {
	if AppEnv == nil {
		AppEnv = SharedEnvironment{}
	}
}
