package main

import (
	"testing"
)

func TestMaybeJust(t *testing.T) {
	maybe := getMaybe(true)
	println(maybe.isDefined())
	println(maybe.Get().(string))
}

func TestMaybeNothing(t *testing.T) {
	maybe := getMaybe(false)
	println(maybe.isDefined())
	println(maybe.Get())
}

func getMaybe(t bool) Maybe {
	if(t) {
		return Just{"Hello Maybe in Golang"}
	}else{
		return Nothing{}
	}
}

type Maybe interface{
	Pure(interface{}) Maybe
	isDefined() bool
	Get() interface {}
}

type Just struct {
	Value interface {}
}

type Nothing struct{}

func (just Just) Pure(i interface{}) Maybe{
	return Just{i}
}

func (just Just) isDefined() bool{
	return true
}

func (just Just) Get() interface{}{
	return just.Value
}

func (n Nothing) Pure(i interface{}) Maybe{
	return Nothing{}
}

func (n Nothing) isDefined() bool{
	return false
}

func (n Nothing) Get() interface{}{
	return nil
}
