package main

import (
	"strconv"
	"fmt"
	"testing"
	"time"
)
/*
In Golang to to run in another thread execution we use [go] operator before the function execution
once the execution is finish it will be clean by GC
*/
func TestGoroutine(t *testing.T) {
	go say("world")
	currentTime := strconv.Itoa(time.Now().Nanosecond())
	fmt.Println(currentTime + " " + "hello")
	time.Sleep(1500)
}

func say(s string) {
	time.Sleep(1000)
	currentTime := strconv.Itoa(time.Now().Nanosecond())
	fmt.Println(currentTime + " " + s)
}
