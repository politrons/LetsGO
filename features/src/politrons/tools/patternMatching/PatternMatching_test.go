package patternMatching

import (
	"fmt"
	"politrons/monads"
	"reflect"
	"strings"
	"testing"
)

/*
Using [switch] expresion in Golang, we can have a very simple pattern matching implementation.
Here we show how we can match the type with the variant type of the Monads, as we would do in Scala or Haskell
*/
func TestPatternMatchingEither(t *testing.T) {
	switch monad := getEither("Hello pattern matching in Golang", true).(type) {
	case monads.Right:
		fmt.Println("Right side of the Monad Either", monad)
	case monads.Left:
		fmt.Println("Left side of the Monad Either", monad)
	}
}

func TestPatternMatchingMaybe(t *testing.T) {
	switch monad := getMaybe("Hello pattern matching in Golang").(type) {
	case monads.Just:
		fmt.Println("Monad Maybe has monad", monad)
	case monads.Nothing:
		fmt.Println("Monad Maybe is empty", monad)
	}
}

func TestPatternMatchingTry(t *testing.T) {
	switch monad := getTry("Hello pattern matching in Golang").(type) {
	case monads.Success:
		fmt.Println("Monad Try finish successful with monad", monad)
	case monads.Failure:
		fmt.Println("Monad Try finish wrong with error", monad)
	}
}

/*
Here using our own type and extended method, we can create this pseudo Pattern matching where
we can match the initial value passed in the [Match] with every operation with the type passed in
the [Case]. Using reflect, if the type match we execute the function with the init value passed into the Match.
*/
func TestOwnPatternMatchingTrySuccess(t *testing.T) {
	value := Match{getTry("Hello pattern matching Try in Golang")}.
		Case(monads.Success{}, func(i interface{}) interface{} {
			return strings.ToUpper(i.(monads.Success).Value.(string))
		}).
		Case(monads.Failure{}, func(i interface{}) interface{} {
			return i.(error).Error() + " append extra error"
		}).Value
	println(value.(string))
}

func TestOwnPatternMatchingEither(t *testing.T) {
	value := Match{getEither("Hello pattern matching Either in Golang", true)}.
		Case(monads.Right{}, func(i interface{}) interface{} {
			return strings.ToUpper(i.(monads.Right).Value.(string))
		}).
		Case(monads.Left{}, func(i interface{}) interface{} {
			return i.(monads.Left).Value.(string) + " append extra info of Left"
		}).Value
	println(value.(string))
}

func TestOwnPatternMatchingMaybe(t *testing.T) {
	value := Match{getMaybe("Hello pattern matching Maybe in Golang")}.
		Case(monads.Just{}, func(i interface{}) interface{} {
			return strings.ToUpper(i.(monads.Just).Value.(string))
		}).
		Case(monads.Nothing{}, func(i interface{}) interface{} {
			return "Default data"
		}).Value
	println(value.(string))
}

//Type of the Pattern matching
type Match struct {
	Value interface{}
}

/*
Implementation of the Case to check if the Match bind with the interface type passed.
In case that match we execute the function passed, passing the value of the Match value, and
we wrap the result of the function into another Match.
*/
func (m Match) Case(i interface{}, fn func(interface{}) interface{}) Match {
	hasSameType := reflect.TypeOf(m.Value) == reflect.TypeOf(i)
	if hasSameType {
		return Match{fn(m.Value)}
	}
	return m
}

// Utils functions
//--------------------

func getEither(value string, right bool) monads.Either {
	if right {
		return monads.Right{value}
	} else {
		return monads.Left{value}
	}
}

func getMaybe(value string) monads.Maybe {
	if value != "" {
		return monads.Just{value}
	} else {
		return monads.Nothing{}
	}
}

func getTry(i interface{}) monads.Try {
	switch value := i.(type) {
	case string:
		return monads.Success{value}
	case error:
		return monads.Failure{value}
	default:
		panic("Not controlled option")
	}
}
