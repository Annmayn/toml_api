package customvalidator

import (
	"fmt"
	"math"
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

1. data map[string]interface{}
2. error map[string]string

*/

//validate type
func ValidateType(toValidate map[string]interface{}, dataTypeOfSchema map[string]string) (map[string]string, bool) {

	typeError := make(map[string]string)

	var hasError bool = false

	fmt.Println(dataTypeOfSchema)

	for i, v := range toValidate {
		typeOfData := fmt.Sprintf("%T", v)

		if typeOfData == "float64" {
			if math.Trunc(v.(float64)) == v.(float64) {

				if dataTypeOfSchema[i] != "int" {
					typeError[i] = dataTypeOfSchema[i] + " required"
				}

			} else {
				if dataTypeOfSchema[i] != "float64" {
					typeError[i] = dataTypeOfSchema[i] + " required"
				}

			}
		} else if dataTypeOfSchema[i] != typeOfData {
			typeError[i] = dataTypeOfSchema[i] + " required"
		}

	}

	if len(typeError) > 0 {
		hasError = true
	}

	return typeError, hasError

}

//main validate function
func Validate(config interface{}, validators []string, schema string, toValidate map[string]interface{}) (map[string]interface{}, map[string]string) {

	//fmt.Println(toRequired["age"].(string))

	//error map
	validityResult := make(map[string]string)

	//type error map
	typeValidityResult := make(map[string]string)

	data := make(map[string]interface{})

	//Compare schema and toValidate data

	resources := strings.Split(schema, ".")

	// resources[0][1:] ::  $schema =>schema
	toRequired := (getresource.GetResource(config, resources[0][1:], resources[1])).(map[string]interface{})

	//data type of schema
	dataTypeOfSchema := make(map[string]string)
	for i, v := range toRequired {

		dataType := v.(string)

		if dataType[len(dataType)-1] == '!' {
			dataTypeOfSchema[i] = dataType[:len(dataType)-1]
		} else {
			dataTypeOfSchema[i] = dataType
		}

	}

	// validate type
	typeValidityResult, hasTypeError := ValidateType(toValidate, dataTypeOfSchema)

	if hasTypeError {
		return data, typeValidityResult
	}

	//get all data within schema
	for i, v := range toValidate {
		if _, ok := toRequired[i]; ok {
			data[i] = v
		}
	}

	//test for required data
	for i, v := range toRequired {
		dataType := v.(string)

		// if data type has ! in last i.e required data type
		if dataType[len(dataType)-1:] == "!" {
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

	return data, validityResult
}
