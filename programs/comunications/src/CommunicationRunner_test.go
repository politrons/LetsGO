package src

import (
	. "gRPC"
	. "rest"
	"testing"
)

/*
We create a flow between different communication technologies and how you can publish and consume those.

We receive a Rest call, we publish an event in Kafka, then Kafka consumer receive the event and invoke a gRPC call to a service,
this service receive the call and return to the client, then the client publish the result into another topic that the rest server is
subscribed, and the return the request.
*/
func TestRun(t *testing.T) {
	go RunKafkaServer() // Kafka consumer that once receive an event it will make a gRPC call using the gRPC client
	go RunGRPCServer()  //gRPC server that when receive a call from gRPC client it will publish a Kafka event that Rest server will receive.
	HttpServer()        //Rest server that when receive Rest call it will subscribe to Kafka topic and publish an event in Kafka
}
