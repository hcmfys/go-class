package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//16
type CONSTANT_MethodType struct {
	tag            uint8
	referenceKind  uint8
	referenceIndex uint16
}

func (cls *CONSTANT_MethodType) Parse(tag int, r *bufio.Reader) {
	lowBytes := make([] byte, 1)
	r.Read(lowBytes)
	cls.referenceKind = uint8(ToInt(lowBytes))

	highBytes := make([] byte, 2)
	r.Read(highBytes)
	cls.referenceIndex = uint16(ToInt(highBytes))
	cls.tag = uint8(tag)
}

func (cls *CONSTANT_MethodType) Print(index int, m map[int] interface{}) {
	cp := m[int(cls.referenceKind)]
	var utf, _ = cp.(CONSTANT_Utf8)
	cpName := m[int(cls.referenceIndex)]
	var utfName, _ = cpName.(CONSTANT_Utf8)
	fmt.Println("#" + strconv.Itoa(index)+ "MethodType=#" + utf.utf8 +" #"+ utfName.utf8)
}