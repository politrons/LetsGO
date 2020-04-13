package lang

import (
	"fmt"
	"testing"
)

type MyFoo struct {
	valueA string
	valueB int
}

//In Go a list are arrays without fixed size, just like in other languages you can have a fixed array, puting the pax size [x]
//when you create
//Using [append] operator we can add/delete elements into the array, and return a copy of a new one, allowing us do
//FP using collection and not mutate previous collection.
//In order to  delete the slice in a FP way, we need to create a function [deleteSliceElementFunc] that recicve
// the slice array and the element to be deleted, and return new collection without that element.
func TestArraysAndSlices(t *testing.T) {
	fixedArray := [3]string{"Hello", "GoLang", "World"}
	fmt.Println(fixedArray)

	myArray := make([]int, 5)
	fmt.Println(myArray)

	slice := []int{1, 2, 3, 4, 5}
	newSlice := append(slice, 6, 7, 8, 9, 10)
	fmt.Println(slice)
	fmt.Println(newSlice)

	newSliceDeleted := deleteSliceElementFunc(3, slice)
	fmt.Println(newSliceDeleted)
}

func deleteSliceElementFunc(_value int, slice []int) []int {
	newSlice := []int{}
	for _, value := range slice {
		if value != _value {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}

//In GO the map collector it's created using map[keyType]valueType format.
//To initialize the map with some values we use the operator [make]
func TestMapMake(t *testing.T) {
	var myMap = make(map[string]MyFoo)
	myMap["MyKey"] = MyFoo{"hello", 1981}
	fmt.Println(myMap["MyKey"])
}

//Another way to initialize maps it's using the technique Literal where after define the types as described before
//We open brackets and we introduce one by one the key:value, key:values.
func TestMapLitertal(t *testing.T) {
	var myMap = map[string]MyFoo{
		"firstKey":  MyFoo{"hello", 1981},
		"secondKey": MyFoo{"golang", 2000},
		"thirdKey":  MyFoo{"!!!!", 2019},
	}
	for key, value := range myMap {
		fmt.Println("Key:", key, " Value:", value)
	}
}

//Map allow operators to add/delete elements in the collection. As OOP Language they dont embrace FP
//So the whole collection remains mutable.
//To get the size of a map we use operator [len]
//When you extract the value from a map by key, it can also return a boolean telling if the argument was
//found or not.
func TestMapFeatures(t *testing.T) {
	var myMap = make(map[string]MyFoo)
	myMap["key1"] = MyFoo{"politron", 38}
	myMap["key2"] = MyFoo{"paul", 21}
	delete(myMap, "key1")
	for k, v := range myMap {
		println("Key", k)
		fmt.Println("Value", v)
	}
	println("Lenght of map", len(myMap))
	value, ok := myMap["key1"]
	fmt.Println("The value:", value, "Present?", ok)
	value1, ok := myMap["key2"]
	fmt.Println("The value:", value1, "Present?", ok)

}
