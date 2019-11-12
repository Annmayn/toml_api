package requesthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"toml_api/authenticator"
	"toml_api/errorresponse"
	"toml_api/getresource"
	"toml_api/methodconfigs"
)

//config for GET Request
var getConfig methodconfigs.GetRequestConfig

//handle all incoming GET requests here

func GetHandler(w http.ResponseWriter, r *http.Request, config interface{}, loc []string) {
	resource := getresource.GetResource(config, loc[0], loc[1], "get")

	b, _ := json.Marshal(resource)
	json.Unmarshal(b, &getConfig)

	//authenticate request
	if !authenticator.IsAuthenticated(r, getConfig.Auth) {
		errorresponse.ThrowError(w, "Request not authorized!")
		return
	}
	//perform next steps
	//result:=queryTable(getConfig.Query,getConfig.QueryParams,getConfig.Attachments)

	/*
		TODO:
			1. Auth [Status : Pseudo Implementation]
			2. Query table with attachments
			3. Get Results and send ```result``` & ```display``` response to client
	*/
	fmt.Fprintf(w, "i am in get handler")

}
