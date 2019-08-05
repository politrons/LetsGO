package publisherSubscriber

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//##################################//
//	PUBLISHER/SUBSCRIBER PATTERN  //
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

Then we create an [Observable] using [just] function, passing the publisher and next we [subscribe] passing three functions that they will be invoked by the Observable,
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

	new(Observable).just(publisher).
		subscribe(
			func(event Event) {
				fmt.Printf("New event reveiced %s \n", event.value)
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
	Publisher Publisher
}

//Subscriber type that contains the functions to execute once we receive an event, error or end of emission.
type Subscriber struct {
	doOnNext     func(event Event)
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

func (publisher Publisher) appendEvent(event Event) {
	go func() { publisher.channel <- event }()
}

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

func (observable Observable) just(publisher Publisher) Observable {
	observable.Publisher = publisher
	return observable
}

//Function to subscribe [Subscriber]
func (observable Observable) subscribe(onNext func(event Event), onError func(error), onComplete func()) {
	//Create subscriber
	subscriber := Subscriber{onNext, onError, onComplete}
	//For provide an infinite loop
	for {
		event, err := observable.Publisher.getNext(subscriber)
		if err != nil {
			subscriber.doOnError(err)
			subscriber.doOnComplete()
			break
		} else {
			subscriber.doOnNext(event)
		}
		time.Sleep(500 * time.Millisecond) // business logic delay (Connection with other backend and so on)
	}
}

func (error NoMoreEvents) Error() string {
	return error.Cause
}
