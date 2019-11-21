package jsonwebtoken

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("df567nm,^&*gbsnaas1212#@")

// Create a new token object, specifying signing method and the claims
var currentTime = time.Now()

//time for refreshtoken in hours
var refreshTimeAdded = 720

//time for accesstoken in minutes
var accessTimeAdded = 5

//generate refresh tokens and access tokens
func GenerateTokens(w http.ResponseWriter, username string) {

	//refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"type": "refresh_token",
		"iat":  currentTime,
		"exp":  time.Now().Add(time.Hour * time.Duration(refreshTimeAdded)).Unix(),
	})

	//access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"type": "access_token",
		"iat":  currentTime,
		"exp":  time.Now().Add(time.Minute * time.Duration(accessTimeAdded)).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	refreshTokenString, err1 := refreshToken.SignedString(secretKey)

	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"refresh_token_generation_failed"}`)
		return
	}

	accessTokenString, err2 := accessToken.SignedString(secretKey)

	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"access_token_generation_failed"}`)
		return
	}

	io.WriteString(w, `{"refresh_token":"`+refreshTokenString+`","access_token":"`+accessTokenString+`"}`)
	return

}

func IsAuthorized(r *http.Request) (bool, error) {

	if len(secretKey) == 0 {
		log.Fatal("HTTP Server unable to start, expected an APP_KEY for JWT auth")
	}

	token, err := jwt.Parse(strings.Split(r.Header["Authorization"][0], " ")[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error with signature check")
		}
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user"], claims["exp"])
		if claims["type"] == "access_token" {
			return true, nil
		} else if claims["type"] == "refresh_token" {
			return false, fmt.Errorf("token invalid")
		}
	}

	return false, err

}
