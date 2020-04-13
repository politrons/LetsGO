package builder

import (
	"fmt"
	"testing"
)

//##########################//
//	BUILDER PATTERN  //
//##########################//

/**
Builder pattern works just like in any other language with a good Strong type system.
The idea is to make a chain of types where you can only create a new type if you're coming
from a previous specific one, providing an order chain to go further.

Finally when we have all the component that our instance needs, we provide the [Build]
function which return the [Server] instance with all the attributes we need.
*/
func TestBuilderPattern(t *testing.T) {
	server := new(ServerFactory).
		WithHost("0.0.0.0").
		WithPort("1981").
		WithEndpoints([]string{"/politrons/foo", "/politrons/bla"}).
		WithTLs(false).
		Build()

	fmt.Println(server)
}

//   Types
//------------
/**
We will use the types to glue the chain from one option from the beginning until the last function [Build] forcing
the consumer to go through all mandatory steps to create the instance [Server]
*/

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

//First function to extend method of [ServerFactory] and extend to the next type [ServerWithHost]
func (_ ServerFactory) WithHost(host string) ServerWithHost {
	server := Server{host, "", nil, false}
	return ServerWithHost{server}
}

//Second function to extend method of [ServerWithHost] and extend to the next type [ServerWithPort]
func (srv ServerWithHost) WithPort(port string) ServerWithPort {
	server := Server{srv.server.Host, port, nil, false}
	return ServerWithPort{server}
}

//third function to extend method of [ServerWithPort] and extend to the next type [ServerWithEndpoints]
func (srv ServerWithPort) WithEndpoints(endpoints []string) ServerWithEndpoints {
	server := Server{srv.server.Host, srv.server.Port, endpoints, false}
	return ServerWithEndpoints{server}
}

//fourth function to extend method of [ServerWithEndpoints] and extend to the next type [ServerWithTls]
func (srv ServerWithEndpoints) WithTLs(tls bool) ServerWithTls {
	server := Server{srv.server.Host, srv.server.Port, srv.server.Endpoints, tls}
	return ServerWithTls{server}
}

//final function to extend method of [ServerWithTls] and return the final instance of type [Server]
func (srv ServerWithTls) Build() Server {
	return Server{srv.server.Host, srv.server.Port, srv.server.Endpoints, srv.server.Tls}
}
