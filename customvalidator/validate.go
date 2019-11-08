package customvalidator

import "strings"

/*
fromValidate map[string]interface{}

{
	"validator" : {
		"name" : [
			{"field" = "name","type" = "string.min_length",value" = "2","error" = 'Minimum required length is $value'},
			{"field" = "name","type" = "string.max_length",value" = "50","error" = 'Maximum required length is $value'}
		]
		"age" : [
			{"field" = "age", "type" = "int.min_value", "value" = 18, "error" = 'Minimum age is $value'}
		]
	}
}

toValidate map[string]interface{}

{
	"name" : "temp_name",
	"age" : 12
}

*/

//Validate : validate name,age
func Validate(fromValidate map[string]interface{}, toValidate map[string]interface{}) map[string]string {

	validityResult := make(map[string]string)

	validMap := fromValidate["validator"].(map[string]interface{})

	for i, v := range validMap {

		temp := v.([]interface{})

		min := temp[0].(map[string]interface{})
		max := temp[1].(map[string]interface{})

		s := strings.Split(min["type"].(string), ".")

		switch s[0] {
		case "string":
			if len(toValidate[i].(string)) >= int(min["value"].(float64)) && len(toValidate[i].(string)) <= int(max["value"].(float64)) {
				validityResult[i] = i + " is valid"
			} else if len(toValidate[i].(string)) < int(min["value"].(float64)) {
				validityResult[i] = min["error"].(string)
			} else {
				validityResult[i] = max["error"].(string)
			}
		case "int":
			if int(toValidate[i].(float64)) >= int(min["value"].(float64)) && int(toValidate[i].(float64)) <= int(max["value"].(float64)) {
				validityResult[i] = i + " is valid"
			} else if int(toValidate[i].(float64)) < int(min["value"].(float64)) {
				validityResult[i] = min["error"].(string)
			} else {
				validityResult[i] = max["error"].(string)
			}

		}

	}

	return validityResult
}
