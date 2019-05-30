package lang

import (
	"fmt"
	"strings"
	"testing"
)

/*
Go has same system as C or C++ with pointers.
We can create types that it will hold a value [Variables], and types that it will hold another type address [Pointers].

In essence, pointers are just like use Observer patter, pointers are subscribed to subjects[Variables] and when
variables change his values, since pointers are sharing memory allocation, we share that change.

* over a type variable, can only hold another variable address. -> var pointerVar *string.
& over a variable, return the address of that variable in memory.
	If we get a pointer variable and we use & over a variable, it will give the address variable, and it will assign to the pointer.
* over a pointer variable, it will give the value of the memory address allocated. -> *pointerVar it will get the value of the variable [var1] which address
	allocated in pointerVar. -> pointerVar = &var1

One really important thing to remember about pointers is, that assign of value is bidirectional, so in case we assign a pointer the address of a variable,
means that both share the same memory allocation, and both can change that variable in memory affecting both
*/
func TestPointers(t *testing.T) {

	var var1 string
	var var2 string
	var pointerVar *string //Marking the type with * means we are creating a type that it will point to another address, but not value.

	//If you uncomment this line, you will receive [panic: runtime error: invalid memory address or nil pointer dereference [recovered]]
	//fmt.Printf("Pointer without assigment: pointerVar = %v\n*pointerVar = %v\n", pointerVar, *pointerVar)

	var1 = "Hello pointers world"
	var2 = var1        //This is just a copy in memory of Var1 to Var2
	pointerVar = &var1 //We point pointerVar [pointer]
	//pointerVar = "It wont compile" Since it's a pointer and not a variable.

	//Since var2 is a copy of var2 they have different address, but var1 and var2 they share address
	fmt.Printf("Address:\n var1 = %v\n var2 = %v\n*pointerVar = %v\n", &var1, &var2, pointerVar)

	fmt.Printf("Values:\n var1 = %v\n var2 = %v\n pointerVar = %v\n*pointerVar = %v\n", var1, var2, pointerVar, *pointerVar)

	var1 = strings.ToUpper(var1) //Since [var1] and [pointerVar] share memory allocation, when you change one of them, affect the other.

	fmt.Printf("Values:\n var1 = %v\n var2 = %v\n pointerVar = %v\n*pointerVar = %v\n", var1, var2, pointerVar, *pointerVar)

	pointerVar = &var2 //We can allocate a new address to the pointer any time. So var1 it's not link with [pointerVar] anymore.

	fmt.Printf("Values:\n var1 = %v\n var2 = %v\n pointerVar = %v\n*pointerVar = %v\n", var1, var2, pointerVar, *pointerVar)

	*pointerVar = "Now I change the pointerVar and then affect also Var2" //Here is [*pointerVar] the one that change the variable, and affect var2

	fmt.Printf("Change Pointer:\n var1 = %v\n var2 = %v\n pointerVar = %v\n*pointerVar = %v\n", var1, var2, pointerVar, *pointerVar)

}
