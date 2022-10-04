package sdl

/*
#include "pixels.h"
*/
import "C"
import (
	"encoding/binary"
	"image/color"
	"unsafe"
)

// define
const (
	SDL_ALPHA_OPAQUE      = 255
	SDL_ALPHA_TRANSPARENT = 0
)

func SDL_DEFINE_PIXELFORMAT[T int32 | uint32](type_, order, layout, bits, bytes T) uint32 {
	ret := (1 << 28) | (type_ << 24) | (order << 20) | (layout << 16) | (bits << 8) | (bytes << 0)
	return uint32(ret)
}

var (
	SDL_DEFINE_PIXELFOURCC = func(a, b, c, d byte) uint32 {
		var endian binary.ByteOrder
		if IsLittleEndian() {
			endian = binary.LittleEndian
		} else {
			endian = binary.BigEndian
		}
		arr := []byte{a, b, c, d}
		return endian.Uint32(arr)
	}
	SDL_PIXELFLAG     = func(x int) uint32 { x2 := uint32(x); return (x2 >> 28) & 0x0F }
	SDL_PIXELTYPE     = func(x int) int32 { x2 := int32(x); return (x2 >> 24) & 0x0F }
	SDL_PIXELORDER    = func(x int) int32 { x2 := int32(x); return (x2 >> 20) & 0x0F }
	SDL_PIXELLAYOUT   = func(x int) uint32 { x2 := uint32(x); return (x2 >> 16) & 0x0F }
	SDL_BITSPERPIXEL  = func(x int) uint32 { x2 := uint32(x); return (x2 >> 8) & 0xFF }
	SDL_BYTESPERPIXEL = func(x int) uint32 {
		x2 := uint32(x)
		if SDL_ISPIXELFORMAT_FOURCC(x) {
			if x2 == SDL_PIXELFORMAT_YUY2 ||
				x2 == SDL_PIXELFORMAT_UYVY ||
				x2 == SDL_PIXELFORMAT_YVYU {
				return 2
			} else {
				return 1
			}
		} else {
			return (x2 >> 0) & 0xFF
		}
	}
	SDL_ISPIXELFORMAT_INDEXED = func(format int) bool {
		if !SDL_ISPIXELFORMAT_FOURCC(format) &&
			(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_INDEX1 ||
				SDL_PIXELTYPE(format) == SDL_PIXELTYPE_INDEX4 ||
				SDL_PIXELTYPE(format) == SDL_PIXELTYPE_INDEX8) {
			return true
		} else {
			return false
		}
	}
	SDL_ISPIXELFORMAT_PACKED = func(format int) bool {
		if !SDL_ISPIXELFORMAT_FOURCC(format) &&
			(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_PACKED8 ||
				SDL_PIXELTYPE(format) == SDL_PIXELTYPE_PACKED16 ||
				SDL_PIXELTYPE(format) == SDL_PIXELTYPE_PACKED32) {
			return true
		} else {
			return false
		}
	}
	SDL_ISPIXELFORMAT_ARRAY = func(format int) bool {
		if !SDL_ISPIXELFORMAT_FOURCC(format) &&
			(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_ARRAYU8 ||
				SDL_PIXELTYPE(format) == SDL_PIXELTYPE_ARRAYU16 ||
				SDL_PIXELTYPE(format) == SDL_PIXELTYPE_ARRAYU32 ||
				SDL_PIXELTYPE(format) == SDL_PIXELTYPE_ARRAYF16 ||
				SDL_PIXELTYPE(format) == SDL_PIXELTYPE_ARRAYF32) {
			return true
		} else {
			return false
		}
	}
	SDL_ISPIXELFORMAT_ALPHA = func(format int) bool {
		if (SDL_ISPIXELFORMAT_PACKED(format) &&
			(SDL_PIXELORDER(format) == SDL_PACKEDORDER_ARGB ||
				SDL_PIXELORDER(format) == SDL_PACKEDORDER_RGBA ||
				SDL_PIXELORDER(format) == SDL_PACKEDORDER_ABGR ||
				SDL_PIXELORDER(format) == SDL_PACKEDORDER_BGRA)) ||
			(SDL_ISPIXELFORMAT_ARRAY(format) &&
				(SDL_PIXELORDER(format) == SDL_ARRAYORDER_ARGB ||
					SDL_PIXELORDER(format) == SDL_ARRAYORDER_RGBA ||
					SDL_PIXELORDER(format) == SDL_ARRAYORDER_ABGR ||
					SDL_PIXELORDER(format) == SDL_ARRAYORDER_BGRA)) {
			return true
		} else {
			return false
		}
	}
	SDL_ISPIXELFORMAT_FOURCC = func(format int) bool {
		if SDL_PIXELFLAG(format) != 1 {
			return true
		} else {
			return false
		}
	}
)

