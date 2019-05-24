package main

import (
	"fmt"
	"testing"
)

//Unit test for function 
//In Golang in order to make the test framework of Go detect the test class, the class it must end with _test
func TestFunctions(t *testing.T) {
	result := callAnotherFunc()
	fmt.Println(result)
	if (result) < 198100 {
		t.Errorf("Error invoking function")
	}
	consumeMultiType()
}
