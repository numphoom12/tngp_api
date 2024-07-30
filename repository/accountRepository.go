package repository

import (
	"API_TRAINING/entity"
	model "API_TRAINING/model/account"
	"errors"
)

func InsertAccount(entity *entity.Accounts) error {
	trx := Repository.DB.Exec(
		Repository.Query.INSERT_ACCOUNT, entity.AccountID, entity.Name, entity.Email, entity.Balance,
	)

	return trx.Error
}

func SelectAccounts() ([]*model.AccountResponse, error){
	var accountsResp []*model.AccountResponse
	var accountEntity entity.Accounts

	rows, _ := Repository.DB.Raw(Repository.Query.SELECT_ACCOUNTS).Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&accountEntity.AccountID, &accountEntity.Name, &accountEntity.Email, &accountEntity.Balance);

		accModel := &model.AccountResponse{
			AccountID: accountEntity.AccountID,
			Name:      accountEntity.Name,
			Email:     accountEntity.Email,
			Balance:   accountEntity.Balance,
		}

		accountsResp = append(accountsResp, accModel)
	}
	return accountsResp, nil
}

func SelectAccountByID(accID string) (*model.AccountResponse, error) {
	var account entity.Accounts

	trx := Repository.DB.Raw(
		Repository.Query.SELECT_ACCOUNT_BY_ID, accID,
	).Scan(&account)

	if trx.Error != nil {
		return nil, trx.Error
	}

	if trx.RowsAffected == 0 {
		return nil, errors.New("ACCOUNT NOT FOUND")
	}

	return &model.AccountResponse{
		AccountID: account.AccountID,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
	}, nil
}

// func CreateAccountRepository(entity *entity.Accounts) error {
// 	trx := repository.Repository.DB.Create(&entity)
// 	if trx.Error != nil {
// 		return trx.Error
// 	}

// 	return nil
// }

// func GetAccounts(entity *entity.Accounts) ([]*model.AccountResponse, error){
//     var accountResponses []*model.AccountResponse

//     // Retrieve all accounts from the database
//     result := repository.Repository.DB.Find(&entity)
//     if result.Error != nil {
//         return nil, result.Error
//     }

//     // Transform database entities to response model
//     for _, acc := range re {
//         accountResponses = append(
//         	accountResponses,
//         	&model.AccountResponse{
//             AccountID:      acc.ID,
//             Name:    		acc.Name,
//             Email:   		acc.Email,
//             Balance: 		acc.Balance,
//         },
//         )
//     }

//     return accountResponses, nil
// }