// typedef
type SDL_PixelType = int32
type SDL_BitmapOrder = int32
type SDL_PackedOrder = int32
type SDL_ArrayOrder = int32
type SDL_PackedLayout = int32
type SDL_PixelFormatEnum = uint32

// enum
const (
	SDL_PIXELTYPE_UNKNOWN  SDL_PixelType = 0
	SDL_PIXELTYPE_INDEX1   SDL_PixelType = 1
	SDL_PIXELTYPE_INDEX4   SDL_PixelType = 2
	SDL_PIXELTYPE_INDEX8   SDL_PixelType = 3
	SDL_PIXELTYPE_PACKED8  SDL_PixelType = 4
	SDL_PIXELTYPE_PACKED16 SDL_PixelType = 5
	SDL_PIXELTYPE_PACKED32 SDL_PixelType = 6
	SDL_PIXELTYPE_ARRAYU8  SDL_PixelType = 7
	SDL_PIXELTYPE_ARRAYU16 SDL_PixelType = 8
	SDL_PIXELTYPE_ARRAYU32 SDL_PixelType = 9
	SDL_PIXELTYPE_ARRAYF16 SDL_PixelType = 10
	SDL_PIXELTYPE_ARRAYF32 SDL_PixelType = 11
)

const (
	SDL_BITMAPORDER_NONE SDL_BitmapOrder = 0
	SDL_BITMAPORDER_4321 SDL_BitmapOrder = 1
	SDL_BITMAPORDER_1234 SDL_BitmapOrder = 2
)

const (
	SDL_PACKEDORDER_NONE SDL_PackedOrder = 0
	SDL_PACKEDORDER_XRGB SDL_PackedOrder = 1
	SDL_PACKEDORDER_RGBX SDL_PackedOrder = 2
	SDL_PACKEDORDER_ARGB SDL_PackedOrder = 3
	SDL_PACKEDORDER_RGBA SDL_PackedOrder = 4
	SDL_PACKEDORDER_XBGR SDL_PackedOrder = 5
	SDL_PACKEDORDER_BGRX SDL_PackedOrder = 6
	SDL_PACKEDORDER_ABGR SDL_PackedOrder = 7
	SDL_PACKEDORDER_BGRA SDL_PackedOrder = 8
)

const (
	SDL_ARRAYORDER_NONE SDL_ArrayOrder = 0
	SDL_ARRAYORDER_RGB  SDL_ArrayOrder = 1
	SDL_ARRAYORDER_RGBA SDL_ArrayOrder = 2
	SDL_ARRAYORDER_ARGB SDL_ArrayOrder = 3
	SDL_ARRAYORDER_BGR  SDL_ArrayOrder = 4
	SDL_ARRAYORDER_BGRA SDL_ArrayOrder = 5
	SDL_ARRAYORDER_ABGR SDL_ArrayOrder = 6
)

const (
	SDL_PACKEDLAYOUT_NONE    SDL_PackedLayout = 0
	SDL_PACKEDLAYOUT_332     SDL_PackedLayout = 1
	SDL_PACKEDLAYOUT_4444    SDL_PackedLayout = 2
	SDL_PACKEDLAYOUT_1555    SDL_PackedLayout = 3
	SDL_PACKEDLAYOUT_5551    SDL_PackedLayout = 4
	SDL_PACKEDLAYOUT_565     SDL_PackedLayout = 5
	SDL_PACKEDLAYOUT_8888    SDL_PackedLayout = 6
	SDL_PACKEDLAYOUT_2101010 SDL_PackedLayout = 7
	SDL_PACKEDLAYOUT_1010102 SDL_PackedLayout = 8
)

