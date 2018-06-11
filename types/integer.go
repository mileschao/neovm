package types

import (
	"math/big"

	"github.com/milechao/neovm/interfaces"
)

type Integer struct {
	value *big.Int
}

func NewInteger(value *big.Int) *Integer {
	var this Integer
	this.value = value
	return &this
}

func (this *Integer) Equals(other StackItems) bool {
	if _, ok := other.(*Integer); !ok {
		return false
	}
	if this.value.Cmp(other.GetBigInteger()) != 0 {
		return false
	}
	return true
}

func (this *Integer) GetBigInteger() *big.Int {
	return this.value
}

func (this *Integer) GetBoolean() bool {
	if this.value.Cmp(big.NewInt(0)) == 0 {
		return false
	}
	return true
}

func (this *Integer) GetByteArray() []byte {
	return ConvertBigIntegerToBytes(this.value)
}

func (this *Integer) GetInterface() interfaces.Interop {
	return nil
}

func (this *Integer) GetArray() []StackItems {
	return []StackItems{this}
}

func (this *Integer) GetStruct() []StackItems {
	return []StackItems{this}
}

func (this *Integer) GetMap() map[StackItems]StackItems {
	return nil
}
