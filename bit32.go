package gluabit32

import (
	"math/bits"

	"github.com/yuin/gopher-lua"
)

func band(ls *lua.LState) uint32 {
	x := ^uint32(0)
	for i, top := 1, ls.GetTop(); i <= top; i += 1 {
		x &= uint32(ls.CheckNumber(i))
	}
	return x
}
func fieldWidthMaskArg(ls *lua.LState, arg int) (uint32, uint32) {
	field := int32(ls.CheckNumber(arg))
	if field < 0 {
		ls.RaiseError("field cannot be negative")
	}
	var width int32
	lwidth := ls.Get(arg + 1)
	if w, ok := lwidth.(lua.LNumber); ok {
		if w <= 0 {
			ls.RaiseError("width must be positive")
		}
		width = int32(w)
	} else if lwidth == lua.LNil {
		width = 1
	} else {
		ls.RaiseError("invalid width type")
		return 0, 0
	}
	if field+width > 32 {
		ls.RaiseError("trying to access non-existent bits")
	}
	return uint32(field), ^(^uint32(0) << uint32(width))
}

func Loader(ls *lua.LState) int {
	m := ls.SetFuncs(ls.NewTable(), map[string]lua.LGFunction{
		"arshift": Bit32arshift,
		"band":    Bit32band,
		"bnot":    Bit32bnot,
		"bor":     Bit32bor,
		"btest":   Bit32btest,
		"bxor":    Bit32bxor,
		"extract": Bit32extract,
		"replace": Bit32replace,
		"lrotate": Bit32lrotate,
		"lshift":  Bit32lshift,
		"rrotate": Bit32rrotate,
		"rshift":  Bit32rshift,
	})
	ls.Push(m)
	return 1
}

func Bit32arshift(ls *lua.LState) int {
	x := int32(ls.CheckNumber(1))
	disp := int32(ls.CheckNumber(2))
	if disp >= 0 {
		ls.Push(lua.LNumber(uint32(x >> disp)))
	} else {
		ls.Push(lua.LNumber(uint32(x) << -disp))
	}
	return 1
}
func Bit32band(ls *lua.LState) int {
	ls.Push(lua.LNumber(band(ls)))
	return 1
}
func Bit32bnot(ls *lua.LState) int {
	ls.Push(lua.LNumber(^uint32(ls.CheckNumber(1))))
	return 1
}
func Bit32bor(ls *lua.LState) int {
	x := uint32(0)
	for i, top := 1, ls.GetTop(); i <= top; i += 1 {
		x |= uint32(ls.CheckNumber(i))
	}
	ls.Push(lua.LNumber(x))
	return 1
}
func Bit32btest(ls *lua.LState) int {
	ls.Push(lua.LBool(band(ls) != 0))
	return 1
}
func Bit32bxor(ls *lua.LState) int {
	x := uint32(0)
	for i, top := 1, ls.GetTop(); i <= top; i += 1 {
		x ^= uint32(ls.CheckNumber(i))
	}
	ls.Push(lua.LNumber(x))
	return 1
}
func Bit32extract(ls *lua.LState) int {
	n := uint32(ls.CheckNumber(1))
	field, m := fieldWidthMaskArg(ls, 2)
	ls.Push(lua.LNumber((n >> field) & m))
	return 1
}
func Bit32lrotate(ls *lua.LState) int {
	x := uint32(ls.CheckNumber(1))
	k := int(ls.CheckNumber(2))
	ls.Push(lua.LNumber(bits.RotateLeft32(x, k)))
	return 1
}
func Bit32lshift(ls *lua.LState) int {
	x := uint32(ls.CheckNumber(1))
	disp := int32(ls.CheckNumber(2))
	if disp >= 0 {
		ls.Push(lua.LNumber(x << disp))
	} else {
		ls.Push(lua.LNumber(x >> -disp))
	}
	return 1
}
func Bit32replace(ls *lua.LState) int {
	n := uint32(ls.CheckNumber(1))
	v := uint32(ls.CheckNumber(2))
	field, m := fieldWidthMaskArg(ls, 3)
	ls.Push(lua.LNumber((n & ^(m << field)) | ((v & m) << field)))
	return 1
}
func Bit32rrotate(ls *lua.LState) int {
	x := uint32(ls.CheckNumber(1))
	k := int(ls.CheckNumber(2))
	ls.Push(lua.LNumber(bits.RotateLeft32(x, -k)))
	return 1
}
func Bit32rshift(ls *lua.LState) int {
	x := uint32(ls.CheckNumber(1))
	disp := int32(ls.CheckNumber(2))
	if disp >= 0 {
		ls.Push(lua.LNumber(x >> disp))
	} else {
		ls.Push(lua.LNumber(x << -disp))
	}
	return 1
}
