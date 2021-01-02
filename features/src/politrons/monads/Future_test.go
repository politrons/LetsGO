package monads

import (
	"testing"
)

func TestFutureSuccess(t *testing.T) {
	futureMonad := FutureSuccess{}.
		Create("Hello Future monad in Go").
		Map(upperCaseFunc).
		Map(func(i interface{}) interface{} {
			return i.(string) + "!!!!"
		}).
		FlatMap(func(i interface{}) Future {
			return FutureSuccess{}.Create(i.(string) + " In a new go routine")
		})
	println("In main routine context execution")
	println(futureMonad.Get().(string))
}
