package model

import "time"

type TransactionRequest struct {
	AccountID string  `json:"AccountID"`
	Amount    float64 `json:"amount"`
}

type TransactionResponse struct {
	TransactionID string    `json:"transactionId"`
	Type          string    `json:"type"`
	Amount        float64   `json:"amount"`
	TimeStamp     time.Time `json:"timestamp"`
	AccountID     string    `json:"AccountID"`
}
