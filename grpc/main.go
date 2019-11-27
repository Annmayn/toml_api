package main

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"os"
	"toml_api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var jsonFile *os.File

func (s *server) ReadFile(ctx context.Context, request *proto.FileName) (*proto.FileContent, error) {
	// fileName := request.GetFileName()
	// jsonFile, err := os.Open(fileName)
	// defer jsonFile.Close()

	// if err != nil {
	// 	panic(err)
	// }
	byteFile, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	// data := string(byteFile)
	// var data map[string]interface{}
	var data io.Reader
	_ = json.Unmarshal(byteFile, &data)
	return &proto.FileContent{Content: byteFile}, nil
}

func init() {
	fileName := "database.json"
	jsonFile, _ = os.Open(fileName)
}

func main() {
	defer jsonFile.Close()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
