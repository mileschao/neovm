package types

import (
	"math/big"

	"github.com/milechao/neovm/interfaces"
)

type StackItems interface {
	Equals(other StackItems) bool
	GetBigInteger() *big.Int
	GetBoolean() bool
	GetByteArray() []byte
	GetInterface() interfaces.Interop
	GetArray() []StackItems
	GetStruct() []StackItems
	GetMap() map[StackItems]StackItems
}

const (
	ByteArrayType byte = 0x00
	BooleanType   byte = 0x01
	IntegerType   byte = 0x02
	InterfaceType byte = 0x40
	ArrayType     byte = 0x80
	StructType    byte = 0x81
	MapType       byte = 0x82
)
