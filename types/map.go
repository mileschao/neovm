package types

import (
	"math/big"

	"github.com/milechao/neovm/interfaces"

	"reflect"
)

type Map struct {
	_map map[StackItems]StackItems
}

func NewMap() *Map {
	var mp Map
	mp._map = make(map[StackItems]StackItems)
	return &mp
}

func (this *Map) Add(key StackItems, value StackItems) error {
	this._map[key] = value
	return nil
}

func (this *Map) Clear() {
	this._map = make(map[StackItems]StackItems)
}

func (this *Map) ContainsKey(key StackItems) bool {
	_, ok := this._map[key]
	return ok
}

func (this *Map) Remove(key StackItems) {
	delete(this._map, key)
}

func (this *Map) Equals(that StackItems) bool {
	return reflect.DeepEqual(this, that)
}

func (this *Map) GetBoolean() bool {
	return true
}

func (this *Map) GetByteArray() []byte {
	return nil
	//return this.ToArray()
}

func (this *Map) GetBigInteger() *big.Int {
	return nil
}

func (this *Map) GetInterface() interfaces.Interop {
	return nil
}

func (this *Map) GetArray() []StackItems {
	return nil
}

func (this *Map) GetStruct() []StackItems {
	return nil
}

func (this *Map) GetMap() map[StackItems]StackItems {
	return this._map
}

func (this *Map) TryGetValue(key StackItems) StackItems {
	for k, v := range this._map {
		if k.Equals(key) {
			return v
		}
	}
	return nil
	//return this._map[key]
}
