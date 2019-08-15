package loan

import (
	"context"
	"database/sql"

	models "tes-project/models"
	pRepo "tes-project/repository"
)

// NewSQLLoanRepo retunrs implement of loan repository interface
func NewSQLLoanRepo(Conn *sql.DB) pRepo.LoanRepo {
	return &mysqlLoanRepo{
		Conn: Conn,
	}
}

type mysqlLoanRepo struct {
	Conn *sql.DB
}

func (m *mysqlLoanRepo) InsertInstallment(ctx context.Context, loanCode string, capital float64, interest float64, total float64, plan int, dueDate string) (int64, error) {
	query := "Insert tbl_installment SET loan_code=?, capital=?, interest=?, total=?, plan=?, due_date=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(
		ctx,
		loanCode,
		capital,
		interest,
		total,
		plan,
		dueDate,
	)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlLoanRepo) GetInstallmentByLoanCode(ctx context.Context, loanCode string) ([]*models.Installment, error) {
	query := "SELECT capital, interest, total, plan, due_date FROM tbl_installment WHERE loan_code=?"

	return m.fetchInstallment(ctx, query, loanCode)
}

func (m *mysqlLoanRepo) fetchInstallment(ctx context.Context, query string, args ...interface{}) ([]*models.Installment, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Installment, 0)
	for rows.Next() {
		data := new(models.Installment)

		err := rows.Scan(
			&data.Capital,
			&data.Interest,
			&data.Total,
			&data.Plan,
			&data.DueDate,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlLoanRepo) GetInterest(ctx context.Context, tenor int) ([]*models.Interest, error) {
	query := "SELECT tenor, interest FROM tbl_loan_interest WHERE tenor=?"

	return m.fetchInterest(ctx, query, tenor)
}

func (m *mysqlLoanRepo) fetchInterest(ctx context.Context, query string, args ...interface{}) ([]*models.Interest, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Interest, 0)
	for rows.Next() {
		data := new(models.Interest)

		err := rows.Scan(
			&data.Tenor,
			&data.Interest,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlLoanRepo) GetLoanTrack(ctx context.Context, from string, to string) ([]*models.LoanTrack, error) {
	query := "SELECT jumlah_pinjaman FROM tbl_request_loans WHERE created_at>=? AND created_at<=?"

	return m.fetchLoanTrack(ctx, query, from, to)
}

func (m *mysqlLoanRepo) fetchLoanTrack(ctx context.Context, query string, args ...interface{}) ([]*models.LoanTrack, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.LoanTrack, 0)
	for rows.Next() {
		data := new(models.LoanTrack)

		err := rows.Scan(
			&data.JumlahPinjaman,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}
