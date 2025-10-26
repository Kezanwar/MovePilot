package user_repo

import (
	"movepilot/pkg/bcrypt"
	"movepilotot/pkg/otp"
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
	ID                 int       `json:"-" db:"id"`
	UUID               string    `json:"-" db:"uuid"`
	FirstName          string    `json:"first_name" db:"first_name"`
	LastName           string    `json:"last_name" db:"last_name"`
	Email              string    `json:"email" db:"email"`
	Password           string    `json:"-" db:"password"`
	TermsAndConditions bool      `json:"terms_and_conditions" db:"terms_and_conditions"`
	EmailConfirmed     bool      `json:"email_confirmed" db:"email_confirmed"`
	OTP                string    `json:"-" db:"otp"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

func (m *Model) IsPassword(to_check string) bool {
	return bcrypt.ValidatePassword(m.Password, to_check)
}

// ValidateOTP checks if the provided OTP matches the user's stored OTP
func (m *Model) ValidateOTP(providedOTP string) bool {
	// First check if the provided OTP has valid format
	if !otp.IsValidFormat(providedOTP) {
		return false
	}
	return otp.Validate(providedOTP, m.OTP)
}
