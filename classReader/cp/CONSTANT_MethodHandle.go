package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//15
type CONSTANT_MethodHandle struct {
	tag              uint8
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (cls *CONSTANT_MethodHandle) Parse(tag int, r *bufio.Reader) {
	lowBytes := make([] byte, 2)
	r.Read(lowBytes)
	cls.classIndex = uint16(ToInt(lowBytes))

	highBytes := make([] byte, 2)
	r.Read(highBytes)
	cls.nameAndTypeIndex = uint16(ToInt(highBytes))
	cls.tag = uint8(tag)
}
func (cls *CONSTANT_MethodHandle) Print( index int ,m map[int] interface{}) {
	cp := m[int(cls.classIndex)]
	var utf, _ = cp.(CONSTANT_Utf8)

	cpName := m[int(cls.nameAndTypeIndex)]
	var utfName, _ = cpName.(CONSTANT_Utf8)
	fmt.Println("#" + strconv.Itoa(index)+  "MethodHandle " + utf.utf8 +" #"+ utfName.utf8)
}