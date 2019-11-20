package requesthandler

import (
	"encoding/json"
	"net/http"
	"toml_api/authenticator"
	"toml_api/errorresponse"
	"toml_api/getresource"
	"toml_api/methodconfigs"
	"toml_api/query"
	"toml_api/responsehandler"
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
	/*
		Current implementation reads only from a database.json file.
		However, a generic QueryHandler has been created for future implementations!
	*/
	result := query.QueryHandler(config, getConfig.Query, getConfig.QueryParams, getConfig.Result, getConfig.Attachments)

	//send response
	responsehandler.SendJSONResponse(w, result, 200)
	return

}
