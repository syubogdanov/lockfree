package pointer

import (
	"sync/atomic"

	"github.com/syubogdanov/lockfree/internal/pair"
)

type AtomicMarkablePointer[T any] struct {
	pointer *atomic.Pointer[pair.Pair[*T, bool]]
}

func NewAtomicMarkablePointer[T any](pointer *T, mark bool) *AtomicMarkablePointer[T] {
	p := &pair.Pair[*T, bool]{
		First:  pointer,
		Second: mark,
	}

	ap := &atomic.Pointer[pair.Pair[*T, bool]]{}
	ap.Store(p)

	return &AtomicMarkablePointer[T]{
		pointer: ap,
	}
}

func (a *AtomicMarkablePointer[T]) Load() (*T, bool) {
	p := a.pointer.Load()
	return p.First, p.Second
}

func (a *AtomicMarkablePointer[T]) Store(pointer *T, mark bool) {
	p := &pair.Pair[*T, bool]{
		First:  pointer,
		Second: mark,
	}
	a.pointer.Store(p)
}

func (a *AtomicMarkablePointer[T]) CompareAndSwap(oldPointer *T, oldMark bool, newPointer *T, newMark bool) bool {
	op := a.pointer.Load()
	if op.First != oldPointer || op.Second != oldMark {
		return false
	}
	np := &pair.Pair[*T, bool]{
		First:  newPointer,
		Second: newMark,
	}
	return a.pointer.CompareAndSwap(op, np)
}
