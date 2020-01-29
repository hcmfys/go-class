package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//1
type CONSTANT_Utf8 struct {
	tag    uint8
	length uint16
	bytes  [] byte
	utf8   string
}

func (cls *CONSTANT_Utf8) Parse(tag int, r *bufio.Reader) {
	bytes := make([] byte, 2)
	r.Read(bytes)
	cls.length = uint16(ToInt(bytes))
	cls.bytes = make([]byte, cls.length)
	r.Read(cls.bytes)
	cls.utf8 = string(cls.bytes)
	cls.tag = uint8(tag)
}
func (cls *CONSTANT_Utf8) Print( index int,m map[int] interface{}) {
	fmt.Println( "#"+strconv.Itoa(index)+"=utf8 " +cls.utf8)
}
