package gRPC

import (
	"context"
	"google.golang.org/grpc"
	. "kafka"
	"log"
)

/*
Function to make a request into gRPC server and print the final result
*/
func MakeGRPCRequest(message string) {
	updatedMessage := message + " and gRPC client"
	log.Printf("gRPC Client:%s \n", updatedMessage)
	conn, err := createConnection()
	if conn != nil {
		defer conn.Close()
	}
	client := NewMessageManagementClient(conn)
	userMessage := &UserMessage{Message: updatedMessage}
	response, err := client.ProcessMessage(context.Background(), userMessage)
	if err != nil {
		log.Fatalf("Error in Login process: %s", err)
	}
	broker := Broker{Value: "localhost:9092"}
	consumeTopic := Topic{Value: "CommunicationRestTopic"}

	PublishEvents(broker, consumeTopic, "myKey", response.Message)
}

/*
We create a new connection using [grpc.Dial] passing the port, and Security strategy, in this case
[WithInsecure]
*/
func createConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn, err
}
