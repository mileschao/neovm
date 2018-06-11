package neovm

import (
	"math/big"
	"testing"

	vtypes "github.com/milechao/neovm/types"
)

func TestOpInvert(t *testing.T) {
	var e ExecutionEngine
	stack := NewRandAccessStack()
	stack.Push(NewStackItem(vtypes.NewInteger(big.NewInt(123456789))))
	e.EvaluationStack = stack

	opInvert(&e)
	i := big.NewInt(123456789)

	if PeekBigInteger(&e).Cmp(i.Not(i)) != 0 {
		t.Fatal("NeoVM OpInvert test failed.")
	}
}

func TestOpEqual(t *testing.T) {
	var e ExecutionEngine
	stack := NewRandAccessStack()
	stack.Push(NewStackItem(vtypes.NewInteger(big.NewInt(123456789))))
	stack.Push(NewStackItem(vtypes.NewInteger(big.NewInt(123456789))))
	e.EvaluationStack = stack

	opEqual(&e)
	if !PopBoolean(&e) {
		t.Fatal("NeoVM OpEqual test failed.")
	}
}
