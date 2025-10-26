package form_repo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, userId int, name string, description *string, formData FormData) (*FormModel, error)
	GetByUUID(ctx context.Context, uuid string) (*FormModel, error)
	GetByID(ctx context.Context, id int) (*FormModel, error)
	GetBasicListingByUserID(ctx context.Context, id int) ([]*FormModel, error)
	GetDetailedListingByUserID(ctx context.Context, id int) ([]*FormModel, error)
	UpdateFormMeta(ctx context.Context, id int, name, description string, status string) (*FormModel, error)
	IncrementViews(ctx context.Context, uuid string) error
	Delete(ctx context.Context, uuid string) error
}

type FormRepository struct {
	db *pgxpool.Pool
}

func NewFormRepo(db *pgxpool.Pool) *FormRepository {
	return &FormRepository{db: db}
}

func (r *FormRepository) Create(ctx context.Context, user_id int, name string, description *string, formData FormData) (*FormModel, error) {
	now := time.Now()

	// Marshal formData to JSON
	formDataJSON, err := json.Marshal(formData)
	if err != nil {
		return nil, fmt.Errorf("form.Create marshal: %w", err)
	}

	query := `
		INSERT INTO forms (user_id, name, description, form_data, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *
	`

	var form FormModel

	err = pgxscan.Get(ctx, r.db, &form, query, user_id, name, description, formDataJSON, now, now)

	if err != nil {
		return nil, fmt.Errorf("form.Create query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) GetByUUID(ctx context.Context, uuid string) (*FormModel, error) {
	var form FormModel

	query := `
	SELECT
		f.*,
		COALESCE(
			jsonb_agg(
				jsonb_build_object(
					'uuid', a.uuid,
					'first_name', a.first_name,
					'last_name', a.last_name
				)
			) FILTER (WHERE a.uuid IS NOT NULL),
			'[]'::jsonb
		) as affiliates,
		COUNT(DISTINCT fs.id) as submission_count
	FROM forms f
	LEFT JOIN form_affiliates fa ON f.id = fa.form_id
	LEFT JOIN affiliates a ON fa.affiliate_id = a.id
	LEFT JOIN form_submissions fs ON f.id = fs.form_id
	WHERE f.uuid=$1
	GROUP BY f.id`

	err := pgxscan.Get(ctx, r.db, &form, query, uuid)
	if err != nil {
		return nil, fmt.Errorf("form.GetByUUID query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) GetByID(ctx context.Context, id int) (*FormModel, error) {
	var form FormModel

	query := `
	SELECT
		f.*,
		COALESCE(
			jsonb_agg(
				jsonb_build_object(
					'uuid', a.uuid,
					'first_name', a.first_name,
					'last_name', a.last_name
				)
			) FILTER (WHERE a.uuid IS NOT NULL),
			'[]'::jsonb
		) as affiliates,
		COUNT(DISTINCT fs.id) as submission_count
	FROM forms f
	LEFT JOIN form_affiliates fa ON f.id = fa.form_id
	LEFT JOIN affiliates a ON fa.affiliate_id = a.id
	LEFT JOIN form_submissions fs ON f.id = fs.form_id
	WHERE f.id=$1
	GROUP BY f.id`

	err := pgxscan.Get(ctx, r.db, &form, query, id)
	if err != nil {
		return nil, fmt.Errorf("form.GetByID query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) GetBasicListingByUserID(ctx context.Context, id int) ([]*FormModel, error) {
	forms := []*FormModel{}

	query := `
	SELECT uuid, name, description, status, views, created_at, updated_at
	FROM forms
	WHERE user_id = $1
	ORDER BY created_at DESC`

	err := pgxscan.Select(ctx, r.db, &forms, query, id)
	if err != nil {
		return nil, fmt.Errorf("form.FetchAll query: %w", err)
	}

	return forms, nil
}

func (r *FormRepository) GetDetailedListingByUserID(ctx context.Context, id int) ([]*FormModel, error) {
	forms := []*FormModel{}

	query := `
	SELECT
	f.uuid,
	f.name,
	f.description,
	f.status,
	f.views,
	f.created_at,
	f.updated_at,
	COALESCE(
		jsonb_agg(
			jsonb_build_object(
				'uuid', a.uuid,
				'first_name', a.first_name,
				'last_name', a.last_name
				)
			) FILTER (WHERE a.uuid IS NOT NULL),
			'[]'::jsonb
		) as affiliates,
		COUNT(DISTINCT fs.id) as submission_count
	FROM forms f
	LEFT JOIN form_affiliates fa ON f.id = fa.form_id
	LEFT JOIN affiliates a ON fa.affiliate_id = a.id
	LEFT JOIN form_submissions fs ON f.id = fs.form_id
	WHERE f.user_id = $1
	GROUP BY f.id, f.uuid, f.name, f.description, f.status, f.views, f.created_at, f.updated_at
	ORDER BY f.created_at DESC`

	err := pgxscan.Select(ctx, r.db, &forms, query, id)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("form.GetDetailedListingByUserID query: %w", err)
	}

	return forms, nil
}
func (r *FormRepository) UpdateFormMeta(ctx context.Context, id int, name, description string, status string) (*FormModel, error) {
	now := time.Now()
	fmt.Println(description)

	var query = `
			UPDATE forms
			SET name=$1, description=$2, status=$3, updated_at=$4
			WHERE id=$5
			RETURNING *
		`

	var form FormModel

	err := pgxscan.Get(ctx, r.db, &form, query, name, description, status, now, id)

	if err != nil {
		return nil, fmt.Errorf("form.Update query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) UpdateFormData(ctx context.Context, uuid, name, description string, formData FormData) (*FormModel, error) {
	now := time.Now()

	// Marshal formData to JSON
	formDataJSON, err := json.Marshal(formData)
	if err != nil {
		return nil, fmt.Errorf("form.Update marshal: %w", err)
	}

	query := `
		UPDATE forms 
		SET name=$1, description=$2, form_data=$3, updated_at=$4 
		WHERE uuid=$5
		RETURNING *
	`

	var form FormModel

	err = pgxscan.Get(ctx, r.db, &form, query, name, description, formDataJSON, now, uuid)
	if err != nil {
		return nil, fmt.Errorf("form.Update query: %w", err)
	}

	return &form, nil
}

func (r *FormRepository) IncrementViews(ctx context.Context, uuid string) error {
	query := `
	    UPDATE forms 
	    SET views = views + 1 
	    WHERE uuid=$1
	`

	_, err := r.db.Exec(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("form.IncrementViews: %w", err)
	}

	return nil
}

func (r *FormRepository) Delete(ctx context.Context, uuid string) error {
	query := `DELETE FROM forms WHERE uuid=$1`

	_, err := r.db.Exec(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("form.Delete: %w", err)
	}

	return nil
}
