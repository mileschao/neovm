package types

import (
	"math/big"

	"github.com/milechao/neovm/interfaces"
)

type Struct struct {
	_array []StackItems
}

func NewStruct(value []StackItems) *Struct {
	var this Struct
	this._array = value
	return &this
}

func (this *Struct) Equals(other StackItems) bool {
	if _, ok := other.(*Struct); !ok {
		return false
	}
	a1 := this._array
	a2 := other.GetStruct()
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

func (this *Struct) GetBigInteger() *big.Int {
	if len(this._array) == 0 {
		return big.NewInt(0)
	}
	return this._array[0].GetBigInteger()
}

func (this *Struct) GetBoolean() bool {
	if len(this._array) == 0 {
		return false
	}
	return this._array[0].GetBoolean()
}

func (this *Struct) GetByteArray() []byte {
	if len(this._array) == 0 {
		return []byte{}
	}
	return this._array[0].GetByteArray()
}

func (this *Struct) GetInterface() interfaces.Interop {
	if len(this._array) == 0 {
		return nil
	}
	return this._array[0].GetInterface()
}

func (s *Struct) GetArray() []StackItems {
	return s._array
}

func (s *Struct) GetStruct() []StackItems {
	return s._array
}

func (s *Struct) Clone() StackItems {
	var arr []StackItems
	for _, v := range s._array {
		if value, ok := v.(*Struct); ok {
			arr = append(arr, value.Clone())
		} else {
			arr = append(arr, v)
		}
	}
	return &Struct{arr}
}

func (this *Struct) GetMap() map[StackItems]StackItems {
	return nil
}
