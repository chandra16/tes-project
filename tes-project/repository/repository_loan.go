package repsitory

import (
	"context"
	"tes-project/models"
)

// LoanRepo explain...
type LoanRepo interface {
	InsertInstallment(ctx context.Context, loanCode string, capital float64, interest float64, total float64, plan int, dueDate string) (int64, error)
	GetInstallmentByLoanCode(ctx context.Context, loanCode string) ([]*models.Installment, error)
	GetInterest(ctx context.Context, tenor int) ([]*models.Interest, error)
	GetLoanTrack(ctx context.Context, from string, to string) ([]*models.LoanTrack, error)
}
