package buildengine

import (
	"encoding/binary"
)

type Stat int

var (
	endian = binary.LittleEndian
)

type Vector struct {
	X     int32
	Y     int32
	Z     int32
	Angle int16
}

type Texture struct {
	Ptr     int
	Angle   int
	Shade   int
	Pallete int
	XOffset int
	YOffset int
}

type Tags struct {
	Lotag int
	Hitag int
	Extra int
}

type Point struct {
	X int
	Y int
}
