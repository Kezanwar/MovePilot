package main

import (
	"context"
	"log"
	"move-pilot/pkg/db"
	_ "move-pilotot/pkg/migrations"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var PORT = os.Getenv("PORT")

func main() {

	ctx := context.Background()

	pool, err := db.Connect(ctx)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	httpclient := &http.Client{}

	db.MigrateUp()

	api, err := NewAPI(ctx, pool, httpclient)

	if err != nil {
		log.Fatalf("Server init failed: %v", err)
	}

	go func() {
		log.Printf("🚀 Server running on %s", PORT)
		if err := api.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("🧼 Shutting down...")

	ctxShutdown, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	api.Shutdown(ctxShutdown)

	db.Close(pool)
}
