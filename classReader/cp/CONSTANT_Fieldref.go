package cp

import (
	"bufio"
	"fmt"
	"strconv"
)

//9
type CONSTANT_Fieldref struct {
	tag              uint8
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (cls *CONSTANT_Fieldref) Parse(tag int, r *bufio.Reader) {
	lowBytes := make([] byte, 2)
	r.Read(lowBytes)
	cls.classIndex = uint16(ToInt(lowBytes))

	highBytes := make([] byte, 2)
	r.Read(highBytes)
	cls.nameAndTypeIndex = uint16(ToInt(highBytes))
	cls.tag = uint8(tag)
}

func (cls *CONSTANT_Fieldref) Print(index int, m map[int] interface{}) {
	cp := m[int(cls.nameAndTypeIndex)]
	var nameAndType, _ = cp.(*CONSTANT_NameAndType)
	classCp := m[int(cls.classIndex)]
	var constrantClass, _ = classCp.(*CONSTANT_Class)
	var curNameIndex=  constrantClass.nameIndex
	cp = m[int(curNameIndex)]
	var utf8, _ = cp.(*CONSTANT_Utf8)
	cp = m[int(nameAndType.nameIndex)]
	var utf, _ = cp.(*CONSTANT_Utf8)
	cp = m[int(nameAndType.descriptorIndex)]
	var desUtf, _ = cp.(*CONSTANT_Utf8)
	fmt.Println("#"+ strconv.Itoa(index) +"=Fieldref    #"+ strconv.Itoa(int(nameAndType.nameIndex))+  ".#"+ strconv.Itoa(int(nameAndType.descriptorIndex))  +" "+ utf8.utf8 + " " +utf.utf8 +" " +desUtf.utf8)
}