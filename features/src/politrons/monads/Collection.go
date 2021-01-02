package monads

/**
We define our interface for this new type to allow have fold operators in collections.
FoldLeft to iterate from the left to the right
*/
type CollectionI interface {
	FoldLeft(interface{}, func(acc interface{}, next interface{}) interface{}) interface{}

	FoldRight(interface{}, func(acc interface{}, next interface{}) interface{}) interface{}

	Map(func(acc interface{}) interface{}) interface{}
}

type Collection []interface{}

func (collection Collection) FoldLeft(
	zero interface{},
	function func(next interface{}, acc interface{}) interface{}) interface{} {
	var init = zero
	for _, value := range collection {
		init = function(init, value)
	}
	return init
}

func (collection Collection) FoldRight(
	zero interface{},
	function func(acc interface{}, next interface{}) interface{}) interface{} {
	var init = zero
	for i := len(collection) - 1; i >= 0; i-- {
		init = function(init, collection[i])
	}
	return init
}

func (collection Collection) Map(function func(b interface{}) interface{}) interface{} {
	var transformCollection []interface{} = nil
	for _, value := range collection {
		transformCollection = append(transformCollection, function(value))
	}
	return transformCollection
}
