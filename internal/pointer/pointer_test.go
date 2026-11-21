package pointer_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/syubogdanov/lockfree/internal/pointer"
)

func TestAtomicMarkablePointer_Load(t *testing.T) {
	ap := pointer.NewAtomicMarkablePointer[int](nil, false)

	p, m := ap.Load()

	require.Nil(t, p)
	require.False(t, m)
}

func TestAtomicMarkablePointer_Store(t *testing.T) {
	ap := pointer.NewAtomicMarkablePointer[int](nil, false)

	ap.Store(nil, true)

	p, m := ap.Load()
	require.Nil(t, p)
	require.True(t, m)
}

func TestAtomicMarkablePointer_CompareAndSwap(t *testing.T) {
	ap := pointer.NewAtomicMarkablePointer[int](nil, false)
	ap.Store(nil, true)

	ok := ap.CompareAndSwap(nil, true, nil, false)
	require.True(t, ok)

	p, m := ap.Load()
	require.Nil(t, p)
	require.False(t, m)

	ok = ap.CompareAndSwap(nil, true, nil, false)
	require.False(t, ok)

	p, m = ap.Load()
	require.Nil(t, p)
	require.False(t, m)
}
