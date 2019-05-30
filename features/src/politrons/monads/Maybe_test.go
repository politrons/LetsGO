package monads

import (
	"strings"
	"testing"
)

func TestMaybeJust(t *testing.T) {
	maybe := getMaybe("Hello Maybe in Golang")
	println(maybe.isDefined())
	println(maybe.Get().(string))
	result := maybe.
		Map(func(i interface{}) interface{} {
			return strings.ToUpper(i.(string))
		}).
		FlatMap(func(i interface{}) Maybe {
			return Just{i.(string) + "!!!"}
		}).
		Get().(string)
	println(result)
}

func TestMaybeNothing(t *testing.T) {
	maybe := getMaybe("")
	println(maybe.isDefined())
	println(maybe.Get())
}

func getMaybe(value string) Maybe {
	if value != "" {
		return Just{value}
	} else {
		return Nothing{}
	}
}
