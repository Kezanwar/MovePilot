package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateFormsTables, downCreateFormsTables)
}

func upCreateFormsTables(ctx context.Context, tx *sql.Tx) error {
	//---- create forms table
	create_forms_table := `CREATE TABLE forms (
		id SERIAL PRIMARY KEY,
		uuid UUID DEFAULT uuid_generate_v7() NOT NULL UNIQUE,
		user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		form_data JSONB NOT NULL,
		status VARCHAR(20) DEFAULT 'draft',
		views INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now()
	)`
	_, err := tx.ExecContext(ctx, create_forms_table)
	if err != nil {
		return err
	}

	//create forms table indexes
	create_forms_uuid_index := `CREATE INDEX IF NOT EXISTS idx_forms_uuid ON forms(uuid)`
	_, err = tx.ExecContext(ctx, create_forms_uuid_index)
	if err != nil {
		return err
	}

	//index on user_id for fast lookups by user
	create_forms_user_index := `CREATE INDEX IF NOT EXISTS idx_forms_user_id ON forms(user_id)`
	_, err = tx.ExecContext(ctx, create_forms_user_index)
	if err != nil {
		return err
	}

	//GIN index on form_data JSONB for better query performance
	create_form_data_index := `CREATE INDEX IF NOT EXISTS idx_forms_data ON forms USING GIN (form_data)`
	_, err = tx.ExecContext(ctx, create_form_data_index)
	if err != nil {
		return err
	}
	//---- end

	//---- create affiliates table
	create_affiliates_table := `CREATE TABLE affiliates (
		id SERIAL PRIMARY KEY,
		uuid UUID DEFAULT uuid_generate_v7() NOT NULL UNIQUE,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		email VARCHAR(255),
		company VARCHAR(255),
		created_at TIMESTAMP DEFAULT now(),
		updated_at TIMESTAMP DEFAULT now()
	)`
	_, err = tx.ExecContext(ctx, create_affiliates_table)
	if err != nil {
		return err
	}

	//create affiliates table indexes
	create_affiliates_uuid_index := `CREATE INDEX IF NOT EXISTS idx_affiliates_uuid ON affiliates(uuid)`
	_, err = tx.ExecContext(ctx, create_affiliates_uuid_index)
	if err != nil {
		return err
	}

	//---- end

	//---- create junction table for many-to-many relationship
	create_form_affiliates_table := `CREATE TABLE form_affiliates (
		form_id INTEGER NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
		affiliate_id INTEGER NOT NULL REFERENCES affiliates(id) ON DELETE CASCADE,
		added_at TIMESTAMP DEFAULT now(),
		PRIMARY KEY (form_id, affiliate_id)
	)`
	_, err = tx.ExecContext(ctx, create_form_affiliates_table)
	if err != nil {
		return err
	}

	//create junction table indexes
	create_form_affiliates_form_index := `CREATE INDEX IF NOT EXISTS idx_form_affiliates_form_id ON form_affiliates(form_id)`
	_, err = tx.ExecContext(ctx, create_form_affiliates_form_index)
	if err != nil {
		return err
	}

	create_form_affiliates_affiliate_index := `CREATE INDEX IF NOT EXISTS idx_form_affiliates_affiliate_id ON form_affiliates(affiliate_id)`
	_, err = tx.ExecContext(ctx, create_form_affiliates_affiliate_index)
	if err != nil {
		return err
	}

	//---- end

	// ---- create form_submissions table
	create_form_submissions_table := `CREATE TABLE form_submissions (
	id SERIAL PRIMARY KEY,
	uuid UUID DEFAULT uuid_generate_v7() NOT NULL UNIQUE,
	form_id INTEGER NOT NULL REFERENCES forms(id) ON DELETE CASCADE,
	full_name VARCHAR(255),
	email VARCHAR(255),
	submission_data JSONB NOT NULL,
	submitted_at TIMESTAMP DEFAULT now()
)`
	_, err = tx.ExecContext(ctx, create_form_submissions_table)
	if err != nil {
		return err
	}

	// create form_submissions table indexes
	create_form_submissions_uuid_index := `CREATE INDEX IF NOT EXISTS idx_form_submissions_uuid ON form_submissions(uuid)`
	_, err = tx.ExecContext(ctx, create_form_submissions_uuid_index)
	if err != nil {
		return err
	}

	create_form_submissions_form_index := `CREATE INDEX IF NOT EXISTS idx_form_submissions_form_id ON form_submissions(form_id)`
	_, err = tx.ExecContext(ctx, create_form_submissions_form_index)
	if err != nil {
		return err
	}

	// Index on email for fast lookups
	create_form_submissions_email_index := `CREATE INDEX IF NOT EXISTS idx_form_submissions_email ON form_submissions(email)`
	_, err = tx.ExecContext(ctx, create_form_submissions_email_index)
	if err != nil {
		return err
	}

	// Optional: GIN index on submission_data if you want to query submission contents
	create_form_submissions_data_index := `CREATE INDEX IF NOT EXISTS idx_form_submissions_data ON form_submissions USING GIN (submission_data)`
	_, err = tx.ExecContext(ctx, create_form_submissions_data_index)
	if err != nil {
		return err
	}
	//---- end

	return nil
}

func downCreateFormsTables(ctx context.Context, tx *sql.Tx) error {

	//drop tables in reverse order (junction table first due to foreign keys)``
	drop_form_submissions := `DROP TABLE IF EXISTS form_submissions`
	_, err := tx.ExecContext(ctx, drop_form_submissions)
	if err != nil {
		return err
	}

	drop_form_affiliates := `DROP TABLE IF EXISTS form_affiliates`
	_, err = tx.ExecContext(ctx, drop_form_affiliates)
	if err != nil {
		return err
	}

	drop_affiliates := `DROP TABLE IF EXISTS affiliates`
	_, err = tx.ExecContext(ctx, drop_affiliates)
	if err != nil {
		return err
	}

	drop_forms := `DROP TABLE IF EXISTS forms`
	_, err = tx.ExecContext(ctx, drop_forms)
	if err != nil {
		return err
	}

	return nil
}
