package entity

import "time"

type Transaction struct {
	TransactionID string    `gorm:"TransactionId"`
	Type          string    `gorm:"Type"`
	Amount        float64   `gorm:"Amount"`
	TimeStamp     time.Time `gorm:"Timestamp"`
	AccountID     string    `gorm:"AccountID"`
	// RecieverAccountID string    `gorm:"RecieverAccountID"`
}

type Deposit struct {
	Transaction
}

type Withdraw struct {
	Transaction
}

type Transfer struct {
	TransactionID     string    `gorm:"TransactionID"`
	Type              string    `gorm:"Type"`
	Amount            float64   `gorm:"Amount"`
	TimeStamp         time.Time `gorm:"Timestamp"`
	SenderAccountID   string    `gorm:"SenderAccountID"`
	RecieverAccountID string    `gorm:"RecieverAccountID"`
}
