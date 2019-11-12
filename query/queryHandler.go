package query

import (
	"strings"
	"toml_api/getresource"
)

func QueryHandler(config interface{}, query string, queryParams string, result string, attachments []string) interface{}{
	//filter query according to resultFields
	var resultFields []string

	result=strings.Replace(result,"$","",1)
	res:=strings.Split(result,".")

	newResultInterface:=getresource.GetResource(config,res[0],res[1])
	for key,_:=range newResultInterface.(map[string]interface{}){
		resultFields=append(resultFields,key)
	}

	//read from db

	//fetch query result
	//return query result

	return newResultInterface

	/*
	   SinkPlugin will receive table name, query params, attachments result
	 */
	//marshal


}