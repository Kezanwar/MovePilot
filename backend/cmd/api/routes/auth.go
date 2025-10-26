package routes

import (
	"movepilot/cmd/api/handlers"
	"movepilot/pkg/middleware"
	"movepilot/pkg/output"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router, h *handlers.AuthHandler, authCached middleware.Middleware) {
	output.MakeRoute(r, "/register", h.Register).Methods("POST", "OPTIONS")
	output.MakeRoute(r, "/sign-in", h.SignIn).Methods("POST", "OPTIONS")
	output.MakeRoute(r, "/initialize", h.Initialize, authCached).Methods("GET", "OPTIONS")
	output.MakeRoute(r, "/confirm-otp/{otp}", h.ConfirmOTP, authCached).Methods("POST", "OPTIONS")
	output.MakeRoute(r, "/resend-otp", h.ResendOTP, authCached).Methods("POST", "OPTIONS")
}
