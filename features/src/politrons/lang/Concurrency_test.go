package lang

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
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
	go composeFromPreviousChannel(channel1, channel2)
	result, isChannelOpen := <-channel2
	if isChannelOpen {
		fmt.Println(result)
	} else {
		fmt.Println("Error reading from channel", isChannelOpen)
	}
}

func someAsyncLogic(channel chan User, user User) {
	newUser := User{userID: user.userID, name: "politrons"}
	channel <- newUser
}

func composeFromPreviousChannel(channelCompose chan User, channel chan User) {
	user := <-channelCompose
	newUser := User{userID: strings.ToUpper(user.userID), name: strings.ToUpper(user.name)}
	channel <- newUser
}

type User struct {
	userID string
	name   string
}

/*
Using for { select { } } structure we can have some sort of switch where we can wait for all the channels in parallel
and assing an action once the process has finish. The order is not establish and it will block until all the process
are done.
*/
func testChannelSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)
	go asyncRandomString(channel1, "Hello")
	go asyncRandomString(channel2, "golang")
	go asyncRandomString(channel3, "world")
	for {
		select {
		case value1 := <-channel1:
			println("A", value1)
		case value2 := <-channel2:
			println("B", value2)
		case value3 := <-channel3:
			println("C", value3)
		}
	}
}

func asyncRandomString(channel chan string, value string) {
	channel <- value
}

/*
Go is not Functional language by design, which means it's provide some mechanism to fight against concurrent access to mutable resources.
Here we create a type that wrap a resource map, that we want to synchronize the access. In order to do that you just need to add
a [sync.Mutex] inside your structure, and implement the extended method where you want to mutate that resource.
In that method since it's an extension of your resource, you can use the [sync.Mutex] to [Lock()] and  [Unlock()] access.
*/
func TestSynchronizeResource(t *testing.T) {
	channel1 := make(chan map[string]int)
	channel2 := make(chan map[string]int)
	channel3 := make(chan map[string]int)
	mySafeMap := SafeMap{myMap: map[string]int{"key1": 1, "key2": 2, "key3": 3}}
	go actionOverMap(channel1, "key1", mySafeMap)
	go actionOverMap(channel2, "key2", mySafeMap)
	go actionOverMap(channel3, "key2", mySafeMap)
	response, response1, response2 := <-channel1, <-channel2, <-channel3
	fmt.Println("Map1:", response, "Map2:", response1, "Map3:", response2)
}

//We need to create a sync.Mutex in our type to provide a mechanism to lock the access into.
type SafeMap struct {
	myMap map[string]int
	mux   sync.Mutex
}

func actionOverMap(channel chan map[string]int, key string, mySafeMap SafeMap) {
	myMap := mySafeMap.deleteElementByID(key)
	channel <- myMap
}

func (mySafeMap SafeMap) deleteElementByID(id string) map[string]int {
	mySafeMap.mux.Lock()
	defer mySafeMap.mux.Unlock()
	delete(mySafeMap.myMap, id)
	return mySafeMap.myMap
}

func TestForChannelComposition(t *testing.T) {
	chan1 := pureChannel("hello golang world")
	chan2 := appendChannel(chan1, "!!!!")
	value2 := <-upperChannel(chan2)
	println(value2)
}

func pureChannel(value string) chan string {
	channel := make(chan string)
	go func() {
		channel <- value
	}()
	return channel
}

func appendChannel(prevChannel chan string, anoherValue string) chan string {
	channel := make(chan string)
	go func() {
		channel <- (<-prevChannel) + anoherValue
	}()
	return channel
}

func upperChannel(prevChannel chan string) chan string {
	channel := make(chan string)
	go func() {
		channel <- strings.ToUpper(<-prevChannel)
	}()
	return channel
}
