package lockfree

type PriorityQueue[K any, V any] struct {
}

func NewPriorityQueue[K any, V any](less func(K, K) bool) *PriorityQueue[K, V] {
	panic("not implemented")
}
