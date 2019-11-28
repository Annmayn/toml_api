package fileio

import (
	"context"
	"encoding/json"
	"toml_api/proto"

	"google.golang.org/grpc"
)

func WriteToFile(data interface{}) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic("here")
	}
	client := proto.NewAddServiceClient(conn)
	defer conn.Close()

	byteData, _ := json.Marshal(data)
	req := &proto.FileData{Content: byteData}
	if _, err := client.WriteToFile(context.Background(), req); err == nil {
		// tmp := make(map[string]interface{})
		return
	}

	panic("hi") //remove this: only for debugging

	// return map[string]interface{}{"error": "couldn't run rpc call"}
}
