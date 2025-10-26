package routes

import (
	"movepilot/cmd/api/handlers"
	"movepilot/pkg/middleware"
	"movepilot/pkg/output"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, h *handlers.AuthHandler, authCached middleware.Middleware) {
	output.MakeRoute(r, "/crm/sign-in", h.CRMSignIn).Methods("POST", "OPTIONS")
	output.MakeRoute(r, "/crm/initialize", h.CRMInitialize, authCached).Methods("GET", "OPTIONS")
}
