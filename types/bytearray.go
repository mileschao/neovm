package types

import (
	"math/big"

	"github.com/milechao/neovm/interfaces"
)

type ByteArray struct {
	value []byte
}

func NewByteArray(value []byte) *ByteArray {
	var this ByteArray
	this.value = value
	return &this
}

func (this *ByteArray) Equals(other StackItems) bool {
	if _, ok := other.(*ByteArray); !ok {
		return false
	}
	a1 := this.value
	a2 := other.GetByteArray()
	l1 := len(a1)
	l2 := len(a2)
	if l1 != l2 {
		return false
	}
	for i := 0; i < l1; i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func (this *ByteArray) GetBigInteger() *big.Int {
	return ConvertBytesToBigInteger(this.value)
}

func (this *ByteArray) GetBoolean() bool {
	for _, b := range this.value {
		if b != 0 {
			return true
		}
	}
	return false
}

func (this *ByteArray) GetByteArray() []byte {
	return this.value
}

func (this *ByteArray) GetInterface() interfaces.Interop {
	return nil
}

func (this *ByteArray) GetArray() []StackItems {
	return []StackItems{this}
}

func (this *ByteArray) GetStruct() []StackItems {
	return []StackItems{this}
}

func (this *ByteArray) GetMap() map[StackItems]StackItems {
	return nil
}
