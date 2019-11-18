package fieldValidator

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateField(fieldValue interface{}, schema map[string]interface{}, errorMap map[string]string) {
	var err error
	cval := schema["value"]

	switch schema["type"].(string) {

	//String handling
	case "string.min_length":
		err = validation.Validate(fieldValue, validation.Length(int(cval.(int64)), 0).Error(schema["error"].(string)))

	case "string.max_length":
		err = validation.Validate(fieldValue, validation.Length(0, int(cval.(int64))).Error(schema["error"].(string)))

	//Int handling
	case "int.min_value":
		//the code to check data type runs before this function is called
		//so we are certain that the value is int at this point
		fv := int(fieldValue.(float64))
		err = validation.Validate(fv, validation.Min(cval).Error(schema["error"].(string)))

	case "int.max_value":
		fv := int(fieldValue.(float64))
		err = validation.Validate(fv, validation.Max(cval).Error(schema["error"].(string)))

	}
	if err != nil {
		errorMap[schema["field"].(string)] = err.Error()
		fmt.Println(errorMap)
	}
}
