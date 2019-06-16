package failFast

//All this implementation has been build in top of https://github.com/sony/gobreaker
import (
	"fmt"
	"github.com/sony/gobreaker"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

//Global Circuit breaker variable
var cb *gobreaker.CircuitBreaker

/*
A fun fact, in Golang a function with name [init()] it will be automatically invoked.
The order of execution in package initialization is:
* Initialization of imported packages
* Declaration of vars.
* Invocation of init() function.

[Name] is the name of the CircuitBreaker.
[ReadyToTrip] bool function tp determine if the Circuit breaker must be open.
[OnStateChange] Callback function to invoke when the Circuit breaker change state.
[MaxRequests] Maximum number of requests allowed to pass through when the CircuitBreaker is half-open
[Timeout] Period of the open state, after which the state of CircuitBreaker becomes half-open.

[NewCircuitBreaker] create the Circuit breaker instance.
*/
func init() {
	var settings gobreaker.Settings
	settings.Name = "Politrons"
	settings.OnStateChange = onChangeStateFunc
	settings.ReadyToTrip = maxConsecutiveFailuresStrategyFunc
	settings.MaxRequests = 1
	settings.Timeout = time.Duration(2000 * time.Millisecond)
	cb = gobreaker.NewCircuitBreaker(settings)
}

/*
Function to be invoked when the state of the Circuit breaker change.
We receive the name of the circuit breaker, previous and new state of the CB.
*/
func onChangeStateFunc(name string, from gobreaker.State, to gobreaker.State) {
	fmt.Printf("Circuit breaker %s change state from %s to %s \n", name, from.String(), to.String())
}

/*
Function to determine if the circuit breaker has to be open.
In this case we determine that, if the number of request are more than 5(Warm up possible failures)
And the failure ratio is equal or higher of 60% then we should open the fail fast mechanism.
*/
func failureRatioStrategyFunc(counts gobreaker.Counts) bool {
	failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
	return counts.Requests >= 5 && failureRatio >= 0.6
}

/*
Here we implement another function strategy for the Circuit breaker
*/
func maxConsecutiveFailuresStrategyFunc(counts gobreaker.Counts) bool {
	return counts.Requests >= 5 && counts.ConsecutiveFailures >= 5
}

/*
For this example we implement a Http client using [net/http] package of Go SDK.

For the circuit breaker, like many implementations, we need to wrap the execution of computation
with the extended function [Execute] where we need to implement an anonymous function with no inputs
and tuple of (interface{}, error) as output.

Now inside we put the request call, and whatever type of error that we return in the tuple, it will
be counted, and once the CB strategy return true, it wont execute the function but it will return a fail
fast error, with the text "circuit breaker is open".

Also since we configure a [OnStateChange] every time the Circuit breaker change state, it will invoke also
that callback function.
*/
func RequestWithCircuitBreaker(url string) ([]byte, error) {
	responseBody, err := cb.Execute(func() (interface{}, error) {
		return makeGetRequest(url)
	})
	if err != nil {
		return nil, err
	}
	return responseBody.([]byte), nil
}

func makeGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

/*
In this test we check how we can pass from succeed request into failure, and at some point the
[maxConsecutiveFailuresStrategyFunc] start taking effect and we return fail fast.
*/
func TestCircuitBreakerOpenAfterSomeSuccess(t *testing.T) {
	//Success request
	MakeRequestLoop("http://www.google.com/search/about")
	//Failure request
	MakeRequestLoop("http://www.notWorkingForSure.com")
}

/*
In this test we check how we can pass from failure request into succeed, and at some point the
[maxConsecutiveFailuresStrategyFunc] and [Timeout] make the circuit breaker back to [half-open]
and from there into [close] state.
*/
func TestCircuitBreakerOpenAndClose(t *testing.T) {
	//Failure request
	MakeRequestLoop("http://www.notWorkingForSure.com")
	//Success request
	MakeRequestLoop("http://www.google.com/search/about")
}

func MakeRequestLoop(url string) {
	for count := 0; count < 15; {
		_, err := RequestWithCircuitBreaker(url)
		if err != nil {
			println(err.Error())
		} else {
			println("Request succeed")
		}
		count++
		time.Sleep(500 * time.Millisecond)
	}
}
