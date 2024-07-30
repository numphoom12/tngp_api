package main

import (
	acc_controller "API_TRAINING/controller/account"
	trx_controller "API_TRAINING/controller/transaction"
	"API_TRAINING/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	err := repository.InitDB()
	if err != nil{
		log.Fatal("DB CONNECTION ERROR")
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Post("/accounts", acc_controller.CreateAccount)
	r.Get("/accounts", acc_controller.GetAccounts)
	r.Get("/accounts/{id}", acc_controller.GetAccountsById)
	r.Put("/accounts/{id}", acc_controller.UpdateAccount)
	r.Patch("/accounts/{id}", acc_controller.UpdateAccountDetail)
	r.Delete("/accounts/{id}", acc_controller.DeleteAccount)

	r.Post("/transactions/deposit", trx_controller.Deposit)
	r.Post("/transactions/withdraw", trx_controller.Withdraw)
	r.Post("/transactions/transfer", trx_controller.Transfer)
	
	fmt.Println("run leaw")
	http.ListenAndServe(":3000", r)
}