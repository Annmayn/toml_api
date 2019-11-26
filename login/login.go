package login

import (
	"fmt"
	"net/http"
)

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
		FacebookLoginHandler(w, r)

		return

	}
	return http.HandlerFunc(fn)
}
