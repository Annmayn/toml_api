package getresource

import (
	"fmt"
	"log"
)

/*
	Receives config, keys as args
	Returns interface{}

	Need to perform typecheck for map[string]interface{}, []map[string]interface{}, string, int64
*/
func GetResource(config interface{}, keys ...interface{})(interface{}){
	//iterates over keys and continuously perform assertion
	for _,p:=range keys{
		switch t:=p.(type){
		case string:
			config=config.(map[string]interface{})[t]
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
