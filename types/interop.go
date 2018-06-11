package types

import (
	"bytes"
	"math/big"

	"github.com/milechao/neovm/interfaces"
)

type Interop struct {
	_object interfaces.Interop
}

func NewInteropInterface(value interfaces.Interop) *Interop {
	var ii Interop
	ii._object = value
	return &ii
}

func (this *Interop) Equals(other StackItems) bool {
	if _, ok := other.(*Interop); !ok {
		return false
	}
	if !bytes.Equal(this._object.ToArray(), other.GetInterface().ToArray()) {
		return false
	}
	return true
}

func (this *Interop) GetBigInteger() *big.Int {
	return big.NewInt(0)
}

func (this *Interop) GetBoolean() bool {
	if this._object == nil {
		return false
	}
	return true
}

func (this *Interop) GetByteArray() []byte {
	return this._object.ToArray()
}

func (this *Interop) GetInterface() interfaces.Interop {
	return this._object
}

func (this *Interop) GetArray() []StackItems {
	return []StackItems{this}
}

func (this *Interop) GetStruct() []StackItems {
	return []StackItems{this}
}

func (this *Interop) GetMap() map[StackItems]StackItems {
	return nil
}
