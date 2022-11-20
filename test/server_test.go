package test

import (
	"testing"

	"github.com/ahrirpc/arpc-go/arpc_package/api"
	"github.com/ahrirpc/arpc-go/arpc_package/test"
	"github.com/ahrirpc/arpc-go/server"
)

type api_client struct{}

func (t *api_client) GetUserV1(request *api.RequestV1) (*api.ResponseV1, error) {
	return &api.ResponseV1{
		UserId:   request.UserId,
		Username: "test",
	}, nil
}

func (t *api_client) PostUserV1(request *api.ResponseV1) (*api.RequestV1, error) {
	return &api.RequestV1{
		UserId: request.UserId,
	}, nil
}

type test_client struct{}

func (t *test_client) GetUserV1(request *test.RequestV1) (*test.ResponseV1, error) {
	return &test.ResponseV1{
		UserId:   request.UserId,
		Username: "test",
	}, nil
}

func (t *test_client) PostUserV1(request *test.ResponseV1) (*test.RequestV1, error) {
	return &test.RequestV1{
		UserId: request.UserId,
	}, nil
}

func TestServer(t *testing.T) {
	s := server.Server{
		Host: "localhost",
		Port: "9000",
	}
	api.Register(&s, &api_client{})
	test.Register(&s, &test_client{})
	s.Start()
}
