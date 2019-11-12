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

//config for POST Request
var postConfig methodconfigs.PostRequestConfig

//handle all incoming POST requests here
func PostHandler(w http.ResponseWriter, r *http.Request, config interface{},loc []string){
	resource:=getresource.GetResource(config,loc[0],loc[1],"post")

	b,_:=json.Marshal(resource)
	json.Unmarshal(b, &postConfig)

	//authenticate request
	if !authenticator.IsAuthenticated(getConfig.Auth){
		errorresponse.ThrowError(w,"Request not authorized!")
		return
	}


	/*
		TODO:
			1. Auth
			2. Query table with attachments
			3. Get Results and send ```result``` & ```display``` response to client
	*/
	fmt.Fprintf(w,"i am in POST handler")


}
