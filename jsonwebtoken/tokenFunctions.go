package jsonwebtoken

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var secretKey []byte

// Create a new token object, specifying signing method and the claims
var currentTime = time.Now()

//time for refreshtoken in hours
var refreshTimeAdded = 720

//time for accesstoken in minutes
var accessTimeAdded = 5

type UserAccessTokenClaims struct {
	User        string   `json:"user"`
	TokenType   string   `json:"type"`
	Permissions []string `json:"permissions"`
	jwt.StandardClaims
}

type UserRefreshTokenClaims struct {
	User      string `json:"user"`
	TokenType string `json:"type"`
	jwt.StandardClaims
}

//generate refresh tokens and access tokens
func GenerateTokens(w http.ResponseWriter, username string) {
	secretKey = []byte(os.Getenv("SECRET_KEY"))

	//define access token claims
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

	// Sign and get the complete encoded token as a string using the secret
	refreshTokenString, err1 := refreshToken.SignedString(secretKey)

	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"refresh_token_generation_failed"}`)
		return
	}
	//get a signed access token string
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
	secretKey = []byte(os.Getenv("SECRET_KEY"))
	if len(secretKey) == 0 {
		log.Fatal("HTTP Server unable to start, expected an APP_KEY for JWT auth")
	}

	token, err := jwt.ParseWithClaims(strings.Split(r.Header["Authorization"][0], " ")[1],&UserAccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Error with signature check")
		}
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*UserAccessTokenClaims); ok && token.Valid {
		if claims.TokenType == "access_token" {

			if exists:=hasPermission(claims.Permissions,r.Method);exists{
				return true,nil
			}else{
				return false, fmt.Errorf("User not authorized")
			}
			//return false, nil
		} else if claims.TokenType == "refresh_token" {
			return false, fmt.Errorf("token invalid")
		}
	}

	return false, err

}


func hasPermission(permissions []string,method string)bool{

	for _,val:=range permissions{
		if method==strings.ToUpper(val){
			return true
		}
	}
	return false
}