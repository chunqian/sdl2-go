package ttf

/*
#include "ttf_wrapper.h"
*/
import "C"

// c basic
type (
	cBool   = C.char   // int8
	cChar   = C.char   // byte
	cInt    = C.int    // int32
	cUint   = C.uint   // uint32
	cFloat  = C.float  // float32
	cDouble = C.double // float64
	cSize_t = C.size_t // uint

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
	SDL_TTF_MAJOR_VERSION = 2
	SDL_TTF_MINOR_VERSION = 0
	SDL_TTF_PATCHLEVEL    = 18

	TTF_STYLE_NORMAL        = 0x00
	TTF_STYLE_BOLD          = 0x01
	TTF_STYLE_ITALIC        = 0x02
	TTF_STYLE_UNDERLINE     = 0x04
	TTF_STYLE_STRIKETHROUGH = 0x08

	TTF_HINTING_NORMAL         = 0
	TTF_HINTING_LIGHT          = 1
	TTF_HINTING_MONO           = 2
	TTF_HINTING_NONE           = 3
	TTF_HINTING_LIGHT_SUBPIXEL = 4
)

// c enum
type (
	cSDL_bool = C.SDL_bool
)

// c struct
type (
	TTF_Font C.TTF_Font
)
