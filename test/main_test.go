package test

import (
	api "arpc-go/arpc_package"
	arpc_net "arpc-go/net"
	"arpc-go/utils"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"
)

func TestCompile(t *testing.T) {
	input, _ := filepath.Abs("../arpc")
	output, _ := filepath.Abs("../arpc_package")
	utils.Compiles(input, output)
}

func TestServer(t *testing.T) {

	server := "127.0.0.1:9000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	st := api.ApiRequestV1{
		UserId: 1,
	}
	body, _ := st.Serialize()

	var length string = fmt.Sprintf("%d", len(body))
	name := "ApiRequestV1"

	var data []byte

	data = append(data, []byte(length)...)
	data = append(data, '\n')
	data = append(data, []byte(name)...)
	data = append(data, '\n')
	data = append(data, body...)

	conn.Write(data)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	st.Deserialize(buf[:n])

	conn.Close()
}

func TestPrintByte(t *testing.T) {
	fmt.Println("===================================")
	a := 123
	b := utils.Int64ToBytes(int64(a))
	fmt.Println(b)
	fmt.Println(utils.BytesToInt64(b))
	fmt.Println("===================================")
}

func TestClient(t *testing.T) {
	fmt.Println("===================================")
	conn := arpc_net.ArpcConn{}
	client := api.NewClient(conn)
	request := &api.ApiRequestV1{
		UserId: 1,
	}
	response, err := client.GetUserV1(request)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("response: %+v\n", response)
	fmt.Println("===================================")
}
