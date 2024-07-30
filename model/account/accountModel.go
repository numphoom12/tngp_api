package model

type AccountRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountResponse struct {
	AccountID string  `json:"accountID"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Balance   float64 `json:"balance"`
}

type UpdateAccountRequest struct {
	AccountRequest
	Balance float64 `json:"balance"`
}

type UpdateAccountResponse struct {
	AccountResponse
}
