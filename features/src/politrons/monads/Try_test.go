package monads

import (
	"strings"
	"testing"
)

func TestSucceedMonad(t *testing.T) {
	tryMonad := getTry("Hello Try monad in Go").
		Map(func(i interface{}) interface{} {
			return strings.ToUpper(i.(string))
		}).
		Map(func(i interface{}) interface{} {
			return i.(string) + "!!!!!"
		})
	println(tryMonad.isSuccess())
	println(tryMonad.isFailure())
	response := tryMonad.Get().(string)
	println(response)
}

func TestFailureMonad(t *testing.T) {
	tryMonad := getTry(MyError{"Custom error in Go"}).
		MapError(func(e error) error {
			return MyError{e.Error() + " Append extra error info"}
		})
	println(tryMonad.isSuccess())
	println(tryMonad.isFailure())
	response := tryMonad.Get().(error)
	println(response.Error())
}

//######################
//#   Utils functions  #
//######################
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

type MyError struct {
	cause string
}

func (e MyError) Error() string {
	return e.cause
}
