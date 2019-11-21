package jsonwebtoken

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TokenRefreshHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		token, err := jwt.Parse(strings.Split(r.Header["Authorization"][0], " ")[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Error with signature check")
			}
			return secretKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//fmt.Println(claims["user"], claims["exp"])
			//w.WriteHeader(http.StatusUnauthorized)
			//io.WriteString(w, `{"error":"refresh_token_expired"}`)
			//return
			if claims["type"] == "refresh_token" {
				//create a new token
				//access token
				access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"user": claims["user"],
					"type": "access_token",
					"iat":  currentTime,
					"exp":  time.Now().Add(time.Minute * time.Duration(accessTimeAdded)).Unix(),
				})

				access_token_string, err := access_token.SignedString(secretKey)

				//create a new refresh token
				refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"user": claims["user"],
					"type": "refresh_token",
					"iat":  currentTime,
					"exp":  time.Now().Add(time.Hour * time.Duration(refreshTimeAdded)).Unix(),
				})

				refresh_token_string, err := refresh_token.SignedString(secretKey)

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					io.WriteString(w, `{"error":"token_generation_failed"}`)
					return
				}

				//send the token to client
				io.WriteString(w, `{"access_token":"`+access_token_string+`","refresh_token":"`+refresh_token_string+`"}`)
				return

			} else {
				w.WriteHeader(http.StatusUnauthorized)
				io.WriteString(w, `{"error":"invalid_refresh_token"}`)
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error":"`+fmt.Sprint("%s", err)+`"}`)
			return

		}

	}
	return http.HandlerFunc(fn)

}
