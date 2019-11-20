package fieldValidator

import (
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateField(fieldValue interface{}, schema map[string]interface{}) string {
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
		fv := int64(fieldValue.(float64))
		cvalConverted := cval.(int64)
		strCval := fmt.Sprintf("%v", cval)
		if fv < cvalConverted {
			err = errors.New("can't be less than " + strCval)
		}

	case "int.max_value":
		fv := int64(fieldValue.(float64))
		cvalConverted := cval.(int64)
		strCval := fmt.Sprintf("%v", cval)
		if fv > cvalConverted {
			err = errors.New("can't be more than " + strCval)
		}
	}
	if err != nil {
		return err.Error()
	}
	return ""
}
