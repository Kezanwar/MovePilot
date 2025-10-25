package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"

	"github.com/pressly/goose/v3"
)

var Conn *pgx.Conn

var DATABASE_URL = os.Getenv("DATABASE_URL")

var db_user = os.Getenv("DB_USER")
var db_name = os.Getenv("DB_NAME")
var db_password = os.Getenv("DB_PASSWORD")

var migrationsDir = "./pkg/migrations"

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(ctx, DATABASE_URL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	fmt.Println("Postgres pool connected 🚀")
	return pool, nil
}

func Close(pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
		fmt.Println("Postgres pool closed 😡")
	}
}

func MigrateUp() {

	connConfig, err := pgx.ParseConfig(DATABASE_URL)
	// Use stdlib.OpenDB to create a database/sql compatible connection
	sqlDB := stdlib.OpenDB(*connConfig)

	defer sqlDB.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := goose.Up(sqlDB, migrationsDir); err != nil {
		log.Fatalf("Failed to run migration up: %v", err)
	}

}

func MigrateDown() {
	connConfig, err := pgx.ParseConfig(DATABASE_URL)
	// Use stdlib.OpenDB to create a database/sql compatible connection
	sqlDB := stdlib.OpenDB(*connConfig)

	defer sqlDB.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := goose.Down(sqlDB, migrationsDir); err != nil {
		log.Fatalf("Failed to run migration down: %v", err)
	}
}

func MigrateDownTo(to int64) {
	connConfig, err := pgx.ParseConfig(DATABASE_URL)
	// Use stdlib.OpenDB to create a database/sql compatible connection
	sqlDB := stdlib.OpenDB(*connConfig)

	defer sqlDB.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := goose.DownTo(sqlDB, migrationsDir, to); err != nil {
		log.Fatalf("Failed to run migration down: %v", err)
	}
}
