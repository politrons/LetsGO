package main

import (
	"fmt"
	"testing"
)

//In Go in order to being detected as Unit test by the test framework you need to start the 
//Test case with [Test] also the file it must end with [_test.go]
func TestTypes(t *testing.T){
	primitiveTypes()
	constTypes()
	typeStruct()
}

//Animal type that define two primitive types
type Animal struct {
	age int
	sex string
}

//Human type that define one primitive type and the Animal type
type Human struct {
	name   string
	animal Animal
}

type Dog struct {
	name   string
	tail   int
	animal Animal
}

//In Go we have the same type like in the rest of the languajes. You can avoid the type declaration since is inferred by the Go compiler.
// Instead of have double we have float of 32 or 64
//We can have multiple declaration with different types like String, int, bool
func primitiveTypes() bool {
	var i int = 1
	var f float64 = 1.1
	var b bool = false
	var s = "Hello Go"
	var stringVal, intVal, boolVal = "String value", 1, true
	println(i, f, b, s, stringVal, intVal, boolVal)
	return true
}

//Const types in Go is like create immutable types in Scala with [val] the compiler
// it wont allow reasign a variable already asigned initiallyl
func constTypes() bool {
	const stringVal, intVal, boolVal = "String value", 1, true
	println(stringVal, intVal, boolVal)
	return true
}

// In Go we can create types adding the key [struct] at the end of the name of the type
// Also to define the types it use a pretty similar syntax as Haskell, we just need to use {} to define it
// using the name of the attributes or witout it, and respeting the order of the arguments.
func typeStruct() bool {
	man := Human{
		name:   "Politrons",
		animal: Animal{age: 10, sex: "male"},
	}
	women := Human{"Esther", Animal{35, "female"}}
	fmt.Println(man)
	fmt.Println(women)

	var dog = Dog{name: "Bingo", tail: 10, animal: Animal{8, "male"}}
	fmt.Println(dog)
	

	fmt.Println(man.name + " - " + women.name)
	return true
}
