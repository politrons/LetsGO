package domain

import (
	uuid "github.com/satori/go.uuid"
)

//Data model of the OrderAggregateRoot, where Product is a model and the rest are just ValueObjects
//######################//
// 		DATA MODEL      //
//######################//
type OrderId struct {
	Id string
}

type Product struct {
	Price       float64
	Description string
}

type Order struct {
	OrderId    OrderId
	Products   map[string]Product
	TotalPrice float64
}

type OrderAggregateRoot struct {
	repository OrderRepository
	order      Order
}

//####################################//
// 		INFRASTRUCTURE INTERFACE      //
//####################################//
/*
In order to do DDD properly we need to have IOC(Inversion of control) where domain layer does not have any dependency with any
other module of your service.

In order to receive an implementation of the Repository, domain is the one that has and expose the interface
[OrderRepository] to be implemented by domain module.
*/
type OrderRepository interface {
	FindOrder(id OrderId) Order
	UpsertOrder(order Order) Order
}

//####################################//
// 		CONSTRUCTOR AND LOGIC         //
//####################################//
//Constructor of the type [OrderAggregateRoot] which contains DI [OrderRepository] and Model [Order]
func CreateOrderAggregateRoot(repository OrderRepository) OrderAggregateRoot {
	return OrderAggregateRoot{repository, Order{}}
}

//Function in AggregateRoot to create an Order
func (aggregateRoot OrderAggregateRoot) CreateOrder() Order {
	orderId, err := uuid.NewV4()
	if err != nil {
		return Order{OrderId{"Error creating OrderId"}, make(map[string]Product, 0), 0.0}
	}
	order := Order{OrderId: OrderId{orderId.String()}, Products: make(map[string]Product, 0), TotalPrice: 0.0}
	return aggregateRoot.repository.UpsertOrder(order)
}

//Function in AggregateRoot to find the Order, create a new Data model Product, and add it in the map of products in order.
func (aggregateRoot OrderAggregateRoot) AddProductInOrder(orderId OrderId, productId string, price float64, productDescription string) Order {
	order := aggregateRoot.repository.FindOrder(orderId)
	product := Product{Price: price, Description: productDescription}
	order.Products[productId] = product
	order.TotalPrice = order.TotalPrice + price
	return aggregateRoot.repository.UpsertOrder(order)
}

/*
Function in AggregateRoot to find the Order, find a Product by id, get discount the total price with the product price,
 and delete it in the map of products in order.
*/
func (aggregateRoot OrderAggregateRoot) RemoveProductInOrder(orderId OrderId, productId string) Order {
	order := aggregateRoot.repository.FindOrder(orderId)
	product := order.Products[productId]
	order.TotalPrice = order.TotalPrice - product.Price
	delete(order.Products, productId)
	return aggregateRoot.repository.UpsertOrder(order)
}
