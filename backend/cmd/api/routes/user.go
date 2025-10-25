package routes

import (
	"move-pilot/cmd/api/handlers"
	"move-pilotot/pkg/output"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	output.MakeRoute(r, "/", handlers.GetUsers).Methods("GET", "OPTIONS")
}
