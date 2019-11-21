package authenticator

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	// username, pwd, _ := r.BasicAuth()
	tokenString := strings.Split(r.Header["Authorization"][0], " ")[1]

	/****************/
	var hmacSampleSecret = []byte("apple")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims["name"], claims["pwd"])
		return true
	}
	fmt.Println(err)
	return false
}
