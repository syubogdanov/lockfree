package lockfree

type Map[K any, V any] struct {
}

func NewMap[K any, V any](less func(K, K) bool) *Map[K, V] {
	panic("not implemented")
}
