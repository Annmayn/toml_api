package requesthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"toml_api/getresource"
	"toml_api/methodconfigs"
)

//config for GET Request
var getConfig methodconfigs.GetRequestConfig

//handle all incoming GET requests here
func GetHandler(w http.ResponseWriter, config interface{}, loc []string) {
	resource := getresource.GetResource(config, loc[0], loc[1], "get")

	b, _ := json.Marshal(resource)
	json.Unmarshal(b, &getConfig)

	fmt.Println(getConfig)

	/*
		TODO:
			1. Auth
			2. Query table with attachments
			3. Get Results and send ```result``` & ```display``` response to client
	*/
	fmt.Fprintf(w, "i am in get handler")

}
