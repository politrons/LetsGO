package main

import (
	"fmt"
	"strings"
)

func callAnotherFunc() int{
	var response = myFirstFunction("Hello Golang", 1981) 
	fmt.Println(response)
	return response
}

func myFirstFunction(value string, intValue int) int {
	fmt.Println(strings.ToUpper(value))
	return intValue * 100
}
