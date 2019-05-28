package main

import (
	"strings"
	"testing"
)

func TestSucceedMonad(t *testing.T) {
	tryMonad := getTry("Hello Try monad in Go").Map(func(i interface{}) interface{}{
		return strings.ToUpper(i.(string))
	}).Map(func(i interface{}) interface{}{
		return i.(string) + "!!!!!"
	})
	println(tryMonad.isSuccess())
	println(tryMonad.isFailure())
	response := tryMonad.Get().(string)
	println(response)
}

func TestFailureMonad(t *testing.T) {
	tryMonad := getTry(MyError{"Custom error in Go"}).MapError(func(e error)error {
		return MyError{e.Error() + " Append extra error info"}
	})
	println(tryMonad.isSuccess())
	println(tryMonad.isFailure())
	response := tryMonad.Get().(error)
	println(response.Error())
}

//###########################
//#	   Monad algebras       #
//###########################

/*
A monad Try has two variants, [Success] and [Failure] here using interface, we can implement
both variants.
*/
type Try interface {
	Map(func(interface{})interface {}) Try
	MapError(func(error) error) Try
	Get() interface{}
	isSuccess() bool
	isFailure() bool
}

//All method implementation of this variant must behave as it would be normal for a Success data
type Success struct {
	Value interface{}
}

//All implementation of this variant must behave as it would be normal for a Failure data
type Failure struct {
	Error error
}

//###########################
//#	 Monad implementation   #
//###########################

//Function to transform the monad applying another function over the monad value
func (s Success) Map(fn func(interface{})interface {}) Try {
	return Success{fn(s.Value)}
}

//Function to transform the monad error applying another function over the monad value
func (s Success) MapError(fn func(error)error) Try {
	return nil
}

//Function to get the monad value
func (s Success) Get() interface{} {
	return s.Value
}

//Function to return if the monad is [Success] variant
func (s Success) isSuccess() bool {
	return true
}

//Function to return if the monad is [Failure] variant
func (s Success) isFailure() bool {
	return false
}

func (f Failure) Map(fn func(interface{}) interface {}) Try {
	return nil
}

func (f Failure) MapError(fn func(error) error) Try {
	return Failure{fn(f.Error)}
}

func (f Failure) Get() interface{} {
	return f.Error
}

func (f Failure) isSuccess() bool {
	return false
}

func (f Failure) isFailure() bool {
	return true
}

//######################
//#   Utils functions  #
//######################
func getTry(i interface{}) Try {
	switch value := i.(type) {
	case string:
		return Success{value}
	case error:
		return Failure{value}
	default:
		panic("Not controlled option")
	}
}

type MyError struct {
	cause string
}

func (e MyError) Error() string {
	return e.cause
}
