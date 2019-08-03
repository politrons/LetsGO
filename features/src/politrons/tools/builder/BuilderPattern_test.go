package builder

import (
	"fmt"
	"testing"
)

/**
Builder pattern works just like in any other language with a good Strong type system.
The idea is to make a chain of types where you can only create a new type if you're coming
from a previous specific one, providing an order chain to go further.

Finally when we have all the component that our instance needs, we provide the Build
function which return the [Server] instance with all the attributes that we need.
*/
func TestBuilderPattern(t *testing.T) {
	server := ServerFactory{}.
		WithHost("0.0.0.0").
		WithPort("1981").
		WithEndpoints([]string{"/politrons/foo", "/politrons/bla"}).
		Build(false)
	fmt.Println(server)
}

type ServerFactory struct {
	Host ServerWithHost
}

type ServerWithHost struct {
	Host string
}

type ServerWithPort struct {
	Host string
	Port string
}

type ServerWithEndpoints struct {
	Host      string
	Port      string
	Endpoints []string
}

type Server struct {
	Host      string
	Port      string
	Endpoints []string
	Tls       bool
}

func (srv ServerFactory) WithHost(host string) ServerWithHost {
	return ServerWithHost{host}
}

func (srv ServerWithHost) WithPort(port string) ServerWithPort {
	return ServerWithPort{Host: srv.Host, Port: port}
}

func (srv ServerWithPort) WithEndpoints(endpoints []string) ServerWithEndpoints {
	return ServerWithEndpoints{Host: srv.Host, Port: srv.Port, Endpoints: endpoints}
}

func (srv ServerWithEndpoints) Build(tls bool) Server {
	return Server{Host: srv.Host, Port: srv.Port, Endpoints: srv.Endpoints, Tls: tls}
}
