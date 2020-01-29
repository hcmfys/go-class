package cp

import (
	"bufio"
	"fmt"
)

type ClassFile struct {
	//u4             magic; //Class文件的标志
	Magic uint32
	//u2             minor_version; //Class的小版本号
	MinorVersion uint16
	//u2             major_version; //Class的大版本号
	MajorVersion uint16
	//u2             constant_pool_count; //常量池的数量
	ConstantPoolCount uint16
	//cp_info        constant_pool[constant_pool_count-1]; //常量池
	ConstantPool map[int]interface{}
	//u2             access_flags; //Class的访问标记
	accessFlags uint16
	//u2             this_class; //当前类
	thisClass uint16
	//u2             super_class; //父类
	superClass uint16
	//u2             interfaces_count; //接口
	interfacesCount uint16
	//u2             interfaces[interfaces_count]; //一个类可以实现多个接口
	interfaces [] uint16
	//u2             fields_count; //Class文件的字段属性
	fieldsCount uint16
	//field_info     fields[fields_count]; //一个类会可以有个字段
	fieldInfo [] Field_Info
	//u2             methods_count; //Class文件的方法数量
	methodsCount uint16
	//method_info    methods[methods_count]; //一个类可以有个多个方法
	methodInfo  [] Method_Info
	// u2             attributes_count; //此类的属性表中的属性数
	attributesCount uint16
	//attributeInfo attributes[attributes_count]; //属性表集合
	attributes []Attribute_Info
}

//魔数
func (cf *ClassFile) ReadMagic(r *bufio.Reader) {
	magics := make([]byte, 4)
	n, _ := r.Read(magics)
	if n != 4 {
		panic("magic[4字节] 魔数,错误！")
	}
	cf.Magic = uint32(ToInt(magics))
}

//次版本号
func (cf *ClassFile) ReadMinorVersion(r *bufio.Reader) {
	//minor_version[2字节] 次版本号
	minorVersion := make([]byte, 2)
	r.Read(minorVersion)
	cf.MinorVersion = uint16(ToInt(minorVersion))
}

func (cf *ClassFile) ReadMajorVersion(r *bufio.Reader) {
	//major_version[2字节] 主版本号，低版本的jdk无法执行高版本的class文件。
	majorVersion := make([]byte, 2)
	r.Read(majorVersion)
	cf.MajorVersion = uint16(ToInt(majorVersion))
}

func (cf *ClassFile) ReadConstantPool(r *bufio.Reader) int {
	//constant_pool_count[2字节] 常量池里的项目个数
	//constant_pool 常量池里每一个项目类型都用一个tag标示。从1开始取值，比如取值为1时，表示info

	constantPoolCount := make([]byte, 2)
	r.Read(constantPoolCount)
	cpCounts := ToInt(constantPoolCount)
	cf.ConstantPoolCount = uint16(cpCounts)
	cpM := make(map[int]interface{})
	for i := 1; i < cpCounts; i++ {
		tag := GetTag(r)
		cpInfo := GetCpInfo(tag)
		cpInfo.Parse(tag, r)
		cpM[i] = cpInfo
	}
	cf.ConstantPool = cpM
	return cpCounts
}

func (cf *ClassFile) ReadAccessFlags(r *bufio.Reader) {
	flags := make([]byte, 2)
	r.Read(flags)
	cf.accessFlags = uint16(ToInt(flags))
}

func (cf *ClassFile) ReadThisClass(r *bufio.Reader) {
	clsses := make([]byte, 2)
	r.Read(clsses)
	cf.thisClass = uint16(ToInt(clsses))
}

func (cf *ClassFile) ReadsuperClass(r *bufio.Reader) {
	superClss := make([]byte, 2)
	r.Read(superClss)
	cf.superClass = uint16(ToInt(superClss))
}

func (cf *ClassFile) ReadInterfaces(r *bufio.Reader) {
	interfaces := make([]byte, 2)
	r.Read(interfaces)
	counts := ToInt(interfaces)
	cf.interfacesCount = uint16(counts)
	cf.interfaces = make([] uint16, counts)
	for i := 0; i < counts; i++ {
		interfaces := make([]byte, 2)
		r.Read(interfaces)
		cf.interfaces[i] = uint16(ToInt(interfaces))
	}
}

func (cf *ClassFile) ReadFieldInfo(r *bufio.Reader) {
	fields := make([]byte, 2)
	r.Read(fields)
	counts := ToInt(fields)
	cf.fieldsCount = uint16(counts)
	cf.fieldInfo = make([] Field_Info, counts)
	for i := 0; i < counts; i++ {
		fieldInfo := Field_Info{}
		fieldInfo.Parse(0, r)
		cf.fieldInfo[i] = fieldInfo
	}
}

func (cf *ClassFile) ReadMethodInfo(r *bufio.Reader) {
	methods := make([]byte, 2)
	r.Read(methods)
	counts := ToInt(methods)
	cf.methodsCount = uint16(counts)
	cf.methodInfo = make([] Method_Info, counts)
	for i := 0; i < counts; i++ {
		methodInfo := Method_Info{}
		methodInfo.Parse(0, r)
		cf.methodInfo[i] = methodInfo
	}
}

func (cf *ClassFile) ReadAttributeInfo(r *bufio.Reader) {

}

/**
dump class file
 */
func (cf *ClassFile) Dump() {
	fmt.Println("Constant pool:")
	cpM := cf.ConstantPool
	for i := 1; i < int(cf.ConstantPoolCount); i++ {
		cpInfo := cpM[i].(CP_Info)
		cpInfo.Print(i, cpM)
	}
	ms:=cf.methodInfo
	for i := 0; i < int(cf.methodsCount); i++ {
		m := ms[i]
		m.Print(i, cpM)
		attrs:=m.attributes
		attrCounts:=int(m.attributesCount)
		if attrCounts>0 {
			for j := 0; j < attrCounts; j++ {
				attr:=attrs[j]
				attr.Print(0,cpM)
			}
		}
	}



}
