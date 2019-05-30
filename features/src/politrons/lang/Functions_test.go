package lang

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type Foo struct {
	intValue int
	strValue string
}

//Unit test for function
//In Golang in order to make the test framework of Go detect the test class, the class it must end with _test
//In Go you can define variable just using [var] or in the [:=] assign a colon before the equals.
func TestFunctions(t *testing.T) {
	var response = myFirstFunction("Hello Golang", 1981)
	fmt.Println(response)
}

func myFirstFunction(value string, intValue int) int {
	fmt.Println(strings.ToUpper(value))
	return intValue * 100
}

//Go function allow just like in Scala multi type.
//The number of types expected from the invoker function must be the same.
func TestMultiTypeReturnFunctions(t *testing.T) {
	a, b, num, foo := returnMultiType()
	println(a + " " + b + " " + strconv.Itoa(num))
	fmt.Printf("%+v", foo)
}

//In go it's also possible return multiple types in one function
func returnMultiType() (string, string, int, Foo) {
	return "hello", "Go", 1981, Foo{1, "politrons"}
}

func TestHighOrderFunctions(t *testing.T) {
	highOrderFunctionInput()
	highOrderFunctionOutput()
}

//In go even as not FP language allow High order functions. In this examples we pass one function from one to another.
func highOrderFunctionInput() {
	concatNameAndAge := func(name string, age int) string {
		return (strings.ToUpper(name) + "*" + strconv.Itoa(age))
	}
	response := getNameAndAgeFunc(concatNameAndAge)
	println(response)
}

func getNameAndAgeFunc(fn func(string, int) string) string {
	return fn("Politrons", 38)
}

//Also it's possible apply High order functions but instead of pass a function we can receive a function
//when we invoke a function.
func highOrderFunctionOutput() {
	concatNameAndAge := concatNameAndAgeFunc()
	response := concatNameAndAge("Paul", 21)
	println(response)
}

func concatNameAndAgeFunc() func(string, int) string {
	return func(name string, age int) string {
		return (strings.ToUpper(name) + "-" + strconv.Itoa(age))
	}
}

//Go even as not FP language allow High order functions, in here, Slice accept an
// interface{} as first argument, and then a function to order the array.
func TestLambdFunctions(t *testing.T) {
	var people = []string{"Politrons", "Jaime", "John"}
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})
	fmt.Println(people)
}
