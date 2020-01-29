package cp

import (
	"bufio"
	"math"
	"strings"
)

func ToInt(bytes []byte) int {
	n := len(bytes)
	ret := 0
	for i := 0; i < n; i++ {
		//fmt.Println(n-i-1, math.Pow10(n-i-1), int(bytes[i : i+1][0]))
		ret += int(math.Pow10(n-i-1)) * int(bytes[i : i+1][0])
	}
	//fmt.Println("ret=", ret)
	return ret
}

func GetTag(r *bufio.Reader) int {
	bytes := make([] byte, 1)
	r.Read(bytes)
	return ToInt(bytes)
}

func  GetCpInfo(tag int) CP_Info {
	mapList := map[int]CP_Info{}
	mapList[1] = &CONSTANT_Utf8{}
	mapList[3] = &CONSTANT_Integer{}
	mapList[4] = &CONSTANT_Float{}
	mapList[5] = &CONSTANT_Long{}
	mapList[7] = &CONSTANT_Class{}
	mapList[8] = &CONSTANT_String{}
	mapList[9] = &CONSTANT_Fieldref{}
	mapList[10] = &CONSTANT_Methodref{}
	mapList[11] = &CONSTANT_InterfaceMethodref{}
	mapList[12] = &CONSTANT_NameAndType{}
	mapList[15] = &CONSTANT_MethodHandle{}
	mapList[16] = &CONSTANT_MethodType{}
	mapList[18] = &CONSTANT_InvokeDynamic{}
	return mapList[tag]
}

func ParseDesc( des string ) (string,string)  {

	/**
	B	byte	signed byte
	C	char	Unicode character code point in the Basic Multilingual Plane, encoded with UTF-16
	D	double	double-precision floating-point value
	F	float	single-precision floating-point value
	I	int	integer
	J	long	long integer
	L ClassName ;	reference	an instance of class ClassName
	S	short	signed short
	Z	boolean	true or false
	[	reference	one array dimension
	 */
	index:=strings.Index(des,")")
	pars:=des[:index+1]
	returns:=des[index+1:]
	return  parsePars(pars) ,parsePars(returns)

}

func parsePars(pars string) string {
	ret := ""
	list := make([]string, 0)
	l := len(pars)
	for i := 0; i < l; i++ {
		c := pars[i]
		if c == 'L' {
			q := i + 1
			for j := q; q < l; j++ {
				d := pars[j]
				if d == ';' {
					ret = pars[q : j]
					list = append(list, strings.ReplaceAll(ret, "/", "."))
					i = j + 1
					break
				}
			}
		} else {
			if c == 'B' {
				ret = "byte"
				list = append(list, ret)
			} else if c == 'C' {
				ret = "char "
				list = append(list, ret)
			} else if c == 'D' {
				ret = "double"
				list = append(list, ret)
			} else if c == 'F' {
				ret = "float"
				list = append(list, ret)
			} else if c == 'I' {
				ret = "int"
				list = append(list, ret)
			} else if c == 'J' {
				ret = "long"
				list = append(list, ret)
			} else if c == 'S' {
				ret = "short"
				list = append(list, ret)
			} else if c == 'Z' {
				ret = "boolean"
				list = append(list, ret)
			} else if c == '[' {
				ret = "[]"
				list = append(list, ret)

			}
		}
	}
	return strings.Join(list, ",")
}
func ParseAccess( des int ) string {

	/**
	ACC_PUBLIC	0x0001	Declared public; may be accessed from outside its package.
	ACC_PRIVATE	0x0002	Declared private; usable only within the defining class.
	ACC_PROTECTED	0x0004	Declared protected; may be accessed within subclasses.
	ACC_STATIC	0x0008	Declared static.
	ACC_FINAL	0x0010	Declared final; never directly assigned to after object construction (JLS §17.5).
	ACC_VOLATILE	0x0040	Declared volatile; cannot be cached.
	ACC_TRANSIENT	0x0080	Declared transient; not written or read by a persistent object manager.
	ACC_SYNTHETIC	0x1000	Declared synthetic; not present in the source code.
	ACC_ENUM	0x4000	Declared as an element of an enum.
	*/
	access := ""
	f1 := des & 0x0f
	f2 := des & 0xf0
	if f1&0x001 != 0 {
		access += " public "
	}
	if f1&0x0002 != 0 {
		access += " private "
	}
	if f1&0x0004 != 0 {
		access += " protected "
	}
	if f2&0x0008 != 0 {
		access += " static "
	}
	if f2&0x0010 != 0 {
		access += " final "
	}
	if f2&0x0040 != 0 {
		access += " volatile "
	}
	if f2&0x0080 != 0 {
		access += " transient "
	}
	if f2&0x1000 != 0 {
		access += " synthetic "
	}
	if f2&0x4000 != 0 {
		access += " enum "
	}

	return access

}