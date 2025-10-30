package client_repo

import (
	"context"
	"fmt"
	"movepilot/pkg/db"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, name, description string, address Address, geo *Geolocation) (*Model, error)
	GetByUUID(ctx context.Context, uuid string) (*Model, error)
	GetByID(ctx context.Context, id int) (*Model, error)
	FetchAll(ctx context.Context) ([]*Model, error)
	FetchActive(ctx context.Context) ([]*Model, error)
	FetchArchived(ctx context.Context) ([]*Model, error)
	FetchDeleted(ctx context.Context) ([]*Model, error)
	Update(ctx context.Context, uuid string, name, description string, address Address, geo *Geolocation) (*Model, error)
	Archive(ctx context.Context, uuid string) error
	Unarchive(ctx context.Context, uuid string) error
	SoftDelete(ctx context.Context, uuid string) error
	Restore(ctx context.Context, uuid string) error
	HardDelete(ctx context.Context, uuid string) error
	FindByRadius(ctx context.Context, lat, lng float64, radiusMiles float64, includeArchived bool) ([]*ModelWithDistance, error)
	FindNearest(ctx context.Context, lat, lng float64, limit int, includeArchived bool) ([]*ModelWithDistance, error)
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	PostalCode   string
	Country      string
}

type Geolocation struct {
	Latitude  float64
	Longitude float64
}

type ClientRepository struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) Create(ctx context.Context, name, description string, address Address, geo *Geolocation) (*Model, error) {
	now := time.Now()

	query := `
		INSERT INTO clients (
			name, description, deleted, archived,
			address_line1, address_line2, city, postal_code, country,
			latitude, longitude,
			created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING *
	`

	var client Model

	var lat, lng *float64
	if geo != nil {
		lat = &geo.Latitude
		lng = &geo.Longitude
	}

	err := pgxscan.Get(ctx, r.db, &client, query,
		name, description, false, false,
		address.AddressLine1, address.AddressLine2, address.City, address.PostalCode, address.Country,
		lat, lng,
		now, now,
	)

	if err != nil {
		return nil, fmt.Errorf("client.Create query: %w", err)
	}

	return &client, nil
}

func (r *ClientRepository) GetByUUID(ctx context.Context, uuid string) (*Model, error) {
	var client Model
	query := `SELECT * FROM clients WHERE uuid=$1`

	err := pgxscan.Get(ctx, r.db, &client, query, uuid)
	if err != nil {
		if db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("client.GetByUUID query: %w", err)
	}

	return &client, nil
}

func (r *ClientRepository) GetByID(ctx context.Context, id int) (*Model, error) {
	var client Model
	query := `SELECT * FROM clients WHERE id=$1`

	err := pgxscan.Get(ctx, r.db, &client, query, id)
	if err != nil {
		if db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("client.GetByID query: %w", err)
	}

	return &client, nil
}

func (r *ClientRepository) FetchAll(ctx context.Context) ([]*Model, error) {
	var clients []*Model
	query := `SELECT * FROM clients ORDER BY created_at DESC`

	err := pgxscan.Select(ctx, r.db, &clients, query)
	if err != nil {
		return nil, fmt.Errorf("client.FetchAll query: %w", err)
	}

	return clients, nil
}

func (r *ClientRepository) FetchActive(ctx context.Context) ([]*Model, error) {
	var clients []*Model
	query := `SELECT * FROM clients WHERE deleted = false AND archived = false ORDER BY created_at DESC`

	err := pgxscan.Select(ctx, r.db, &clients, query)
	if err != nil {
		return nil, fmt.Errorf("client.FetchActive query: %w", err)
	}

	return clients, nil
}

func (r *ClientRepository) FetchArchived(ctx context.Context) ([]*Model, error) {
	var clients []*Model
	query := `SELECT * FROM clients WHERE deleted = false AND archived = true ORDER BY created_at DESC`

	err := pgxscan.Select(ctx, r.db, &clients, query)
	if err != nil {
		return nil, fmt.Errorf("client.FetchArchived query: %w", err)
	}

	return clients, nil
}

func (r *ClientRepository) FetchDeleted(ctx context.Context) ([]*Model, error) {
	var clients []*Model
	query := `SELECT * FROM clients WHERE deleted = true ORDER BY created_at DESC`

	err := pgxscan.Select(ctx, r.db, &clients, query)
	if err != nil {
		return nil, fmt.Errorf("client.FetchDeleted query: %w", err)
	}

	return clients, nil
}

