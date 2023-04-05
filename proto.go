package cart

type It[T any] interface {
	Next() T
	Peek() T
	HasNext() bool
}
