package builder

import (
	"fmt"
	"testing"
)

/**
Builder pattern works just like in any other language with a good Strong type system.
The idea is to make a chain of types where you can only create a new type if you're coming
from a previous specific one, providing an order chain to go further.

Finally when we have all the component that our instance needs, we provide the WithTLs
function which return the [Server] instance with all the attributes that we need.
*/
func TestBuilderPattern(t *testing.T) {
	server := ServerFactory{}.
		WithHost("0.0.0.0").
		WithPort("1981").
		WithEndpoints([]string{"/politrons/foo", "/politrons/bla"}).
		WithTLs(false).
		Build()

	fmt.Println(server)
}

type ServerFactory struct {
	server Server
}

type ServerWithHost struct {
	server Server
}

type ServerWithPort struct {
	server Server
}

type ServerWithEndpoints struct {
	server Server
}

type ServerWithTls struct {
	server Server
}

type Server struct {
	Host      string
	Port      string
	Endpoints []string
	Tls       bool
}

func (_ ServerFactory) WithHost(host string) ServerWithHost {
	server := Server{host, "", nil, false}
	return ServerWithHost{server}
}

func (srv ServerWithHost) WithPort(port string) ServerWithPort {
	server := Server{srv.server.Host, port, nil, false}
	return ServerWithPort{server}
}

func (srv ServerWithPort) WithEndpoints(endpoints []string) ServerWithEndpoints {
	server := Server{srv.server.Host, srv.server.Port, endpoints, false}
	return ServerWithEndpoints{server}
}

func (srv ServerWithEndpoints) WithTLs(tls bool) ServerWithTls {
	server := Server{srv.server.Host, srv.server.Port, srv.server.Endpoints, tls}
	return ServerWithTls{server}
}

func (srv ServerWithTls) Build() Server {
	return Server{srv.server.Host, srv.server.Port, srv.server.Endpoints, srv.server.Tls}
}