func (r *ClientRepository) Update(ctx context.Context, uuid string, name, description string, address Address, geo *Geolocation) (*Model, error) {
	now := time.Now()

	query := `
		UPDATE clients
		SET name = $1,
		    description = $2,
		    address_line1 = $3,
		    address_line2 = $4,
		    city = $5,
		    postal_code = $6,
		    country = $7,
		    latitude = $8,
		    longitude = $9,
		    updated_at = $10
		WHERE uuid = $11
		RETURNING *
	`

	var client Model

	var lat, lng *float64
	if geo != nil {
		lat = &geo.Latitude
		lng = &geo.Longitude
	}

	err := pgxscan.Get(ctx, r.db, &client, query,
		name, description,
		address.AddressLine1, address.AddressLine2, address.City, address.PostalCode, address.Country,
		lat, lng,
		now, uuid,
	)

	if err != nil {
		if db.IsNoRowsError(err) {
			return nil, fmt.Errorf("client.Update not found: %s", uuid)
		}
		return nil, fmt.Errorf("client.Update query: %w", err)
	}

	return &client, nil
}

func (r *ClientRepository) Archive(ctx context.Context, uuid string) error {
	query := `UPDATE clients SET archived = true, updated_at = $1 WHERE uuid = $2`

	result, err := r.db.Exec(ctx, query, time.Now(), uuid)
	if err != nil {
		return fmt.Errorf("client.Archive query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("client.Archive not found: %s", uuid)
	}

	return nil
}

func (r *ClientRepository) Unarchive(ctx context.Context, uuid string) error {
	query := `UPDATE clients SET archived = false, updated_at = $1 WHERE uuid = $2`

	result, err := r.db.Exec(ctx, query, time.Now(), uuid)
	if err != nil {
		return fmt.Errorf("client.Unarchive query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("client.Unarchive not found: %s", uuid)
	}

	return nil
}

func (r *ClientRepository) SoftDelete(ctx context.Context, uuid string) error {
	query := `UPDATE clients SET deleted = true, updated_at = $1 WHERE uuid = $2`

	result, err := r.db.Exec(ctx, query, time.Now(), uuid)
	if err != nil {
		return fmt.Errorf("client.SoftDelete query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("client.SoftDelete not found: %s", uuid)
	}

	return nil
}

func (r *ClientRepository) Restore(ctx context.Context, uuid string) error {
	query := `UPDATE clients SET deleted = false, updated_at = $1 WHERE uuid = $2`

	result, err := r.db.Exec(ctx, query, time.Now(), uuid)
	if err != nil {
		return fmt.Errorf("client.Restore query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("client.Restore not found: %s", uuid)
	}

	return nil
}

func (r *ClientRepository) HardDelete(ctx context.Context, uuid string) error {
	query := `DELETE FROM clients WHERE uuid=$1`

	result, err := r.db.Exec(ctx, query, uuid)
	if err != nil {
		return fmt.Errorf("client.HardDelete query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("client.HardDelete not found: %s", uuid)
	}

	return nil
}

func (r *ClientRepository) FindByRadius(ctx context.Context, lat, lng float64, radiusMiles float64, includeArchived bool) ([]*ModelWithDistance, error) {
	whereClause := "WHERE latitude IS NOT NULL AND longitude IS NOT NULL AND deleted = false"
	if !includeArchived {
		whereClause += " AND archived = false"
	}

	query := fmt.Sprintf(`
		SELECT id, uuid, name, description, deleted, archived,
		       address_line1, address_line2, city, postal_code, country,
		       latitude, longitude, created_at, updated_at,
		       (3959 * acos(
		           cos(radians($1)) * cos(radians(latitude)) * 
		           cos(radians(longitude) - radians($2)) + 
		           sin(radians($1)) * sin(radians(latitude))
		       )) AS distance
		FROM clients
		%s
		  AND (3959 * acos(
		          cos(radians($1)) * cos(radians(latitude)) * 
		          cos(radians(longitude) - radians($2)) + 
		          sin(radians($1)) * sin(radians(latitude))
		      )) <= $3
		ORDER BY distance ASC
	`, whereClause)

	var clients []*ModelWithDistance
	err := pgxscan.Select(ctx, r.db, &clients, query, lat, lng, radiusMiles)
	if err != nil {
		return nil, fmt.Errorf("client.FindByRadius query: %w", err)
	}

	return clients, nil
}

func (r *ClientRepository) FindNearest(ctx context.Context, lat, lng float64, limit int, includeArchived bool) ([]*ModelWithDistance, error) {
	whereClause := "WHERE latitude IS NOT NULL AND longitude IS NOT NULL AND deleted = false"
	if !includeArchived {
		whereClause += " AND archived = false"
	}

	query := fmt.Sprintf(`
		SELECT id, uuid, name, description, deleted, archived,
		       address_line1, address_line2, city, postal_code, country,
		       latitude, longitude, created_at, updated_at,
		       (3959 * acos(
		           cos(radians($1)) * cos(radians(latitude)) * 
		           cos(radians(longitude) - radians($2)) + 
		           sin(radians($1)) * sin(radians(latitude))
		       )) AS distance
		FROM clients
		%s
		ORDER BY distance ASC
		LIMIT $3
	`, whereClause)

	var clients []*ModelWithDistance
	err := pgxscan.Select(ctx, r.db, &clients, query, lat, lng, limit)
	if err != nil {
		return nil, fmt.Errorf("client.FindNearest query: %w", err)
	}

	return clients, nil
}
