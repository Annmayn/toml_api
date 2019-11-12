package fileio

import (
	"encoding/json"
	"io/ioutil"
)

func WriteToFile(data interface{}) {
	b, _ := json.Marshal(data)
	ioutil.WriteFile("database.json", b, 777)
}
