# Arpc - go

> A framework of remote procedure call.

# Quick Start

#### arpc

```bash
# ${project}/arpc/api.arpc
arpc: 1.0

package {
    python: api
    go: api
}

procedures {
    procedure GetUserV1(RequestV1): ResponseV1
}

param RequestV1 {
    UserId: integer = 1
}

param ResponseV1 {
    UserId: integer = 1
    Username: string = 2
}
```

#### Compile

```bash
arpc-go -i ./arpc -o ./api
```

#### Server

```go
package main

import (
	"github.com/ahriroot/arpc-go/server"

	"project/api"
)

type test struct{}

func (t *test) GetUserV1(request *api.RequestV1) (*api.ResponseV1, error) {
	return &api.ResponseV1{
		UserId:   request.UserId,
		Username: "test",
	}, nil
}

func main() {
	s := server.Server{}
	api.RegisterGetUserV1(&s, &test{})
	s.Start()
}
```

#### Client

```go
package main

import (
    "context"
    "fmt"

    "github.com/ahriroot/arpc-go/client"
	"github.com/ahriroot/arpc-go/net"

    "project/api"
)

func main() {
	conn := net.ArpcConn{}
	client := api.NewClient(conn)
	request := &api.RequestV1{
		UserId: 1,
	}
	response, err := client.GetUserV1(request)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("response: %+v\n", response)
}
```
