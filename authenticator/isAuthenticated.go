package authenticator

import (
	"io"
	"net/http"
	"toml_api/jsonwebtoken"
)

func IsAuthenticated(w http.ResponseWriter, r *http.Request, auth bool) bool {
	if !auth {
		return true
	}

	authenticated := performAuthentication(w, r)
	return authenticated
}

func performAuthentication(w http.ResponseWriter, r *http.Request) bool {

	//verify token

	checkVerification, err := jsonwebtoken.IsAuthorized(r)

	if !checkVerification {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"`+err.Error()+`"}`)

		return false
	}

	return true

}
