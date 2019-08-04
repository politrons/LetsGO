package observer

import (
	"fmt"
	"testing"
)

//##########################//
//	OBSERVER PATTERN  //
//##########################//

/**
Observer pattern divide in two parts, one is the type of the [Observer] which basically is an instance that want to
subscribe to the second part, the type [Channel].

When we create a new instance of the observer we provide an [id] of this instance which it will be used as key in the channel later
to subscribe and unsubscribe.

Then we create the instance [Channel] which basically contains an [id] of the channel and the map of [observers] subscribed to that channel

*/
func TestObserverPattern(t *testing.T) {
	observer1 := Observer{"1000"}
	observer2 := Observer{"2000"}
	observer3 := Observer{"3000"}

	channel1 := Channel{"1981", make(map[string]Observer)}
	channel2 := Channel{"2981", make(map[string]Observer)}

	channel1.subscribe(observer1, observer2, observer3)
	channel1.unsubscribe(observer2)
	channel1.sendEvent(Event{channel1.id, "Event1: Hello Observer pattern world"})

	channel2.subscribe(observer2)
	channel2.sendEvent(
		Event{channel2.id, "Event2: Hello Observer pattern world"},
		Event{channel2.id, "Event3: Hello again buddy!"},
		Event{channel2.id, "Event4: Last message promise"})

}

/*
Type that define the event that it will be used in communications between channel and observers.
It contains the id of the channel since the event it could arrive from different channels where
an observer might be subscribed, and the message,
*/
type Event struct {
	channelId string
	value     string
}

//Simple Observer type with the id to identify his instance.
type Observer struct {
	id string
}

/*
Channel type with the id to identify his instance, and then map of observers to make
broadcast later as we will see.
*/
type Channel struct {
	id        string
	observers map[string]Observer
}

/**
The API of the Observer, the only function that must implement is the callback function to be invoked from
the channel, when a new [Event] arrive.
*/
type ObserverAPI interface {
	notify(event Event)
}

/**
The API of Channel, it contains three functions:
	subscribe: It receive 0 to N [Observer] to add in the map that Channel type contains
	unsubscribe: It receive 0 to N [Observer] to delete in the map that Channel type contains
	sendEvent: It receive 0 to N [Events] to broadcast to all observers in the map subscribed.
*/
type ChannelAPI interface {
	subscribe(observer ...Observer)
	unsubscribe(observer ...Observer)
	sendEvent(event ...Event)
}

//Implementation of callback [notify] when a new Event arrive.
func (observer Observer) notify(event Event) {
	fmt.Printf("Observer %s: New event received from channel %s with message %s \n", observer.id, event.channelId, event.value)
}

/*
Implementation of API, we check if the map of observers, if does not exist we created, and then we
iterate over the array of observers to subscribe all of them.
*/
func (channel Channel) subscribe(observers ...Observer) {
	channel.checkChannelMap()
	for _, observer := range observers {
		channel.observers[observer.id] = observer
		fmt.Printf("New observer %s subscribed in channel %s \n", observer.id, channel.id)
	}
}

/*
Implementation of API, we check if the map of observers, if does not exist we created, and then we
iterate over the array of observers to unsubscribe all of them.
*/
func (channel Channel) unsubscribe(observers ...Observer) {
	channel.checkChannelMap()
	for _, observer := range observers {
		delete(channel.observers, observer.id)
		fmt.Printf("Observer %s unsubscribed from channel %s \n", observer.id, channel.id)
	}
}

/*
Implementation of API, we iterate over all array of events to be send, and for each
we iterate over the array of observers to call [notify] passing the [Event] to all of them.
*/
func (channel Channel) sendEvent(events ...Event) {
	fmt.Printf("Boradcasting event to all observers from channel %s \n", channel.id)
	for _, event := range events {
		for _, observer := range channel.observers {
			observer.notify(event)
		}
	}
}

func (channel Channel) checkChannelMap() {
	if channel.observers == nil {
		channel.observers = make(map[string]Observer)
	}
}
