package routes

import (
	"movepilot/cmd/api/handlers"
	"movepilotot/pkg/output"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	output.MakeRoute(r, "/", handlers.GetUsers).Methods("GET", "OPTIONS")
}
