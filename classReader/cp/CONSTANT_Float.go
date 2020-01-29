package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//4
type CONSTANT_Float struct {
	tag   uint8
	//4 bit
	bytes [] byte
}

func (cls *CONSTANT_Float) Parse(tag int, r *bufio.Reader) {
	bytes := make([] byte, 4)
	r.Read(bytes)
	cls.tag = uint8(tag)
	cls.bytes=bytes
}

func (cls *CONSTANT_Float) Print(index int, m map[int]   interface{} ) {
	fmt.Println( "#"+strconv.Itoa(index)+"=float ", ToInt(cls.bytes))
}