package models

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
	ExpiresAt time.Time
}

type PasswordResetService struct {
	DB            *sql.DB
	BytesPerToken int
	Duration      time.Duration
}

func (p *PasswordResetService) Create(email string) (*PasswordReset, error) {
	return &PasswordReset{}, fmt.Errorf("")
}

func (p *PasswordResetService) Consume(token string) (*User, error) {
	return nil, fmt.Errorf("")
}