var (
	SDL_PIXELFORMAT_UNKNOWN     SDL_PixelFormatEnum = 0
	SDL_PIXELFORMAT_INDEX1LSB   SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX1, SDL_BITMAPORDER_4321, 0, 1, 0)
	SDL_PIXELFORMAT_INDEX1MSB   SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX1, SDL_BITMAPORDER_1234, 0, 1, 0)
	SDL_PIXELFORMAT_INDEX4LSB   SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX4, SDL_BITMAPORDER_4321, 0, 4, 0)
	SDL_PIXELFORMAT_INDEX4MSB   SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX4, SDL_BITMAPORDER_1234, 0, 4, 0)
	SDL_PIXELFORMAT_INDEX8      SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX8, 0, 0, 8, 1)
	SDL_PIXELFORMAT_RGB332      SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED8, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_332, 8, 1)
	SDL_PIXELFORMAT_XRGB4444    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_4444, 12, 2)
	SDL_PIXELFORMAT_RGB444      SDL_PixelFormatEnum = SDL_PIXELFORMAT_XRGB4444
	SDL_PIXELFORMAT_XBGR4444    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_4444, 12, 2)
	SDL_PIXELFORMAT_BGR444      SDL_PixelFormatEnum = SDL_PIXELFORMAT_XBGR4444
	SDL_PIXELFORMAT_XRGB1555    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_1555, 15, 2)
	SDL_PIXELFORMAT_RGB555      SDL_PixelFormatEnum = SDL_PIXELFORMAT_XRGB1555
	SDL_PIXELFORMAT_XBGR1555    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_1555, 15, 2)
	SDL_PIXELFORMAT_BGR555      SDL_PixelFormatEnum = SDL_PIXELFORMAT_XBGR1555
	SDL_PIXELFORMAT_ARGB4444    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_4444, 16, 2)
	SDL_PIXELFORMAT_RGBA4444    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_RGBA, SDL_PACKEDLAYOUT_4444, 16, 2)
	SDL_PIXELFORMAT_ABGR4444    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ABGR, SDL_PACKEDLAYOUT_4444, 16, 2)
	SDL_PIXELFORMAT_BGRA4444    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_BGRA, SDL_PACKEDLAYOUT_4444, 16, 2)
	SDL_PIXELFORMAT_ARGB1555    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_1555, 16, 2)
	SDL_PIXELFORMAT_RGBA5551    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_RGBA, SDL_PACKEDLAYOUT_5551, 16, 2)
	SDL_PIXELFORMAT_ABGR1555    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ABGR, SDL_PACKEDLAYOUT_1555, 16, 2)
	SDL_PIXELFORMAT_BGRA5551    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_BGRA, SDL_PACKEDLAYOUT_5551, 16, 2)
	SDL_PIXELFORMAT_RGB565      SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_565, 16, 2)
	SDL_PIXELFORMAT_BGR565      SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_565, 16, 2)
	SDL_PIXELFORMAT_RGB24       SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU8, SDL_ARRAYORDER_RGB, 0, 24, 3)
	SDL_PIXELFORMAT_BGR24       SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU8, SDL_ARRAYORDER_BGR, 0, 24, 3)
	SDL_PIXELFORMAT_XRGB8888    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_8888, 24, 4)
	SDL_PIXELFORMAT_RGB888      SDL_PixelFormatEnum = SDL_PIXELFORMAT_XRGB8888
	SDL_PIXELFORMAT_RGBX8888    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_RGBX, SDL_PACKEDLAYOUT_8888, 24, 4)
	SDL_PIXELFORMAT_XBGR8888    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_8888, 24, 4)
	SDL_PIXELFORMAT_BGR888      SDL_PixelFormatEnum = SDL_PIXELFORMAT_XBGR8888
	SDL_PIXELFORMAT_BGRX8888    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_BGRX, SDL_PACKEDLAYOUT_8888, 24, 4)
	SDL_PIXELFORMAT_ARGB8888    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_8888, 32, 4)
	SDL_PIXELFORMAT_RGBA8888    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_RGBA, SDL_PACKEDLAYOUT_8888, 32, 4)
	SDL_PIXELFORMAT_ABGR8888    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ABGR, SDL_PACKEDLAYOUT_8888, 32, 4)
	SDL_PIXELFORMAT_BGRA8888    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_BGRA, SDL_PACKEDLAYOUT_8888, 32, 4)
	SDL_PIXELFORMAT_ARGB2101010 SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_2101010, 32, 4)
	SDL_PIXELFORMAT_RGBA32      SDL_PixelFormatEnum = func() uint32 {
		if !IsLittleEndian() {
			return SDL_PIXELFORMAT_RGBA8888
		} else {
			return SDL_PIXELFORMAT_ABGR8888
		}
	}()
	SDL_PIXELFORMAT_ARGB32 SDL_PixelFormatEnum = func() uint32 {
		if !IsLittleEndian() {
			return SDL_PIXELFORMAT_ARGB8888
		} else {
			return SDL_PIXELFORMAT_BGRA8888
		}
	}()
	SDL_PIXELFORMAT_BGRA32 SDL_PixelFormatEnum = func() uint32 {
		if !IsLittleEndian() {
			return SDL_PIXELFORMAT_BGRA8888
		} else {
			return SDL_PIXELFORMAT_ARGB8888
		}
	}()
	SDL_PIXELFORMAT_ABGR32 SDL_PixelFormatEnum = func() uint32 {
		if !IsLittleEndian() {
			return SDL_PIXELFORMAT_ABGR8888
		} else {
			return SDL_PIXELFORMAT_RGBA8888
		}
	}()
	SDL_PIXELFORMAT_YV12         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('Y', 'V', '1', '2')
	SDL_PIXELFORMAT_IYUV         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('I', 'Y', 'U', 'V')
	SDL_PIXELFORMAT_YUY2         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('Y', 'U', 'Y', '2')
	SDL_PIXELFORMAT_UYVY         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('U', 'Y', 'V', 'Y')
	SDL_PIXELFORMAT_YVYU         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('Y', 'V', 'Y', 'U')
	SDL_PIXELFORMAT_NV12         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('N', 'V', '1', '2')
	SDL_PIXELFORMAT_NV21         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('N', 'V', '2', '1')
	SDL_PIXELFORMAT_EXTERNAL_OES SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('O', 'E', 'S', ' ')
)

