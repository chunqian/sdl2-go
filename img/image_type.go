package img

/*
#include "image_wrapper.h"
*/
import "C"

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

// c #define
const (
	SDL_IMAGE_MAJOR_VERSION = 2
	SDL_IMAGE_MINOR_VERSION = 0
	SDL_IMAGE_PATCHLEVEL    = 5
)

// enum
const (
	IMG_INIT_JPG  IMG_InitFlags = 0x00000001
	IMG_INIT_PNG  IMG_InitFlags = 0x00000002
	IMG_INIT_TIF  IMG_InitFlags = 0x00000004
	IMG_INIT_WEBP IMG_InitFlags = 0x00000008
)

type IMG_InitFlags = int32
