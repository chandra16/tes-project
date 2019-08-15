package models

// Loan type details
type Loan struct {
	Date   string `json:"date"`
	Amount int    `json:"amount"`
	Tenor  int    `json:"tenor"`
}
