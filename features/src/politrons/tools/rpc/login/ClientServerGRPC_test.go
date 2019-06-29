package login

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func TestCommunication(t *testing.T) {
	go createServer()
	createClient()
}

func createClient() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":1981", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := NewAccountClient(conn)
	response, err := c.LoginUser(context.Background(), &LoginMessage{Username: "politrons", Password: "12345"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response User with info Name: %s", response.Name)
	log.Printf("Response User with info Age: %s", response.Age)
	log.Printf("Response User with info Sex: %s", response.Sex)
}

// main start a gRPC server and waits for connection
func createServer() {
	// create a listener on TCP port 1981
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1981))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	server := Server{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	RegisterAccountServer(grpcServer, &server)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %server", err)
	}
}

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) LoginUser(ctx context.Context, in *LoginMessage) (*UserMessage, error) {
	log.Printf("Login with username %s", in.Username)
	return &UserMessage{
		Name: "Pablo",
		Age:  "38",
		Sex:  "Male",
	}, nil
}
