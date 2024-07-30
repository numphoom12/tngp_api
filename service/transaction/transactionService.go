package service

import (
	"API_TRAINING/entity"
	common "API_TRAINING/model/transaction"
	deposit_model "API_TRAINING/model/transaction/deposit"
	transfer_model "API_TRAINING/model/transaction/transfer"
	withdraw_model "API_TRAINING/model/transaction/withdraw"
	"API_TRAINING/repository"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var counter int

func DepositService(r *http.Request, in *deposit_model.DepositRequest) (*deposit_model.DepositResponse, error) {
	if in.Amount <= 0 {
		return nil, errors.New("AMOUNT MUST MORE THAN 0")
	}

	counter++
	now := time.Now()

	// fmt.Println(now)
	acc, err := repository.SelectAccountByID(in.AccountID)
	if err != nil{
		// repository.Repository.DB.Rollback()
		return nil, err
	}

	newBalance := acc.Balance
	newBalance += in.Amount

	trxID := fmt.Sprintf("transaction%03d", counter)
	trxEntity := &entity.Transaction{
		TransactionID: trxID,
		Type:          "deposit",
		Amount:        in.Amount,

		AccountID:     in.AccountID,
	}

	// repository.Repository.DB.Begin()
	err = repository.InsertTransaction(trxEntity)
	if err != nil{
		// repository.Repository.DB.Rollback()
		return nil, err
	}

	err = repository.UpdateAccountBalance(in.AccountID, newBalance)
	if err != nil{
		// repository.Repository.DB.Rollback()
		return nil, err
	}
	// repository.Repository.DB.Commit()
	
	return &deposit_model.DepositResponse{
		TransactionResponse: common.TransactionResponse{
			TransactionID: trxID,
			Type:          "deposit",
			Amount:        in.Amount,
			TimeStamp:     now,
			AccountID:     in.AccountID,
		},
	}, err
}

func WithDrawService(r *http.Request, in *withdraw_model.WithdrawRequest) (*withdraw_model.WithdrawResponse, error) {
	if in.Amount <= 0 {
		return nil, errors.New("AMOUNT MUST MORE THAN 0")
	}

	counter++
	now := time.Now()

	// withdrawResponse, err := repository(in.AccountID, in.Amount)
	acc, err := repository.SelectAccountByID(in.AccountID)
	if err != nil{
		// repository.Repository.DB.Rollback()
		return nil, err
	}

	newBalance := acc.Balance - in.Amount

	if newBalance < 0{
		return nil, errors.New("BALANCE NOT ENOUGE")
	}

	trxID := fmt.Sprintf("transaction%03d", counter)
	trxEntity := &entity.Transaction{
		TransactionID: trxID,
		Type:          "withdraw",
		Amount:        in.Amount,
		AccountID:     in.AccountID,
	}

	err = repository.InsertTransaction(trxEntity)
	if err != nil{
		return nil, err
	}

	err = repository.UpdateAccountBalance(in.AccountID, newBalance)
	if err != nil{
		return nil, err
	}

	return &withdraw_model.WithdrawResponse{
		TransactionResponse: common.TransactionResponse{
			TransactionID: trxID,
			Type:          "withdraw",
			Amount:        in.Amount,
			TimeStamp:     now,
			AccountID:     in.AccountID,
		},
	}, err
}

func TransferService(r *http.Request, in *transfer_model.TransferRequest) (*transfer_model.TransferResponse, error) {
	counter++
	now := time.Now()

	if in.Amount <= 0 {
		return nil, errors.New("AMOUNT MUST MORE THAN 0")
	}

	if in.SenderAccountID == in.RecieverAccountID {
		return nil, errors.New("RECIEVER ACCOUNT IS SAME AS SENDER ACCOUNT")
	}

	senderAcc, err := repository.SelectAccountByID(in.SenderAccountID)
	if err != nil {
		return nil, err
	}

	recieveAcc, err :=  repository.SelectAccountByID(in.RecieverAccountID)
	if err != nil {
		return nil, err
	}

	trxID := fmt.Sprintf("transaction%03d", counter)
	trxEntity := &entity.Transaction{
		TransactionID: trxID,
		Type:          "transfer",
		Amount:        in.Amount,
		AccountID:     in.RecieverAccountID,
	}

	err = repository.InsertTransaction(trxEntity)
	if err != nil {
		return nil, err
	}

	if senderAcc.Balance - in.Amount < 0 {
		return nil, errors.New("BALANCE NOT ENOUGE")
	}

	newSenderBalance := senderAcc.Balance - in.Amount
	newRecieveBalance := recieveAcc.Balance + in.Amount
	
	err = repository.UpdateAccountBalance(in.SenderAccountID, newSenderBalance)
	if err != nil{
		return nil, err
	}

	err = repository.UpdateAccountBalance(in.RecieverAccountID, newRecieveBalance)
	if err != nil{
		return nil, err
	}

	return &transfer_model.TransferResponse{
		TransactionID:     trxID,
		Type:              "transfer",
		Amount:            in.Amount,
		TimeStamp:         now,
		SenderAccountID:   in.SenderAccountID,
		RecieverAccountID: in.RecieverAccountID,
	}, nil
	
	// transactionResponse, error := transfer(in.SenderAccountID, in.RecieverAccountID, in.Amount)
	// return transactionResponse, error
}
