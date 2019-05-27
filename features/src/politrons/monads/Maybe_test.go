package main

import (
	"strings"
	"testing"
)
/*
Go cannot escape of FP, it's a powerful tool, and even having to do it without lambdas the concept of 
transformation and composition can being applyied.
Here applying interfarces type with method extensions allow us that Maybe interface can have variant 
[Just] or [Nothing]. This approach works pretty similar like type classes of Haskell.
*/
func TestMaybeJust(t *testing.T) {
	maybe := getMaybe("Hello Maybe in Golang")
	println(maybe.isDefined())
	println(maybe.Get().(string))
	result := 
	maybe.Then(func(i interface{}) interface{} {
		return strings.ToUpper(i.(string))
	}).Then(func(i interface{}) interface{} {
		return i.(string) + "!!!"
	}).Get().(string)
	println(result)
}

func TestMaybeNothing(t *testing.T) {
	maybe := getMaybe("")
	println(maybe.isDefined())
	println(maybe.Get())
}

//Maybe monad with variant Just or Nothing
type Maybe interface {
	Pure(interface{}) Maybe
	isDefined() bool
	Get() interface{}
	Then(func(interface{}) interface{}) Maybe
}

//Allegra of the Maybe monad
type Just struct {
	Value interface{}
}

//Allegra of the Maybe monad
type Nothing struct{}

//Function to wrap a value i into the Just[interface{}] monad
func (just Just) Pure(i interface{}) Maybe {
	return Just{i}
}

//Function to let know if the monad is full or empty
func (just Just) isDefined() bool {
	return true
}

//Function just to extract the value of the Monad
func (just Just) Get() interface{} {
	return just.Value
}

//Composition operator, having a Just we get the value from it, and we return another Just.
func (just Just) Then(fn func(interface{}) interface{}) Maybe {
	return Just{fn(just.Get())}
}

func (n Nothing) Pure(i interface{}) Maybe {
	return Nothing{}
}

func (n Nothing) isDefined() bool {
	return false
}

func (n Nothing) Get() interface{} {
	return nil
}

func (n Nothing) Then(fn func(interface{}) interface{}) Maybe {
	return nil
}

func getMaybe(value string) Maybe {
	if value != "" {
		return Just{value}
	} else {
		return Nothing{}
	}
}
