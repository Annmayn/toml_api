package login

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"toml_api/errorresponse"
	"toml_api/jsonwebtoken"
)

var config = &oauth2.Config{
	ClientID:     "431310981126615",
	ClientSecret: "2b4360d7f2247d664aed53bd945725dc",
	Endpoint: oauth2.Endpoint{
		AuthURL:   "https://graph.facebook.com/oauth/authorize",
		TokenURL:  "https://graph.facebook.com/oauth/access_token",
	},
	RedirectURL:  "http://localhost:8080/api/auth/callback",
	Scopes: []string{"email","user_likes","user_birthday","user_gender"},
}


//OAuth handler for facebook OAuth
func FacebookLoginHandler(w http.ResponseWriter, r *http.Request){

	//config=oauth2.Config{
	//	ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
	//	ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
	//	Endpoint:     oauth2.Endpoint{
	//		AuthURL:   os.Getenv("FACEBOOK_AUTH_URL"),
	//		TokenURL:  os.Getenv("FACEBOOK_TOKEN_URL"),
	//	},
	//	RedirectURL:   os.Getenv("FACEBOOK_REDIRECT_URL"),
	//	Scopes:       []string{"email","user_likes","user_birthday","user_gender"},
	//}

	url:=config.AuthCodeURL("state")
	http.Redirect(w,r,url,http.StatusTemporaryRedirect)
}

//facebook auth callback
func FacebookCallbackHandler(w http.ResponseWriter, r *http.Request){
	state:=r.FormValue("state")
	if state!="state"{
		errorresponse.ThrowError(w,"Invalid oauth state")
		return
		//http.Redirect(w,r,"/api",http.StatusTemporaryRedirect)  //307
	}


	token,err:=config.Exchange(oauth2.NoContext,r.URL.Query().Get("code"))
	if err!=nil{
		errorresponse.ThrowError(w,fmt.Sprintf("%s",err))
		return
		//http.Redirect(w,r,"/api",http.StatusTemporaryRedirect)
	}

	response,err:=http.Get("https://graph.facebook.com/me?fields=name,email,gender&access_token=" + token.AccessToken)
	if err!=nil{
		//throw error here
		errorresponse.ThrowError(w,fmt.Sprintf("%s",err))

		return
	}

	defer response.Body.Close()

	var facebookResponse interface{}
	contents,err:=ioutil.ReadAll(response.Body)
	json.Unmarshal(contents,&facebookResponse)

	jsonwebtoken.GenerateTokens(w, facebookResponse.(map[string]interface{})["name"].(string))
	return
}


