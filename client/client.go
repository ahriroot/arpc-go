package client

import (
	"net"
	"time"
)

type ArpcConn struct {
	Url  string
	Pool *Pool
}

func NewArpcConn(url string) (*ArpcConn, error) {
	return &ArpcConn{
		Url:  url,
		Pool: nil,
	}, nil
}

func NewArpcConnPool(url string) (*ArpcConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", url)
	if err != nil {
		return nil, err
	}

	factory := func() (interface{}, error) { return net.DialTCP("tcp", nil, tcpAddr) }
	//close 关闭连接的方法
	close := func(v interface{}) error { return v.(net.Conn).Close() }
	//创建一个连接池： 初始化2，最大连接5，空闲连接数是4
	poolConfig := &Config{
		InitialCap: 2,
		MaxIdle:    4,
		MaxCap:     5,
		Factory:    factory,
		Close:      close,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	pool, err := NewChannelPool(poolConfig)
	if err != nil {
		return nil, err
	}

	return &ArpcConn{
		Pool: &pool,
	}, nil
}
