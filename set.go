package lockfree

type Set[T any] struct {
}

func NewSet[T any](less func(T, T) bool) *Set[T] {
	panic("not implemented")
}
