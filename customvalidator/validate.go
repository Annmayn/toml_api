package customvalidator

import (
	"strings"

	"toml_api/getresource"
)

/*

parameters::

1. config interface{}


2. validators []string

	eg.
	["$validator.name","$validator.age"]

3. schema string

	eg.
	"$schema.user"


4. toValidate map[string]interface{}

	data
	{
		"name" : "temp_name",
		"age" : 12
	}

returns::

1. data
2. error

*/

//Validate
func Validate(config interface{}, validators []string, schema string, toValidate map[string]interface{}) (map[string]interface{}, map[string]string) {

	//fmt.Println(toRequired["age"].(string))

	errorResult := make(map[string]string)

	data := make(map[string]interface{})

	//Compare schema and toValidate data

	resources := strings.Split(schema, ".")

	// resources[0][1:] ::  $schema =>schema
	toRequired := (getresource.GetResource(config, resources[0][1:], resources[1])).(map[string]interface{})

	//get all data within schema
	for i, v := range toValidate {
		if _, ok := toRequired[i]; ok {
			data[i] = v
		}
	}

	for i, v := range toRequired {
		temp := v.(string)
		if temp[len(temp)-1:] == "!" {
			if _, ok := toValidate[i]; !ok {
				validityResult[i] = "required"
			}
		}
	}

	for _, v := range validators {
		validator := strings.Split(v, ".")

		//type of validation
		i := validator[1]

		validMap := getresource.GetValidator(config, i).([]map[string]interface{})
		for _, x := range validMap {

			title := x["type"].(string)

			// Split type eg. string.min_length

			s := strings.Split(title, ".")

			if _, ok := toValidate[i]; ok {
				switch s[0] {
				case "string":
					if s[1] == "min_length" {
						if len(toValidate[i].(string)) < int(x["value"].(int64)) {
							validityResult[i] = x["error"].(string)
						}
					} else if s[1] == "max_length" {
						if len(toValidate[i].(string)) > int(x["value"].(int64)) {
							validityResult[i] = x["error"].(string)
						}

					}
				case "int":
					if s[1] == "min_value" {
						if int(toValidate[i].(float64)) < int(x["value"].(int64)) {
							validityResult[i] = x["error"].(string)
						}
					} else if s[1] == "max_value" {
						if int(toValidate[i].(float64)) > int(x["value"].(int64)) {
							validityResult[i] = x["error"].(string)
						}

					}
				}

			}

		}

	}

	return data, errorResult
}
