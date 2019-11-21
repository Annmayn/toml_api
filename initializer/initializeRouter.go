package initializer

import (
	"encoding/json"
	"net/http"
	"strings"
	"toml_api/authenticator"
	"toml_api/errorresponse"
	"toml_api/getresource"
	"toml_api/methodconfigs"
	"toml_api/requesthandler"

	"github.com/gorilla/mux"
)

//todo: validate url
//validates url and maps variable values of config endpoint for mapping
func parseEndpoint(endpoint string) string {
	epArr := strings.Split(endpoint, "/")[1:] //drop first "/"
	fixedEndpointArr := make([]string, len(epArr))
	for i, val := range epArr {
		if val[0] == ':' {
			val = "{" + val[1:] + "}"
		}
		fixedEndpointArr[i] = val
	}
	return "/" + strings.Join(fixedEndpointArr[:], "/")
}

//todo: error handling
/*
	Parses, validates and maps url to customHandler
		- recursively parses url to expand the url

	Returns the new key-value pair mapping enpoint to schema location
*/
func InitializeRouter(r *mux.Router, config interface{}, apiEndPoint map[string][]string) map[string][]string {
	var getConfig methodconfigs.GetRequestConfig
	var postConfig methodconfigs.PostRequestConfig
	var deleteConfig methodconfigs.DeleteRequestConfig
	var putConfig methodconfigs.PutRequestConfig
	var patchConfig methodconfigs.PatchRequestConfig
	var hasError bool

	//kv = map : endpoint-> root path
	/*
		eg:
			kv : {
					"api/foo/bar/baz": ["api" "root"],
					"api/...":[api detail]
				}
	*/
	kv := make(map[string][]string)
	for endpoint, loc := range apiEndPoint {
		ep := parseEndpoint(endpoint) // maps "api/:foo/:bar" to "api/{foo}/{bar}"

		//store apiEndPoint values of corresponding keys in apiEndPoint in kv
		kv[ep] = loc //eg: kv["/api/{foo}/{bar}/{baz}"] = ["api" "root"]

		//get the resources for every methods
		resources := make(map[string]interface{})
		methods := []string{"get", "post", "put", "patch", "delete"}

		for _, method := range methods {
			hasError, resources[method] = getresource.GetResource(config, loc[0], loc[1], method)
			// if len(resources[method]) != 0{}
			b, _ := json.Marshal(resources[method])
			if !hasError {
				if method == "get" {
					json.Unmarshal(b, &getConfig)
					r.HandleFunc(ep, requesthandler.GetHandler(config, getConfig, loc)).Methods("GET")
				} else if method == "post" {
					json.Unmarshal(b, &postConfig)
					r.HandleFunc(ep, requesthandler.PostHandler(config, postConfig, loc)).Methods("POST")
				} else if method == "put" {
					json.Unmarshal(b, &putConfig)
					r.HandleFunc(ep, requesthandler.PutHandler(config, putConfig, loc)).Methods("PUT")
				} else if method == "delete" {
					json.Unmarshal(b, &deleteConfig)
					r.HandleFunc(ep, requesthandler.DeleteHandler(config, deleteConfig, loc)).Methods("DELETE")
				} else if method == "patch" {
					json.Unmarshal(b, &patchConfig)
					r.HandleFunc(ep, requesthandler.PatchHandler(config, patchConfig, loc)).Methods("PATCH")
				}
			}
		}

		// r.HandleFunc(ep, handler.CustomHandler(config, loc))

	}
	//Handle all the undefined endpoints
	r.HandleFunc("/auth", authenticator.CreateNew()).Methods("POST")
	r.PathPrefix("/").HandlerFunc(notDefined)
	return kv
}

func notDefined(w http.ResponseWriter, r *http.Request) {
	errorresponse.ThrowError(w, "Endpoint not defined")
}
