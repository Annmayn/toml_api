package requesthandler

import (
	"encoding/json"
	"net/http"
	"toml_api/authenticator"
	"toml_api/customvalidator"
	"toml_api/errorresponse"
	"toml_api/fileio"
	"toml_api/methodconfigs"
	"toml_api/responsehandler"
)

//handle all incoming POST requests here
func PostHandler(config interface{}, postConfig methodconfigs.PostRequestConfig, loc []string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//authenticate request

		if !authenticator.IsAuthenticated(w, r, postConfig.Auth) {
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
	return http.HandlerFunc(fn)
}
