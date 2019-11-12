package errorresponse

import (
	"encoding/json"
	"net/http"
)

type ErrResponse struct{
	Error_message string `json:"error"`
}

func ThrowError(w http.ResponseWriter, err string){
	var errResponse ErrResponse
	errResponse.Error_message = err
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(errResponse)
}
