package tools

import (
	"fmt"
	"testing"
	"time"
)

//In this strategy we login the user using username password using same [loginUserFunc] function contract
func TestStrategyLogin(t *testing.T) {
	//Load strategy
	applyStrategy(true)
	//Normal request
	response := loginUserFunc("politrons", "mypassword")
	println(response)
}

//In this strategy we login the user using token and login time using same [loginUserFunc] function contract
func TestStrategyToken(t *testing.T) {
	//Load strategy
	applyStrategy(false)
	//Normal request
	response := loginUserFunc("FDdsgasdg787sda87987sdfa909808SADF", time.Now().UnixNano()/int64(time.Millisecond))
	println(response)
}

/**
This pattern represent a function to be consumed that it might change the state to one from another dependening of an external factor,
that we can anticipate, so we can avoid perform a condition during the client request, improving performance.
Here the system can change the strategy changing one function implementation for other.
The main class it continue working since [loginUserFunc] it's a generic function that only need to respect, two generics inputs,
and one string as output.
*/
type LoginUserFunc = func(interface{}, interface{}) string

//We create this global variable that contains the strategy to apply.
var loginUserFunc LoginUserFunc

func applyStrategy(strategy bool) {
	if strategy {
		loginUserFunc = func(_username interface{}, _password interface{}) string {
			username := _username.(string)
			password := _password.(string)
			return fmt.Sprintf("User login with username %s and password %s ", username, password)
		}
	} else {
		loginUserFunc = func(_token interface{}, _time interface{}) string {
			token := _token.(string)
			loginTime := _time.(int64)
			return fmt.Sprintf("User login with token %s at loginTime %d ", token, loginTime)
		}
	}
}
