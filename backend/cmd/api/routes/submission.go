package routes

import (
	"movepilot/cmd/api/handlers"
	"movepilot/pkg/output"

	"github.com/gorilla/mux"
)

func SubmissionRoutes(r *mux.Router, h *handlers.SubmissionHandler) {
	output.MakeRoute(r, "/{uuid}", h.GetForm).Methods("GET", "OPTIONS")
	output.MakeRoute(r, "/{uuid}/submit", h.GetForm).Methods("POST", "OPTIONS")
}
