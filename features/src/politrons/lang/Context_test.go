package lang

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*

Context type, allow us contains information together with deadlines, cancelation signals and timeouts.

When we receive a call in our program, we should create a Context, and outgoing calls to servers should accept a Context.

The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created
using [WithCancel], [WithDeadline], [WithTimeout], or [WithValue]. When a Context is canceled, all Contexts derived from it are also canceled.
*/

/*
Using [WithValue] function over context we are able to create a context, as part of the function we need to pass
an empty context, for that we use [Background()] which returns a non-nil, empty Context, and a key/value entry
in case you want to share some info in the context.
*/
func TestContextValue(t *testing.T) {
	key := contextKey("myKey")
	ctx := context.WithValue(context.Background(), key, "Go value in context")

	extractContextValue(ctx, key)
	extractContextValue(ctx, contextKey("anotherKey"))
}

type contextKey string

func extractContextValue(ctx context.Context, k contextKey) {
	if v := ctx.Value(k); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found:", k)
}

/*
Using [WithTimeout] function, we can create a context with a TTL, after that time, the resource it will be close
and clean, and the context channel [Done] it will be ready to be consumed. Then the context it will have
the [Error] attribute.
*/
func TestContextTimeout(t *testing.T) {
	ctx := context.WithValue(context.Background(), "myKey", "Go value in context")
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel() // releases resources if maybeOperationSlow completes before timeout elapses
	value, err := maybeOperationSlow(ctx)
	if err != nil {
		fmt.Printf("Error:%e", err)
	} else {
		fmt.Printf("Response:%s", value)
	}
}

func maybeOperationSlow(ctx context.Context) (string, error) {
	select {
	case <-time.After(time.Duration(random(0, 1000)) * time.Millisecond):
		return ctx.Value("myKey").(string), nil
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
		return "", ctx.Err()
	}
}

/*
Another handy way to use [context] with Timeout, is when you want to set a maximum timeout for a goroutine.
Here in case the process take more than 500ms the <-ctx.Done function it will be invoked, and the async process it will
be end and clean normally.

Then from the invoker since you cannot trust that the channel return something, the normal pattern is to use a
[select] where you can pattern matching what to do in case channel return value, or some timeout happens.
*/
func TestContextTimeoutReleasingGoRoutines(t *testing.T) {
	ctx := context.WithValue(context.Background(), "myKey", "Go value in context")
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel() // releases resources if asyncFunc completes before timeout elapses
	select {
	case response := <-asyncFunc(ctx):
		fmt.Printf("Response:%s \n", response)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Go routine took too long. Timed out")
	}
}

func asyncFunc(ctx context.Context) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			select {
			case <-time.After(time.Duration(random(0, 1000)) * time.Millisecond):
				channel <- ctx.Value("myKey").(string)
			case <-ctx.Done():
				fmt.Println("Cleaning resources")
				fmt.Println(ctx.Err()) // prints "context deadline exceeded"
				return
			}
		}
	}()
	return channel
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

/*
Using [WithCancel] you can create a Context that return a Context together with CancelFunc which it can be used to
tells an operation to abandon its work.

In this example we create some goroutines every time we run the function [generateChannel], using [range] for
we can iterate over the channels, here we have the limit of this infinite loop iteration in the body of the [for]
where in case we reach the 10th iteration, we just break out from the loop.

Then once the [cancel] function it's invoked, after return from the function, it will make that all goroutines
created that contain that context to finish, leaking any possible open resource.

Using [for] inside the [select] you can make run the loop forever, until some case just make return.
*/
func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for responseNumber := range runAsyncProgram(ctx) {
		if responseNumber == 10 {
			break
		}
	}
}

//Specifying [<-] before a channel return definition it means the channel can only receive data, but not send.
type returnChannel <-chan int

func runAsyncProgram(ctx context.Context) returnChannel {
	number := 1
	channel := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				println("Cleaning leaks")
				return // returning not to leak the goroutine
			case channel <- number:
				fmt.Println(number)
				number++
			}
		}
	}()
	return channel
}

/*
Using [WithDeadline] it's pretty much the same than Timeout, but instead use type Duration, you just use time.

Here we use in the same context, just to have a deadline in the execution of a async thread process, in case the
process exceed the deadline, the [Done] is invoked, and the whole process is stop and clean.

CancelFunc is not mandatory to execute, but it always a good practice add the defer Cancel() to clean all context after the function
ends.
*/
func TestContextDeadLine(t *testing.T) {
	deadlineTime := time.Now().Add(500 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadlineTime)
	defer cancel()
	channel := asyncWithDeadline(ctx)
	select {
	case response := <-channel:
		fmt.Println(response)
	case <-time.After(510 * time.Millisecond):
		fmt.Println("Seems the async action did not finish on time")
	}
}

func asyncWithDeadline(ctx context.Context) <-chan string {
	channel := make(chan string)
	select {
	case <-time.After(time.Duration(random(0, 1000)) * time.Millisecond):
		fmt.Println("Hello Async go with Deadline")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	return channel
}
