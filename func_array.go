package neovm

import (
	"math/big"

	"github.com/milechao/neovm/errors"
	"github.com/milechao/neovm/types"
)

func opArraySize(e *ExecutionEngine) (VMState, error) {
	item := PopStackItem(e)
	if _, ok := item.(*types.Array); ok {
		PushData(e, len(item.GetArray()))
	} else {
		PushData(e, len(item.GetByteArray()))
	}

	return NONE, nil
}

func opPack(e *ExecutionEngine) (VMState, error) {
	size := PopInt(e)
	var items []types.StackItems
	for i := 0; i < size; i++ {
		items = append(items, PopStackItem(e))
	}
	PushData(e, items)
	return NONE, nil
}

func opUnpack(e *ExecutionEngine) (VMState, error) {
	arr := PopArray(e)
	l := len(arr)
	for i := l - 1; i >= 0; i-- {
		Push(e, arr[i])
	}
	PushData(e, l)
	return NONE, nil
}

func opPickItem(e *ExecutionEngine) (VMState, error) {
	index := PopStackItem(e)
	items := PopStackItem(e)

	switch items.(type) {
	case *types.Array:
		i := int(index.GetBigInteger().Int64())
		if i < 0 || i >= len(items.GetArray()) {
			return FAULT, errors.NewErr("opPickItem invalid array.")
		}
		PushData(e, items.GetArray()[i])
	case *types.Map:
		value := items.(*types.Map).TryGetValue(index)
		//TODO should return a nil type when not exist?
		if value == nil {
			return FAULT, errors.NewErr("opPickItem map element not exist.")
		}
		PushData(e, value)

	default:
		return FAULT, errors.NewErr("opPickItem unknown item type.")
	}

	return NONE, nil
}

func opSetItem(e *ExecutionEngine) (VMState, error) {
	newItem := PopStackItem(e)
	if value, ok := newItem.(*types.Struct); ok {
		newItem = value.Clone()
	}

	index := PopStackItem(e)
	item := PopStackItem(e)

	switch item.(type) {
	case *types.Map:
		item.GetMap()[index] = newItem

	case *types.Array:
		items := item.GetArray()
		i := int(index.GetBigInteger().Int64())
		if i < 0 || i >= len(items) {
			return FAULT, errors.NewErr("opSetItem invalid array.")
		}
		items[i] = newItem
	default:
		return FAULT, errors.NewErr("opSetItem unknown item type.")
	}

	return NONE, nil
}

func opNewArray(e *ExecutionEngine) (VMState, error) {
	count := PopInt(e)
	var items []types.StackItems
	for i := 0; i < count; i++ {
		items = append(items, types.NewBoolean(false))
	}
	PushData(e, types.NewArray(items))
	return NONE, nil
}

func opNewStruct(e *ExecutionEngine) (VMState, error) {
	count := PopBigInt(e)
	var items []types.StackItems
	for i := 0; count.Cmp(big.NewInt(int64(i))) > 0; i++ {
		items = append(items, types.NewBoolean(false))
	}
	PushData(e, types.NewStruct(items))
	return NONE, nil
}

func opNewMap(e *ExecutionEngine) (VMState, error) {
	PushData(e, types.NewMap())
	return NONE, nil
}

func opAppend(e *ExecutionEngine) (VMState, error) {
	newItem := PopStackItem(e)
	if value, ok := newItem.(*types.Struct); ok {
		newItem = value.Clone()
	}
	itemArr := PopArray(e)
	itemArr = append(itemArr, newItem)
	return NONE, nil
}

func opReverse(e *ExecutionEngine) (VMState, error) {
	itemArr := PopArray(e)
	for i, j := 0, len(itemArr)-1; i < j; i, j = i+1, j-1 {
		itemArr[i], itemArr[j] = itemArr[j], itemArr[i]
	}
	return NONE, nil
}
