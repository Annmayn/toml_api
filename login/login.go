package login

import (
	"io"
	"net/http"
	"toml_api/jsonwebtoken"
)

type LoginUser interface {
	Authenticate() (int, string)
}

type BasicAuth struct {
	username string
	password string
}

func (basic BasicAuth) Authenticate() (int, string) {

	uname := "admin"
	pwd := "admin"

	var unauthorized_error string = "none"

	if basic.username != uname || basic.password != pwd {

		unauthorized_error := "invalid_credentials"

		return http.StatusUnauthorized, unauthorized_error
	}

	return 0, unauthorized_error

}

func Login() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		var authentication LoginUser

		//define authentication type
		username, pwd, _ := r.BasicAuth()

		//assign to interface
		authentication = BasicAuth{username, pwd}

		error_code, error_message := authentication.Authenticate()

		if error_code == http.StatusUnauthorized {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `{"error":"`+error_message+`"}`)
			return

		} else {
			jsonwebtoken.GenerateTokens(w, username)
			return

		}

	}
	return http.HandlerFunc(fn)
}
