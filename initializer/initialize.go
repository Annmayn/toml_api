package initializer

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

// func evaluateURL(url string) int {
// 	urlSplit := strings.Split(url, "/")
// 	for ind, dir := range urlSplit {
// 		if dir[0] == '{' {
// 			return ind
// 		}
// 	}
// 	return 0
// }

//////////////////////////////////////////////////////////e
//Can be used to check schema and redirect (not complete)
//////////////////////////////////////////////////////////
// var rInter interface{}
// 	if _, err := toml.DecodeFile(resourceLocation, &rInter); err != nil {
// 		fmt.Println(err)
// 	} else {
// 		r := rInter.(map[string]interface{})
// 		var endPoints map[string]interface{}

// 		endPoints = r["api"].(map[string]interface{}) //endPoints : root, detail, approve

// 		for method := range endPoints {
// 			var url string
// 			if method == "endpoint" {
// 				url, ok := endPoints[method].(string)
// 				if !ok {
// 					fmt.Println("Endpoint not defined")
// 				} else {

// 				}
// 			} else {
// 				//handle all validations and what not
// 				//store schema in memory
// 				r.HandleFunc()
// 			}
// 		}
// 		// }
// 	}
////////////************************************///////////

//InitializeRoutes : route initialization
func InitializeRoutes(resourceLocation string) (map[string]interface{}, map[string][]string, error) {
	var configInterface interface{}
	kv := make(map[string][]string)

	//load configuration file in memory
	if _, err := toml.DecodeFile(resourceLocation, &configInterface); err != nil {
		//return error
		return make(map[string]interface{}), make(map[string][]string), errors.New("Can't read from file")
	}
	//make a key-value pair that maps endpoints to its schema in config
	root := configInterface.(map[string]interface{})
	var endPoints map[string]interface{}

	endPoints = root["api"].(map[string]interface{}) //endPoints : root, detail, approve

	for endPoint := range endPoints {
		//eg. m = config of api.root
		m := endPoints[endPoint]
		//get the endpoint and assert to string
		endPointURL, ok := m.(map[string]interface{})["endpoint"].(string)
		if !ok {
			//error if no endpoint is found
			fmt.Println("Endpoint not defined")
		} else {
			//add the resource location to the key-value pair
			kv[endPointURL] = []string{"api", endPoint}
		}
	}

	//return configuration and key-value pair
	return configInterface.(map[string]interface{}), kv, nil
}
