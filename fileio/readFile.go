package fileio

import (
	"context"
	"encoding/json"
	"fmt"
	"toml_api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func ReadFromFile() interface{} {
	// var b interface{}
	// jsonFile, _ := os.Open("database.json")
	// defer jsonFile.Close()
	// byteFile, _ := ioutil.ReadAll(jsonFile)
	// _ = json.Unmarshal(byteFile, &b)
	// return b

	crt := "cert/server.crt"
	creds, e := credentials.NewClientTLSFromFile(crt, "")
	if e != nil {
		fmt.Println("Couldn't read crt file")
		var tmpRes interface{}
		tmpRes = map[string]interface{}{
			"name": "error",
		}
		return tmpRes
	}

	conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(creds))
	// conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Couldn't connect\n\n\n")
		// panic("here")
		var a interface{}
		a = map[string]interface{}{
			"name": "error",
		}
		return a
	}
	client := proto.NewAddServiceClient(conn)
	defer conn.Close()
	fmt.Println("1")

	req := &proto.FileName{FileName: "database.json"}
	if response, err := client.ReadFile(context.Background(), req); err == nil {
		var tmp map[string]interface{}
		json.Unmarshal(response.Content, &tmp)
		return tmp
	}
	fmt.Println(2)

	// panic("readFile panic") //remove this: only for debugging

	var a interface{} = map[string]interface{}{
		"name": "couldn't run rpc call",
	}
	fmt.Println(a)
	return a
}
