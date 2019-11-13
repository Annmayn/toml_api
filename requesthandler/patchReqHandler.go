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
)

//config for Patch Request
var patchConfig methodconfigs.PatchRequestConfig

//handle all incoming GET requests here
func PatchHandler(w http.ResponseWriter, r *http.Request, config interface{}, loc []string) {
	resource := getresource.GetResource(config, loc[0], loc[1], "patch")

	b, _ := json.Marshal(resource)
	json.Unmarshal(b, &patchConfig)

	//authenticate request
	if !authenticator.IsAuthenticated(r, getConfig.Auth) {
		errorresponse.ThrowError(w, "Request not authorized!")
		return
	}

	var schema string
	if len(patchConfig.Data)==0{
		schema=patchConfig.Schema
	}else{
		schema=patchConfig.Data
	}

	s:=strings.Replace(schema,"$","",1)
	sch:=strings.Split(s,".")

	schemaData:=getresource.GetResource(config,sch[0],sch[1])

	//verify incoming data according to schema and patch
	data:=make(map[string]interface{})
	//
	if strings.Compare(r.Header.Get("Content-Type"),strings.Trim(r.Header.Get("Content-Type"),"\n"))!=0{
		errorresponse.ThrowError(w,"Content-Type Header is not application/json")
		return
	}

	if err:=json.NewDecoder(r.Body).Decode(&data);err!=nil{
		errorresponse.ThrowError(w,fmt.Sprint("%s",err))
		return
	}

	//necessary data and error result
	necessaryData, validityResult := customvalidator.Validate(config, patchConfig.Validator, schema, data)

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
	} else if len(validityResult) > 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(validityResult)
		return
	}


	//read data from file
	queryResults:=fileio.ReadFromFile()
	//create a map for response
	response:=make(map[string]interface{})
	//get ```result``` value from config file
	resultWithoutDollar:=strings.Replace(patchConfig.Result,"$","",1)
	res:=strings.Split(resultWithoutDollar,".")
	//send response as per schema defined in ```result```
	responseResult:=getresource.GetResource(config,res[0],res[1])
	//check which keys are required in response
	resultFields:=make(map[string]bool)
	//create a map of required keys in resultFields
	for key,_:=range responseResult.(map[string]interface{}){
		resultFields[key]=true
	}

	//determine which fields need to be patched
	patchFields:=make(map[string]bool)
	for key,_:=range necessaryData{
		patchFields[key]=true
	}

	//patch data
	dataToWrite:=make(map[string]interface{})
	for key,_:=range schemaData.(map[string]interface{}){
		if _,ok:=patchFields[key];ok{
			v:=necessaryData[key]
			dataToWrite[key]=v
		}else{
			for k,v:=range queryResults.(map[string]interface{}){
				if k==key{
					dataToWrite[key]=v
				}
			}
		}

	}

	//write patched data
	fileio.WriteToFile(dataToWrite)

	//send response according to schema specified in ```result```
	//read again
	queryResults=fileio.ReadFromFile()
	for key,val:=range queryResults.(map[string]interface{}){
		if _,ok:=resultFields[key];ok {
			response[key]=val
		}
	}

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(response)
	return




}
