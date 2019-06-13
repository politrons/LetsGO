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
	Product    Product
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
	CreateOrder(order Order) Order
}

//####################################//
// 		CONSTRUCTOR AND LOGIC         //
//####################################//
//Constructor of the type [OrderAggregateRoot] which contains DI [OrderRepository] and Model [Order]
func CreateOrderAggregateRoot(repository OrderRepository) OrderAggregateRoot {
	return OrderAggregateRoot{repository, Order{}}
}

func (aggregateRoot OrderAggregateRoot) CreateOrder() Order {
	orderId, err := uuid.NewV4()
	if err != nil {
		return Order{OrderId{"Error creating OrderId"}, Product{}, 0.0}
	}
	order := Order{OrderId: OrderId{orderId.String()}, Product: Product{}, TotalPrice: 0.0}
	return aggregateRoot.repository.CreateOrder(order)
}
