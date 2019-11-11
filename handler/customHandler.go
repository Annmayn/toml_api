package handler

import (
	"fmt"
	"net/http"
	"strings"
	"toml_api/getresource"
)

//CustomHandler : redirects according to request method
func CustomHandler(config interface{}, apiEndPoint map[string][]string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		if url[len(url)-1] == '/' {
			url = url[:len(url)-1]
		}
		if _, ok := apiEndPoint[url]; ok {
			loc := apiEndPoint[url]

			//eg: loc -> api/root
			//skip api and send only "root" to getresource
			methods := getresource.GetApiMethods(config, loc[1]) //returns map[string]bool
			if _, ok := methods[strings.ToLower(r.Method)]; ok {
				//method allowed
				fmt.Fprint(w, "Get access allowed")
			} else {
				//method not allowed
				fmt.Fprint(w, "Not allowed")
			}
		} else {
			//handle error because endpoint doesn't exist
			fmt.Fprint(w, "Endpoint not defined")
		}
	}
	return http.HandlerFunc(fn)

	// //handle according to method
	// switch r.Method {
	// case "GET":

	// case "POST":
	// case "DELETE":
	// case "PUT":
	// case "PATCH":
	// default:
	// 	fmt.Fprint(w, "Method not recognized")
	// }
}
