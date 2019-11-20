package requesthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"toml_api/authenticator"
	"toml_api/customvalidator"
	"toml_api/errorresponse"
	"toml_api/fileio"
	"toml_api/getresource"
	"toml_api/methodconfigs"
	"toml_api/responsehandler"
)

var mutex = sync.Mutex{}

//config for POST Request
var postConfig methodconfigs.PostRequestConfig

//handle all incoming POST requests here
func PostHandler(w http.ResponseWriter, r *http.Request, config interface{}, loc []string) {
	resource := getresource.GetResource(config, loc[0], loc[1], "post")

	b, _ := json.Marshal(resource)
	json.Unmarshal(b, &postConfig)

	fmt.Println(postConfig.Auth)

	//authenticate request
	if !authenticator.IsAuthenticated(r, postConfig.Auth) {
		errorresponse.ThrowError(w, "Request not authorized!")
		return
	}

	data := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		errorresponse.ThrowError(w, "Error while decoding incoming data")
		return
	}

	//necessary data and error result
	necessaryData, validityResult := customvalidator.Validate(config, postConfig.Validator, postConfig.Data, data)

	dataRequired := make(map[string]string)

	for k, v := range validityResult {
		if v == "required" {
			dataRequired[k] = v
		}
	}

	if len(dataRequired) > 0 {
		//send JSON response
		responsehandler.SendJSONResponse(w, dataRequired, 404)
		return
	} else if len(validityResult) > 0 {
		//send JSON response
		responsehandler.SendJSONResponse(w, validityResult, 404)
		return
	}

	fileio.WriteToFile(necessaryData)

	//send JSON response
	responsehandler.SendJSONResponse(w, necessaryData, 201)
	return
}
