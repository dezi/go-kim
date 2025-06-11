package types

import (
	"errors"
	"fmt"
	"github.com/dezi/go-kim/utils/log"
)

var (
	localStaticInt  int
	GlobalStaticInt int = 12
)

type OurTypes struct {
	VarInt     int
	VarByte    byte
	VarInt8    int8
	VarInt16   int16
	VarInt32   int32
	VarInt64   int64
	VarUInt8   uint8
	VarUInt16  uint16
	VarUInt32  uint32
	VarUInt64  uint64
	VarString  string
	VarBool    bool
	VarFloat32 float32
	VarFloat64 float64
}

type Address struct {
	FirstName string
	LastName  string
	Street    string
}

func Test() {

	log.Printf("In types test...")
	defer log.Printf("In types test fettich.")

	var ot OurTypes

	ot1 := OurTypes{
		VarInt:     12,
		VarByte:    0,
		VarInt8:    0,
		VarInt16:   0,
		VarInt32:   -2,
		VarInt64:   0,
		VarUInt8:   0,
		VarUInt16:  0,
		VarUInt32:  0,
		VarUInt64:  0,
		VarString:  "huhu",
		VarBool:    false,
		VarFloat32: 0,
		VarFloat64: 0,
	}

	ot1.VarFloat64 = 23
	ot.VarInt = 3

	ot1.VarInt16 = int16(ot.VarInt32)
	ot1.VarInt32 = int32(ot.VarInt16)
	ot1.VarInt = int(ot.VarInt64)

	ot1.VarFloat32 = float32(ot1.VarFloat64)
	ot1.VarFloat64 = float64(ot1.VarFloat32)

	ot1.VarString = "12.4"

	var parsedInt int

	_, err := fmt.Sscanf(ot1.VarString, "%d", &parsedInt)
	if err != nil {
		log.Cerror(err)
		return
	}

	log.Printf("err=%v parsedInt=%d", err, parsedInt)

	var int1 = 17
	var int2 = 23

	CallBy(int1, &int2)

	log.Printf("int1=%d int2=%d", int1, int2)

	res, err := Divide(12, 3)
	if err != nil {
		log.Cerror(err)
		return
	}

	log.Printf("res=%d", res)

	localStaticInt = res
	GlobalStaticInt = res
}

func Test2() {
	log.Printf("Local Static int=%d", localStaticInt)
	log.Printf("Global Static int=%d", GlobalStaticInt)
}

func CallBy(value int, reference *int) {
	*reference = value
	value = value * 2
}

func Divide(value1, value2 int) (res int, err error) {

	if value2 == 0 {
		err = errors.New("divide by zero")
		log.Cerror(err)
		return
	}

	res = value1 / value2

	return
}
