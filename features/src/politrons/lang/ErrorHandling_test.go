package lang

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
Go instead of return a throwable in case of error it make explicit returning in case of effects of throwble a tuple
with convention (value, error), a poor choice since the right value normally it's the "right" value.
*/
func TestError(t *testing.T) {
	result, err := runLogic()
	if err != nil {
		fmt.Println(err)
	} else {
		println(result)
	}
}

type MyCustomError struct {
	Time  time.Time
	Cause string
}

/*
In order to define a type of interface [error] we need to make the [type] implement method interface
[Error()], then it will automatically candidate to be used by error system of GO.
*/
func (e MyCustomError) Error() string {
	return fmt.Sprintf("at %v, %s", e.Time, e.Cause)
}

/*
This function it has a similar type return a tuple(value, error) which it would be a very rudimentary
[Either[T,R]] here we return as error the [error] type used by GO.
*/
func runLogic() (string, error) {
	if randFunc() {
		return "", MyCustomError{time.Now(), "Error found in logic"}
	} else {
		return "Success result", nil
	}
}

func randFunc() bool {
	return rand.Intn(2) == 0
}
