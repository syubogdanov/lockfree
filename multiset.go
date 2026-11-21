package lockfree

type Multiset[T any] struct {
}

func NewMultiset[T any](less func(T, T) bool) *Multiset[T] {
	panic("not implemented")
}
