package functionalOptions

//https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
import (
	"fmt"
	"testing"
)

//######################//
//	FUNCTIONAL OPTIONS  //
//######################//

/**
In this example since the functional options patter allow 0 to N arguments we can create the server
using [CreateServer] without have to specify any configuration.
*/
func TestOptionalFuncNoArguments(t *testing.T) {
	server, err := ServerConfig{}.CreateServer()
	if err != nil {
		panic("Error creating server")
	}
	fmt.Println(server.ToString())
}

/**
But what about if we want to pass extra configuration?. Shall we create a function per argument?,
put all of them mandatory and then force to pass nulls? yikes!.

What about use optional function pattern, that as the name already explain is just pass multiple functions 0 to N
which contains as entry argument the server itself, and each function know how to configure the server instance.

For this example, we create three functions that internally configure the ServerConfig type, to configure
the instance.
Then we pass this three functions to the function [CreateServer] which internally invoke all of them passing
the type class instance configured as extended method.
*/
func TestOptionalFuncWithArguments(t *testing.T) {
	configHostAndPortFunc := func(server ServerConfig) ServerConfig {
		server.Host = Host{"0.0.0.0"}
		server.Port = Port{"1981"}
		return server
	}

	configEndpointFunc := func(server ServerConfig) ServerConfig {
		server.Endpoints = []Endpoint{{"/politrons/foo"}, {"/politrons/bla"}}
		return server
	}

	configTlsFunc := func(server ServerConfig) ServerConfig {
		server.Tls = Tls{true}
		return server
	}

	server, err := ServerConfig{}.CreateServer(configHostAndPortFunc, configEndpointFunc, configTlsFunc)
	if err != nil {
		panic("Error creating server")
	}
	fmt.Println(server.ToString())
}

//   Types
//------------

//Some optional types needed to configure the server
type Host struct{ value string }
type Port struct{ value string }
type Endpoint struct{ value string }
type Tls struct{ value bool }

/**
Type [ServerConfig] contains some attributes that are optional, so no need to be passed
to the creation of the server function.
*/
type ServerConfig struct {
	Host      Host
	Port      Port
	Endpoints []Endpoint
	Tls       Tls
}

//   Contract & Implementation
//------------------------------

/**
Interface to declare the contract of how the Server must be created, and the type that must return.
*/
type MyServer interface {
	CreateServer(options ...func(ServerConfig) ServerConfig) (MyServer, error)

	ToString() (Host, Port, []Endpoint, Tls)
}

/**
Implementation of the interface where as usual we have the implementation type [ServerConfig] and also
we define argument the N possible option functions, where receiving those functions we apply over the serverConfig,
helping to create the instance of the server with the dynamic configuration.
*/
func (srcConf ServerConfig) CreateServer(options ...func(ServerConfig) ServerConfig) (MyServer, error) {
	var newSrvConf = srcConf
	for _, opt := range options {
		newSrvConf = opt(newSrvConf) //Here we apply the function passing the server
	}
	//We return always error nil, but you should expect a more realistic Server creation with side-effects in real world
	return newSrvConf, nil
}

//Util function to return a tuple of all elements of the Server
func (srcConf ServerConfig) ToString() (Host, Port, []Endpoint, Tls) {
	return srcConf.Host, srcConf.Port, srcConf.Endpoints, srcConf.Tls
}