package getresource

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
	Receives config, keys as args
	Returns interface{}

	Need to perform typecheck for map[string]interface{}, []map[string]interface{}, string, int64
*/
func GetResource(config interface{}, keys ...interface{}) interface{} {
	//iterates over keys and continuously perform assertion
	for _, p := range keys {
		switch t := p.(type) {
		case string:
			config = config.(map[string]interface{})[t]
		case int:
			log.Println("received integer")
		default:
			fmt.Println(t)
		}

	}
	return config
}

/*

Needs to typecheck like this

val:=getresource.GetResource(c,"api","root","put","validator")

	switch vv:=val.(type){
		case []map[string]interface{}:
			fmt.Println("type []interface{}")
			for i,v:=range vv{
				res[i]=v
			}
		case map[string]interface{}:
			fmt.Println("type interface")
			fmt.Println(vv)
		case string:
			fmt.Println(vv)
		case int64:
			fmt.Println("int aayo")
		default:
			fmt.Println(reflect.TypeOf(vv))
	}
	fmt.Println(res)
*/

//get validators in map[string]interface{}
//replaces $value by their corresponding values
//function returns nil if validator config is not in array form
func GetValidator(config interface{}, validatorName string) interface{} {
	config = config.(map[string]interface{})["validator"]

	if len(validatorName) == 0 {
		//this is not working as of now 11:11
		switch newConfig := config.(type) {
		case []map[string]interface{}:
			for _, v := range newConfig {
				switch v["value"].(type) {
				case string:
					v["error"] = strings.Replace(v["error"].(string), "$value", v["value"].(string), 1)

				case int64:
					str := strconv.Itoa(int(v["value"].(int64)))
					v["error"] = strings.Replace(v["error"].(string), "$value", str, 1)
				}
			}
			return newConfig
		}

	} else {
		config = config.(map[string]interface{})[validatorName]
		switch newConfig := config.(type) {
		case []map[string]interface{}:
			for _, v := range newConfig {
				switch v["value"].(type) {
				case string:
					v["error"] = strings.Replace(v["error"].(string), "$value", v["value"].(string), 1)

				case int64:
					str := strconv.Itoa(int(v["value"].(int64)))
					v["error"] = strings.Replace(v["error"].(string), "$value", str, 1)
				}
			}
			return newConfig
		}

	}
	return nil
}

//get set of methods for requestLevel
/*
	#Args: -config		(reader instance of toml file)
          -requestLevel   (e.g.: root,detail,approve)

	#Response: map[string]bool with set of allowed methods on that request level

	e.g.: map[delete:true get:true patch:true post:true put:true]


	#Example:
			GetApiMethods(config,"root")
		#Response:
			map[delete:true get:true patch:true post:true put:true]


	Note: Function will return nil if requestLevel doesn't match!
*/
func GetApiMethods(config interface{}, requestLevel string) map[string]bool {
	response := make(map[string]bool) //set of methods allowed for that requestLevel

	config = config.(map[string]interface{})["api"]
	config, ok := config.(map[string]interface{})[requestLevel]
	if ok == false {
		return nil
	}

	switch newConfig := config.(type) {
	case map[string]interface{}:
		for i, _ := range newConfig {
			if i != "endpoint" { //do not add endpoint to response string slice
				response[i] = true
			}
		}
	}
	return response

}
