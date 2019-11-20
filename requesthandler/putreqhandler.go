package requesthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"toml_api/authenticator"
	"toml_api/customvalidator"
	"toml_api/errorresponse"
	"toml_api/fileio"
	"toml_api/getresource"
	"toml_api/methodconfigs"
	"toml_api/responsehandler"
)

//config for PUT Request
var putConfig methodconfigs.PutRequestConfig

//handle all incoming PUT requests here
func PutHandler(w http.ResponseWriter, r *http.Request, config interface{}, loc []string) {

	resource := getresource.GetResource(config, loc[0], loc[1], "put")

	b, _ := json.Marshal(resource)
	json.Unmarshal(b, &putConfig)

	//authenticate request
	if !authenticator.IsAuthenticated(r, putConfig.Auth) {
		errorresponse.ThrowError(w, "Request not authorized!")
		return
	}

	var schema string = putConfig.Data

	s := strings.Replace(schema, "$", "", 1)
	sch := strings.Split(s, ".")

	schemaData := getresource.GetResource(config, sch[0], sch[1])

	//verify incoming data according to schema and patch
	data := make(map[string]interface{})
	//
	if strings.Compare(r.Header.Get("Content-Type"), strings.Trim(r.Header.Get("Content-Type"), "\n")) != 0 {
		errorresponse.ThrowError(w, "Content-Type Header is not application/json")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errorresponse.ThrowError(w, fmt.Sprint("%s", err))
		return
	}

	//validate data
	//necessary data and error result
	necessaryData, validityResult := customvalidator.Validate(config, putConfig.Validator, schema, data)

	dataRequired := make(map[string]string)

	for k, v := range validityResult {
		if v == "required" {
			dataRequired[k] = v
		}
	}

	if len(dataRequired) > 0 {
		responsehandler.SendJSONResponse(w, dataRequired, 404)
		return
	} else if len(validityResult) > 0 {
		responsehandler.SendJSONResponse(w, validityResult, 404)
		return
	}

	//read data from file
	queryResults := fileio.ReadFromFile()
	//create a map for response
	response := make(map[string]interface{})
	//get ```result``` value from config file
	resultWithoutDollar := strings.Replace(putConfig.Result, "$", "", 1)
	resultArr := strings.Split(resultWithoutDollar, ".")
	//send response as per schema defined in ```result```
	responseResult := getresource.GetResource(config, resultArr[0], resultArr[1])
	//check which keys are required in response
	resultFields := make(map[string]bool)
	//create a map of required keys in resultFields
	for key, _ := range responseResult.(map[string]interface{}) {
		resultFields[key] = true
	}

	//put new data
	dataToWrite := make(map[string]interface{})

	for key, _ := range schemaData.(map[string]interface{}) {
		if _, ok := necessaryData[key]; ok {
			dataToWrite[key] = necessaryData[key]
		} else if resultFields[key] {
			dataToWrite[key] = necessaryData[key]
		}
	}

	//write new data
	fileio.WriteToFile(dataToWrite)

	//send response according to schema specified in ```result```
	//read again
	queryResults = fileio.ReadFromFile()
	for key, val := range queryResults.(map[string]interface{}) {
		if _, ok := resultFields[key]; ok {
			response[key] = val
		}
	}

	responsehandler.SendJSONResponse(w, response, 201)
	return
}
