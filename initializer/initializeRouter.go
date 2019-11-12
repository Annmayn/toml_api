package initializer

import (
	"github.com/gorilla/mux"
	"strings"
	"toml_api/handler"
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
	kv := make(map[string][]string)
	for endpoint, loc := range apiEndPoint {
		ep := parseEndpoint(endpoint)

		//store apiEndPoint values of corresponding keys in apiEndPoint in kv
		kv[ep] = loc
		r.HandleFunc(ep, handler.CustomHandler(config, kv, loc))
	}
	return kv
}
