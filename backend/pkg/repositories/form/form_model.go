package form_repo

import (
	"encoding/json"
	"time"
)

type FormModel struct {
	ID              int             `json:"-" db:"id"`
	UUID            string          `json:"uuid" db:"uuid"`
	UserID          int             `json:"-" db:"user_id"`
	Name            string          `json:"name" db:"name"`
	Description     *string         `json:"description" db:"description"`
	FormData        json.RawMessage `json:"form_data,omitempty" db:"form_data"` // Use json.RawMessage for JSONB
	Status          string          `json:"status" db:"status"`
	Views           int             `json:"views" db:"views"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
	Affiliates      json.RawMessage `json:"affiliates,omitempty" db:"affiliates"`
	SubmissionCount int             `json:"submission_count" db:"submission_count"`
}

const (
	StatusInactive = "inactive"
	StatusDraft    = "draft"
	StatusActive   = "active"
)

var ValidStatuses = []string{StatusInactive, StatusDraft, StatusActive}

// Helper method to unmarshal FormData into a specific struct
func (m *FormModel) UnmarshalFormData(v interface{}) error {
	return json.Unmarshal(m.FormData, v)
}

// Helper method to unmarshal Affiliates into AffiliateInfo slice
func (m *FormModel) GetAffiliates() ([]AffiliateInfo, error) {
	var affiliates []AffiliateInfo
	err := json.Unmarshal(m.Affiliates, &affiliates)
	return affiliates, err
}

type AffiliateInfo struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
