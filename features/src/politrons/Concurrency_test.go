package main

import (
	"fmt"
	"strconv"
	"strings"
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
	value := <-channel
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
The order of execution of the goroutines cannot being guarantee.
*/
func TestChannelMultiple(t *testing.T) {
	channel := make(chan string)
	go runAsyncWithChannel(channel, "Hello")
	go runAsyncWithChannel(channel, "async")
	go runAsyncWithChannel(channel, "Golang")
	go runAsyncWithChannel(channel, "world")

	d, c, b, a := <-channel, <-channel, <-channel, <-channel
	println(a + " " + b + " " + c + " " + d)
}

func runAsyncWithChannel(channel chan string, s string) {
	time.Sleep(1000)
	channel <- s
}

/*
We can have Goroutine composition like [flatMap] with Scala futures creating two channels once, for the first Goroutine, 
and another to the second one, and the second routine receive not only his own channel but the previous routine function
to start working, once receive the end of the previous one.
Doing this, we manage to work sequentially doing composition.
*/
func TestGoroutineComposition(t *testing.T) {
	channel1 := make(chan User)
	channel2 := make(chan User)
	go someAsyncLogic(channel1, User{"customUserId", ""})
	go composeFromPreviousChannel(channel1,channel2)
	result :=<-channel2
	fmt.Println(result)
}

func someAsyncLogic(channel chan User, user User) {
	newUser := User{userID: user.userID, name: "politrons"}
	channel <- newUser
}

func composeFromPreviousChannel(channelCompose chan User,channel chan User) {
	user := <-channelCompose
	newUser := User{userID: strings.ToUpper(user.userID), name: strings.ToUpper(user.name)}
	channel <- newUser
}

type User struct {
	userID string
	name   string
}
