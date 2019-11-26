package jsonwebtoken

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
)

func TokenRefreshHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		secretKey = []byte(os.Getenv("SECRET_KEY"))

		token, err := jwt.ParseWithClaims(strings.Split(r.Header["Authorization"][0], " ")[1], &UserRefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Error with signature check")
			}
			return secretKey, nil
		})

		if claims, ok := token.Claims.(*UserRefreshTokenClaims); ok && token.Valid {

			if claims.TokenType == "refresh_token" {
				//create a new token
				//define access token claims
				username := claims.User
				accessTokenClaims := UserAccessTokenClaims{
					username,
					"access_token",
					[]string{"get", "post"},
					jwt.StandardClaims{
						ExpiresAt: time.Now().Add(time.Minute * time.Duration(accessTimeAdded)).Unix(),
						Issuer:    "rara",
						IssuedAt:  currentTime.Unix(),
					},
				}

				//define refresh token claims
				refreshTokenClaims := UserRefreshTokenClaims{
					username,
					"refresh_token",
					jwt.StandardClaims{
						ExpiresAt: time.Now().Add(time.Hour * time.Duration(refreshTimeAdded)).Unix(),
						Issuer:    "rara",
						IssuedAt:  currentTime.Unix(),
					},
				}

				//refresh token
				refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
				//access token
				accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

				access_token_string, err := accessToken.SignedString(secretKey)
				refresh_token_string, err := refreshToken.SignedString(secretKey)

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
