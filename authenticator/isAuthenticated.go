package authenticator

func IsAuthenticated(auth bool)	bool {
	if !auth{
		return true
	}else{
		authenticated := performAuthentication()
		return authenticated
	}
}

func performAuthentication() bool {
	//logic for authentication

	//return true if authorized else false
	return true

}
