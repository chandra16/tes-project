package main

import (
	"fmt"
	// "go-mysql-crud-master/driver"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"tes-project/driver"
	ph "tes-project/handler/http"
)

func main() {
	dbName := "test-tunaiku"
	dbUsername := "root"
	dbPass := "novendra1"
	dbHost := "localhost"
	dbPort := "3306"

	connection, err := driver.ConnectSQL(dbHost, dbPort, dbUsername, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandlerUser := ph.NewUserHandler(connection)
	pHandlerLoan := ph.NewLoanHandler(connection)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/users", userRouter(pHandlerUser))
		rt.Mount("/loans", loanRouter(pHandlerLoan))
	})

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)
}

// A completely separate router for posts routes
func userRouter(pHandler *ph.User) http.Handler {
	r := chi.NewRouter()
	r.Post("/insert", pHandler.Create)

	return r
}

// A completely separate router for posts routes
func loanRouter(pHandler *ph.Loan) http.Handler {
	r := chi.NewRouter()
	r.Post("/insert-installment", pHandler.Create)
	r.Post("/track-loan", pHandler.TrackLoan)

	return r
}
