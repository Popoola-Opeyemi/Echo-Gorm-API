package controller

import "popstar/boilerplate/util"

type Handler interface {
	Start(util.Environment, string) error
}

func InitRouter(env util.Environment, prefix string) {

	routesHandlers := []Handler{
		&Auth{},
	}

	for _, routes := range routesHandlers {

		if err := routes.Start(env, prefix); err != nil {
			env.Log.Debug("router error", err)
		}
	}

	env.Log.Debug("routes initialized")
}
