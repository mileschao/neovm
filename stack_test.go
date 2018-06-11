package neovm

import (
	"testing"

	"math/big"

	"github.com/milechao/neovm/types"
)

func TestRandomAccessStack_Count(t *testing.T) {
	r := NewRandAccessStack()
	r.Push(types.NewInteger(big.NewInt(9999)))
	r.Push(types.NewInteger(big.NewInt(8888)))

	if r.Count() != 2 {
		t.Fatalf("stack count test failed: expected 2, got %d ", r.Count())
	}
}

func TestRandomAccessStack_Pop(t *testing.T) {
	r := NewRandAccessStack()
	r.Push(types.NewInteger(big.NewInt(9999)))
	r.Push(types.NewInteger(big.NewInt(8888)))

	ret := r.Remove(0)
	ret.GetBigInteger()
	if ret.GetBigInteger().Int64() != 8888 {
		t.Fatalf("stack remove test failed: expect aaaa, got %d.", ret.GetBigInteger().Int64())
	}
}

func TestRandomAccessStack_Swap(t *testing.T) {
	r := NewRandAccessStack()
	r.Push(types.NewInteger(big.NewInt(9999)))
	r.Push(types.NewInteger(big.NewInt(8888)))
	r.Push(types.NewInteger(big.NewInt(7777)))

	r.Swap(0, 2)

	e0 := r.Pop().GetBigInteger().Int64()
	r.Pop()
	e2 := r.Pop().GetBigInteger().Int64()

	if e0 != 9999 || e2 != 7777 {
		t.Fatal("stack swap test failed.")
	}
}

func TestRandomAccessStack_Peek(t *testing.T) {
	r := NewRandAccessStack()
	r.Push(types.NewInteger(big.NewInt(9999)))
	r.Push(types.NewInteger(big.NewInt(8888)))

	e0 := r.Peek(0).GetBigInteger().Int64()
	e1 := r.Peek(1).GetBigInteger().Int64()

	if e0 != 8888 || e1 != 9999 {
		t.Fatal("stack peek test failed.")
	}
}
