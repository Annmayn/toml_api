package requesthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"toml_api/getresource"
	"toml_api/methodconfigs"
)

//config for PUT Request
var putConfig methodconfigs.PutRequestConfig

//handle all incoming PUT requests here
func PutHandler(w http.ResponseWriter, config interface{}, loc []string) {

	fmt.Println(loc)
	resource := getresource.GetResource(config, loc[0], loc[1], "put")

	fmt.Println(resource)

	b, _ := json.Marshal(resource)
	json.Unmarshal(b, &putConfig)

	fmt.Println(putConfig)

	/*
		TODO:
			1. Auth
			2. Query table with attachments
			3. Get Results and send ```result``` & ```display``` response to client
	*/
	fmt.Fprintf(w, "i am in put handler")

}
