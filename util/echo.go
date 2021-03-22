package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitEcho() *echo.Echo {
	e := echo.New()

	port := Getkey("server", "http_port")
	origins := LoadList("server", "origins")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     origins,
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	lc := middleware.LoggerConfig{
		Format: `[${method}] ${status} - ${uri}` +
			` - ${latency_human}, rx:${bytes_in}, tx:${bytes_out}` + "\n",
	}
	e.Use(middleware.LoggerWithConfig(lc))

	e.Use(middleware.Recover())

	e.Use(session.Middleware(sessions.NewFilesystemStore("./tmp", []byte("!yours_truly_secrets_incantations_7070_7070_7070!"))))

	e.Validator = &CustomValidator{
		Validator: validator.New(),
	}

	controller.DefineRoutes(e, db, logger, "/api", cfg)
	logger.Debug("routes defined")

	port := fmt.Sprintf(":%s", cfg.Section("").Key("http_port").String())
	e.Logger.Fatal(e.Start(port))

	return e
}
