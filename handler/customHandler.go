package handler

import (
	"fmt"
	"net/http"
	"strings"
	"toml_api/requestconfig"
)

//CustomHandler : redirects according to request method
func CustomHandler(apiEndPoint map[string][]string) http.HandlerFunc{
	fn := func(w http.ResponseWriter, r *http.Request){
		url := r.URL.String()
		if url[len(url)-1] == '/' {
			url = url[:len(url)-1]
		}
		if _, ok := apiEndPoint[url]; ok{
			loc := apiEndPoint[url]
			
			//eg: loc -> api/root
			//skip api and send only "root" to getresource
			methods := getresource.GetAnishCode(loc[1])	//returns map[string]
			if _, ok := methods[r.Method]{
				//method allowed
			}else{
				//method not allowed
			}
		} else{
			//handle error because endpoint doesn't exist
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
