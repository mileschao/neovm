package types

import (
	"math/big"

	"github.com/milechao/neovm/interfaces"
)

type Array struct {
	_array []StackItems
}

func NewArray(value []StackItems) *Array {
	var this Array
	this._array = value
	return &this
}

func (this *Array) Equals(other StackItems) bool {
	if _, ok := other.(*Array); !ok {
		return false
	}
	a1 := this._array
	a2 := other.GetArray()
	l1 := len(a1)
	l2 := len(a2)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		if !a1[i].Equals(a2[i]) {
			return false
		}
	}
	return true
}

func (this *Array) GetBigInteger() *big.Int {
	if len(this._array) == 0 {
		return big.NewInt(0)
	}
	return this._array[0].GetBigInteger()
}

func (this *Array) GetBoolean() bool {
	if len(this._array) == 0 {
		return false
	}
	return this._array[0].GetBoolean()
}

func (this *Array) GetByteArray() []byte {
	if len(this._array) == 0 {
		return []byte{}
	}
	return this._array[0].GetByteArray()
}

func (this *Array) GetInterface() interfaces.Interop {
	if len(this._array) == 0 {
		return nil
	}
	return this._array[0].GetInterface()
}

func (this *Array) GetArray() []StackItems {
	return this._array
}

func (this *Array) GetStruct() []StackItems {
	return this._array
}

func (this *Array) GetMap() map[StackItems]StackItems {
	return nil
}
