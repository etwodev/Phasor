package redirect

import (
	"github.com/Etwodev/ramchi/router"
)

func Routes() []router.Route {
	return []router.Route{
		router.NewPostRoute("/", true, false, RedirectPostRoute),
		router.NewGetRoute("/{id}", true, false, RedirectGetRoute),
	}
}
