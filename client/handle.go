package client

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
)

func (c *ArpcConn) Handle(req_name string, req_body []byte) ([]byte, error) {
	var conn net.Conn
	if c.Pool == nil {
		var err error
		var tcpAddr *net.TCPAddr
		tcpAddr, err = net.ResolveTCPAddr("tcp4", c.Url)
		if err != nil {
			return nil, err
		}
		conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			return nil, err
		}
	} else {
		pool, err := (*c.Pool).Get()
		if err != nil {
			return nil, err
		}
		conn = pool.(net.Conn)
	}

	req_length := len(req_body)

	var req_data []byte = make([]byte, 0)
	req_data = append(req_data, []byte(fmt.Sprintf("%d", req_length))...)
	req_data = append(req_data, '\n')
	req_data = append(req_data, []byte(req_name)...)
	req_data = append(req_data, '\n')
	req_data = append(req_data, req_body...)

	conn.Write(req_data)

	var res_buf = make([]byte, 1024)

	var res_length = 0
	var res_body []byte

	for {
		n, err := conn.Read(res_buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if res_length == 0 {
			res := bytes.SplitN(res_buf[:n], []byte{'\n'}, 3)
			if len(res) > 2 {
				res_length, err = strconv.Atoi(string(res[0]))
				if err != nil {
					panic(err)
				}
				res_body = append(res_body, res[2]...)
			}
		} else {
			res_body = append(res_body, res_buf[:n]...)
		}

		if len(res_body) == res_length {
			break
		}
	}

	conn.Close()

	return res_body, nil
}
