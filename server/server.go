package server

import (
	api "arpc-go/arpc_package"
	"bytes"
	"fmt"
	"io"
	"net"
)

func Start() {
	// socket listen 127.0.0.1:9000
	var listener, err = net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// accept client connection
	for {
		var conn, err = listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("handleConnection")

	var buf = make([]byte, 1024)

	var length = 0
	var name = ""
	var body []byte

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if length == 0 {
			// split 2 次
			res := bytes.SplitN(buf[:n], []byte{'\n'}, 2)
			fmt.Printf("res: %+v \n", res)
			if len(res) > 2 {
				length = int(buf[0])
				name = string(buf[1:n])
				body = append(body, res[2]...)
			}
		} else {
			body = append(body, buf[:n]...)
		}

		if len(body) == length {
			break
		}
	}
	fmt.Println("receive data:", length, name, body)

	st := api.ApiRequestV1{
		UserId: 1,
	}
	st.Deserialize(body)
	st.UserId = 2
	res_length := len(body)
	res_name := "ApiRequestV1"
	data := append([]byte{byte(res_length)}, res_name...)
	data = append(data, body...)
	conn.Write(data)
	fmt.Println("receive struct:", st)

	conn.Close()
}
