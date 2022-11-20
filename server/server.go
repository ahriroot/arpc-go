package server

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"

	net_ "github.com/ahrirpc/arpc-go/client"
)

type Server struct {
	handles map[string]interface{}
	Host    string
	Port    string
}

func (s *Server) Register(name string, f interface{}) {
	if s.handles == nil {
		s.handles = make(map[string]interface{})
	}
	s.handles[name] = f
}

func (s *Server) Start() error {
	if s.Host == "" {
		s.Host = "127.0.0.1"
	}
	if s.Port == "" {
		s.Port = "9000"
	}
	listener, err := net.Listen("tcp", s.Host+":"+s.Port)
	if err != nil {
		return err
	}
	defer listener.Close()
	for {
		var conn, err = listener.Accept()
		if err != nil {
			return err
		}
		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {

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
		}

		if length == 0 {
			res := bytes.SplitN(buf[:n], []byte{'\n'}, 3)
			if len(res) > 2 {
				length, err = strconv.Atoi(string(res[0]))
				if err != nil {
					panic(err)
				}
				name = string(res[1])
				fmt.Println(name)
				body = append(body, res[2]...)
			}
		} else {
			body = append(body, buf[:n]...)
		}

		if len(body) == length {
			break
		}
	}

	function, has := s.handles[name]
	if !has {
		panic("function not found")
	}
	res, err := function.(func([]byte, net_.ArpcConn) ([]byte, error))(body, net_.ArpcConn{})
	if err != nil {
		panic(err)
	}
	var data []byte
	res_length := len(res)
	data = append(data, []byte(fmt.Sprintf("%d", res_length))...)
	data = append(data, '\n')
	data = append(data, []byte(name)...)
	data = append(data, '\n')
	data = append(data, res...)
	conn.Write(data)

	conn.Close()
}
