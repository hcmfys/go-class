package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//	3
type CONSTANT_Integer struct {
	tag   uint8
	//4 bit
	bytes [] byte
}

func (cls *CONSTANT_Integer) Parse(tag int, r *bufio.Reader) {
	bytes := make([] byte, 4)
	r.Read(bytes)
	cls.tag = uint8(tag)
	cls.bytes=bytes
}
func (cls *CONSTANT_Integer) Print(index int,  m map[int]  interface{}) {
	fmt.Println( "#"+strconv.Itoa(index)+"=Integer ", ToInt(cls.bytes))
}
