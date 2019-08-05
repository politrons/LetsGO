package publisherSubscriber

import (
	"fmt"
	"testing"
	"time"
)

func TestPublisherSubscriberPattern(t *testing.T) {
	events := []Event{{"Hello"}, {"Publisher"}, {"subscriber"}, {"world"}}
	publisher := Publisher{"1000", make(chan Event), events}
	/*	newPublisher := publisher.appendEvent(Event{"last message"})
	 */
	index := 0
	go func() {
		for index < 100 {
			publisher.appendEvent(Event{"message:" + string(index)})
			index++
			time.Sleep(500 * time.Millisecond)
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
	id         string
	channel    chan Event
	eventQueue []Event
}

type Observable struct {
	Id        string
	Publisher Publisher
}

type Subscriber struct {
	Offset       int
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
	/*	publisher.eventQueue = append(publisher.eventQueue, event)
		return publisher*/
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
	/*if len(publisher.eventQueue) > subscriber.Offset {
		event := publisher.eventQueue[subscriber.Offset]
		return event, nil
	} else {
		return Event{}, NoMoreEvents{"No more events from publisher"}
	}*/
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
	subscriber := Subscriber{0, onNext, onError, onComplete}
	//For provide an infinite loop
	for {
		event, err := observable.Publisher.getNext(subscriber)
		if err != nil {
			subscriber.doOnError(err)
			subscriber.doOnComplete()
			break
		} else {
			subscriber.doOnNext(event)
			subscriber.Offset += 1 //Increase offset
		}
	}
}

func (error NoMoreEvents) Error() string {
	return error.Cause
}

func (publisher Publisher) checkMessageQueue() {
	if publisher.eventQueue == nil {
		publisher.eventQueue = []Event{}
	}
}
