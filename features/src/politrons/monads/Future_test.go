package main

import (
	"testing"
)

func TestFutureSuccess(t *testing.T) {
	futureMonad := FutureSuccess{nil}.Create(func() interface{} {
		return "Hello Future monad in Go"
	})
	result := futureMonad.OnComplete().(string)
	println(result)
}

type Future interface {
	Create(func() interface{}) Future
	OnComplete() interface{}
}

type FutureSuccess struct {
	Channel chan interface{}
}

func (fs FutureSuccess) Create(fn func() interface{}) Future {
	channel := make(chan interface{})
	go func (){
		channel <- fn()
	}()
	return FutureSuccess{channel}
}

func (fs FutureSuccess) OnComplete() interface{} {
	return <-fs.Channel
}
