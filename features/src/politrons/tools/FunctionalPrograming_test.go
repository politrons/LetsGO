package tools

import (
	"strings"
	"testing"
)

/*
If you want immutability of primitives types in Go, you can always use const as final types.
*/
func TestFPConstants(t *testing.T) {
	const var1 = "hello immutable world"
	const var2 = var1 + "!!!"
	println("Never mutate types: " + var2)
}

/*
[struct] cannot being constant, but you can force yourself to avoid mutate a struct and create a new one
*/
func TestFPStructImmutable(t *testing.T) {
	user := User{"politrons", 38}
	userUpdated := User{strings.ToUpper(user.name), user.age}
	println(userUpdated)
}

type User struct {
	name string
	age  int
}

/*func (u User) copy() User {
	s,ok := reflect.TypeOf(u).FieldByName("")
	s.
}*/
