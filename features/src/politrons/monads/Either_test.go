package main

import (
	"strings"
	"testing"
)

func TestRightMonad(t *testing.T) {
	eitherMonad := getEither("Hello Either Right monad in Go", true).Map(func(i interface{}) interface{}{
   		return strings.ToUpper(i.(string))
   	}).Map(func(i interface{}) interface{}{
   		return i.(string) + "!!!!!"
   	})
	println("Right:",eitherMonad.isRight())
	println("Left:",eitherMonad.isLeft())
   	response := eitherMonad.Get().(string)
   	println(response)
}

func TestLeftMonad(t *testing.T) {
	eitherMonad := getEither("Hello Either Left monad in Go", false).Map(func(i interface{}) interface{}{
		return strings.ToUpper(i.(string))
	}).Map(func(i interface{}) interface{}{
		return i.(string) + "!!!!!"
	})
	println("Right:",eitherMonad.isRight())
	println("Left:",eitherMonad.isLeft())
	response := eitherMonad.Get().(string)
	println(response)
}

//###########################
//#	     Monad algebras     #
//###########################

/*
A monad Either has two variants, [Right] and [Left] here using interface, we can implement
both variants.
*/
type Either interface {
	Map(func(interface{}) interface{}) Either
	Get() interface{}
	isRight() bool
	isLeft() bool
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
//#	 Monad implementation   #
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
func (r Right) isRight() bool {
	return true
}

//Function to return if the monad is [Left] variant
func (r Right) isLeft() bool {
	return false
}

func (l Left) Map(fn func(interface{}) interface{}) Either {
	return Left{fn(l.Value)}
}

func (l Left) Get() interface{} {
	return l.Value
}

func (l Left) isRight() bool {
	return false
}

func (l Left) isLeft() bool {
	return true
}

//######################
//#   Utils functions  #
//######################
func getEither(value string, right bool) Either {
	if right {
		return Right{value}
	} else {
		return Left{value}
	}
}
