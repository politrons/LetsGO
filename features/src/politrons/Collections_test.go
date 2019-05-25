package main

import (
	"fmt"
	"testing"
)

func TestMapMake(t *testing.T) {
	mapMake()
}

type MyFoo struct {
	valueA string
	valueB int
}

//In GO the map collector it's created using map[keyType]valueType format.
//To initialize the map with some values we use the operator [make]
func mapMake() {
	var myMap = make(map[string]MyFoo)
	myMap["MyKey"] = MyFoo{"hello", 1981}
	fmt.Println(myMap["MyKey"])
}

func TestMapLitertal(t *testing.T) {
	mapLitertal()
}
//Another way to initialize maps it's using the technique Literal where after define the types as described before
//We open brackets and we introduce one by one the key:value, key:values.
func mapLitertal() {
	var myMap = map[string]MyFoo{
		"firstKey":MyFoo{"hello", 1981},
		"secondKey":MyFoo{"golang", 2000},
		"thirdKey":MyFoo{"!!!!", 2019},
	}
	for key, value := range myMap {
		fmt.Println("Key:", key, " Value:", value)
	}
}

func TestMapFeatures(t *testing.T) {
	mapFeatures()
}
//Map allow operators to add/delete elements in the collection. As OOP Language they dont embrace FP
//So the whole collection remains mutable.
//To get the size of a map we use operator [len] 
//When you extract the value from a map by key, it can also return a boolean telling if the argument was
//found or not.
func mapFeatures() {
	var myMap = make(map[string]MyFoo)
	myMap["key1"] = MyFoo{"politron", 38}
	myMap["key2"] = MyFoo{"paul", 21}
	delete(myMap,"key1")
	for k,v := range myMap {
		println("Key", k)
		fmt.Println("Value", v)
	}
	println("Lenght of map", len(myMap))

	value, ok := myMap["key1"]
	fmt.Println("The value:", value, "Present?", ok)
	value1, ok := myMap["key2"]
	fmt.Println("The value:", value1, "Present?", ok)
	
}