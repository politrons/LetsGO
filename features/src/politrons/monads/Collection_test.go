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

func TestMapMonad(t *testing.T) {
	words := Collection{"hello", "Monads", "world", "!!!"}.
		Map(func(value interface{}) interface{} {
			return strings.ToUpper(value.(string))
		})
	fmt.Printf("%v", words.([]interface{}))
}

func TestFlatMapMonad(t *testing.T) {
	total := Collection{1, 2, 3, 4, 5}.
		FlatMap(func(value interface{}) []interface{} {
			return []interface{}{value.(int) * 100}
		})
	fmt.Printf("%v", total.([]interface{}))
}

func TestIntFilterMonad(t *testing.T) {
	total := Collection{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}.
		Filter(func(value interface{}) bool {
			return value.(int) > 5
		})
	fmt.Printf("%v", total)
}

func TestIntFindMonad(t *testing.T) {
	total := Collection{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}.
		Find(func(value interface{}) bool {
			return value.(int) > 5
		})
	fmt.Printf("%v", total.(interface{}))
}

func TestIntUntilMonad(t *testing.T) {
	total := Collection{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}.
		Until(func(value interface{}) bool {
			return value.(int) <= 5
		})
	fmt.Printf("%v", total)
}

func TestIntTakeMonad(t *testing.T) {
	total := Collection{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}.
		Take(5)
	fmt.Printf("%v", total)
}
