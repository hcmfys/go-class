package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//18
type CONSTANT_InvokeDynamic struct {
	tag                      uint8
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (cls *CONSTANT_InvokeDynamic) Parse(tag int, r *bufio.Reader) {
	lowBytes := make([] byte, 2)
	r.Read(lowBytes)
	cls.bootstrapMethodAttrIndex = uint16(ToInt(lowBytes))

	highBytes := make([] byte, 2)
	r.Read(highBytes)
	cls.nameAndTypeIndex = uint16(ToInt(highBytes))
	cls.tag = uint8(tag)
}

func (cls *CONSTANT_InvokeDynamic) Print(index int, m map[int] interface{}) {
	cp := m[int(cls.bootstrapMethodAttrIndex)]
	var utf, _ = cp.(CONSTANT_Utf8)

	cpName := m[int(cls.nameAndTypeIndex)]
	var utfName, _ = cpName.(CONSTANT_Utf8)
	fmt.Println("#" + strconv.Itoa(index)+"=InvokeDynamic #"+  string(cls.bootstrapMethodAttrIndex) +".#"+string(cls.nameAndTypeIndex) + " "  +utf.utf8 +" "+ utfName.utf8)
}
