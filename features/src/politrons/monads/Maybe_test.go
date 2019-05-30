package monads

import (
	"strings"
	"testing"
)

/*
Go cannot escape of FP, it's a powerful tool, and even having to do it without lambdas the concept of
transformation and composition can being applyied.
Here applying interfarces type with method extensions allow us that Maybe interface can have variant
[Just] or [Nothing]. This approach works pretty similar like type classes of Haskell.
*/
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
