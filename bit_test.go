package golangbit

import (
	"testing"
	"fmt"
)

func TestGetBit(t *testing.T) {
	fmt.Println(GetBits(   0x1))
	fmt.Println(GetBits(   0x2))
	fmt.Println(GetBits(   0x3))
	fmt.Println(GetBits(   0x4))
	fmt.Println(GetBits(   0x5))
	fmt.Println(GetBits(   0x6))
	fmt.Println(GetBits(   0x7))
	fmt.Println(GetByte(GetBits(   0x7)))
	fmt.Println(GetBits(   46))
	fmt.Println(GetByte(GetBits( 46)))
	fmt.Println(GetByteBit(46,0))
	fmt.Println(GetByteBit(46,1))
	fmt.Println(GetByteBit(46,2))
	fmt.Println(GetByteBit(46,3))

}