// struct
// Color represents a color. This implements image/color.Color interface.
type SDL_Color color.RGBA

// type SDL_Color struct {
// 	R uint8
// 	G uint8
// 	B uint8
// 	A uint8
// }

type SDL_Palette struct {
	Ncolors  int32
	Colors   *SDL_Color
	Version  uint32
	Refcount int32
}

type SDL_PixelFormat struct {
	Format        uint32
	Palette       *SDL_Palette
	BitsPerPixel  uint8
	BytesPerPixel uint8
	Padding       [2]uint8
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	Refcount      int32
	Next          *SDL_PixelFormat
}

func cPixelFormat(format *SDL_PixelFormat) *C.SDL_PixelFormat {
	return (*C.SDL_PixelFormat)(unsafe.Pointer(format))
}

func cPalette(palette *SDL_Palette) *C.SDL_Palette {
	return (*C.SDL_Palette)(unsafe.Pointer(palette))
}

func cColor(c *SDL_Color) *C.SDL_Color {
	return (*C.SDL_Color)(unsafe.Pointer(c))
}

func pixelColor(c SDL_Color) uint32 {
	var v uint32
	v |= uint32(c.R) << 24
	v |= uint32(c.G) << 16
	v |= uint32(c.B) << 8
	v |= uint32(c.A)
	return v
}

func SDL_GetPixelFormatName(format uint) string {
	cName := C.SDL_GetPixelFormatName(cUint32(format))
	if cName == nil {
		return ""
	}
	return SDL_GoString(cName)
}

