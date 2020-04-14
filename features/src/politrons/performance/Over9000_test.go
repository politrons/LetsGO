package performance

//All this test are done in top of this awesome library https://github.com/tsenart/vegeta
//Once we import the [vegeta] library we can specify an alias to detect in our code, when
// we invoke something related with this library
import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/lib"
	"testing"
	"time"
)

/*
This test using the type and interface that we create, we invoke the implementation
[PerformanceRequestInfo] which use internally [vegeta] DSL to run a performance process.
The [RunVegetaPerformance] return a vegeta [Metrics] type with all the information of the
performance process response.
*/
func TestVegetaHttp(t *testing.T) {

	warmUp := PerformanceRequestInfo{
		numberOfRequest: 10,
		duration:        5,
		method:          "GET",
		url:             "http://google.com/",
		name:            "Big bang attack!",
	}

	warmUp.RunVegetaPerformance()

	metrics := PerformanceRequestInfo{
		numberOfRequest: 100,
		duration:        10,
		method:          "GET",
		url:             "http://google.com/",
		name:            "Big bang attack!",
	}.RunVegetaPerformance()

	/*
		Telemetry
		----------
		Vegeta library provide [Metrics] type, in which DSL has all the typical values that you must control in your
		endpoints performance test.(avg, max, min, percentiles 50,90, 95, 99)
	*/
	fmt.Println("Total request:", metrics.Requests)
	fmt.Println("Min latency:", metrics.Latencies.Min)
	fmt.Println("Max latency:", metrics.Latencies.Max)
	fmt.Println("Mean latency:", metrics.Latencies.Mean)
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Println("Success percentile:", metrics.Success)
	fmt.Println("Error percentile:", metrics.Errors)

}

/*
Data type to contain all the info that Vegeta DSL needs
*/
type PerformanceRequestInfo struct {
	numberOfRequest int
	duration        int
	method          string
	url             string
	name            string
}

/*
We implement an interface where the implementation must return the metrics of
Vegeta library
*/
type VegetaInterface interface {
	RunVegetaPerformance() vegeta.Metrics
}

/*
Implementation of [VegetaInterface] that we create where passing the [PerformanceRequestInfo] type
we can use the [vegeta] DSL to perform a test.

Use
----
Vegeta DSL require several parts to fill the response values [Metrics] of the process that run.
- We create a [Rate] which is how many request we will perform by time.
- We create a [Targeter] which contains the typical request info (method, url, headers, body)
- We create a [Duration] to specify how much time the performance process must take
- We create an [Attacker] which is the client that run the request using all the information we provide before.
- We create a [Metrics] which we will use to fill with the metrics info obtained by the [Result] of the Attacker client.

Once we have the DSL formed, we use [Attack] function passing all the arguments created before, and it will return a Response
value per each request which we will add into the [Metrics] type created previously.
This Metrics type it will generate all the telemetry info of all request info together.
*/
func (requestInfo PerformanceRequestInfo) RunVegetaPerformance() vegeta.Metrics {
	rate := requestInfo.createRate()
	targeter := requestInfo.createTargeter()
	duration := requestInfo.createDuration()
	metrics := requestInfo.runPerformanceAttack(vegeta.NewAttacker(), targeter, rate, duration)
	return metrics
}

func (requestInfo PerformanceRequestInfo) runPerformanceAttack(
	attacker *vegeta.Attacker,
	targeter vegeta.Targeter,
	rate vegeta.Rate,
	duration time.Duration) vegeta.Metrics {
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, requestInfo.name) {
		metrics.Add(res)
	}
	metrics.Close()
	return metrics
}

func (requestInfo PerformanceRequestInfo) createDuration() time.Duration {
	return time.Duration(requestInfo.duration) * time.Second
}

func (requestInfo PerformanceRequestInfo) createTargeter() vegeta.Targeter {
	return vegeta.NewStaticTargeter(vegeta.Target{
		Method: requestInfo.method,
		URL:    requestInfo.url,
	})
}

func (requestInfo PerformanceRequestInfo) createRate() vegeta.Rate {
	return vegeta.Rate{Freq: requestInfo.numberOfRequest, Per: time.Second}
}
