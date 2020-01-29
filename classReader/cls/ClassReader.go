package cls

import (
	"bufio"
	"os"
	"springbus.org/classReader/cp"
)

type ClassReader struct {
}

func (c *ClassReader) ReadClass(fileName string)  cp.ClassFile {
	cf := cp.ClassFile{}
	f, e := os.Open(fileName)
	defer f.Close()
	if e != nil {
		panic(e.Error())
	}
	r := bufio.NewReader(f)
	cf.ReadMagic(r)
	cf.ReadMinorVersion(r)
	cf.ReadMajorVersion(r)

	//constant_pool_count[2字节] 常量池里的项目个数
	//constant_pool 常量池里每一个项目类型都用一个tag标示。从1开始取值，比如取值为1时，表示info
	 cf.ReadConstantPool(r)


	//u2             access_flags; //Class的访问标记
	cf.ReadAccessFlags(r)
	//u2             this_class; //当前类
	cf.ReadThisClass(r)
	//u2             super_class; //父类
	cf.ReadsuperClass(r)
	//u2             interfaces_count; //接口
	cf.ReadInterfaces(r)
	//field_info     fields[fields_count]; //一个类会可以有个字段
	cf.ReadFieldInfo(r)
	//u2             methods_count; //Class文件的方法数量
	//method_info    methods[methods_count]; //一个类可以有个多个方法
	cf.ReadMethodInfo(r)
	// u2             attributes_count; //此类的属性表中的属性数
	//attributeInfo attributes[attributes_count]; //属性表集合
	cf.ReadAttributeInfo(r)

	cf.Dump()
	return cf

}
