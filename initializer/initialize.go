package initializer

import (
	"errors"
	"strings"

	"github.com/BurntSushi/toml"
)

func evaluateURL(url string) int {
	urlSplit := strings.Split(url, "/")
	for ind, dir := range urlSplit {
		if dir[0] == '{' {
			return ind
		}
	}
	return 0
}

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
func InitializeRoutes(resourceLocation string) (map[string]interface{}, error) {
	var configInterface interface{}
	if _, err := toml.DecodeFile(resourceLocation, &configInterface); err != nil {
		return make(map[string]interface{}), errors.New("Can't read from file")
	}
	return configInterface.(map[string]interface{}), nil
}
