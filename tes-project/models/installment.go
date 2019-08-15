package models

// Installment type details
type Installment struct {
	Capital  float64 `json:"capital"`
	Interest float64 `json:"interest"`
	Total    float64 `json:"total"`
	Plan     int     `json:"plan"`
	DueDate  string  `json:"duedate"`
}
