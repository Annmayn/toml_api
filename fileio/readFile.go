package fileio

import (
	"context"
	"encoding/json"
	"toml_api/proto"

	"google.golang.org/grpc"
)

func ReadFromFile() interface{} {
	// var b interface{}
	// jsonFile, _ := os.Open("database.json")
	// defer jsonFile.Close()
	// byteFile, _ := ioutil.ReadAll(jsonFile)
	// _ = json.Unmarshal(byteFile, &b)
	// return b

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic("here")
	}
	client := proto.NewAddServiceClient(conn)
	defer conn.Close()

	req := &proto.FileName{FileName: "database.json"}
	if response, err := client.ReadFile(context.Background(), req); err == nil {
		var tmp map[string]interface{}
		json.Unmarshal(response.Content, &tmp)
		return tmp
	}

	panic("hi") //remove this: only for debugging

	return map[string]interface{}{"error": "couldn't run rpc call"}
}
