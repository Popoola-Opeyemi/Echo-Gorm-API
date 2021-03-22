package main

import (
	"fmt"
	"os"
	"popstar/boilerplate/util"

	"gopkg.in/ini.v1"
)

func main() {

	cfg, err := ini.Load("my.ini")

	if err != nil {
		fmt.Printf("cant find config file my.ini: %s", err)
		os.Exit(1)
	}

	// init zap logger
	logger := util.InitLogger()
	defer logger.Sync()

	echo := util.InitEcho()

	_ = echo

}
