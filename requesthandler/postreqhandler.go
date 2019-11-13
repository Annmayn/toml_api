package requesthandler

import (
	"encoding/json"
	"log"
	"net/http"
	"toml_api/authenticator"
	"toml_api/customvalidator"
	"toml_api/errorresponse"
	"toml_api/fileio"
	"toml_api/getresource"
	"toml_api/methodconfigs"
)

//config for POST Request
var postConfig methodconfigs.PostRequestConfig

//handle all incoming POST requests here
func PostHandler(w http.ResponseWriter, r *http.Request, config interface{}, loc []string) {
	resource := getresource.GetResource(config, loc[0], loc[1], "post")

	b, _ := json.Marshal(resource)
	json.Unmarshal(b, &postConfig)

	r.ParseForm()

	data := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		log.Println(err)
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
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(dataRequired)
		return
	}

	fileio.WriteToFile(necessaryData)

	//authenticate request
	if !authenticator.IsAuthenticated(r, getConfig.Auth) {
		errorresponse.ThrowError(w, "Request not authorized!")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(necessaryData)

}
