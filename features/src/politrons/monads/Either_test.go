package monads

import (
	"testing"
)

func TestRightMonad(t *testing.T) {
	eitherMonad := getEither("Hello Either Right monad in Go", true).
		Map(upperFunc).
		Map(appendFunc)
	println("Right:", eitherMonad.IsRight())
	println("Left:", eitherMonad.IsLeft())
	response := eitherMonad.Get().(string)
	println(response)
}

func TestLeftMonad(t *testing.T) {
	eitherMonad := getEither("Hello Either Left monad in Go", false).
		Map(upperFunc).
		Map(appendFunc)
	println("Right:", eitherMonad.IsRight())
	println("Left:", eitherMonad.IsLeft())
	response := eitherMonad.Get().(string)
	println(response)
}

//######################
//#   Utils functions  #
//######################
func getEither(value string, right bool) Either {
	if right {
		return Right{value}
	} else {
		return Left{value}
	}
}
