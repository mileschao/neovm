package types

import (
	"math/big"

	"github.com/milechao/neovm/interfaces"
)

type Boolean struct {
	value bool
}

func NewBoolean(value bool) *Boolean {
	var this Boolean
	this.value = value
	return &this
}

func (this *Boolean) Equals(other StackItems) bool {
	if _, ok := other.(*Boolean); !ok {
		return false
	}
	if this.value != other.GetBoolean() {
		return false
	}
	return true
}

func (this *Boolean) GetBigInteger() *big.Int {
	if this.value {
		return big.NewInt(1)
	}
	return big.NewInt(0)
}

func (this *Boolean) GetBoolean() bool {
	return this.value
}

func (this *Boolean) GetByteArray() []byte {
	if this.value {
		return []byte{1}
	}
	return []byte{0}
}

func (this *Boolean) GetInterface() interfaces.Interop {
	return nil
}

func (this *Boolean) GetArray() []StackItems {
	return []StackItems{this}
}

func (this *Boolean) GetStruct() []StackItems {
	return []StackItems{this}
}

func (this *Boolean) GetMap() map[StackItems]StackItems {
	return nil
}
