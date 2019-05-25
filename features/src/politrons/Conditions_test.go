package main

import (
	"fmt"
	"testing"
)
// Unit test for the for loop functionality
func TestForLoop(t *testing.T) {
	forLoops()
}

// Pretty much the same like in other languajes, quite extend and simple to use like in Scala
func forLoops(){
	foo := 0
	for i := 0; i < 10; i++ {
		foo += i
	}
	fmt.Println(foo)

	for foo = 0; foo < 100; {
		foo = +1
	}
	fmt.Println(foo)

	for foo < 100 {
		foo = +1
	}
	fmt.Println(foo)
}
func TestForRange(t *testing.T) {
	forRange()
}

//Using [range] operator we can extract and interate elements from an array and get index and value.
//You can also do it to iterte over a map and extract pair of key/value per entry.
func forRange(){
	var array = []int{1,4,30,10,5,68}
    for index,value := range array {
		println("Index" ,index) 
		println("Value" ,value)
	}
	myMap := map[int]string{1:"value1",2:"value2",3:"value3"}
	for key, value := range myMap{
		println("Key", key, "Value", value)
	}
}