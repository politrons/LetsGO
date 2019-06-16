package infrastructure

import (
	. "domain"
	"github.com/sony/gobreaker"
	"time"
)

//Global Circuit breaker variable
var cb *gobreaker.CircuitBreaker

/*
[Name] is the name of the CircuitBreaker.
[ReadyToTrip] bool function tp determine if the Circuit breaker must be open.
[OnStateChange] Callback function to invoke when the Circuit breaker change state.
[MaxRequests] Maximum number of requests allowed to pass through when the CircuitBreaker is half-open
[Timeout] Period of the open state, after which the state of CircuitBreaker becomes half-open.

[NewCircuitBreaker] create the Circuit breaker instance.
*/
func init() {
	var settings gobreaker.Settings
	settings.Name = "Politrons"
	settings.ReadyToTrip = maxConsecutiveFailuresStrategyFunc
	settings.MaxRequests = 1
	settings.Timeout = time.Duration(2000 * time.Millisecond)
	cb = gobreaker.NewCircuitBreaker(settings)
}

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

In this code we run the access to the "backend" through a Circuit breaker, but we dont control the error response.
*/
func (repository OrderRepositoryImpl) UpsertOrder(order Order) chan Order {
	chanOrder := make(chan Order)
	go func() {
		persistedOrder, _ := cb.Execute(func() (interface{}, error) {
			repository.database[order.OrderId.Id] = order
			return order, nil
		})
		chanOrder <- persistedOrder.(Order)
	}()
	return chanOrder
}

/*
Here we implement the [FindOrder] of interface [OrderRepository] that is defined in the Domain module to have IOC.
The implementation of the Repository contains a Database in memory, so we search in there for the Order using the Id
Again this operation is async so we will use channel to share the output of the go routine computation.

In this code we run the access to the "backend" through a Circuit breaker, but we dont control the error response
*/
func (repository OrderRepositoryImpl) FindOrder(orderId OrderId) chan Order {
	chanOrder := make(chan Order)
	go func() {
		order, _ := cb.Execute(func() (interface{}, error) {
			return repository.database[orderId.Id], nil
		})
		chanOrder <- order.(Order)
	}()
	return chanOrder
}

/*
Here we implement another function strategy for the Circuit breaker
*/
func maxConsecutiveFailuresStrategyFunc(counts gobreaker.Counts) bool {
	return counts.Requests >= 5 && counts.ConsecutiveFailures >= 5
}
