package main

import (
	"context"
	"log"
	"movepilot/cmd/api/handlers"
	"movepilot/cmd/api/routes"
	user_memory_cache "movepilot/pkg/cache/user_memory"
	"movepilot/pkg/email"
	"movepilot/pkg/middleware"
	crm_user_repo "movepilot/pkg/repositories/crm_user"

	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewAPI(ctx context.Context, pool *pgxpool.Pool, client *http.Client) (*http.Server, error) {

	TWO_HOURS := 2 * time.Hour

	emailClient, err := email.NewClient()

	if err != nil {
		log.Fatalf("Email client failed to init: %v", err)
	}

	//memory cache
	userCache := user_memory_cache.New(TWO_HOURS)

	//repositories
	crmUserRepo := crm_user_repo.NewUserRepo(pool)

	//handlers
	authHandlers := handlers.NewAuthHandler(crmUserRepo, userCache, emailClient)

	crmAuthFresh := middleware.CRMAuthAlwaysFreshMiddleware(crmUserRepo, userCache)
	crmAuthCached := middleware.CRMAuthCachedMiddleware(crmUserRepo, userCache)

	//router
	r := mux.NewRouter()
	r.Use(middleware.Cors)
	api := r.PathPrefix("/api").Subrouter()

	//apply routes
	routes.Register(
		//main router
		api,
		//handlers
		authHandlers,

		//middleware
		crmAuthFresh,
		crmAuthCached,
	)

	return &http.Server{
		Addr:    PORT,
		Handler: r,
	}, nil
}
