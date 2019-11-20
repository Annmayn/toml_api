package query

import (
	"fmt"
	"strings"
	"toml_api/fileio"
	"toml_api/getresource"
)

func QueryHandler(config interface{}, query string, queryParams string, result string, attachments []string) interface{} {
	//filter query according to resultFields
	resultFields := make(map[string]bool)
	response := make(map[string]interface{})

	result = strings.Replace(result, "$", "", 1)
	res := strings.Split(result, ".")

	_, newResultInterface := getresource.GetResource(config, res[0], res[1])
	for key := range newResultInterface.(map[string]interface{}) {
		resultFields[key] = true
	}

	//read from db
	queryResults := fileio.ReadFromFile()
	fmt.Println(queryResults)

	//filter query results according to resultFields
	for key, val := range queryResults.(map[string]interface{}) {
		if _, ok := resultFields[key]; ok {
			response[key] = val
		}
	}
	//return query result
	return response

	/*
	   TODO: SinkPlugin will receive table name, query params, attachments result
	*/

}
