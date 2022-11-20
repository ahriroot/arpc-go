package test

import (
	"fmt"
	"path/filepath"
	"testing"

	api "github.com/ahrirpc/arpc-go/arpc_package/api"
	net_ "github.com/ahrirpc/arpc-go/net"
	"github.com/ahrirpc/arpc-go/utils"
)

func TestCompile(t *testing.T) {
	input, _ := filepath.Abs("../arpc")
	output, _ := filepath.Abs("../arpc_package")
	utils.Compiles(input, output)
}

func TestClient(t *testing.T) {
	fmt.Println("===================================")
	// conn := net_.ArpcConn{}
	conn, err := net_.NewArpcConn("localhost:9000")
	if err != nil {
		fmt.Println(err)
	}

	client := api.NewClient(conn)
	request := &api.RequestV1{
		UserId: 1,
	}
	response, err := client.GetUserV1(request)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("response: %+v\n", response)
	fmt.Println("===================================")
}

func TestTcpPool(t *testing.T) {
	fmt.Println("===================================")
	conn, err := net_.NewArpcConn("localhost:9000")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("response: %+v\n", conn)
	fmt.Println("===================================")
}
