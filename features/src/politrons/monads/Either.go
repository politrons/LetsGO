package monads

import (
	"reflect"
	"strings"
)

var upperFunc = func(i interface{}) interface{} {
	return strings.ToUpper(i.(string))
}

var appendFunc = func(i interface{}) interface{} {
	return i.(string) + "!!!!!"
}

//###########################
//    Monad algebras
//###########################

/*
A monad Either has two variants, [Right] and [Left] here using interface, we can implement
both variants.
*/
type Either interface {
	Map(func(interface{}) interface{}) Either
	Get() interface{}
	IsRight() bool
	IsLeft() bool
	IsTypeOf(interface{}) bool
}

//All method implementation of this variant must behave as it would be normal for a Right type data
type Right struct {
	Value interface{}
}

//All implementation of this variant must behave as it would be normal for a Left type data
type Left struct {
	Value interface{}
}

//###########################
// Monad implementation
//###########################

//Function to transform the monad applying another function over the monad value
func (r Right) Map(fn func(interface{}) interface{}) Either {
	return Right{fn(r.Value)}
}

//Function to get the monad value
func (r Right) Get() interface{} {
	return r.Value
}

//Function to return if the monad is [Right] variant
func (r Right) IsRight() bool {
	return true
}

//Function to return if the monad is [Left] variant
func (r Right) IsLeft() bool {
	return false
}

func (r Right) IsTypeOf(i interface{}) bool {
	return reflect.TypeOf(r.Get()) == reflect.TypeOf(i)
}

func (l Left) IsTypeOf(i interface{}) bool {
	return reflect.TypeOf(l.Get()) == reflect.TypeOf(i)
}

func (l Left) Map(fn func(interface{}) interface{}) Either {
	return Left{fn(l.Value)}
}

func (l Left) Get() interface{} {
	return l.Value
}

func (l Left) IsRight() bool {
	return false
}

func (l Left) IsLeft() bool {
	return true
}
