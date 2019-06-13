package infrastructure

import . "domain"

//Implementation type of [OrderRepository]
type OrderRepositoryImpl struct {
	database map[string]Order
}

//Constructor of the Type [OrderRepositoryImpl]
func CreateOrderRepository() OrderRepositoryImpl {
	return OrderRepositoryImpl{database: make(map[string]Order)}
}

//Here we implement the [UpsertOrder] of interface [OrderRepository] that is defined in the Domain module to have IOC.
func (repository OrderRepositoryImpl) UpsertOrder(order Order) Order {
	repository.database[order.OrderId.Id] = order
	return order
}

/*
Here we implement the [FindOrder] of interface [OrderRepository] that is defined in the Domain module to have IOC.
The implementation of the Repository contains a Database in memory, so we search in there for the Order using the Id
*/
func (repository OrderRepositoryImpl) FindOrder(orderId OrderId) Order {
	return repository.database[orderId.Id]
}
