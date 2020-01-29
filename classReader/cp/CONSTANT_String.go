package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//8
type CONSTANT_String struct {
	tag       uint8
	nameIndex uint16
}

func (cls *CONSTANT_String) Parse(tag int, r *bufio.Reader) {
	lowBytes := make([] byte, 2)
	r.Read(lowBytes)
	cls.nameIndex = uint16(ToInt(lowBytes))
	cls.tag = uint8(tag)
}
func (cls *CONSTANT_String) Print(index int, m map[int] interface{}) {
	cp := m[int(cls.nameIndex)]
	utf,_ := cp.(*CONSTANT_Utf8)
	fmt.Println("#"+strconv.Itoa(index)+"=String #" +strconv.Itoa(int(cls.nameIndex)) +" " + utf.utf8)
}