func SDL_PixelFormatEnumToMasks(format uint, bpp *int, rmask, gmask, bmask, amask *uint32) bool {
	cBool := C.SDL_PixelFormatEnumToMasks(
		cUint32(format),
		(*cInt)(unsafe.Pointer(bpp)),
		(*cUint32)(rmask),
		(*cUint32)(gmask),
		(*cUint32)(bmask),
		(*cUint32)(amask))
	if int(cBool) == 0 {
		return false
	}
	return true
}

func MasksToPixelFormatEnum(bpp int, rmask, gmask, bmask, amask uint32) uint {
	cRet := C.SDL_MasksToPixelFormatEnum(
		cInt(bpp),
		cUint32(rmask),
		cUint32(gmask),
		cUint32(bmask),
		cUint32(amask))
	return uint(cRet)
}

func SDL_AllocFormat(format uint32) *SDL_PixelFormat {
	cPixelFormat := C.SDL_AllocFormat(cUint32(format))
	pixelFormat := (*SDL_PixelFormat)(unsafe.Pointer(cPixelFormat))
	return pixelFormat
}

func SDL_FreeFormat(format *SDL_PixelFormat) {
	C.SDL_FreeFormat((*C.SDL_PixelFormat)(unsafe.Pointer(format)))
}

func SDL_AllocPalette(ncolors int) *SDL_Palette {
	cPalette := C.SDL_AllocPalette(cInt(ncolors))
	palette := (*SDL_Palette)(unsafe.Pointer(cPalette))
	return palette
}

func SDL_SetPixelFormatPalette(format *SDL_PixelFormat, palette *SDL_Palette) int {
	cRet := C.SDL_SetPixelFormatPalette(cPixelFormat(format), cPalette(palette))
	return int(cRet)
}

func SDL_SetPaletteColors(palette *SDL_Palette, colors []SDL_Color) int {
	if colors == nil {
		return -1
	}
	var ptr *SDL_Color
	if len(colors) > 0 {
		ptr = &colors[0]
	}

	cRet := C.SDL_SetPaletteColors(cPalette(palette), cColor(ptr), 0, cInt(len(colors)))
	return int(cRet)
}

func SDL_FreePalette(palette *SDL_Palette) {
	C.SDL_FreePalette(cPalette(palette))
}

func SDL_MapRGB(format *SDL_PixelFormat, r, g, b uint8) uint32 {
	cRet := C.SDL_MapRGB(cPixelFormat(format), cUint8(r), cUint8(g), cUint8(b))
	return uint32(cRet)
}

func SDL_MapRGBA(format *SDL_PixelFormat, r, g, b, a uint8) uint32 {
	cRet := C.SDL_MapRGBA(cPixelFormat(format), cUint8(r), cUint8(g), cUint8(b), cUint8(a))
	return uint32(cRet)
}

func SDL_GetRGB(pixel uint32, format *SDL_PixelFormat, r, g, b *uint8) {
	C.SDL_GetRGB(
		cUint32(pixel),
		cPixelFormat(format),
		(*cUint8)(r),
		(*cUint8)(g),
		(*cUint8)(b))
}

func SDL_GetRGBA(pixel uint32, format *SDL_PixelFormat, r, g, b, a *uint8) {
	C.SDL_GetRGBA(
		cUint32(pixel),
		cPixelFormat(format),
		(*cUint8)(r),
		(*cUint8)(g),
		(*cUint8)(b),
		(*cUint8)(a))
}

func SDL_CalculateGammaRamp(gamma float32, ramp *[256]uint16) {
	C.SDL_CalculateGammaRamp(cFloat(gamma), (*cUint16)(unsafe.Pointer(&ramp[0])))
}

func SDL_BytesPerPixel(format uint32) int {
	return int(C.SDL_BytesPerPixel(cUint32(format)))
}

func SDL_BitsPerPixel(format uint32) int {
	return int(C.SDL_BitsPerPixel(cUint32(format)))
}

type RGB444 struct {
	R, G, B byte
}

func (c RGB444) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 12
	g = uint32(c.G) << 12
	b = uint32(c.B) << 12
	return
}

