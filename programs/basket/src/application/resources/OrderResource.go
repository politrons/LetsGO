package resources

import (
	"application/commands"
	"application/handle"
	"application/service"
	. "domain"
	"encoding/json"
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
HttpServer method of our Go application to run the server in a port and configure the router to redirect endpoints
invocations to the specific handlers.

In Go in order to control and route request into handle, we use [http] package and function [HandleFunc]
where we pass the endpoint to bind, and the function handle to call once we receive the request.

*/
func HttpServer() {
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
	order := <-orderService.FindOrder(OrderId{Id: orderId})
	writeResponse(response, order)
}

/*
Function to create a new Order and return the uuid of that order for the rest of the transactions regarding that order.
Having this orderId we can make the API idempotent.
*/
func createOrderHandle(response http.ResponseWriter, request *http.Request) {
	log.Println("Create new Order in System")
	createOrderCommand := commands.CreateOrder{}
	orderId := <-orderHandler.CreateOrder(createOrderCommand)
	writeResponse(response, orderId)
}

/*
Function to get the orderId from uri param, and the body information of the product to add into the order.
We use [decoder.Decode] function to decode the body into the Type defined in our service.
In order to make it work the name of the attributes of your Type must be equal to the once defined in the json.
*/
func addProductHandle(response http.ResponseWriter, request *http.Request) {
	orderId := strings.Split(request.URL.Path, "/")[3]
	log.Printf("Add product for order %s!", orderId)
	decoder := json.NewDecoder(request.Body)
	addProductCommand := commands.AddProduct{}
	err := decoder.Decode(&addProductCommand)
	if err != nil {
		writeErrorResponse(response, err)
	}
	order := <-orderHandler.AddProductInOrder(OrderId{Id: orderId}, addProductCommand)
	writeResponse(response, order)
}

/*
Function to get the orderId from uri param, and the body information of the productId to remove it from the order.
We use [decoder.Decode] function to decode the body into the Type defined in our service.
In order to make it work the name of the attributes of your Type must be equal to the once defined in the json.
*/
func removeProductHandle(response http.ResponseWriter, request *http.Request) {
	orderId := strings.Split(request.URL.Path, "/")[3]
	log.Printf("Remove product for order %s!", orderId)
	decoder := json.NewDecoder(request.Body)
	removeProductCommand := commands.RemoveProduct{}
	err := decoder.Decode(&removeProductCommand)
	if err != nil {
		writeErrorResponse(response, err)
	}
	order := <-orderHandler.RemoveProductInOrder(OrderId{Id: orderId}, removeProductCommand)
	writeResponse(response, order)
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
