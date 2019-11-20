package requesthandler

import (
	"net/http"
	"strings"
	"toml_api/authenticator"
	"toml_api/errorresponse"
	"toml_api/fileio"
	"toml_api/methodconfigs"
	"toml_api/responsehandler"
)

//DeleteHandler : handles all incoming DELETE requests
func DeleteHandler(config interface{}, deleteConfig methodconfigs.DeleteRequestConfig, loc []string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		var deleteKey string

		//authenticate request
		if !authenticator.IsAuthenticated(r, deleteConfig.Auth) {
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
			fileio.WriteToFile(map[string]string{})
			m := map[string]interface{}{deleteKey: fileKV[deleteKey]}
			//send JSON response
			responsehandler.SendJSONResponse(w, m, 200)
			return
		}
		//throw error
		errorresponse.ThrowError(w, "item not found")
	}
	return http.HandlerFunc(fn)
}