var (
	RGB444Model   color.Model = color.ModelFunc(rgb444Model)
	RGB332Model   color.Model = color.ModelFunc(rgb332Model)
	RGB565Model   color.Model = color.ModelFunc(rgb565Model)
	RGB555Model   color.Model = color.ModelFunc(rgb555Model)
	BGR565Model   color.Model = color.ModelFunc(bgr565Model)
	BGR555Model   color.Model = color.ModelFunc(bgr555Model)
	RGB888Model   color.Model = color.ModelFunc(rgb888Model)
	BGR888Model   color.Model = color.ModelFunc(bgr888Model)
	ARGB4444Model color.Model = color.ModelFunc(argb4444Model)
	ABGR4444Model color.Model = color.ModelFunc(abgr4444Model)
	RGBA4444Model color.Model = color.ModelFunc(rgba4444Model)
	BGRA4444Model color.Model = color.ModelFunc(bgra4444Model)
	ARGB1555Model color.Model = color.ModelFunc(argb1555Model)
	RGBA5551Model color.Model = color.ModelFunc(rgba5551Model)
	ABGR1555Model color.Model = color.ModelFunc(abgr1555Model)
	BGRA5551Model color.Model = color.ModelFunc(bgra5551Model)
	RGBA8888Model color.Model = color.ModelFunc(rgba8888Model)
	BGRA8888Model color.Model = color.ModelFunc(bgra8888Model)
	ARGB8888Model color.Model = color.ModelFunc(argb8888Model)
	ABGR8888Model color.Model = color.ModelFunc(abgr8888Model)
)

func rgb444Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB444{uint8(r >> 12), uint8(g >> 12), uint8(b >> 12)}
}

type RGB332 struct {
	R, G, B byte
}

func (c RGB332) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 13
	g = uint32(c.G) << 13
	b = uint32(c.B) << 14
	return
}

func rgb332Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB332{uint8(r >> 13), uint8(g >> 13), uint8(b >> 14)}
}

type RGB565 struct {
	R, G, B byte
}

func (c RGB565) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 11
	g = uint32(c.G) << 10
	b = uint32(c.B) << 11
	return
}

func rgb565Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB565{uint8(r >> 11), uint8(g >> 10), uint8(b >> 11)}
}

type RGB555 struct {
	R, G, B byte
}

func (c RGB555) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 11
	g = uint32(c.G) << 11
	b = uint32(c.B) << 11
	return
}

func rgb555Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB555{uint8(r >> 11), uint8(g >> 11), uint8(b >> 11)}
}

type BGR565 struct {
	B, G, R byte
}

func (c BGR565) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 11
	g = uint32(c.G) << 10
	b = uint32(c.B) << 11
	return
}

func bgr565Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return BGR565{uint8(b >> 11), uint8(g >> 10), uint8(r >> 11)}
}

type BGR555 struct {
	B, G, R byte
}

func (c BGR555) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 11
	g = uint32(c.G) << 11
	b = uint32(c.B) << 11
	return
}

func bgr555Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return BGR555{uint8(b >> 11), uint8(g >> 11), uint8(r >> 11)}
}

type RGB888 struct {
	R, G, B byte
}

func (c RGB888) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	g = uint32(c.G)
	b = uint32(c.B)
	a = 255
	return
}

func rgb888Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB888{uint8(r), uint8(g), uint8(b)}
}

type BGR888 struct {
	B, G, R byte
}

func (c BGR888) RGBA() (r, g, b, a uint32) {
	b = uint32(c.B)
	g = uint32(c.G)
	r = uint32(c.R)
	a = 255
	return
}

func bgr888Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return BGR888{uint8(b), uint8(g), uint8(r)}
}

type ARGB4444 struct {
	A, R, G, B byte
}

func (c ARGB4444) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A) << 4
	r = uint32(c.R) << 4
	g = uint32(c.G) << 4
	b = uint32(c.B) << 4
	return
}

func argb4444Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return ARGB4444{uint8(a >> 4), uint8(r >> 4), uint8(g >> 4), uint8(b >> 4)}
}

type ABGR4444 struct {
	A, B, G, R byte
}

func (c ABGR4444) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A) << 4
	r = uint32(c.R) << 4
	g = uint32(c.G) << 4
	b = uint32(c.B) << 4
	return
}

