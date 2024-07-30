package repository

import (
	"API_TRAINING/entity"
	// "time"
)

func UpdateAccountBalance(accId string, balance float64) error {
	trx := Repository.DB.Exec(Repository.Query.UPDATE_ACCOUNT_BALANCE, balance, accId)

	return trx.Error
}

func InsertTransaction(entity *entity.Transaction) error {
	// now := time.Now()
	trx := Repository.DB.Exec(Repository.Query.INSERT_INTO_TRANSACTION, entity.TransactionID, entity.Type, entity.Amount,  entity.AccountID) 

	return trx.Error
}