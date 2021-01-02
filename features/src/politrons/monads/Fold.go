package monads

type Fold interface {
	FoldLeft(interface{}, func(acc interface{}, next interface{}) interface{}) interface{}

	FoldRight(interface{}, func(acc interface{}, next interface{}) interface{}) interface{}
}

type IntCollection []int

type StringCollection []string

func (collection IntCollection) FoldLeft(
	zero interface{},
	function func(next interface{}, acc interface{}) interface{}) interface{} {
	return collection.processIntFunction(zero, function)
}

func (collection IntCollection) FoldRight(
	zero interface{},
	function func(acc interface{}, next interface{}) interface{}) interface{} {
	return collection.processIntFunction(zero, function)
}

func (collection IntCollection) processIntFunction(
	zero interface{},
	function func(next interface{}, acc interface{}) interface{}) int {
	var init = zero.(int)
	for _, value := range collection {
		init = function(init, value).(int)
	}
	return init
}

func (collection StringCollection) FoldLeft(
	zero interface{},
	function func(next interface{}, acc interface{}) interface{}) interface{} {
	return collection.processStringFunction(zero, function)
}

func (collection StringCollection) FoldRight(
	zero interface{},
	function func(acc interface{}, next interface{}) interface{}) interface{} {
	return collection.processStringFunction(zero, function)
}

func (collection StringCollection) processStringFunction(
	zero interface{},
	function func(next interface{}, acc interface{}) interface{}) string {
	var init = zero.(string)
	for _, value := range collection {
		init = function(init, value).(string)
	}
	return init
}
