package authenticator

import (
	"net/http"
)

func IsAuthenticated(r *http.Request, auth bool) bool {
	if !auth {
		return true
	}
	authenticated := performAuthentication(r)
	return authenticated
}

func performAuthentication(r *http.Request) bool {
	//logic for authentication
	username, pwd, _ := r.BasicAuth()
	if username == "admin" && pwd == "admin" {
		return true
	}
	return false

}
