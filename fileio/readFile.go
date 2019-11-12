package fileio

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadFromFile() interface{} {
	var b interface{}
	jsonFile, _ := os.Open("database.json")
	defer jsonFile.Close()
	byteFile, _ := ioutil.ReadAll(jsonFile)
	_ = json.Unmarshal(byteFile, &b)
	return b
}
