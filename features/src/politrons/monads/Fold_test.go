package monads

import (
	"testing"
)

func TestIntFoldRightMonad(t *testing.T) {
	total := IntCollection{1, 2, 3, 4, 5}.FoldRight(0, func(acc interface{}, next interface{}) interface{} {
		return acc.(int) + next.(int)
	})
	println(total.(int))
}

func TestIntFoldLeftMonad(t *testing.T) {
	total := IntCollection{1, 2, 3}.FoldLeft(0, func(acc interface{}, next interface{}) interface{} {
		return acc.(int) + next.(int)
	})
	println(total.(int))
}

func TestStringFoldRightMonad(t *testing.T) {
	total := StringCollection{"hello", "world"}.FoldRight("", func(acc interface{}, next interface{}) interface{} {
		return acc.(string) + " " + next.(string)
	})
	println(total.(string))
}
