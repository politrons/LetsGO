package login

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

/*
Runner of the server and client. Since the server get block listening for communications,
we run in another thread using go routines.
*/
func TestCommunication(t *testing.T) {
	go createServer()
	createClient()
}

/*
Run gRPC request against a server following the next steps:

* Create connection
* Create the client, associated with the type of service we define in the proto file. In this case using factory with format
	[New[service name]Client] providing a client with format [service name[Client]] which means AccountClient.
* Create the [LoginMessage] defined in proto file as transport message.
* Make the gRPC call using the client previously created pointing to [LoginUser] which is the function defied in the service
  in the proto file.
*/
func createClient() {
	conn, err := createConnection()
	if conn != nil {
		defer conn.Close()
	}
	accountClient := NewAccountClient(conn)
	loginMessage := &LoginMessage{Username: "politrons", Password: "12345"}
	response, err := accountClient.LoginUser(context.Background(), loginMessage)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response User with info:")
	log.Printf("Name: %s", response.Name)
	log.Printf("Age: %s", response.Age)
	log.Printf("Sex: %s", response.Sex)
}

/*
We create a new connection using [grpc.Dial] passing the port, and Security strategy, in this case
[WithInsecure]
*/
func createConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(":1981", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn, err
}

/*
Start a gRPC server following next steps:

* Create a listener on TCP port 1981
* Create a server instance
* Create a gRPC server instance
* Attach the Service to the server using the class generate with the proto file. Format[Register[service]Server]
  Now all the communication from client that reach the ip/port of server it will be redirect to all extend methods,
  of Server[Account]
* Start the server using the [net.Listen] specifying the network protocol and port
*/
func createServer() {
	listener := createServerListener()
	server := Server{}
	grpcServer := grpc.NewServer()
	RegisterAccountServer(grpcServer, &server)
	startServer(grpcServer, listener)
}

/*
We use this function to create a Listener in a network protocol and port.
*/
func createServerListener() net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1981))
	if err != nil {
		log.Fatalf("Error to listen: %v", err)
	}
	return lis
}

/*
Using the instance of the Server and the listener we run the server and we control the error in
the initialization.
*/
func startServer(grpcServer *grpc.Server, listener net.Listener) {
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error running server: %server", err)
	}
}

// Server represents the gRPC server
type Server struct {
}

// Implementation of the service [Account]
func (s *Server) LoginUser(ctx context.Context, in *LoginMessage) (*UserMessage, error) {
	log.Printf("Login with username %s", in.Username)
	return &UserMessage{
		Name: "Pablo",
		Age:  "38",
		Sex:  "Male",
	}, nil
}
