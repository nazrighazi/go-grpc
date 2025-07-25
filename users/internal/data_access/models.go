// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package dataaccess

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID                   pgtype.UUID        `json:"id"`
	Email                string             `json:"email"`
	PasswordHash         string             `json:"password_hash"`
	IsActive             pgtype.Bool        `json:"is_active"`
	IsVerified           pgtype.Bool        `json:"is_verified"`
	Role                 pgtype.Text        `json:"role"`
	CreatedAt            pgtype.Timestamptz `json:"created_at"`
	UpdatedAt            pgtype.Timestamptz `json:"updated_at"`
	DeletedAt            pgtype.Timestamptz `json:"deleted_at"`
	LastLogin            pgtype.Timestamptz `json:"last_login"`
	PasswordResetToken   pgtype.Text        `json:"password_reset_token"`
	PasswordResetExpires pgtype.Timestamptz `json:"password_reset_expires"`
}
