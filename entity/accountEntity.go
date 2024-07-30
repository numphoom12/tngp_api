package entity

type Accounts struct {
	AccountID string  `gorm:"AccountID"`
	Name      string  `gorm:"Name"`
	Email     string  `gorm:"Email"`
	Balance   float64 `gorm:"Balance"`
}

type UpdateAccount struct {
	Accounts
}
