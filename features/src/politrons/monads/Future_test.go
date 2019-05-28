package main

import (
	"strings"
	"testing"
	"time"
)
func TestFutureSuccess(t *testing.T) {
	futureMonad := FutureSuccess{nil}.
		Create(func() interface{} {
			return "Hello Future monad in Go"
		}).
		Map(upperCaseFunc).
		Map(func(i interface{}) interface{} {
			return i.(string) + "!!!!"
		})
	println("In main routine context execution")
	println(futureMonad.Get().(string))
}

var upperCaseFunc = func(i interface{}) interface{} {
	time.Sleep(2000 * time.Millisecond)
	println("In go routine context execution")
	return strings.ToUpper(i.(string))
}

//###########################
//    Monad algebras    
//###########################

type Future interface {
	Create(func() interface{}) Future
	Map(func(interface{}) interface{}) Future
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
func (fs FutureSuccess) Create(fn func() interface{}) Future {
	channel := make(chan interface{})
	go func() {
		channel <- fn()
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

func (fs FutureSuccess) Get() interface{} {
	return <-fs.Channel
}
