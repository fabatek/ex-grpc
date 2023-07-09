package dao

import (
	"github.com/faba/server/pkg/adapter"
	"google.golang.org/grpc"
)

// FabaClient is client wrapper of FabaServiceClient.
type FabaClient struct {
	*adapter.FabaServiceClient
}

// NewAdaptedClient creates a new ApiClient.
func NewAdaptedClient(conn *grpc.ClientConn) (*FabaClient, error) {

	// create client
	newClient, err := adapter.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &FabaClient{
		&newClient,
	}, nil
}
