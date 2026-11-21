package lockfree

type Multimap[K any, V any] struct {
}

func NewMultimap[K any, V any](less func(K, K) bool) *Multimap[K, V] {
	panic("not implemented")
}
