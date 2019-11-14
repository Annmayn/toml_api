package responsehandler

import (
	"net/http"
	"encoding/json"
)
/*
This package is for sending response to client as per
schema defined in ```result```.

It takes map[string]interface{} as input, checks its keys as
per schema in ```result``` and returns a map[string]interface{}
*/



func SendJSONResponse(w http.ResponseWriter, response interface{},statusCode int){
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
	return
}
