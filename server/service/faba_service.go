package service

import (
	"context"

	"github.com/faba/server/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type FabaService struct {
	proto.UnimplementedFabaServiceServer
}

func NewFabaService() *FabaService {
	return &FabaService{}
}

func Apply(server *grpc.Server) {
	proto.RegisterFabaServiceServer(
		server,
		NewFabaService(),
	)
}

// RegisterEmployee register employee
func (fb *FabaService) RegisterEmployee(ctx context.Context, in *proto.RegisterEmployeeRequest) (*proto.RegisterEmployeeResponse, error) {
	protoRegisterEmployeeResponse := proto.RegisterEmployeeResponse{
		Uuid:    uuid.New().String(),
		Message: "create employee successfully",
	}
	return &protoRegisterEmployeeResponse, nil
}
