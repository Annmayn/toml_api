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

//config for DELETE Request
var deleteConfig methodconfigs.DeleteRequestConfig

//handle all incoming DELETE requests here
func DeleteHandler(w http.ResponseWriter, r *http.Request, config interface{}, loc []string) {
	resource := getresource.GetResource(config, loc[0], loc[1], "delete")

	b, _ := json.Marshal(resource)
	json.Unmarshal(b, &deleteConfig)

	//authenticate request
	if !authenticator.IsAuthenticated(r, getConfig.Auth) {
		errorresponse.ThrowError(w, "Request not authorized!")
		return
	}

	/*
		TODO:
			1. Auth
			2. Query table with attachments
			3. Get Results and send ```result``` & ```display``` response to client
	*/
	fmt.Fprintf(w, "i am in delete handler")

}
