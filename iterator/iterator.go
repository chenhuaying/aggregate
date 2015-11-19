package iterator

type Iterator interface {
	Next()
	Prev()
	// the index of an array or key of a map, like c++
	First() interface{}
	// the value of the container
	Second() interface{}
	End() bool
}
