package main

import (
	api "arpc-go/arpc_package"
	"arpc-go/server"
)

type test struct{}

func (t *test) GetUserV1(request *api.ApiRequestV1) (*api.ApiResponseV1, error) {
	return &api.ApiResponseV1{
		UserId: request.UserId,
	}, nil
}

func main() {
	println("Hello, world!")
	s := server.Server{}
	api.RegisterApiServer(&s, &test{})
	s.Start()
}
