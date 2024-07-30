package transfer

import "time"

type TransferRequest struct {
	SenderAccountID   string  `json:"senderAccountID"`
	RecieverAccountID string  `json:"recieverAccountID"`
	Amount            float64 `json:"amount"`
}

type TransferResponse struct {
	TransactionID     string    `json:"transactionId"`
	Type              string    `json:"type"`
	Amount            float64   `json:"amount"`
	TimeStamp         time.Time `json:"timestamp"`
	SenderAccountID   string    `json:"senderAccountID"`
	RecieverAccountID string    `json:"recieverAccountID"`
}
