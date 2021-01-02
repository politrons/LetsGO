package rest

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	. "kafka"
	"log"
	"net/http"
	"strings"
	"time"
)

/*
HttpServer method of our Go application to run the server in a port and configure the router to redirect endpoints
invocations to the specific handlers.

In Go in order to control and route request into handle, we use [http] package and function [HandleFunc]
where we pass the endpoint to bind, and the function handle to call once we receive the request.

*/
func HttpServer() {
	println("Running Rest server on port 4000")
	mux := http.NewServeMux()
	mux.HandleFunc("/communication/restKafkaGRPC/", processRequest)
	log.Fatal(http.ListenAndServe(":4000", mux))
}

/*
Function to receive  a rest call, publish to Kafka the message received and subscribe into Kafka to receive the response.
Once the event has been processed by another two services(Kafka, gRPC) we subscribe into a new [Kafka] topic where
we receive the final message.

We use context to set the timeout of the request in 5 seconds.
*/
func processRequest(response http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	message := "hello world from rest"
	fmt.Printf("REST request: %s \n", message)
	broker := Broker{Value: "localhost:9092"}
	publishTopic := Topic{Value: "CommunicationTopic"}
	consumeTopic := Topic{Value: "CommunicationRestTopic"}

	channel := make(chan string)
	go SubscribeConsumer(broker, consumeTopic, func(str string) {
		fmt.Printf("REST response: %s \n", str)
		channel <- strings.ToUpper(str)
	})
	time.Sleep(1 * time.Second) //Time enough to subscribe and avoid RC(better improve with channels)
	PublishEvents(
		broker,
		publishTopic,
		"myKey", message)

	select {
	case messageResponse := <-channel:
		fmt.Printf("############################################################################################")
		fmt.Printf("End of transaction with Message:")
		fmt.Printf("%s", messageResponse)
		fmt.Printf("############################################################################################")
		writeResponse(response, messageResponse)
	case <-ctx.Done():
		writeResponse(response, "Error:Timeout request")
	}

}

/*
Function to marshal the generic type [interface{}] into json response, then if everything is fine
we return a 200 status code with the response, otherwise a 500 error response
*/
func writeResponse(response http.ResponseWriter, t interface{}) {
	jsonResponse, err := json.Marshal(t)
	if err != nil {
		writeErrorResponse(response, err)
	} else {
		writeSuccessResponse(response, jsonResponse)
	}
}

func writeSuccessResponse(response http.ResponseWriter, jsonResponse []byte) {
	response.Header().Set("Content-Type", "application/jsonResponse")
	response.WriteHeader(http.StatusOK)
	_, _ = response.Write(jsonResponse)
}

func writeErrorResponse(response http.ResponseWriter, err error) {
	response.Header().Set("Content-Type", "application/jsonResponse")
	response.WriteHeader(http.StatusServiceUnavailable)
	errorResponse, _ := json.Marshal("Error in request since " + err.Error())
	_, _ = response.Write(errorResponse)
}
