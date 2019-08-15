package repsitory

import (
	"context"

	"tes-project/models"
)

// UserRepo explain...
type UserRepo interface {
	InsertUser(ctx context.Context, p *models.User) (int64, error)
	InsertLoan(ctx context.Context, p *models.User) (int64, error)
}
