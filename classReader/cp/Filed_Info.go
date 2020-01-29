package cp

import (
    "bufio"
)

type   Field_Info struct {
    accessFlags     uint16
    nameIndex       uint16
    descriptorIndex uint16
    attributesCount uint16
    //attribute_info attributes[attributes_count];
    attributes []  Attribute_Info
}

func (cls *Field_Info) Parse(tag int, r *bufio.Reader) {
    highBytes := make([] byte, 2)
    r.Read(highBytes)
    lowBytes := make([] byte, 2)
    r.Read(lowBytes)
    cls.accessFlags = uint16(ToInt(highBytes))
    cls.nameIndex = uint16(ToInt(lowBytes))

    highBytes = make([] byte, 2)
    r.Read(highBytes)
    lowBytes = make([] byte, 2)
    r.Read(lowBytes)
    cls.descriptorIndex = uint16(ToInt(highBytes))
    cls.attributesCount = uint16(ToInt(lowBytes))

}
func (cls *Field_Info) Print(index int, m map[int]  interface{} ) {

}