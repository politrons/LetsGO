package main

import (
	"testing"
	"fmt"
)

type MyFirstInterface interface {
	contractToImplement()
}

type MyName string

type MyCustomType struct {
	name string
	sex string
}

/*
Interface in GO are a tuple of (value,type) the type it's as implementation of the interface,
 once we use an extend method with the same nme than the interface.

 Just like in other languajes, one interface can have multiple implementations, and also a type 
 it might have also multiple interface to implement.

 In this example we have types [Name] and [MyCustomType] and sincxe both has an extended method implemented 
 with the same name than the interface, both indeed implement MyFirstInterface.

 This design it could be somehow similar to type class of Haskell, where depending of the type it will
 be redirected to a specific implementation.
*/
func TestInterfaces(t *testing.T){
	var myfirstInterfaceImpl MyFirstInterface = MyCustomType{"politrons", "male"}
	myfirstInterfaceImpl.contractToImplement()
	fmt.Printf("(%v, %T)\n", myfirstInterfaceImpl, myfirstInterfaceImpl)

	var nameType MyFirstInterface = MyName("Golang has cool features")
	nameType.contractToImplement()
	fmt.Printf("(%v, %T)\n", nameType, nameType)
}

/*
In Go the [Any] type of Scala it's called [interface{}] which means it can be any possible type.
*/
func TestInterfaceGeneric(t *testing.T){
	var anyVal interface{}
	anyVal = "politrons"
	fmt.Printf("(%v, %T)\n", anyVal, anyVal)
	anyVal = 1981
	fmt.Printf("(%v, %T)\n", anyVal, anyVal)
}

// This method means type MyCustomType implements the interface MyFirstInterface,
func (customType MyCustomType) contractToImplement() {
	fmt.Println(customType.name + " " + customType.sex)
}

// This method means type Name implements the interface MyFirstInterface,
func (name MyName) contractToImplement() {
	fmt.Println(name)
}

