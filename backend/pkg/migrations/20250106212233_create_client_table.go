package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateClientTable, downCreateClientTable)
}

func upCreateClientTable(ctx context.Context, tx *sql.Tx) error {
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
	create_index := `CREATE INDEX IF NOT EXISTS idx_users_uuid ON users(uuid)`
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

func downCreateClientTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	query := `DROP TABLE users`
	_, err := tx.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
