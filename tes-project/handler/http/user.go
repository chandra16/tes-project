package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"tes-project/driver"
	models "tes-project/models"
	repository "tes-project/repository"
	user "tes-project/repository/user"
)

// NewUserHandler ...
func NewUserHandler(db *driver.DB) *User {
	return &User{
		repo: user.NewSQLUserRepo(db.SQL),
	}
}

// User ...
type User struct {
	repo repository.UserRepo
}

// Create a new post
func (p *User) Create(w http.ResponseWriter, r *http.Request) {
	post := models.User{}
	json.NewDecoder(r.Body).Decode(&post)

	newID, err := p.repo.InsertUser(r.Context(), &post)
	newIDLoan, errLoan := p.repo.InsertLoan(r.Context(), &post)

	fmt.Println("Success insert User : ID => ", newID)
	fmt.Println("Success insert Loan : ID => ", newIDLoan)

	if err != nil || errLoan != nil {
		respondWithErrorUser(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSONUser(w, http.StatusCreated, map[string]string{"KTP": strconv.FormatInt(post.KTP, 10), "Status": "Valid"})
}

// respondwithJSON write json response format
func respondwithJSONUser(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithErrorUser(w http.ResponseWriter, code int, msg string) {
	respondwithJSONUser(w, code, map[string]string{"message": msg})
}
