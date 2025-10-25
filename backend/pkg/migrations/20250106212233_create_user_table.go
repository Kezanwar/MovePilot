package migrations

import (
	"context"
	"database/sql"
	"move-pilot/pkg/bcrypt"
	user_repo "move-pilotot/pkg/repositories/user"
	"time"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUserTable, downCreateUserTable)
}

func upCreateUserTable(ctx context.Context, tx *sql.Tx) error {
	create_table := `CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		uuid UUID DEFAULT uuid_generate_v7() NOT NULL UNIQUE,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		email VARCHAR(120),
		password VARCHAR(120),
		terms_and_conditions BOOLEAN DEFAULT false,
		email_confirmed BOOLEAN DEFAULT false,
		otp VARCHAR(255),
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now()
	)`
	_, err := tx.Exec(create_table)
	if err != nil {
		return err
	}

	// Optional: UNIQUE constraint already creates an index
	create_index := `CREATE INDEX IF NOT EXISTS idx_users_uuid ON users(uuid)`
	_, err = tx.Exec(create_index)
	if err != nil {
		return err
	}

	err = insertDummyUser(ctx, tx)
	if err != nil {
		return err
	}

	return nil
}

func insertDummyUser(ctx context.Context, tx *sql.Tx) error {

	hashed_password, err := bcrypt.HashPassword("hashed_password")

	if err != nil {
		return err
	}

	kez := user_repo.Model{
		FirstName: "Kez",
		LastName:  "Anwar",
		Email:     "kezanwar@gmail.com",
		Password:  hashed_password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	insert_user := `
		INSERT INTO users 
		(first_name, last_name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = tx.ExecContext(ctx, insert_user,
		kez.FirstName, kez.LastName, kez.Email, kez.Password, kez.CreatedAt, kez.UpdatedAt,
	)

	return err
}
func downCreateUserTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	query := `DROP TABLE users`
	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
