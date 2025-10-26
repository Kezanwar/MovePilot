package routes

import (
	"movepilot/cmd/api/handlers"
	"movepilotot/pkg/middleware"
	"movepilotot/pkg/output"

	"github.com/gorilla/mux"
)

func FormRoutes(r *mux.Router, h *handlers.FormHandler, authCached middleware.Middleware) {
	output.MakeRoute(r, "/list", h.GetDetailedListing, authCached).Methods("GET", "OPTIONS")
	output.MakeRoute(r, "/new", h.NewForm, authCached).Methods("POST", "OPTIONS")
	output.MakeRoute(r, "/view/{uuid}", h.GetForm, authCached).Methods("GET", "OPTIONS")
	output.MakeRoute(r, "/update/{uuid}/data", h.UpdateFormData, authCached).Methods("PUT", "OPTIONS")
	output.MakeRoute(r, "/update/{uuid}/meta", h.UpdateFormMeta, authCached).Methods("PUT", "OPTIONS")
	output.MakeRoute(r, "/update/{uuid}/affiliates", h.UpdateFormAffiliates, authCached).Methods("PUT", "OPTIONS")
	output.MakeRoute(r, "/delete/{uuid}", h.DeleteForm, authCached).Methods("DELETE", "OPTIONS")
}
