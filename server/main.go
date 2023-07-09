package main

import (
	"fmt"
	"log"
	"net"

	"github.com/faba/server/service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to listen").Error())
	}

	server := grpc.NewServer()
	service.Apply(server)

	fmt.Println(fmt.Sprintf("start grpc server port: %s", "8080"))

	if err := server.Serve(lis); err != nil {
		log.Fatal(fmt.Sprintf("failed to serve: %v", err))
	}
}
