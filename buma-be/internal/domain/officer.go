package domain

import "time"

type Officer struct {
	OfficerID   string    `json:"officer_id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PairingTeam string    `json:"pairing_team"`
	Phone       string    `json:"phone,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AccountAssignment struct {
	ID         string    `json:"id"`
	AccountNo  string    `json:"account_no"`
	OfficerID  string    `json:"officer_id"`
	AssignedAt time.Time `json:"assigned_at"`
}

type CreateOfficerRequest struct {
	OfficerID   string `json:"officer_id,omitempty"`
	FullName    string `json:"full_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PairingTeam string `json:"pairing_team" validate:"required"`
	Phone       string `json:"phone,omitempty"`
}

