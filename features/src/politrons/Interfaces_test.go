package main

import (
	"testing"
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func TestInterfaces(t *testing.T){
	var i I = T{"hello"}
	i.M()
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}


