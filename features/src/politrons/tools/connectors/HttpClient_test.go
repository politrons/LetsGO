package connectors

/*
I have seen far by standing on the shoulders of giants.
Example implemented in top of ["github.com/gojektech/heimdall/httpclient"]
*/
import (
	"fmt"
	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/httpclient"
	"github.com/gojektech/heimdall/hystrix"
	"io/ioutil"
	"strconv"
	"testing"
	"time"
)

func TestClientRetryStrategy(t *testing.T) {
	for index := 10; index > 0; {
		runClientWithRetryStrategy(index)
		index--
	}
}

func TestClientCircuitBreaker(t *testing.T) {
	for index := 10; index > 0; {
		runClientWithCircuitBreaker(index)
		index--
	}
}

/*
First of all we create a retry strategy and we pass to the creation of the http client.
Then using the clients GET method to create and execute the request

Once we are able to return the response on time, Heimdall returns the standard *http.Response object,
ans using the deserializer operator [ioutil.ReadAll] we transform the [Body] into [[]byte]
*/
func runClientWithRetryStrategy(delay int) {
	delay = delay * 200
	retryStrategy := createRetryStrategy()
	client := createRetryStrategyClient(retryStrategy)
	response, err := client.Get("http://slowwly.robertomurray.co.uk/delay/"+strconv.Itoa(delay)+"/url/http://www.google.com", nil)
	if err != nil {
		println("Error response:", err.Error())
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}

/*
First of all we create a retry strategy and we pass to the creation of the client with Hystrix Circuit breaker.
Then using the clients GET method to create and execute the request

Having a Circuit breaker client we test how hystrix return timeouts, until open the Circuit breaker.

Once we are able to return the response on time, Heimdall returns the standard *http.Response object,
ans using the deserializer operator [ioutil.ReadAll] we transform the [Body] into [[]byte]

Special mention to test the delays the web page [http://slowwly.robertomurray.co.uk] really handy to test this delay case.
*/
func runClientWithCircuitBreaker(delay int) {
	delay = delay * 200
	retryStrategy := createRetryStrategy()
	client := createRetryStrategyCircuitBreakerClient(retryStrategy)
	response, err := client.Get("http://slowwly.robertomurray.co.uk/delay/"+strconv.Itoa(delay)+"/url/http://www.google.com", nil)
	if err != nil {
		println("Error response:", err.Error())
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}

/*
We create an Http Client with Retry strategy and Circuit breaker using [httpclient] package and [NewClient]
factory which accept N options which can be found in [httpclient] package.
Here we configure:
	[WithHTTPTimeout] having a max time for response.
	[WithCommandName] name of the circuit breaker command
	[WithHystrixTimeout] timeout of Circuit breaker
	[WithMaxConcurrentRequests] how many concurrent request we allow
	[WithErrorPercentThreshold] percentage of Error before we activate Circuit breaker
	[WithRetrier] having a retry strategy previously created.
	[WithRetryCount] having a number of times that we're going to retry in case of error.
*/
func createRetryStrategyCircuitBreakerClient(retryStrategy heimdall.Retriable) *hystrix.Client {
	return hystrix.NewClient(
		hystrix.WithCommandName("My custom request"),
		hystrix.WithHystrixTimeout(500),
		hystrix.WithMaxConcurrentRequests(5),
		hystrix.WithErrorPercentThreshold(1),
		hystrix.WithRetrier(retryStrategy),
		hystrix.WithRetryCount(2),
	)
}

/*
We create an Http Client with Retry strategy using [httpclient] package and [NewClient] factory which accept N options
which can be found in [httpclient] package.
Here we configure:
	[WithHTTPTimeout] having a max time for response.
	[WithRetrier] having a retry strategy previously created.
	[WithRetryCount] having a number of times that we're going to retry in case of error.
*/
func createRetryStrategyClient(retryStrategy heimdall.Retriable) *httpclient.Client {
	return httpclient.NewClient(
		httpclient.WithHTTPTimeout(1000*time.Millisecond),
		httpclient.WithRetrier(retryStrategy),
		httpclient.WithRetryCount(4),
	)
}

/*
We create an strategy for the http client using [NewRetrier] passing a [backoff] which it's created using
[NewConstantBackoff] where we set the Duration time for [backoffInterval] which means how much time do we
increase between retries, and also the max jitter interval.
*/
func createRetryStrategy() heimdall.Retriable {
	backoff := heimdall.NewConstantBackoff(2*time.Millisecond, 5*time.Millisecond)
	return heimdall.NewRetrier(backoff)
}
