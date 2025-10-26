package crm_user_repo

import (
	"movepilot/pkg/bcrypt"
	"time"
)

/*
	`CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	uuid UUID DEFAULT uuid_generate_v4() UNIQUE,
	first_name VARCHAR(50),
	last_name VARCHAR(50),
	email VARCHAR(120),
	password VARCHAR(120),
	created_at TIMESTAMP DEFAULT now(),
	updated_at TIMESTAMP DEFAULT now()
    )`
*/

type Model struct {
	ID        int       `json:"-" db:"id"`
	UUID      string    `json:"-" db:"uuid"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (m *Model) IsPassword(to_check string) bool {
	return bcrypt.ValidatePassword(m.Password, to_check)
}
