package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//12
type CONSTANT_NameAndType struct {
	tag             uint8
	nameIndex       uint16
	descriptorIndex uint16
}

func (cls *CONSTANT_NameAndType) Parse(tag int, r *bufio.Reader) {
	lowBytes := make([] byte, 2)
	r.Read(lowBytes)
	cls.nameIndex = uint16(ToInt(lowBytes))

	highBytes := make([] byte, 2)
	r.Read(highBytes)
	cls.descriptorIndex = uint16(ToInt(highBytes))
	cls.tag = uint8(tag)
}

func (cls *CONSTANT_NameAndType) Print(index int, m map[int] interface{}) {
	cp := m[int(cls.descriptorIndex)]
	var utf, _ = cp.(*CONSTANT_Utf8)

	fmt.Println("#" + strconv.Itoa(index)+"=NameAndType #" +strconv.Itoa(int(cls.descriptorIndex)) + " "+ utf.utf8)
}
