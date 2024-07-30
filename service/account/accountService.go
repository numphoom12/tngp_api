package service

import (
	"API_TRAINING/entity"
	model "API_TRAINING/model/account"
	repository "API_TRAINING/repository"
	// service "API_TRAINING/service/account"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var Accounts []*model.AccountResponse

func CreateAccountService(r *http.Request, in *model.AccountRequest) (*model.AccountResponse, error) {
	id := uuid.New()

	accEntity := &entity.Accounts{
		AccountID: id.String(),
		Name:      in.Name,
		Email:     in.Email,
		Balance:   0,
	}

	err := repository.InsertAccount(accEntity)
	if err != nil {
		return nil, err
	}

	accountResp := &model.AccountResponse{
		AccountID: id.String(),
		Name:      in.Name,
		Email:     in.Email,
		Balance:   0,
	}

	return accountResp, nil
}

func GetAccountsService(w http.ResponseWriter, r *http.Request) ([]*model.AccountResponse, error) {

	accounts, err := repository.SelectAccounts()
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetAccountsById(accID string) (*model.AccountResponse, error) {
	for _, account := range Accounts {
		if account.AccountID == accID {
			return account, nil
		}
	}

	return nil, errors.New("ACCOUNT NOT FOUND")
}

func GetAccountsByIdService(w http.ResponseWriter, r *http.Request) (*model.AccountResponse, error) {
	accID := chi.URLParam(r, "id")
	accResp, err := repository.SelectAccountByID(accID)
	if err != nil {
		return nil, err
	}

	return accResp, nil
}

func updateAccount(accRequest *model.UpdateAccountRequest, accID string) (*model.UpdateAccountResponse, error) {
	for _, account := range Accounts {
		if account.AccountID == accID {
			account.Name = accRequest.Name
			account.Email = accRequest.Email
			account.Balance = accRequest.Balance

			return &model.UpdateAccountResponse{
				AccountResponse: model.AccountResponse{
					AccountID: accID,
					Name:      accRequest.Name,
					Email:     accRequest.Email,
					Balance:   accRequest.Balance,
				},
			}, nil
		}
	}
	return nil, errors.New("ACCOUNT NOT FOUND")
}

func UpdateAccountService(r *http.Request, in *model.UpdateAccountRequest) (*model.UpdateAccountResponse, error) {
	accID := chi.URLParam(r, "id")
	accountResponse, err := updateAccount(in, accID)
	if err != nil {
		return nil, err
	}

	return accountResponse, nil
}

func updateAccountDetail(in map[string]interface{}, accID string) (*model.UpdateAccountResponse, error) {
	acc, err := GetAccountsById(accID)
	if err != nil {
		return nil, err
	}

	updateAccReq := &model.UpdateAccountRequest{
		AccountRequest: model.AccountRequest{
			Name:  acc.Name,
			Email: acc.Email,
		},
		Balance: acc.Balance,
	}

	if name, ok := in["name"].(string); ok {
		updateAccReq.Name = name
	}
	if email, ok := in["email"].(string); ok {
		updateAccReq.Email = email
	}
	if balance, ok := in["balance"].(float64); ok {
		updateAccReq.Balance = balance
	}

	updateAccountDetailResp, err := updateAccount(updateAccReq, accID)
	if err != nil {
		return nil, err
	}

	return updateAccountDetailResp, nil
}

func UpdateAccountDetailService(r *http.Request, in map[string]interface{}) (*model.UpdateAccountResponse, error) {
	accID := chi.URLParam(r, "id")
	accountResponse, err := updateAccountDetail(in, accID)
	if err != nil {
		return nil, err
	}

	return accountResponse, nil
}

func removeSliceElem(slice []*model.AccountResponse, index int) error {
	if index >= len(slice) || index < 0 {
		return errors.New("INDEX OUT OF RANGE")
	}
	slice[index] = slice[len(slice)-1]
	Accounts = slice[:len(slice)-1]

	return nil
}

func deleteAccount(accID string) error {
	accIndex := -1
	for i, account := range Accounts {
		if accID == account.AccountID {
			accIndex = i
		}
	}

	if accIndex == -1 {
		return errors.New("ACCOUNT ID NOT FOUND")
	}

	err := removeSliceElem(Accounts, accIndex)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAccountService(r *http.Request) error {
	accID := chi.URLParam(r, "id")

	err := deleteAccount(accID)
	if err != nil {
		return err
	}

	return nil
}
