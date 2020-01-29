package cp

import "bufio"

type CP_Info interface {
	Parse(tag int, r *bufio.Reader)
	Print( index int,m map[int] interface{})
}
