package ttf

/*
#include "ttf_wrapper.h"
*/
import "C"
import (
	"reflect"
	"unsafe"

	. "github.com/chunqian/memory"
)

// c basic
type (
	cChar   = C.char   // byte
	cInt    = C.int    // int32
	cUint   = C.uint   // uint32
	cFloat  = C.float  // float32
	cDouble = C.double // float64

	cInt8   = C.int8_t   // int8
	cInt16  = C.int16_t  // int16
	cInt32  = C.int32_t  // int32
	cInt64  = C.int64_t  // int64
	cIntPtr = C.intptr_t // uintptr

	cUint8   = C.uint8_t   // uint8
	cUint16  = C.uint16_t  // uint16
	cUint32  = C.uint32_t  // uint32
	cUint64  = C.uint64_t  // uint64
	cUintPtr = C.uintptr_t // uintptr

	cSize  = C.size_t  // uint
	cWchar = C.wchar_t // windows uint16, unix int32

	cSchar  = C.schar  // int8
	cUchar  = C.uchar  // uint8
	cShort  = C.short  // int16
	cUshort = C.ushort // uint16

	cLong      = C.long      // int32
	cUlong     = C.ulong     // int32
	cLonglong  = C.longlong  // int64
	cUlonglong = C.ulonglong // uint64
)

// c enum
type cBool = C.SDL_bool

// c struct
type TTF_Font C.TTF_Font

func createCString(mp *PX_memorypool, s string) *cChar {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	n := PX_uint(sh.Len)
	p := MP_Malloc(mp, n+1)
	PX_memcpy(p, unsafe.Pointer(sh.Data), PX_int(n))
	*(*cChar)(unsafe.Pointer(uintptr(p) + uintptr(n))) = '\x00'
	return (*cChar)(p)
}

func destroyCString(mp *PX_memorypool, cStr *cChar) {
	MP_Free(mp, unsafe.Pointer(cStr))
}

func createGoString(cStr *cChar) string {
	len := PX_strlen((*PX_char)(unsafe.Pointer(cStr)))

	var slice []uint8
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sh.Data = uintptr(unsafe.Pointer(cStr))
	sh.Len = int(len)
	sh.Cap = int(len)
	src := *(*[]byte)(unsafe.Pointer(sh))

	return string(src)
}
