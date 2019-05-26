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
	go runAsyncAction("world")
	currentTime := strconv.Itoa(time.Now().Nanosecond())
	fmt.Println(currentTime + " " + "hello")
	time.Sleep(1500)
}

func runAsyncAction(s string) {
	time.Sleep(1000)
	currentTime := strconv.Itoa(time.Now().Nanosecond())
	fmt.Println(currentTime + " " + s)
}

/*
In GO when we execute a process in a goroutine and we want to pass the process result into the main thread
we need to use channels [chan]. Channel it's created using [make] operator passing the type of the channel.
Once we have it, we need to pass to the func to be run in a goroutine, and once we are ready to pass the
result into the channel we use [channel <- result] from the goroutine process. And from the other main 
thread we wait for the result to be set into the channel [response:=<-channel]
*/
func TestChannel(t *testing.T) {
	channel := make(chan string)
	go runAsyncActionWithChannel(channel, "Hello async Golang world")
	value := <- channel
	println(value)
}

func runAsyncActionWithChannel(channel chan string, s string) {
	time.Sleep(1000)
	currentTime := strconv.Itoa(time.Now().Nanosecond())
	fmt.Println(currentTime + " " + s)
	channel <- "process done"
}

/*
We can reuse the same channel between gotoutines and run in parallel every routine and then
write result into the channel and wait for the result of all routines in the main thread.
The order of execution of the goroutines cannot being garantee.
*/
func TestChannelMultiple(t *testing.T) {
	channel := make(chan string)
	go runAsyncWithChannel(channel, "Hello")
	go runAsyncWithChannel(channel, "async")
	go runAsyncWithChannel(channel, "Golang")
	go runAsyncWithChannel(channel, "world")

	d,c,b,a := <- channel,<-channel,<-channel,<-channel
	println(a + " " + b + " " + c + " " + d)
}

func runAsyncWithChannel(channel chan string, s string) {
	time.Sleep(1000)
	channel <- s
}
