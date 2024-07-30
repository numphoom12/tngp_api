package model

import (
	model "API_TRAINING/model/transaction"
)

type DepositRequest struct {
	model.TransactionRequest
}

type DepositResponse struct {
	model.TransactionResponse
}