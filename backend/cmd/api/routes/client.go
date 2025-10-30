package routes

import (
	"movepilot/cmd/api/handlers"
	"movepilot/pkg/middleware"
	"movepilot/pkg/output"

	"github.com/gorilla/mux"
)

func ClientRoutes(r *mux.Router, h *handlers.ClientHandler, authCached middleware.Middleware) {
	output.MakeRoute(r, "/active", h.GetActive, authCached).Methods("GET", "OPTIONS")
	output.MakeRoute(r, "/archived", h.GetArchived, authCached).Methods("GET", "OPTIONS")
	output.MakeRoute(r, "/create", h.Create, authCached).Methods("POST", "OPTIONS")
	output.MakeRoute(r, "/view/{uuid}", h.View, authCached).Methods("GET", "OPTIONS")
	output.MakeRoute(r, "/update/{uuid}", h.Update, authCached).Methods("PUT", "OPTIONS")
	output.MakeRoute(r, "/archive/{uuid}", h.Archive, authCached).Methods("PUT", "OPTIONS")
	output.MakeRoute(r, "/unarchive/{uuid}", h.Unarchive, authCached).Methods("PUT", "OPTIONS")
	output.MakeRoute(r, "/delete/{uuid}", h.Delete, authCached).Methods("DELETE", "OPTIONS")
}
