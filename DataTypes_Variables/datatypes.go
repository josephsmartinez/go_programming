// types and assignments
package main

import (
	"fmt"
	"math/cmplx"
)

// Truth - Capitial variable can be accessed outside of the package
// type constant - forced to obey the strict rules and prevent combining values
const Truth bool = true

// untype constant - can be implicitly converted by the conpiler
const x = 42

/*
Iota
Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants. It is reset to 0 whenever the reserved word const appears in the source. It can be used to construct a set of related constants
*/
const (
	c0 = iota
	c1 = iota
	c2 = iota
)

var (

	// boolean
	tobe    bool = true
	notTobe      = false

	// integers unsigned
	maxint8  uint8  = 1<<8 - 1
	maxint16 uint16 = 1<<16 - 1
	maxint32 uint32 = 1<<32 - 1
	maxint64 uint64 = 1<<64 - 1

	// intergers signed
	minint8  int8  = ((1 << 8 * -1) / 2)
	minint16 int16 = ((1 << 16 * -1) / 2)
	minint32 int32 = ((1 << 32 * -1) / 2)
	minint64 int64 = ((1 << 64 * -1) / 2)

	// bytes alias for uint8
	maxbyteVal byte = 1<<8 - 1

	// rune alias for int32
	minRuneVal rune = ((1 << 32 * -1) / 2)

	// floats
	floatVal32 float32 = 1
	floatVal64 float64 = 1

	// complex
	complexVal64  complex64  = (3 + 2i)
	complexVal128 complex128 = cmplx.Sqrt(-5 + 12i)

	// string
	name string = "hello world!"

	// array
	arraySting = [...]string{"Joe", "Mike", "Chris", "Sam"}
	arrayInt   = [2]uint8{(1 >> 8), (maxint8)}

	// slice
	sliceMatrixInt = [2][2]int{
		{0, 1},
		{2, 4},
	}
	sliceFloat32 = [3]float32{3.86, 66.6666, 99.99}

	// map
	timeZones = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	customers = map[int]string{
		1: "Tim",
		2: "Jen",
		3: "Abi",
	}
)

//  two dimensional slices
type matrixSlice [5][5]int

func main() {

	// String literals
	//sentence := `This is a string literal and will print as is with new line. \n will not work`

	// prints max values for unsigned ints
	fmt.Print("Unsigned Int-------------\n")
	fmt.Printf("%T %v\n", maxint8, maxint8)
	fmt.Printf("%T %v\n", maxint16, maxint16)
	fmt.Printf("%T %v\n", maxint32, maxint32)
	fmt.Printf("%T %v\n", maxint64, maxint64)

	// prints min values for signed ints
	fmt.Print("Signed Int-------------\n")
	fmt.Printf("%T %v\n", minint8, minint8)
	fmt.Printf("%T %v\n", minint16, minint16)
	fmt.Printf("%T %v\n", minint32, minint32)
	fmt.Printf("%T %v\n", minint64, minint64)

	// print other numiric types
	fmt.Print("Boolean-------------\n")
	fmt.Printf("%T %v\n", tobe, tobe)

	// complex values
	fmt.Print("Complex-------------\n")
	fmt.Printf("%T %v\n", complexVal64, complexVal64)
	fmt.Printf("%T %v\n", complexVal128, complexVal128)

	// bytes and runes
	fmt.Print("Bytes and Runes-------------\n")
	fmt.Printf("%T %v\n", maxbyteVal, maxbyteVal)
	fmt.Printf("%T %v\n", minRuneVal, minRuneVal)

	// string
	fmt.Print("String-------------\n")
	fmt.Printf("%T %v\n", name, name)
	fmt.Printf("%T %v\n", name, name[0:5])
	fmt.Printf("%T %v\n", name, name[6:11])

	// array
	fmt.Print("Array-------------\n")
	fmt.Printf("%T %v\n", arraySting, arraySting)
	for _, v := range arraySting {
		fmt.Printf("%T %v\n", v, v)
	}
	for _, v := range arrayInt {
		fmt.Printf("%T %v\n", v, v)
	}

	// slices
	fmt.Print("Array-------------\n")
	fmt.Printf("%T %v\n", sliceMatrixInt, sliceMatrixInt)
	fmt.Printf("%T %v\n", sliceMatrixInt[1][1], sliceMatrixInt[1][1])
	for i := range sliceMatrixInt {
		for _, v := range sliceMatrixInt[i] {
			fmt.Printf("%T %v\n", v, v)
		}
	}

	// map
	fmt.Print("Map-------------\n")
	fmt.Printf("%T %v\n", timeZones, timeZones)
	fmt.Printf("%T %v\n", customers, customers)
	for k, v := range timeZones {
		fmt.Printf("{\"%v\" : %v }\n", k, v)
	}
	for k, v := range customers {
		fmt.Printf("{\"%v\" : %v }\n", k, v)
	}

}
