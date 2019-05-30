package lang

import (
	"fmt"
	"strings"
	"testing"
)

type MethodFoo struct {
	age  int
	name string
}

type MyList []string

/*
[Methods] are called in GO all functions that receice a [receiver] argument. To me it's just like
the method extensions system that Scala use with implicits. This type of sugar syntax methods
introduce a really powerful mechanism to create DSL and also to make the code more extensible.

Method extensions only works with [type] so it force us to create Types, which it's good.

it's also possible combine multiple method extensions.
*/
func TestMethodFeatures(t *testing.T) {
	foo := MethodFoo{38, "politrons"}
	newFoo := foo.upperCaseMethod()
	fmt.Println(newFoo)

	myList := MyList{"hello", "golang", "bla", "world", "extension", "methods", "it`s", "cool"}
	newList := myList.deleteByIndex(2).toUpprCase()
	fmt.Println(newList)
}

func (foo MethodFoo) upperCaseMethod() MethodFoo {
	return MethodFoo{age: 38, name: strings.ToUpper(foo.name) + "!!"}
}

//Function that extend the functionality of the type List and allow delete elements in the list by index.
func (myList MyList) deleteByIndex(i int) MyList {
	newList := []string{}
	for index, value := range myList {
		if i != index {
			newList = append(newList, value)
		}
	}
	return newList
}

//Function that iterate the internal type that extend and put every element in upperCase.
func (myList MyList) toUpprCase() MyList {
	newList := []string{}
	for _, value := range myList {
		newList = append(newList, strings.ToUpper(value))
	}
	return newList
}
