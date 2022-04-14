package main

import "fmt"

type noKTP string

const (
	number1 int8=8
	number2 int16=16
	number3 int32=32
	number4 int64=64

	unumber1 uint8=8
	unumber2 uint16=16
	unumber3 uint32=32
	unumber4 uint64=64
	noKTPEko noKTP= "tes"
)

var test bool = number1 == int8(number2)
func main() {

	fmt.Println(test)
}