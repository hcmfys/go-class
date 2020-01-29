package cp

import (
	"bufio"
	"fmt"
)

type   Method_Info struct {
	accessFags      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	//attribute_info attributes[attributes_count];
	attributes []Attribute_Info
}


func (cls *Method_Info) Parse(tag int, r *bufio.Reader) {
	highBytes := make([] byte, 2)
	r.Read(highBytes)
	lowBytes := make([] byte, 2)
	r.Read(lowBytes)
	cls.accessFags = uint16(ToInt(highBytes))
	cls.nameIndex = uint16(ToInt(lowBytes))

	highBytes = make([] byte, 2)
	r.Read(highBytes)
	lowBytes = make([] byte, 2)
	r.Read(lowBytes)
	cls.descriptorIndex = uint16(ToInt(highBytes))
	attrCounts := ToInt(lowBytes)
	cls.attributesCount = uint16(attrCounts)
	cls.attributes = make([]Attribute_Info, attrCounts)
	for i := 0; i < attrCounts; i++ {
		attr := Attribute_Info{}
		attr.Parse(0, r)
		cls.attributes[i] = attr
	}

}
func (cls *Method_Info) Print(index int, m map[int]  interface{} ) {

	cp := m[int(cls.nameIndex)]
	var utf, _ =(cp).(*CONSTANT_Utf8)
	cpDes := m[int(cls.descriptorIndex)]
	var utfDes, _ =(cpDes).(*CONSTANT_Utf8)
	fmt.Print("method==>")

	p,r:=ParseDesc(utfDes.utf8)
	fmt.Println(ParseAccess(int(cls.accessFags))  +" "+ r + " "+ utf.utf8 +" ("+p +")")

}
