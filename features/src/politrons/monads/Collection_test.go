package monads

import (
	"fmt"
	"strings"
	"testing"
)

func TestIntFoldRightMonad(t *testing.T) {
	total := Collection{1, 2, 3, 4, 5}.
		FoldRight(0, func(acc interface{}, next interface{}) interface{} {
			return acc.(int) + next.(int)
		})
	println(total.(int))
}

func TestIntFoldLeftMonad(t *testing.T) {
	total := Collection{1, 2, 3}.
		FoldLeft(0, func(acc interface{}, next interface{}) interface{} {
			return acc.(int) + next.(int)
		})
	println(total.(int))
}

func TestStringFoldRightMonad(t *testing.T) {
	total := Collection{"hello", "world"}.
		FoldRight("", func(acc interface{}, next interface{}) interface{} {
			return acc.(string) + " " + next.(string)
		})
	println(total.(string))
}

func TestStringMapMonad(t *testing.T) {
	words := Collection{"hello", "Monads", "world", "!!!"}.
		Map(func(value interface{}) interface{} {
			return strings.ToUpper(value.(string))
		})
	fmt.Printf("%v", words.([]interface{}))
}
