package routes

import (
	"movepilot/cmd/api/handlers"
	"movepilotot/pkg/middleware"
	"movepilotot/pkg/output"

	"github.com/gorilla/mux"
)

func Register(
	//router
	r *mux.Router,
	//handlers
	authHandlers *handlers.AuthHandler,
	formHandlers *handlers.FormHandler,
	submissionHandlers *handlers.SubmissionHandler,

	//middlewares
	authFresh middleware.Middleware,
	authCached middleware.Middleware) {

	output.MakeSubRouter(r, "/auth", func(sr *mux.Router) {
		AuthRoutes(sr, authHandlers, authCached)
	})
	output.MakeSubRouter(r, "/form", func(sr *mux.Router) {
		FormRoutes(sr, formHandlers, authCached)
	})
	output.MakeSubRouter(r, "/submission", func(sr *mux.Router) {
		SubmissionRoutes(sr, submissionHandlers)
	})

}
