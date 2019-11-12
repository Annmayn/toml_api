package handler

import (
	"fmt"
	"net/http"
	"strings"
	"toml_api/getresource"
	"toml_api/errorresponse"
)



//CustomHandler : redirects according to request method
func CustomHandler(config interface{}, apiEndPoint map[string][]string,loc []string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.String()
		fmt.Println("location",loc[1])
		fmt.Println("url",url)
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
				fmt.Println(getresource.GetResource(config,loc[0],loc[1],strings.ToLower(r.Method)))
			case "POST":

			case "DELETE":

			case "PUT":

			case "PATCH":

			default:
				errorresponse.ThrowError(w,"Method not recognized!")
			}
		} else {
			//method not allowed  : JSON response
			errorresponse.ThrowError(w,"Request Method not allowed!")

		}

	}
	return http.HandlerFunc(fn)


}
