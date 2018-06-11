package neovm

import (
	"bytes"
	"math/big"
	"testing"

	vtypes "github.com/milechao/neovm/types"
)

func TestOpCat(t *testing.T) {
	var e ExecutionEngine
	stack := NewRandAccessStack()
	stack.Push(NewStackItem(vtypes.NewByteArray([]byte("aaa"))))
	stack.Push(NewStackItem(vtypes.NewByteArray([]byte("bbb"))))
	e.EvaluationStack = stack

	opCat(&e)
	if Count(&e) != 1 || !bytes.Equal(PeekNByteArray(0, &e), []byte("aaabbb")) {
		t.Fatalf("NeoVM OpCat test failed, expect aaabbb, got %s.", string(PeekNByteArray(0, &e)))
	}
}

func TestOpSubStr(t *testing.T) {
	var e ExecutionEngine
	stack := NewRandAccessStack()
	stack.Push(NewStackItem(vtypes.NewByteArray([]byte("12345"))))
	stack.Push(NewStackItem(vtypes.NewInteger(big.NewInt(1))))
	stack.Push(NewStackItem(vtypes.NewInteger(big.NewInt(4))))
	e.EvaluationStack = stack

	opSubStr(&e)
	if !bytes.Equal(PeekNByteArray(0, &e), []byte("2345")) {
		t.Fatalf("NeoVM OpSubStr test failed, expect 234, got %s.", string(PeekNByteArray(0, &e)))
	}
}

func TestOpLeft(t *testing.T) {
	var e ExecutionEngine
	stack := NewRandAccessStack()
	stack.Push(NewStackItem(vtypes.NewByteArray([]byte("12345"))))
	stack.Push(NewStackItem(vtypes.NewInteger(big.NewInt(4))))
	e.EvaluationStack = stack

	opLeft(&e)
	if !bytes.Equal(PeekNByteArray(0, &e), []byte("1234")) {
		t.Fatalf("NeoVM OpLeft test failed, expect 1234, got %s.", string(PeekNByteArray(0, &e)))
	}
}

func TestOpRight(t *testing.T) {
	var e ExecutionEngine
	stack := NewRandAccessStack()
	stack.Push(NewStackItem(vtypes.NewByteArray([]byte("12345"))))
	stack.Push(NewStackItem(vtypes.NewInteger(big.NewInt(3))))
	e.EvaluationStack = stack

	opRight(&e)
	if !bytes.Equal(PeekNByteArray(0, &e), []byte("345")) {
		t.Fatalf("NeoVM OpRight test failed, expect 345, got %s.", string(PeekNByteArray(0, &e)))
	}
}

func TestOpSize(t *testing.T) {
	var e ExecutionEngine
	stack := NewRandAccessStack()
	stack.Push(NewStackItem(vtypes.NewByteArray([]byte("12345"))))
	e.EvaluationStack = stack

	opSize(&e)
	if PeekInt(&e) != 5 {
		t.Fatalf("NeoVM OpSize test failed, expect 5, got %d.", PeekInt(&e))
	}
}
