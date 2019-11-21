package authenticator

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func CreateNew() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var hmacSampleSecret = []byte("apple")
		body := make(map[string]interface{})
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			fmt.Println("Error decoding body in authenticator")
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(body))

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)

		fmt.Println(tokenString, err)
	}
	return http.HandlerFunc(fn)
}
