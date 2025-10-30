package routes

import (
	"movepilot/cmd/api/handlers"
	"movepilot/pkg/middleware"
	"movepilot/pkg/output"

	"github.com/gorilla/mux"
)

func Register(
	//router
	r *mux.Router,
	//handlers
	authHandlers *handlers.AuthHandler,

	//middlewares
	authFresh middleware.Middleware,
	authCached middleware.Middleware) {

	output.MakeSubRouter(r, "/auth", func(sr *mux.Router) {
		AuthRoutes(sr, authHandlers, authCached)
	})
	output.MakeSubRouter(r, "/client", func(sr *mux.Router) {
		AuthRoutes(sr, authHandlers, authCached)
	})

}
