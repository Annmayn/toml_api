package handler

import (
	"fmt"
	"net/http"
	"strings"
	"toml_api/errorresponse"
	"toml_api/getresource"
	"toml_api/requesthandler"
)

//CustomHandler : redirects according to request method
func CustomHandler(config interface{}, apiEndPoint map[string][]string, loc []string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.String()
		fmt.Println("location", loc[1])
		fmt.Println("url", url)
		if url[len(url)-1] == '/' {
			url = url[:len(url)-1]
		}

		//eg: loc -> api/root
		//skip api and send only "root" to getresource
		methods := getresource.GetApiMethods(config, loc[1]) //returns map[string]bool
		if _, ok := methods[strings.ToLower(r.Method)]; ok {
			//method allowed
			//handle according to method
			switch r.Method {
			case "GET":
				//pass request to GET Request Handler
				requesthandler.GetHandler(w, config, loc)
			case "POST":
				requesthandler.PostHandler(w,config,loc)
			case "DELETE":
				requesthandler.DeleteHandler(w, config, loc)
			case "PUT":
				//pass request to PUT Request Handler
				requesthandler.PutHandler(w, config, loc)

			case "PATCH":
				requesthandler.PatchHandler(w, config, loc)
			default:
				errorresponse.ThrowError(w, "Method not recognized!")
			}
		} else {
			//method not allowed  : JSON response
			errorresponse.ThrowError(w, "Request Method not allowed!")

		}

	}
	return http.HandlerFunc(fn)

}
