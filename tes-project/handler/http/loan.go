package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"tes-project/driver"
	models "tes-project/models"
	repository "tes-project/repository"
	loan "tes-project/repository/loan"
)

// NewLoanHandler ...
func NewLoanHandler(db *driver.DB) *Loan {
	return &Loan{
		repo: loan.NewSQLLoanRepo(db.SQL),
	}
}

// Loan ...
type Loan struct {
	repo repository.LoanRepo
}

// Create a new installment
func (p *Loan) Create(w http.ResponseWriter, r *http.Request) {
	post := models.Loan{}
	json.NewDecoder(r.Body).Decode(&post)

	// Loan DATA
	Amount := post.Amount
	Tenor := post.Tenor
	RequestDate, _ := time.Parse("2006-01-02", post.Date)

	// GET Interest
	TempInterest, _ := p.repo.GetInterest(r.Context(), Tenor)
	Interest := TempInterest[0].Interest

	TempAmountInstallment := Amount / Tenor
	TempAmountInterest := float64(Amount) * Interest / 100

	totalInstallment := float64(TempAmountInstallment) + TempAmountInterest

	// Loan CODE
	DateTimeNow := time.Now()
	TempDate := DateTimeNow.Format("01-02-2006")
	TempTime := DateTimeNow.Format("15:04:05")

	LoanCode := "LOAN" + strings.Replace(string(TempDate), "-", "", -1) + strings.Replace(string(TempTime), ":", "", -1)

	for i := 1; i <= Tenor; i++ {
		// fmt.Println("============================================================")
		// fmt.Println("== Temp Amount Installment 	: ", TempAmountInstallment)
		// fmt.Println("== Temp Amount Interest		: ", TempAmountInterest)
		// fmt.Println("== Total Installment		: ", totalInstallment)
		// fmt.Println("== Request Date			: ", RequestDate.AddDate(0, i, 0).Format("2006-01-02"))
		// fmt.Println("============================================================")

		newID, err := p.repo.InsertInstallment(r.Context(), LoanCode, float64(TempAmountInstallment), TempAmountInterest, totalInstallment, i, RequestDate.AddDate(0, i, 0).Format("2006-01-02"))

		if err != nil {
			respondWithErrorLoan(w, http.StatusInternalServerError, "Server Error")
		}

		fmt.Println("Success insert installment : ID => ", newID)
	}
	payload, _ := p.repo.GetInstallmentByLoanCode(r.Context(), LoanCode)

	respondwithJSONLoan(w, http.StatusCreated, payload)
}

// TrackLoan to Track average of loan application
func (p *Loan) TrackLoan(w http.ResponseWriter, r *http.Request) {
	loan := models.LoanTrack{}
	json.NewDecoder(r.Body).Decode(&loan)

	RequestDate, _ := time.Parse("2006-01-02", loan.Date)

	payload, _ := p.repo.GetLoanTrack(r.Context(), RequestDate.AddDate(0, 0, -7).Format("2006-01-02"), RequestDate.Format("2006-01-02"))

	var totalLoan float64

	for index := 0; index < len(payload); index++ {
		totalLoan = totalLoan + payload[0].JumlahPinjaman
	}

	average := totalLoan / 7
	countLoan := len(payload)

	respondwithJSONLoan(w, http.StatusCreated, map[string]string{"Count": strconv.Itoa(countLoan), "Summary": fmt.Sprintf("%.0f", totalLoan), "7-day avg": fmt.Sprintf("%.0f", average)})
}

// respondwithJSON write json response format
func respondwithJSONLoan(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithErrorLoan(w http.ResponseWriter, code int, msg string) {
	respondwithJSONLoan(w, code, map[string]string{"message": msg})
}
