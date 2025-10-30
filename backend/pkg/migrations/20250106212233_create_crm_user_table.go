package migrations

import (
	"context"
	"database/sql"
	"movepilot/pkg/bcrypt"
	crm_user_repo "movepilot/pkg/repositories/crm_user"
	"time"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateCRMUserTable, downCreateCRMUserTable)
}

func upCreateCRMUserTable(ctx context.Context, tx *sql.Tx) error {
	create_table := `CREATE TABLE crm_users (
		id SERIAL PRIMARY KEY,
		uuid UUID DEFAULT uuid_generate_v7() NOT NULL UNIQUE,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		email VARCHAR(120),
		password VARCHAR(120),
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now()
	)`
	_, err := tx.Exec(create_table)
	if err != nil {
		return err
	}

	// Optional: UNIQUE constraint already creates an index
	create_index := `CREATE INDEX IF NOT EXISTS idx_crm_users_uuid ON crm_users(uuid)`
	_, err = tx.Exec(create_index)
	if err != nil {
		return err
	}

	err = insertBaseCRMUsers(ctx, tx)
	if err != nil {
		return err
	}

	return nil
}

func insertBaseCRMUsers(ctx context.Context, tx *sql.Tx) error {

	hashed_password, err := bcrypt.HashPassword("hashed_password")

	if err != nil {
		return err
	}

	kez := crm_user_repo.Model{
		FirstName: "Kez",
		LastName:  "Anwar",
		Email:     "kezanwar@gmail.com",
		Password:  hashed_password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	insert_user := `
		INSERT INTO crm_users 
		(first_name, last_name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = tx.ExecContext(ctx, insert_user,
		kez.FirstName, kez.LastName, kez.Email, kez.Password, kez.CreatedAt, kez.UpdatedAt,
	)

	return err
}
func downCreateCRMUserTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	query := `DROP TABLE crm_users`
	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
