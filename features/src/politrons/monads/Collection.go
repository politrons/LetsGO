package monads

/**
We define our [CollectionMonad] interface for this new type to allow to have
[fold, map and flatMap] operators in collections.

*/
type CollectionMonad interface {
	/**
	[FoldLeft] to iterate from the left to right of the collection, apply the function where
	we accumulate the result and pass in each iteration of the collection.
	*/
	FoldLeft(interface{}, func(acc interface{}, next interface{}) interface{}) interface{}

	/**
	[FoldRight] to iterate from the right to left of the collection, apply the function where
	we accumulate the result and pass in each iteration of the collection.
	*/
	FoldRight(interface{}, func(acc interface{}, next interface{}) interface{}) interface{}

	/**
	[Map] to iterate the collection, apply the function where
	we pass the transform value, into a new array, that we return once we finish to iterate all elements.
	*/
	Map(func(acc interface{}) interface{}) interface{}

	/**
	[FlatMap] to iterate the collection, apply the function, which it return another collection, then we flat them
	passing the transform value, into a new array, that we return once we finish to iterate all elements.
	*/
	FlatMap(func(acc interface{}) []interface{}) interface{}
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
	for _, a := range collection {
		transformCollection = append(transformCollection, function(a))
	}
	return transformCollection
}

func (collection Collection) FlatMap(function func(b interface{}) []interface{}) interface{} {
	var transformCollection []interface{} = nil
	for _, a := range collection {
		for _, b := range function(a) {
			transformCollection = append(transformCollection, b)
		}
	}
	return transformCollection
}
