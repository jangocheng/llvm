// Code generated by "string2enum -linecomment -type UnnamedAddr ../../ir/enum"; DO NOT EDIT.

package enum

import "fmt"
import "github.com/llir/llvm/ir/enum"

const _UnnamedAddr_name = "nonelocal_unnamed_addrunnamed_addr"

var _UnnamedAddr_index = [...]uint8{0, 4, 22, 34}

func UnnamedAddrFromString(s string) enum.UnnamedAddr {
	if len(s) == 0 {
		return 0
	}
	for i := range _UnnamedAddr_index[:len(_UnnamedAddr_index)-1] {
		if s == _UnnamedAddr_name[_UnnamedAddr_index[i]:_UnnamedAddr_index[i+1]] {
			return enum.UnnamedAddr(i)
		}
	}
	panic(fmt.Errorf("unable to locate UnnamedAddr enum corresponding to %q", s))
}
