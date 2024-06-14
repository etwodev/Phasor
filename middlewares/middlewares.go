package middlewares

import (
	"github.com/Etwodev/ramchi/middleware"
)

func Middlewares() []middleware.Middleware {
	return []middleware.Middleware{
		middleware.NewMiddleware(Logger(), "logger", true, false),
	}
}
