package login

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	createServer()
}

func TestClient(t *testing.T) {
	createClient()
}

func createClient() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":1981", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
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
	RegisterPingServer(grpcServer, &server)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %server", err)
	}
}

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}
