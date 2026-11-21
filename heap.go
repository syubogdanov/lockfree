package lockfree

type Heap[T any] struct {
}

func NewHeap[T any](less func(T, T) bool) *Heap[T] {
	panic("not implemented")
}
