package monads

import (
	"strings"
	"time"
)

var upperCaseFunc = func(i interface{}) interface{} {
	time.Sleep(2000 * time.Millisecond)
	println("In go routine context execution")
	return strings.ToUpper(i.(string))
}

//###########################
//    Monad algebras
//###########################

/*
Future monad provide the next functions:
	* [Create] To create a new Future Monad.
	* [Map] To transform the value that contains one channel into another after apply a function.
	* [Get] To get the value from the channel that the Future contains.
*/
type Future interface {
	Create(interface{}) Future
	Map(func(interface{}) interface{}) Future
	FlatMap(func(interface{}) Future) Future
	Get() interface{}
}

//The [FutureSuccess] contains the channel where we set the async execution results.
type FutureSuccess struct {
	Channel chan interface{}
}

//###########################
//  Monad implementation
//###########################

//Function that create a Monad Future where set the channel where the async execution set the return value
func (fs FutureSuccess) Create(i interface{}) Future {
	channel := make(chan interface{})
	go func() {
		channel <- i
	}()
	return FutureSuccess{channel}
}

/*
Function that transform a Monad Future passing to the function the return of a previous execution. where the async execution.
Finally we create a new Future and we set the new channel with the response of the function.
*/
func (fs FutureSuccess) Map(fn func(interface{}) interface{}) Future {
	channel := make(chan interface{})
	go func() {
		channel <- fn(<-fs.Channel)
	}()
	return FutureSuccess{channel}
}

/*
With [FlatMap] we can do composition of Future Monads, so we can execute every computation async sequentally
without have to do any block.
*/
func (fs FutureSuccess) FlatMap(fn func(interface{}) Future) Future {
	channel := make(chan interface{})
	go func() {
		channel <- fn(<-fs.Channel).Get()
	}()
	return FutureSuccess{channel}
}

func (fs FutureSuccess) Get() interface{} {
	return <-fs.Channel
}
