package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//6
type CONSTANT_Double struct {
	tag       uint8
	highBytes [] byte
	lowBytes  [] byte
}

func (cls *CONSTANT_Double) Parse(tag int, r *bufio.Reader) {
	highBytes := make([] byte, 4)
	r.Read(highBytes)
	lowBytes := make([] byte, 4)
	r.Read(lowBytes)
	cls.tag = uint8(tag)
}
func (cls *CONSTANT_Double) Print(index int, m map[int]  interface{} ) {
	fmt.Println( "#"+strconv.Itoa(index)+"double=", ToInt(cls.lowBytes))
}