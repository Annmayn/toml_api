package login

import (
	"io"
	"net/http"
	"toml_api/jsonwebtoken"
)

func Login() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		username, pwd, _ := r.BasicAuth()
		if username != "admin" || pwd != "admin" {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `{"error":"invalid_credentials"}`)
			return
		}

		jsonwebtoken.GenerateTokens(w, username)
		return

	}
	return http.HandlerFunc(fn)
}
