package resources

import (
	"application/commands"
	"application/handle"
	"application/service"
	"domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//DI of OrderService to handle Queries
var orderService = service.CreateOrderService()

//DI of OrderHandler to handle Commands
var orderHandler = handle.CreateOrderHandler()

/*
Main method of our Go application to run the server in a port and configure the router to redirect endpoints
invocations to the specific handlers.

In Go in order to control and route request into handle, we use [http] package and function [HandleFunc]
where we pass the endpoint to bind, and the function handle to call once we receive the request.

*/
func Main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/order/", findOrderHandle)
	mux.HandleFunc("/order/create/", createOrderHandle)
	mux.HandleFunc("/order/addProduct/", addProductHandle)
	mux.HandleFunc("/order/removeProduct/", removeProductHandle)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

/*
Function that receive the request and response to handle the communication.

We extract from the request the argument [orderId] which we will use to find in the database in memory.
*/
func findOrderHandle(response http.ResponseWriter, request *http.Request) {
	log.Printf("Finding order %s!", request.URL.Path[1:])
	orderId := request.URL.Path[1:]
	order := orderService.FindOrder(domain.OrderId{orderId})
	jsonResponse, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/jsonResponse")
	response.WriteHeader(http.StatusOK)
	_, _ = response.Write(jsonResponse)
}

/*
Function to create a new Order and return the uuid of that order for the rest of the transactions regarding that order.
Having this orderId we can make the API idempotent.
*/
func createOrderHandle(response http.ResponseWriter, request *http.Request) {
	log.Println("Create new Order in System")
	createOrderCommand := commands.CreateOrder{}
	orderId := orderHandler.CreateOrder(createOrderCommand)

	jsonResponse, err := json.Marshal(orderId)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/jsonResponse")
	response.WriteHeader(http.StatusOK)
	_, _ = response.Write(jsonResponse)
}

func addProductHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("Add product in Order with id")
	fmt.Fprintf(w, "Adding product  %s!", r.URL.Path[1:])

}

func removeProductHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove product in Order with id")
	fmt.Fprintf(w, "Removing product  %s!", r.URL.Path[1:])

}
