package patternMatching

import (
	"fmt"
	. "politrons/monads"
	"testing"
)

/*
Using [switch] expresion in Golang, we can have a very simple pattern matching implementation.
Here we show how we can match the type with the variant type of the Monads, as we would do in Scala or Haskell
*/
func TestPatternMatchingEither(t *testing.T) {
	switch monad := getEither("Hello pattern matching in Golang", true).(type) {
	case Right:
		fmt.Println("Right side of the Monad Either", monad)
	case Left:
		fmt.Println("Left side of the Monad Either", monad)
	}
}

func TestPatternMatchingMaybe(t *testing.T) {
	switch monad := getMaybe("Hello pattern matching in Golang").(type) {
	case Just:
		fmt.Println("Monad Maybe has monad", monad)
	case Nothing:
		fmt.Println("Monad Maybe is empty", monad)
	}
}

func TestPatternMatchingTry(t *testing.T) {
	switch monad := getTry("Hello pattern matching in Golang").(type) {
	case Success:
		fmt.Println("Monad Try finish successful with monad", monad)
	case Failure:
		fmt.Println("Monad Try finish wrong with error", monad)
	}
}

func getEither(value string, right bool) Either {
	if right {
		return Right{value}
	} else {
		return Left{value}
	}
}

func getMaybe(value string) Maybe {
	if value != "" {
		return Just{value}
	} else {
		return Nothing{}
	}
}

func getTry(i interface{}) Try {
	switch value := i.(type) {
	case string:
		return Success{value}
	case error:
		return Failure{value}
	default:
		panic("Not controlled option")
	}
}
