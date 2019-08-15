package user

import (
	"context"
	"database/sql"
	"strings"
	"time"

	models "tes-project/models"
	pRepo "tes-project/repository"
)

// NewSQLUserRepo retunrs implement of user repository interface
func NewSQLUserRepo(Conn *sql.DB) pRepo.UserRepo {
	return &mysqlUserRepo{
		Conn: Conn,
	}
}

type mysqlUserRepo struct {
	Conn *sql.DB
}

func (m *mysqlUserRepo) InsertUser(ctx context.Context, p *models.User) (int64, error) {
	query := "Insert tbl_user SET user_code=?, no_ktp=?, nama=?, gender=?, tanggal_lahir=?"

	TempBirthday := strings.Replace(string(p.Birthdate), "-", "", -1)
	TempName := p.Name[:3]

	UserCode := TempName + TempBirthday

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(
		ctx,
		UserCode,
		p.KTP,
		p.Name,
		p.Gender,
		p.Birthdate,
	)

	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlUserRepo) InsertLoan(ctx context.Context, p *models.User) (int64, error) {
	query := "Insert tbl_request_loans SET user_code=?, loan_code=?, jumlah_pinjaman=?, lama_tenor=?"

	TempBirthday := strings.Replace(string(p.Birthdate), "-", "", -1)
	TempName := p.Name[:3]

	UserCode := TempName + TempBirthday

	DateTimeNow := time.Now()
	TempDate := DateTimeNow.Format("01-02-2006")
	TempTime := DateTimeNow.Format("15:04:05")

	LoanCode := "LOAN" + strings.Replace(string(TempDate), "-", "", -1) + strings.Replace(string(TempTime), ":", "", -1)

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(
		ctx,
		UserCode,
		LoanCode,
		p.Amount,
		p.Tenor,
	)

	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
