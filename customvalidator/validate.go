package customvalidator

import (
	"errors"
	"fmt"
	"strings"

	"toml_api/fieldValidator"
	"toml_api/getresource"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
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

//validates the variable type in request body against schema definition
func ValidateType(toValidate map[string]interface{}, schema map[string]string, errorMap map[string]string) bool {

	var hasError bool = false

	for k, v := range schema {
		if _, ok := toValidate[k]; ok {
			// fmt.Println(toValidate[k])
			var err error
			switch v {
			case "string":
				err = validation.Validate(toValidate[k], is.Alpha)

			case "int":
				tmp := fmt.Sprintf("%g", toValidate[k])
				err = validation.Validate(tmp, is.Int)

			case "float64":
				tmp := fmt.Sprintf("%g", toValidate[k])
				err = validation.Validate(tmp, is.Float)

			default:
				err = errors.New("couldn't decode type")
			}

			if err != nil {
				hasError = true
				errorMap[k] = err.Error()
			}
		}

	}
	return hasError
}

func checkIfExists(field string, toValidate map[string]interface{}, errorMap map[string]string) error {
	var err error
	if _, ok := toValidate[field]; !ok {
		errorMap[field] = "cannot be blank"
		err = errors.New("cannot be blank")
	}
	return err
}

func schemaValidation(config interface{}, data map[string]interface{}, validators []string, errorMap map[string]string) bool {
	var hasSchemaError bool = false
	for _, v := range validators { //eg validators: [$validator.name, $validator.age]
		validatorField := strings.Split(v, ".")[1]

		fieldSchema := getresource.GetValidator(config, validatorField).([]map[string]interface{})


		for _, schema := range fieldSchema { //for each schema in array of schemas such as $validator.name ([]interface{})

			fieldValidator.ValidateField(data[validatorField], schema, errorMap)
		}
	}

	if len(errorMap) != 0 {
		hasSchemaError = true
	}
	return hasSchemaError

	// 	for _, v := range validators {
	// 		validator := strings.Split(v, ".")

	// 		//type of validation
	// 		i := validator[1]

	// 		validMap := getresource.GetValidator(config, i).([]map[string]interface{})
	// 		for _, x := range validMap {

	// 			title := x["type"].(string)

	// 			// Split type eg. string.min_length

	// 			s := strings.Split(title, ".")

	// 			if _, ok := toValidate[i]; ok {
	// 				switch s[0] {
	// 				case "string":
	// 					if s[1] == "min_length" {
	// 						if len(toValidate[i].(string)) < int(x["value"].(int64)) {
	// 							validityResult[i] = x["error"].(string)
	// 						}
	// 					} else if s[1] == "max_length" {
	// 						if len(toValidate[i].(string)) > int(x["value"].(int64)) {
	// 							validityResult[i] = x["error"].(string)
	// 						}

	// 					}
	// 				case "int":
	// 					if s[1] == "min_value" {
	// 						if int(toValidate[i].(float64)) < int(x["value"].(int64)) {
	// 							validityResult[i] = x["error"].(string)
	// 						}
	// 					} else if s[1] == "max_value" {
	// 						if int(toValidate[i].(float64)) > int(x["value"].(int64)) {
	// 							validityResult[i] = x["error"].(string)
	// 						}

	// 					}
	// 				}

	// 			}

	// 		}

	// 	}
}

//main validate function
//schema: $schema.father, $schema.mother etc
func Validate(config interface{}, validators []string, schema string, toValidate map[string]interface{}) (map[string]interface{}, map[string]string) {

	//fmt.Println(toRequired["age"].(string))

	//error map
	errorMap := make(map[string]string)

	//Compare schema and toValidate data
	//schema : $schema.user
	resources := strings.Split(schema, ".") //[$schema, user]

	// resources[0][1:] ::  $schema =>schema
	toRequired := (getresource.GetResource(config, resources[0][1:], resources[1])).(map[string]interface{})

	fmt.Println(toRequired)

	//data type of schema
	dataTypeOfSchema := make(map[string]string)

	for field, dType := range toRequired {
		dataType := dType.(string)
		if dataType[len(dataType)-1] == '!' {
			dataTypeOfSchema[field] = dataType[:len(dataType)-1]
			//check if required values exists in body
			//todo
			_ = checkIfExists(field, toValidate, errorMap)
		} else {
			dataTypeOfSchema[field] = dataType
		}
	}
	if len(errorMap) != 0 {
		return map[string]interface{}{}, errorMap
	}

	// validate type : returns error and updates errorMap
	hasTypeError := ValidateType(toValidate, dataTypeOfSchema, errorMap)

	if hasTypeError {
		return map[string]interface{}{}, errorMap
	}

	//get all data within schema
	data := make(map[string]interface{})
	for i, v := range toValidate {
		if _, ok := toRequired[i]; ok {
			data[i] = v
		}
	}

	hasSchemaError := schemaValidation(config, data, validators, errorMap)
	if hasSchemaError {
		return data, errorMap
	}

	return data, errorMap
}
