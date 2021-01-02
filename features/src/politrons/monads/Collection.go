package monads

/**
We define our [CollectionMonad] interface for this new type to allow to have
[Find, Filter, FoldLeft, FoldRight, Map and flatMap] operators in collections.

*/
type CollectionMonad interface {

	/**
	[Find] to iterate over of the collection, apply the predicate function
	and return return the first element of the collection that the function return true.
	*/
	Find(func(a interface{}) bool) interface{}

	/**
	[Filter] to iterate over of the collection, apply the predicate function
	and return a new collection with the elements that the function return true.
	*/
	Filter(func(a interface{}) bool) interface{}

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
	Map(func(a interface{}) interface{}) interface{}

	/**
	[FlatMap] to iterate the collection, apply the function, which it return another collection, then we flat them
	passing the transform value, into a new array, that we return once we finish to iterate all elements.
	*/
	FlatMap(func(a interface{}) []interface{}) interface{}
}

type Collection []interface{}

func (collection Collection) Find(function func(a interface{}) bool) interface{} {
	var element interface{} = nil
	for _, value := range collection {
		if function(value) {
			element = value
			break
		}
	}
	return element
}

func (collection Collection) Filter(function func(a interface{}) bool) interface{} {
	var newCollection []interface{} = nil
	for _, value := range collection {
		if function(value) {
			newCollection = append(newCollection, value)
		}
	}
	return newCollection
}

func (collection Collection) FoldLeft(
	init interface{},
	function func(next interface{}, acc interface{}) interface{}) interface{} {
	for _, value := range collection {
		init = function(init, value)
	}
	return init
}

func (collection Collection) FoldRight(
	init interface{},
	function func(acc interface{}, next interface{}) interface{}) interface{} {
	for i := len(collection) - 1; i >= 0; i-- {
		init = function(init, collection[i])
	}
	return init
}

func (collection Collection) Map(function func(b interface{}) interface{}) interface{} {
	var newCollection []interface{} = nil
	for _, a := range collection {
		newCollection = append(newCollection, function(a))
	}
	return newCollection
}

func (collection Collection) FlatMap(function func(b interface{}) []interface{}) interface{} {
	var newCollection []interface{} = nil
	for _, a := range collection {
		for _, b := range function(a) {
			newCollection = append(newCollection, b)
		}
	}
	return newCollection
}
