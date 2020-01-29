package cp

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type Attribute_Info struct {

	//u2 attribute_name_index;
	//u4 attribute_length;
	//u1 info[attribute_length];
	attributeNameIndex uint16
	attributeLength    uint32
	info               []uint8
}


func (cls *Attribute_Info) Parse(tag int, r *bufio.Reader) {
	highBytes := make([] byte, 2)
	r.Read(highBytes)
	lowBytes := make([] byte, 4)
	r.Read(lowBytes)
	cls.attributeNameIndex = uint16(ToInt(highBytes))
	attributeLength := ToInt(lowBytes)
	cls.attributeLength = uint32(attributeLength)

	cls.info = make([] uint8, attributeLength)
	b := make([] byte, attributeLength)
	r.Read(b)
	for i := 0; i < attributeLength; i++ {
		cls.info[i] = uint8(b[i])
	}

}
func (cls *Attribute_Info) Print(index int, m map[int]  interface{} ) {

	var curNameIndex = cls.attributeNameIndex
	cp := m[int(curNameIndex)]
	var utf8, _ = cp.(*CONSTANT_Utf8)
	fmt.Println("attr==>" + utf8.utf8)
	attrName := utf8.utf8
	if strings.Compare(attrName, "Code") == 0 {

		reader:=bytes.NewReader(cls.info)
		r:=bufio.NewReader(reader)
		code:=Code_attribute{}
		code.attributeNameIndex=cls.attributeNameIndex
		code.attributeLength=cls.attributeLength
		code.Parse(0,r)
		code.Print(0,m)
	}

}
