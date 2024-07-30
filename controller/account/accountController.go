package controller

import (
	model "API_TRAINING/model/account"
	service "API_TRAINING/service/account"
	"API_TRAINING/util"
	"encoding/json"
	"net/http"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	accountsResp, err := service.GetAccountsService(w, r)
	if err != nil{
		// http.Error(w, err.Error(), http.StatusBadRequest)
		// return

		response := util.ResponseWrapper{
            Code:  "500",
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
	}

	response := util.ResponseWrapper{
        Code:    "200",
        Message: "success",
		Result: accountsResp,
    }

	bytes, err := json.Marshal(response)
	if err != nil{
		response := util.ResponseWrapper{
            Code:  "500",
            Message: util.ErrServer.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func GetAccountsById(w http.ResponseWriter, r *http.Request) {
	accountsResp, err := service.GetAccountsByIdService(w, r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := util.ResponseWrapper{
        Code:    "200",
        Message: "success",
		Result: accountsResp,
    }

	bytes, err := json.Marshal(response)
	if err != nil{
		response := util.ResponseWrapper{
            Code:  "500",
            Message: util.ErrServer.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var accountRequest *model.AccountRequest
	err := json.NewDecoder(r.Body).Decode(&accountRequest)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		// return
		response := util.ResponseWrapper{
            Code:  "500",
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
	}

	accountResp, err := service.CreateAccountService(r, accountRequest)
	if err != nil{
		// http.Error(w, err.Error(), http.StatusBadRequest)
		// return
		response := util.ResponseWrapper{
            Code:  "500",
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
	}

	response := util.ResponseWrapper{
        Code:    "200",
        Message: "success",
		Result: accountResp,
    }

	bytes, err := json.Marshal(response)
	if err != nil{
		response := util.ResponseWrapper{
            Code:  "500",
            Message: util.ErrServer.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var accountRequest *model.UpdateAccountRequest
	err := json.NewDecoder(r.Body).Decode(&accountRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountsResp, err := service.UpdateAccountService(r, accountRequest)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := util.ResponseWrapper{
        Code:    "200",
        Message: "success",
		Result: accountsResp,
    }

	bytes, err := json.Marshal(response)
	if err != nil{
		response := util.ResponseWrapper{
            Code:  "500",
            Message: util.ErrServer.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func UpdateAccountDetail(w http.ResponseWriter, r *http.Request) {
	var accountRequest map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&accountRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountsResp, err := service.UpdateAccountDetailService(r, accountRequest)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := util.ResponseWrapper{
        Code:    "200",
        Message: "success",
		Result: accountsResp,
    }

	bytes, err := json.Marshal(response)
	if err != nil{
		response := util.ResponseWrapper{
            Code:  "500",
            Message: util.ErrServer.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	err := service.DeleteAccountService(r)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := util.ResponseWrapper{
        Code:    "200",
        Message: "success",
    }

	bytes, err := json.Marshal(response)
	if err != nil{
		response := util.ResponseWrapper{
            Code:  "500",
            Message: util.ErrServer.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
	}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}