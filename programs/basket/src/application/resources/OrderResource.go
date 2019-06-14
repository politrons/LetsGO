package resources

import (
	"application/commands"
	"application/handle"
	"application/service"
	"domain"
	"encoding/json"
	"fmt"
	"infrastructure"
	"log"
	"net/http"
	"strings"
)

var repository = infrastructure.CreateOrderRepository()

//DI of OrderService to handle Queries
var orderService = service.CreateOrderService(repository)

//DI of OrderHandler to handle Commands
var orderHandler = handle.CreateOrderHandler(repository)

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
	orderId := strings.Split(request.URL.Path, "/")[2]
	log.Printf("Finding order %s!", orderId)
	order := orderService.FindOrder(domain.OrderId{Id: orderId})
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

func addProductHandle(response http.ResponseWriter, request *http.Request) {
	orderId := strings.Split(request.URL.Path, "/")[3]
	log.Printf("Add product for order %s!", orderId)
	decoder := json.NewDecoder(request.Body)
	addProductCommand := commands.AddProduct{}
	err := decoder.Decode(&addProductCommand)
	if err != nil {
		panic(err)
		println("Error decoding add product command")
	}
	order := orderHandler.UpdateOrder(domain.OrderId{Id: orderId}, addProductCommand)
	jsonResponse, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/jsonResponse")
	response.WriteHeader(http.StatusOK)
	_, _ = response.Write(jsonResponse)

}

func removeProductHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove product in Order with id")
	fmt.Fprintf(w, "Removing product  %s!", r.URL.Path[1:])

}
