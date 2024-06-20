package auth

import (
	"github.com/Etwodev/ramchi/router"
)

func Routes() []router.Route {
	return []router.Route{
		router.NewGetRoute("/auth/id_rsa.pub", true, false, PubKeyGetRoute),
	}
}
