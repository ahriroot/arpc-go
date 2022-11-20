//119189-110-11 20:00:00

package test

import (
	"encoding/json"
	
	"github.com/ahrirpc/arpc-go/net"
	"github.com/ahrirpc/arpc-go/server"
)

type RequestV1 struct {
    UserId int `json:"user_id"`
}

func (b *RequestV1) New(user_id int) {
	b.UserId = user_id
}

func (b *RequestV1) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

func (b *RequestV1) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}

type ResponseV1 struct {
    UserId   int    `json:"user_id"`
    Username string `json:"username"`
}

func (b *ResponseV1) New(user_id int, username string) {
	b.UserId = user_id
    b.Username = username
}

func (b *ResponseV1) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

func (b *ResponseV1) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}

type client struct {
	conn *net.ArpcConn
}

type Client interface {
    GetUserV1(*RequestV1) (*ResponseV1, error)
    PostUserV1(*ResponseV1) (*RequestV1, error)
}

func (c *client) GetUserV1(request *RequestV1) (*ResponseV1, error) {
	req_bytes, err := request.Serialize()
	if err != nil {
		return nil, err
	}
	data, err := net.Handle("baiCMR.1", req_bytes, c.conn)
	if err != nil {
		return nil, err
	}
	response := &ResponseV1{}
	err = response.Deserialize(data)
	if err != nil {
		return nil, err
	}
	return response, nil
}


func (c *client) PostUserV1(request *ResponseV1) (*RequestV1, error) {
	req_bytes, err := request.Serialize()
	if err != nil {
		return nil, err
	}
	data, err := net.Handle("baiCMR.2", req_bytes, c.conn)
	if err != nil {
		return nil, err
	}
	response := &RequestV1{}
	err = response.Deserialize(data)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func Register(s *server.Server, i Client) {
    s.Register("baiCMR.1", func(request []byte, _ net.ArpcConn) ([]byte, error) {
		req := &RequestV1{}
		err := req.Deserialize(request)
		if err != nil {
			return nil, err
		}
		response, err := i.GetUserV1(req)
		if err != nil {
			return nil, err
		}
		return response.Serialize()
	})
    s.Register("baiCMR.2", func(request []byte, _ net.ArpcConn) ([]byte, error) {
		req := &ResponseV1{}
		err := req.Deserialize(request)
		if err != nil {
			return nil, err
		}
		response, err := i.PostUserV1(req)
		if err != nil {
			return nil, err
		}
		return response.Serialize()
	})
}

func NewClient(c *net.ArpcConn) Client {
	return &client{c}
}
