package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitEcho() *echo.Echo {
	e := echo.New()

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

	e.Use(session.Middleware(sessions.NewFilesystemStore("./tmp", []byte("ilovego!angtodm00n!"))))

	e.Validator = &CustomValidator{
		Validator: validator.New(),
	}

	return e
}
