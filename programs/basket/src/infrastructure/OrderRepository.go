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

/*
Here we implement the [UpsertOrder] of interface [OrderRepository] that is defined in the Domain module to have IOC.
Since the operation to database normally are blocked we make it asynchronously using channels and go routines.
We create the channel where we will use to communicate from the new thread into the main or other thread.
*/
func (repository OrderRepositoryImpl) UpsertOrder(order Order) chan Order {
	chanOrder := make(chan Order)
	go func() {
		repository.database[order.OrderId.Id] = order
		chanOrder <- order
	}()
	return chanOrder
}

/*
Here we implement the [FindOrder] of interface [OrderRepository] that is defined in the Domain module to have IOC.
The implementation of the Repository contains a Database in memory, so we search in there for the Order using the Id
Again this operation is async so we will use channel to share the output of the go routine computation.
*/
func (repository OrderRepositoryImpl) FindOrder(orderId OrderId) chan Order {
	chanOrder := make(chan Order)
	go func() { chanOrder <- repository.database[orderId.Id] }()
	return chanOrder
}
