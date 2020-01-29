package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//5
type CONSTANT_Long struct {
	tag       uint8
	highBytes [] byte
	lowBytes  [] byte
}

func (cls *CONSTANT_Long) Parse(tag int, r *bufio.Reader) {
	highBytes := make([] byte, 4)
	r.Read(highBytes)
	lowBytes := make([] byte, 4)
	r.Read(lowBytes)
	cls.tag = uint8(tag)
	cls.lowBytes=lowBytes
	cls.highBytes=highBytes
}
func (cls *CONSTANT_Long) Print(index int, m map[int]  interface{} ) {
	fmt.Println( "#"+strconv.Itoa(index)+" long=", ToInt(cls.lowBytes))
}