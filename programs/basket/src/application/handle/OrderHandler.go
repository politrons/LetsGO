package handle

import (
	"application/commands"
	. "domain"
	"infrastructure"
)

//Type of OrderHandler that define all types that contains.
type OrderHandler struct {
	orderAggregateRoot OrderAggregateRoot
}

/*
Constructor of the Type [OrderHandler]
/*
DI of [OrderAggregateRoot], which is the responsible for business logic and persis the entity in the infrastructure.
Also the [OrderAggregateRoot] require a DI for the repository in order to persist the Model in the infra, so we
have to pass the dependency in the constructor of the type.
*/
func CreateOrderHandler() OrderHandler {
	orderAggregateRoot := CreateOrderAggregateRoot(infrastructure.CreateOrderRepository())
	return OrderHandler{orderAggregateRoot}
}

/*
Extended method defined OrderHandler, now all classes that contains an instance of [OrderHandler]
can use this method.
*/
func (handler OrderHandler) CreateOrder(command commands.CreateOrder) Order {
	return handler.orderAggregateRoot.CreateOrder()
}
