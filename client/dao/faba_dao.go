package dao

import (
	"context"

	"github.com/faba/client/domain/model"
	rep "github.com/faba/client/domain/repository"
	"github.com/faba/server/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Faba is a dao to call rpc to Faba server.
type Faba struct {
	// endpoint client object
	client     *FabaClient
	serverHost string
}

// NewFabaClient create a new Faba client.
func NewFabaClient(serverHost string) (rep.FabaRepository, error) {
	return &Faba{
		&FabaClient{},
		serverHost,
	}, nil
}

func (fb *Faba) initClient(ctx context.Context, conn *grpc.ClientConn) (context.Context, error) {
	adaptedClient, err := NewAdaptedClient(conn)
	if err != nil {
		return ctx, err
	}
	fb.client = adaptedClient

	return ctx, nil
}

// CreateEmployee create employee
func (fb *Faba) CreateEmployee(ctx context.Context, employeeRequestParam *model.EmployeeRequestParam) (model.EmployeeResponseData, error) {
	// start connect rpc
	conn, err := grpc.Dial(fb.serverHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return model.EmployeeResponseData{}, err
	}
	defer conn.Close()

	// init rpc client
	ctx, err = fb.initClient(ctx, conn)
	if err != nil {
		return model.EmployeeResponseData{}, err
	}

	protoEmployeeRequestParam := convertParamToProtoRegisterEmployeeRequest(employeeRequestParam)
	responseData, err := fb.client.RegisterEmployee(ctx, protoEmployeeRequestParam)
	if err != nil {
		return model.EmployeeResponseData{}, err
	}

	registerEmployeeResponse := convertProtoRegisterEmployeeResponseToParam(responseData)
	return registerEmployeeResponse, nil
}

// convertParamToProtoRegisterEmployeeRequest convert model.EmployeeRequestParam to proto.RegisterEmployeeRequest
func convertParamToProtoRegisterEmployeeRequest(employeeRequestParam *model.EmployeeRequestParam) *proto.RegisterEmployeeRequest {
	protoRegisterEmployeeRequest := proto.RegisterEmployeeRequest{
		Name:    &wrappers.StringValue{Value: employeeRequestParam.Name},
		Address: &wrappers.StringValue{Value: employeeRequestParam.Address},
		Phone:   &wrappers.StringValue{Value: employeeRequestParam.Phone},
		BirthDateTime: &proto.TimestampValue{
			Value: &timestamp.Timestamp{Seconds: employeeRequestParam.BirthOfDate.Unix(), Nanos: int32(employeeRequestParam.BirthOfDate.Nanosecond())},
		},
	}

	return &protoRegisterEmployeeRequest
}

// convertProtoRegisterEmployeeResponseToParam convert from proto.RegisterEmployeeResponse to model.EmployeeResponseData
func convertProtoRegisterEmployeeResponseToParam(response *proto.RegisterEmployeeResponse) model.EmployeeResponseData {
	employeeResponseData := model.EmployeeResponseData{
		UUID:    response.Uuid,
		Message: response.Message,
	}

	return employeeResponseData
}
