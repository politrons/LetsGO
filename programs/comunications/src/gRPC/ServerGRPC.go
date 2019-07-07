package gRPC

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunGRPCServer() {
	println("Running gRPC server in port 3000")
	listener := createServerListener()
	server := Server{}
	grpcServer := grpc.NewServer()
	RegisterMessageManagementServer(grpcServer, &server)
	startServer(grpcServer, listener)
}

/*
We use this function to create a Listener in a network protocol and port.
*/
func createServerListener() net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
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

/*
We receive the message from the client and we just attach an extra text in the message and return to the client.

If you want to understand how works all the pieces of a gRPC server go to the section [gRPC] of this project
*/
func (server *Server) ProcessMessage(ctx context.Context, userMessage *UserMessage) (*UserMessage, error) {
	updatedMessage := userMessage.Message + " and gRPC server"
	fmt.Printf("gRPC Server:%s \n", updatedMessage)
	message := &UserMessage{Message: updatedMessage}
	return message, nil
}
