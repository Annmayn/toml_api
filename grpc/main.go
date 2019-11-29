package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"toml_api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	/////
	byteFile, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	return &proto.FileContent{Content: byteFile}, nil
}

func (s *server) WriteToFile(ctx context.Context, request *proto.FileData) (*proto.EmptyResponse, error) {
	data := request.GetContent()
	// b, _ := json.Unmarshal(data)
	ioutil.WriteFile("database.json", data, 777)
	return &proto.EmptyResponse{}, nil
}

func main() {
	fileName := "database.json"
	jsonFile, _ = os.Open(fileName)

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	crt := "cert/server.crt"
	key := "cert/server.key"
	creds, e := credentials.NewServerTLSFromFile(crt, key)
	if e != nil {
		fmt.Println("Couldn't read files")
		return
	}

	fmt.Println("Server started...")

	srv := grpc.NewServer(grpc.Creds(creds))
	// srv := grpc.NewServer()

	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

	// we don't want to close the "jsonFile" prematurely; only after the grpc server is terminated
	jsonFile.Close()
}
