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

// #define
const (
	SDL_ALPHA_OPAQUE      = 255
	SDL_ALPHA_TRANSPARENT = 0
)

var (
	SDL_DEFINE_PIXELFORMAT = func(type_, order, layout, bits, bytes uint32) uint32 {
		ret := (1 << 28) | (type_ << 24) | (order << 20) | (layout << 16) | (bits << 8) | (bytes << 0)
		return uint32(ret)
	}
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
	SDL_PIXELFLAG     = func(x int) int { return (x >> 28) & 0x0F }
	SDL_PIXELTYPE     = func(x int) int { return (x >> 24) & 0x0F }
	SDL_PIXELORDER    = func(x int) int { return (x >> 20) & 0x0F }
	SDL_PIXELLAYOUT   = func(x int) int { return (x >> 16) & 0x0F }
	SDL_BITSPERPIXEL  = func(x int) int { return (x >> 8) & 0xFF }
	SDL_BYTESPERPIXEL = func(x int) int {
		if SDL_ISPIXELFORMAT_FOURCC(x) {
			if x == SDL_PIXELFORMAT_YUY2 ||
				x == SDL_PIXELFORMAT_UYVY ||
				x == SDL_PIXELFORMAT_YVYU {
				return 2
			} else {
				return 1
			}
		} else {
			return (x >> 0) & 0xFF
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
		if format && SDL_PIXELFLAG(format) != 1 {
			return true
		} else {
			return false
		}
	}
)

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
