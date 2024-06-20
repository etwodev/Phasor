package main

import (
	"github.com/Etwodev/Phasor/crypto"
	"github.com/Etwodev/Phasor/middlewares"
	"github.com/Etwodev/Phasor/routes/auth"
	"github.com/Etwodev/Phasor/routes/redirect"

	"github.com/Etwodev/ramchi"
	"github.com/Etwodev/ramchi/middleware"
	"github.com/Etwodev/ramchi/router"
)

func main() {
	crypto.GenerateRSAKey(4096)
	s := ramchi.New()
	s.LoadRouter(Routers())
	s.LoadMiddleware(Middlewares())
	s.Start()
}

func Routers() []router.Router {
	return []router.Router{
		router.NewRouter(auth.Routes(), true),
		router.NewRouter(redirect.Routes(), true),
	}
}

func Middlewares() []middleware.Middleware {
	return middlewares.Middlewares()
}
