package rest

import (
	"fmt"
	"gRPC"
	. "kafka"
)

/*
Kafka service that subscribe into a topic and once it receive the event make a gRPC call to another gRPC server.

In the subscription of the Topic we also pass the function to execute once the event it's received, making the consumer
completely generic for other topics and actions
*/
func RunKafkaServer() {
	topic := "CommunicationTopic"
	broker := "localhost:9092"
	fmt.Printf("Running Kafka service, Subscribing to broker %s and topic:%s \n", broker, topic)
	SubscribeConsumer(
		Broker{Value: broker},
		Topic{Value: topic}, func(str string) {
			gRPC.MakeGRPCRequest(str)
		})
}
