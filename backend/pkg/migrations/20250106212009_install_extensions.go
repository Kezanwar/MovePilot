package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upInstallExtensions, downInstallExtensions)
}

func upInstallExtensions(ctx context.Context, tx *sql.Tx) error {
	// On Postgres 13+, gen_random_uuid() is built-in
	// On Postgres 12 and below, we need pgcrypto
	// This is safe to run on all versions
	query := `CREATE EXTENSION IF NOT EXISTS "pgcrypto";`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}

	// Create UUID v7 generation function
	create_uuid_v7_function := `
	CREATE OR REPLACE FUNCTION uuid_generate_v7() RETURNS uuid AS $$
	DECLARE
		unix_ts_ms bytea;
		uuid_bytes bytea;
	BEGIN
		unix_ts_ms = substring(int8send(floor(extract(epoch from clock_timestamp()) * 1000)::bigint) from 3);
		
		uuid_bytes = uuid_send(gen_random_uuid());
		
		uuid_bytes = overlay(uuid_bytes placing unix_ts_ms from 1 for 6);
		
		uuid_bytes = set_byte(uuid_bytes, 6, (b'0111' || get_byte(uuid_bytes, 6)::bit(4))::bit(8)::int);
		uuid_bytes = set_byte(uuid_bytes, 8, (b'10' || substring(get_byte(uuid_bytes, 8)::bit(8) from 3))::bit(8)::int);
		
		return encode(uuid_bytes, 'hex')::uuid;
	END
	$$ LANGUAGE plpgsql VOLATILE;
	`
	_, err = tx.Exec(create_uuid_v7_function)
	if err != nil {
		return err
	}

	return nil
}

func downInstallExtensions(ctx context.Context, tx *sql.Tx) error {
	// Drop UUID v7 function first
	drop_function := `DROP FUNCTION IF EXISTS uuid_generate_v7();`
	_, err := tx.Exec(drop_function)
	if err != nil {
		return err
	}

	// Drop extension
	query := `DROP EXTENSION IF EXISTS "pgcrypto";`
	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
