package main

import (
	"fmt"
	"os"
	"popstar/boilerplate/controller"
	"popstar/boilerplate/util"

	"gopkg.in/ini.v1"
)

func main() {

	cfg, err := ini.Load("my.ini")

	if err != nil {
		fmt.Printf("cant find config file my.ini: %s", err)
		os.Exit(1)
	}

	logger := util.InitLogger().Sugar()
	defer logger.Sync()

	db := util.InitDatabase(cfg, true, true)

	// setting the environment variables
	util.AppEnv.Config = cfg
	util.AppEnv.Log = logger
	util.AppEnv.Db = db

	echo := util.InitEcho()

	environment := util.Environment{
		Log:    logger,
		Config: cfg,
		Db:     db,
		Echo:   echo,
	}

	controller.InitRouter(environment, "/api")

	port := util.Getkey("server", "http_port")
	echo.Logger.Fatal(echo.Start(port))

}
