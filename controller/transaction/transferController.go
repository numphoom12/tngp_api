package controller

import (
	model "API_TRAINING/model/transaction/transfer"
	service "API_TRAINING/service/transaction"
	"API_TRAINING/util"
	"encoding/json"
	"net/http"
)

func Transfer(w http.ResponseWriter, r *http.Request) {
	var transferRequest *model.TransferRequest
	err := json.NewDecoder(r.Body).Decode(&transferRequest)
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

	transferResp, err := service.TransferService(r, transferRequest)
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
		Result: transferResp,
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