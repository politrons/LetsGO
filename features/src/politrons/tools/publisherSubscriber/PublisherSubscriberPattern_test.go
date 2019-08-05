package publisherSubscriber

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestPublisherSubscriberPattern(t *testing.T) {
	publisher := Publisher{"1000", make(chan Event)}

	index := 0
	go func() {
		for index < 100 {
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

type Event struct {
	value string
}

type Publisher struct {
	id      string
	channel chan Event
}

type Observable struct {
	Id        string
	Publisher Publisher
}

type Subscriber struct {
	doOnNext     func(event Event)
	doOnError    func(err error)
	doOnComplete func()
}

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
		time.Sleep(200 * time.Millisecond) // business logic delay (Connection with other backend and so on)
	}
}

func (error NoMoreEvents) Error() string {
	return error.Cause
}
