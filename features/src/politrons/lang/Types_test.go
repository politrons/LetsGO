package lang

import (
	. "fmt"
	"testing"
)

/*
In Go all primitive types can be wrapped in types, preventing to use just the primitive type in your functions
making the code more readable and type-safe.
To create a type of a primitive, you just need to wrap the value in the type Name("politrons")
*/
type Name string
type Age int
type Sex string

//In Go in order to be detected as Unit test by the test framework you need to start the
//Test case with [Test] also the file name must end with [_test.go]
func TestTypes(t *testing.T) {
	primitiveTypes()
	constTypes()
	typeStruct()
}

//Animal type that define two primitive types
type Animal struct {
	age Age
	sex Sex
}

//Human type that define one primitive type and the Animal type
type Human struct {
	name   Name
	animal Animal
}

type Dog struct {
	name   Name
	tail   int
	animal Animal
}

//In Go we have the same type like in the rest of the languajes. You can avoid the type declaration since is inferred by the Go compiler.
// Instead of have double we have float of 32 or 64
//We can have multiple declaration with different types like String, int, bool
func primitiveTypes() {
	var i int = 1
	var f float64 = 1.1
	var b bool = false
	s := "Hello Go"
	var stringVal, intVal, boolVal = "String value", 1, true
	println(i, f, b, s, stringVal, intVal, boolVal)
}

//Const types in Go are like creation of immutable types in Scala with [val] the compiler
// it won’t allow reassignment of a variable already asigned initially
func constTypes() {
	const stringVal, intVal, boolVal = "String value", 1, true
	println(stringVal, intVal, boolVal)
}

// In Go we can create types adding the key [struct] at the end of the name of the type
// Also to define the types it use a pretty similar syntax as Haskell, we just need to use {} to define it
// using the name of the attributes or without it, and respecting the order of the arguments.
func typeStruct() {
	man := Human{
		name:   Name("Politrons"),
		animal: Animal{age: Age(10), sex: Sex("male")},
	}
	women := Human{"Esther", Animal{35, "female"}}
	Println(man)
	Println(women)

	var dog = Dog{name: Name("Bingo"), tail: 10, animal: Animal{8, "male"}}
	Println(dog)
	Println(man.name + " - " + women.name)
}
