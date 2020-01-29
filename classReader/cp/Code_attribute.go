package cp

import (
	"bufio"
	"fmt"
	"springbus.org/classReader/ldc"
)

type  Exception_table struct {
	start_pc   uint16
	end_pc     uint16
	handler_pc uint16
	catch_type uint16
}

type   Code_attribute struct {
	attributeNameIndex     uint16
	attributeLength       uint32
	maxStack              uint16
	maxLocals             uint16
	codeLength            uint32
	code                   [] uint8
	exceptionTableLength uint16
	exceptionTable        []Exception_table
	attributesCount       uint16
	attributeInfo         [] Attribute_Info
}

func (cls *Code_attribute) Parse(tag int, r *bufio.Reader) {

	highBytes := make([] byte, 2)
	r.Read(highBytes)
	cls.maxStack = uint16(ToInt(highBytes))

	highBytes = make([] byte, 2)
	r.Read(highBytes)
	cls.maxLocals = uint16(ToInt(highBytes))

	highBytes = make([] byte, 4)
	r.Read(highBytes)
	cls.codeLength = uint32(ToInt(highBytes))

	highBytes = make([] byte, int(cls.codeLength))
	r.Read(highBytes)
	cls.code = make([]uint8, int(cls.codeLength))
	for i := 0; i < int(cls.codeLength); i++ {
		cls.code[i] = uint8(highBytes[i])
	}
}

func (cls *Code_attribute) Print( index int,m map[int] interface{}) {

	fmt.Println("maxLocals=", cls.maxLocals)
	fmt.Println("maxStack=", cls.maxStack)
	fmt.Println("code=", cls.code)
	codeArr := cls.code
	ldcCode := ldc.Ldc_Code{}
	codeMap := ldcCode.GetIdcMap()
	len := len(codeArr)
	for i := 0; i < len; i++ {
		code := codeArr[i]
		ldc := codeMap[int(code)]
		offX := ldc.CodeBytes
		fmt.Print(ldc.CodeTxt + " ")
		if offX > 0 {
			bts := codeArr[i+1 : i+offX+1]
			cpIndex := ToInt(bts)
			cp := m[cpIndex]
			if cp != nil {
				cpInfo := cp.(CP_Info)
				cpInfo.Print(cpIndex, m)
			}
		}
		fmt.Println("")
		i = i + offX
	}

}

