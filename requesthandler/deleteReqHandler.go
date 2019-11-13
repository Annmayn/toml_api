package requesthandler

import (
	"encoding/json"
	"net/http"
	"strings"
	"toml_api/authenticator"
	"toml_api/errorresponse"
	"toml_api/fileio"
	"toml_api/getresource"
	"toml_api/methodconfigs"
)

//config for DELETE Request
var deleteConfig methodconfigs.DeleteRequestConfig

//DeleteHandler : handles all incoming DELETE requests
func DeleteHandler(w http.ResponseWriter, r *http.Request, config interface{}, loc []string) {
	var deleteKey string
	resource := getresource.GetResource(config, loc[0], loc[1], "delete")

	b, _ := json.Marshal(resource)
	json.Unmarshal(b, &deleteConfig)

	//authenticate request
	if !authenticator.IsAuthenticated(r, getConfig.Auth) {
		errorresponse.ThrowError(w, "Request not authorized!")
		return
	}

	fileContent := fileio.ReadFromFile()
	fileKV := fileContent.(map[string]interface{})
	urlArr := strings.Split(r.URL.String(), "/")
	if urlArr[len(urlArr)-1] == "/" {
		deleteKey = urlArr[len(urlArr)-2]
	} else {
		deleteKey = urlArr[len(urlArr)-1]

	}

	if _, ok := fileKV[deleteKey]; ok {
		fileio.WriteToFile("{}")
	} else {
		errorresponse.ThrowError(w, "item not found")
	}
}
