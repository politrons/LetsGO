package publisherSubscriber

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//##################################//
//PUBLISHER/SUBSCRIBER PATTERN  //
//##################################//

/**
Publisher subscriber pattern is also know as event-driven, is a new reactive way to communicate between
[producer/publisher] and [consumer/subscriber].

The pattern it was design due of the need to solve one of the most commons problems in backend, (Out of memory) basically overwhelm a backend
with messages when it's not possible to process those messages. This need brings the concept of [Back-pressure] which basically solve this problem
changing the old fashion [iterable] by [pull] in the communications. Where instead the publisher send all the message that has, they wait until
the client make pull and tell him that he is ready to receive more messages.

The design it divide in two parts, the [publisher], [observable] and [subscriber].
*/

/**
For this test, we create a [Publisher] and we start appending Events in the [channel] of Events with a production of 100 ms (Higher than what consumer can process)

Then we create an [Observable] using [Just] function, passing the publisher and next we [Subscribe] passing three functions that they will be invoked by the Observable,
for a specific state.
Those functions are:
	[doOnNext(event Event)]: A consumer function that receive the event passed from the publisher to do something.
	[doOnError(error Error)]: A consumer function that receive an error in the emission and process of one of the events
	[doOnComplete()]: A function with no arguments that it's invoked when there's no more events from the publisher and the emissions is finish.
*/
func TestPublisherSubscriberPattern(t *testing.T) {
	publisher := Publisher{"1000", make(chan Event)}

	index := 0
	go func() {
		for index < 10 {
			publisher.appendEvent(Event{"message:" + strconv.Itoa(index)})
			index++
			time.Sleep(100 * time.Millisecond)
		}
	}()

	new(Observable).Just(publisher).
		Map(func(i interface{}) interface{} {
			return "[" + i.(Event).value + "]"
		}).
		Map(func(i interface{}) interface{} {
			return "**" + i.(string) + "**"
		}).
		Subscribe(
			func(value interface{}) {
				fmt.Printf("New event recieved %s \n", value.(string))
			}, func(e error) {
				fmt.Printf("Error in pipeline %e \n", e)
			}, func() {
				fmt.Println("Pipeline complete")
			})
}

//Event type passed between publisher/subscriber
type Event struct {
	value string
}

//Type that receive events and pass to the subscriber when he ask for.
type Publisher struct {
	id      string
	channel chan Event
}

//Observable type that contains the publisher and create the [Subscriber]
type Observable struct {
	Id        string
	doOnMaps  []func(value interface{}) interface{}
	Publisher Publisher
}

//Subscriber type that contains the functions to execute once we receive an event, error or end of emission.
type Subscriber struct {
	doOnNext     func(value interface{})
	doOnError    func(err error)
	doOnComplete func()
}

//Error type to specify the emission has finish
type NoMoreEvents struct {
	Cause string
}

//#################//
//	PUBLISHER   //
//#################//

//Simple async function that write the received event into the channel
func (publisher Publisher) appendEvent(event Event) {
	go func() { publisher.channel <- event }()
}

/*
This function is invoked by the observable when he can process new messages.
It will try to read from the channel using [for] we create this infinite loop and using [select]
We're able to Subscribe to the channel waiting for the element, and also we define a max wait time [2000]
where after that we return the error that no more events on channel found.
*/
func (publisher Publisher) getNext(subscriber Subscriber) (Event, error) {
	for {
		select {
		case event := <-publisher.channel:
			return event, nil
		case <-time.After(time.Duration(2000 * time.Millisecond)):
			return Event{}, NoMoreEvents{"No more events from publisher"}

		}
	}
}

//#################//
//	OBSERVABLE   //
//#################//

//Simple factory function to create the [Observable] with the [Publisher] inside.
func (observable Observable) Just(publisher Publisher) Observable {
	observable.Publisher = publisher
	return observable
}

func (observable Observable) Map(mapFunc func(event interface{}) interface{}) Observable {
	observable.doOnMaps = append(observable.doOnMaps, mapFunc)
	return observable
}

/*
Function to create [Subscriber] with the functions we receive, and then in a infinite [for] loop
we ask to the publisher for the next message, which return a tuple of (Event, Error).

In case the tuple Error is nil, we iterate over the array of functions [doOnMap], and we invoke the [doOnMap] to transform the event,
and then [doOnNext] function of the subscriber passing the transformed event.
Otherwise we invoke the [doOnError] and [doOnComplete] to finish the emission of the stream.

Finally to prove the [Back-pressure] concept here, we make delay of [500 ms] which represent some business logic delay,
which is higher than the production of events in the producer [100 ms] proving than is the subscriber the one
that mark the speed of consumption in the stream.
*/
func (observable Observable) Subscribe(onNext func(value interface{}), onError func(error), onComplete func()) {
	subscriber := Subscriber{onNext, onError, onComplete}
	for {
		event, err := observable.Publisher.getNext(subscriber)
		var transformedValue interface{} = event
		if err == nil {
			for _, doOnMap := range observable.doOnMaps {
				transformedValue = doOnMap(transformedValue)
			}
			subscriber.doOnNext(transformedValue)
		} else {
			subscriber.doOnError(err)
			subscriber.doOnComplete()
			break
		}
		time.Sleep(500 * time.Millisecond) // business logic delay (Connection with other backend and so on)
	}
}

func (error NoMoreEvents) Error() string {
	return error.Cause
}
