package handle

import (
	"application/commands"
	. "domain"
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
func CreateOrderHandler(repository OrderRepository) OrderHandler {
	orderAggregateRoot := CreateOrderAggregateRoot(repository)
	return OrderHandler{orderAggregateRoot}
}

/*
Extended method defined OrderHandler, now all classes that contains an instance of [OrderHandler]
can use this method.
Create a new Order using the orderEntityAggregateRoot
*/
func (handler OrderHandler) CreateOrder(command commands.CreateOrder) Order {
	return handler.orderAggregateRoot.CreateOrder()
}

/*
This method it find the order using the id and it update it adding a new product, and updating the totalPrice
*/
func (handler OrderHandler) UpdateOrder(orderId OrderId, command commands.AddProduct) Order {
	return handler.orderAggregateRoot.UpdateOrder(orderId, command.ProductId, command.Price, command.ProductDescription)
}
