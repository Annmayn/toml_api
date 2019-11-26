package login

import (
	"fmt"
	"net/http"
)

<<<<<<< HEAD
//type login interface{
//	Authenticate()
//	Callback()
//	GenerateToken()
//}
//
//type basicAuth struct{
//	username string
//	password  string
//}
//
//func (auth *basicAuth) Authenticate(){
//	//username password validate
//}
//
//func (auth *basicAuth) GenerateToken(){
//
//}
//
//type OAuth struct{
//	//initialization
//}
//
//func (oAuth *OAuth) Authenticate(){
//
//}
//
//func (oAuth *OAuth) Callback(){
//
//}
//
//func (oAuth *OAuth) GenerateToken(){
//
//}


func Login() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//username, pwd, _ := r.BasicAuth()
		//if username != "admin" || pwd != "admin" {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	io.WriteString(w, `{"error":"invalid_credentials"}`)
		//	return
		//}
		fmt.Println("come to login")
		//use OAuth here
		FacebookLoginHandler(w,r)

		return
=======
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
>>>>>>> 68995691f405f17b4384e848af34ac83cba48908

	}
	return http.HandlerFunc(fn)
}
