package customvalidator

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

	var validityResult map[string]string

	validMap := fromValidate["validator"].(map[string]interface{})

	//Name array interface

	tempNameMap := validMap["name"].([]interface{})

	//Age map interface
	ageValidity := validMap["age"].(map[string]interface{})

	nameMinLengthValidity := tempNameMap[0].(map[string]interface{})
	nameMaxLengthValidity := tempNameMap[1].(map[string]interface{})

	//Name validity parameters
	nameMinLengthValue := int(nameMinLengthValidity["value"].(float64))
	nameMaxLengthValue := int(nameMaxLengthValidity["value"].(float64))

	//Age validity parameters
	ageMinLenghtValue := ageValidity["value"].(float64)

	//Name and age to check for validity
	nameToValid := toValidate["name"].(string)
	ageToValid := toValidate["age"].(float64)

	//Validate Name
	if len(nameToValid) >= nameMinLengthValue && len(nameToValid) <= nameMaxLengthValue {
		validityResult["name"] = "name valid"
	} else if len(nameToValid) < nameMinLengthValue {
		validityResult["name"] = nameMinLengthValidity["error"].(string)
	} else if len(nameToValid) > nameMaxLengthValue {
		validityResult["name"] = nameMaxLengthValidity["error"].(string)
	}

	//Validate Age
	if ageToValid >= ageMinLenghtValue {
		validityResult["age"] = "age valid"
	} else {
		validityResult["age"] = ageValidity["error"].(string)
	}

	return validityResult
}
