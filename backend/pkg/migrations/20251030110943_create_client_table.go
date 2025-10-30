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
	create_table := `CREATE TABLE clients (
		id SERIAL PRIMARY KEY,
		uuid UUID DEFAULT uuid_generate_v7() NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		deleted BOOLEAN NOT NULL DEFAULT false,
		archived BOOLEAN NOT NULL DEFAULT false,
		address_line1 VARCHAR(255),
		address_line2 VARCHAR(255),
		city VARCHAR(100),
		postal_code VARCHAR(20),
		country VARCHAR(100),
		latitude DOUBLE PRECISION,
		longitude DOUBLE PRECISION,
		created_at TIMESTAMP DEFAULT now() NOT NULL,
		updated_at TIMESTAMP DEFAULT now() NOT NULL
	)`
	_, err := tx.Exec(create_table)
	if err != nil {
		return err
	}

	// Create index on UUID
	create_uuid_index := `CREATE INDEX IF NOT EXISTS idx_clients_uuid ON clients(uuid)`
	_, err = tx.Exec(create_uuid_index)
	if err != nil {
		return err
	}

	// Create index for filtering active clients (both false)
	create_active_index := `CREATE INDEX IF NOT EXISTS idx_clients_active ON clients(deleted, archived) WHERE deleted = false AND archived = false`
	_, err = tx.Exec(create_active_index)
	if err != nil {
		return err
	}

	// Create index for deleted clients
	create_deleted_index := `CREATE INDEX IF NOT EXISTS idx_clients_deleted ON clients(deleted) WHERE deleted = true`
	_, err = tx.Exec(create_deleted_index)
	if err != nil {
		return err
	}

	// Create index for archived clients
	create_archived_index := `CREATE INDEX IF NOT EXISTS idx_clients_archived ON clients(archived) WHERE archived = true`
	_, err = tx.Exec(create_archived_index)
	if err != nil {
		return err
	}

	return nil
}
func downCreateClientTable(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS clients`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
