
package main

import (
	"strconv"
	"fmt"
	"strings"
)

type  Foo struct{
	intValue int
	strValue string
}

func callAnotherFunc() int{
	var response = myFirstFunction("Hello Golang", 1981) 
	fmt.Println(response)
	return response
}

func myFirstFunction(value string, intValue int) int {
	fmt.Println(strings.ToUpper(value))
	return intValue * 100
}

//Go function allow just like in Scala multi type 
func consumeMultiType(){
	 a, b, num, foo := returnMultiType()
	println(a + " " + b + " " + strconv.Itoa(num))
	fmt.Printf("%+v", foo)
}

func returnMultiType() (string, string, int, Foo) {
	return "hello", "Go", 1981, Foo{1,"politrons"}
}