func abgr4444Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return ABGR4444{uint8(a >> 4), uint8(b >> 4), uint8(g >> 4), uint8(r >> 4)}
}

type RGBA4444 struct {
	R, G, B, A byte
}

func (c RGBA4444) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 4
	g = uint32(c.G) << 4
	b = uint32(c.B) << 4
	a = uint32(c.A) << 4
	return
}

func rgba4444Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA4444{uint8(r >> 4), uint8(g >> 4), uint8(b >> 4), uint8(a >> 4)}
}

type BGRA4444 struct {
	B, G, R, A byte
}

func (c BGRA4444) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 4
	g = uint32(c.G) << 4
	b = uint32(c.B) << 4
	a = uint32(c.A) << 4
	return
}

func bgra4444Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return BGRA4444{uint8(b >> 4), uint8(g >> 4), uint8(r >> 4), uint8(a >> 4)}
}

type ARGB1555 struct {
	A, R, G, B byte
}

func (c ARGB1555) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 3
	g = uint32(c.G) << 3
	b = uint32(c.B) << 3
	if c.A > 0 {
		tmp := int32(-1)
		a = uint32(tmp)
	}
	return
}

func argb1555Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a > 0 {
		a = 1
	} else {
		a = 0
	}
	return ARGB1555{uint8(a), uint8(r >> 3), uint8(g >> 3), uint8(b >> 3)}
}

type RGBA5551 struct {
	R, G, B, A byte
}

func (c RGBA5551) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 3
	g = uint32(c.G) << 3
	b = uint32(c.B) << 3
	if c.A > 0 {
		tmp := int32(-1)
		a = uint32(tmp)
	}
	return
}

func rgba5551Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a > 0 {
		a = 1
	} else {
		a = 0
	}
	return RGBA5551{uint8(r >> 3), uint8(g >> 3), uint8(b >> 3), uint8(a)}
}

type ABGR1555 struct {
	A, R, G, B byte
}

func (c ABGR1555) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 3
	g = uint32(c.G) << 3
	b = uint32(c.B) << 3
	if c.A > 0 {
		tmp := int32(-1)
		a = uint32(tmp)
	}
	return
}

func abgr1555Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a > 0 {
		a = 1
	} else {
		a = 0
	}
	return ABGR1555{uint8(a), uint8(r >> 3), uint8(g >> 3), uint8(b >> 3)}
}

type BGRA5551 struct {
	B, G, R, A byte
}

func (c BGRA5551) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 3
	g = uint32(c.G) << 3
	b = uint32(c.B) << 3
	if c.A > 0 {
		tmp := int32(-1)
		a = uint32(tmp)
	}
	return
}

func bgra5551Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a > 0 {
		a = 1
	} else {
		a = 0
	}
	return BGRA5551{uint8(b >> 3), uint8(g >> 3), uint8(r >> 3), uint8(a)}
}

type RGBA8888 struct {
	R, G, B, A byte
}

func (c RGBA8888) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	g = uint32(c.G)
	b = uint32(c.B)
	a = uint32(c.A)
	return
}

func rgba8888Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA8888{uint8(r), uint8(g), uint8(b), uint8(a)}
}

type BGRA8888 struct {
	B, G, R, A byte
}

func (c BGRA8888) RGBA() (r, g, b, a uint32) {
	b = uint32(c.B)
	g = uint32(c.G)
	r = uint32(c.R)
	a = uint32(c.A)
	return
}

func bgra8888Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return BGRA8888{uint8(b), uint8(g), uint8(r), uint8(a)}
}

type ARGB8888 struct {
	A, R, G, B byte
}

func (c ARGB8888) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A)
	r = uint32(c.R)
	g = uint32(c.G)
	b = uint32(c.B)
	return
}

func argb8888Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return ARGB8888{uint8(a), uint8(r), uint8(g), uint8(b)}
}

type ABGR8888 struct {
	A, B, G, R byte
}

func (c ABGR8888) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A)
	r = uint32(c.R)
	g = uint32(c.G)
	b = uint32(c.B)
	return
}

func abgr8888Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return ABGR8888{uint8(a), uint8(b), uint8(g), uint8(r)}
}
