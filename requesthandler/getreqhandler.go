package requesthandler

import (
	"net/http"
	"toml_api/authenticator"
	"toml_api/methodconfigs"
	"toml_api/query"
	"toml_api/responsehandler"
)

//handle all incoming GET requests here

func GetHandler(config interface{}, getConfig methodconfigs.GetRequestConfig, loc []string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {

		//authenticate request
		if !authenticator.IsAuthenticated(w, r, getConfig.Auth) {
			return
		}

		//may come in handy, look up google for more info (request header connection close)
		r.Close = true

		//perform next steps
		/*
			Current implementation reads only from a database.json file.
			However, a generic QueryHandler has been created for future implementations!
		*/
		result := query.QueryHandler(config, getConfig.Query, getConfig.QueryParams, getConfig.Result, getConfig.Attachments)

		//send response
		responsehandler.SendJSONResponse(w, result, 200)
	}
	return http.HandlerFunc(fn)
}
