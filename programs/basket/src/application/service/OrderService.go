package service

import (
	. "domain"
)

//Type of [OrderService] which contains all DI types that needs to work.
type OrderService struct {
	repository OrderRepository
}

//Constructor of [OrderService]
func CreateOrderService(repository OrderRepository) OrderService {
	return OrderService{repository: repository}
}

/*
Extended method defined [OrderService], now all classes that contains an instance of [OrderService]
can use this method.
Internally since [OrderService] has a dependency with [OrderRepository] we use it, to reach the
infrastructure layer to get the order.
In case of Queries, since we need to be fast, we dont delegate the obtain of the Order to the domain,
 but we do an interaction between application and infrastructure layer directly.
*/
func (service OrderService) FindOrder(orderId OrderId) chan Order {
	return service.repository.FindOrder(orderId)
}
