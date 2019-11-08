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

		for _, v := range temp {
			x := v.(map[string]interface{})

			title := x["type"].(string)

			// Split type eg. string.min_length

			s := strings.Split(title, ".")

			if _, ok := toValidate[i]; ok {
				switch s[0] {
				case "string":
					if s[1] == "min_length" {
						if len(toValidate[i].(string)) < int(x["value"].(float64)) {
							validityResult[i] = x["error"].(string)
						}
					} else if s[1] == "max_length" {
						if len(toValidate[i].(string)) > int(x["value"].(float64)) {
							validityResult[i] = x["error"].(string)
						}

					}
				case "int":
					if s[1] == "min_value" {
						if int(toValidate[i].(float64)) < int(x["value"].(float64)) {
							validityResult[i] = x["error"].(string)
						}
					} else if s[1] == "max_value" {
						if int(toValidate[i].(float64)) > int(x["value"].(float64)) {
							validityResult[i] = x["error"].(string)
						}

					}
				}

			}

		}

	}

	return validityResult
}
