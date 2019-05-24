package main

import (
	"fmt"
)

// Pretty much the same like in other languajes, quite extend and simple to use like in Scala
func forLoops() {

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
