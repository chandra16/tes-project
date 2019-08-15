package models

// User type details
type User struct {
	KTP       int64  `json:"ktp"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Gender    string `json:"gender"`
	Amount    int64  `json:"amount"`
	Tenor     int64  `json:"tenor"`
}
