package controller

import (
	model "API_TRAINING/model/transaction/withdraw"
	service "API_TRAINING/service/transaction"
	"API_TRAINING/util"
	"encoding/json"
	"net/http"
)

func Withdraw(w http.ResponseWriter, r *http.Request) {
	var withdrawRequest *model.WithdrawRequest
	err := json.NewDecoder(r.Body).Decode(&withdrawRequest)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		// return
		response := util.ResponseWrapper{
			Code:    "400",
			Message: util.ErrInvalidInput.Error(),
		}
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
	}
	
	withdrawResp, err := service.WithDrawService(r, withdrawRequest)
	if err != nil{
		// http.Error(w, err.Error(), http.StatusBadRequest)
		// return
		response := util.ResponseWrapper{
            Code:  "500",
            Message: util.ErrServer.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
	}

	response := util.ResponseWrapper{
        Code:    "200",
        Message: "success",
		Result: withdrawResp,
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