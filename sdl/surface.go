package sdl

/*
#include "surface.h"
*/
import "C"
import (
	"image"
	"image/color"
	"reflect"
	"unsafe"
)

func (surface *SDL_Surface) cSDL_Surface() *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(surface))
}

// MustLock reports whether the surface must be locked for access.
func (surface *SDL_Surface) MustLock() bool {
	if (surface.Flags & SDL_RLEACCEL) != 0 {
		return true
	}
	return false
}

// CreateRGBSurface allocates a new RGB surface.
func SDL_CreateRGBSurface(flags uint32, width, height, depth int32, Rmask, Gmask, Bmask, Amask uint32) *SDL_Surface {
	cSurface := C.SDL_CreateRGBSurface(
		cUint32(flags),
		cInt(width),
		cInt(height),
		cInt(depth),
		cUint32(Rmask),
		cUint32(Gmask),
		cUint32(Bmask),
		cUint32(Amask))
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

// CreateRGBSurfaceFrom allocate a new RGB surface with existing pixel data.
func SDL_CreateRGBSurfaceFrom(pixels unsafe.Pointer, width, height int32, depth, pitch int, Rmask, Gmask, Bmask, Amask uint32) *SDL_Surface {
	cSurface := C.SDL_CreateRGBSurfaceFrom(
		pixels,
		cInt(width),
		cInt(height),
		cInt(depth),
		cInt(pitch),
		cUint32(Rmask),
		cUint32(Gmask),
		cUint32(Bmask),
		cUint32(Amask))
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

// CreateRGBSurfaceWithFormat allocates an RGB surface.
func SDL_CreateRGBSurfaceWithFormat(flags uint32, width, height, depth int32, format uint32) *SDL_Surface {
	cSurface := C.SDL_CreateRGBSurfaceWithFormat(
		cUint32(flags),
		cInt(width),
		cInt(height),
		cInt(depth),
		cUint32(format))
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

// CreateRGBSurfaceWithFormatFrom allocates an RGB surface from provided pixel data.
func SDL_CreateRGBSurfaceWithFormatFrom(pixels unsafe.Pointer, width, height, depth, pitch int32, format uint32) *SDL_Surface {
	cSurface := C.SDL_CreateRGBSurfaceWithFormatFrom(
		pixels,
		cInt(width),
		cInt(height),
		cInt(depth),
		cInt(pitch),
		cUint32(format))
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

// SetYUVConversionMode sets the YUV conversion mode
func SetYUVConversionMode(mode SDL_YUV_CONVERSION_MODE) {
	cMode := C.SDL_YUV_CONVERSION_MODE(mode)
	C.SDL_SetYUVConversionMode(cMode)
}

// GetYUVConversionMode gets the YUV conversion mode
func GetYUVConversionMode() SDL_YUV_CONVERSION_MODE {
	return SDL_YUV_CONVERSION_MODE(C.SDL_GetYUVConversionMode())
}

// GetYUVConversionModeForResolution gets the YUV conversion mode
func GetYUVConversionModeForResolution(width, height int) SDL_YUV_CONVERSION_MODE {
	cWidth := cInt(width)
	cHeight := cInt(height)
	return SDL_YUV_CONVERSION_MODE(C.SDL_GetYUVConversionModeForResolution(cWidth, cHeight))
}

// Free frees the RGB surface.
func SDL_FreeSurface(surface *SDL_Surface) {
	C.SDL_FreeSurface(surface.cSDL_Surface())
}

// SetPalette sets the palette used by the surface.
func SDL_SetSurfacePalette(surface *SDL_Surface, palette *SDL_Palette) int {
	cRet := C.SDL_SetSurfacePalette(surface.cSDL_Surface(), palette.cSDL_Palette())
	return int(cRet)
}

// Lock sets up the surface for directly accessing the pixels.
func SDL_LockSurface(surface *SDL_Surface) int {
	cRet := C.SDL_LockSurface(surface.cSDL_Surface())
	return int(cRet)
}

// Unlock releases the surface after directly accessing the pixels.
func SDL_UnlockSurface(surface *SDL_Surface) {
	C.SDL_UnlockSurface(surface.cSDL_Surface())
}

// LoadBMPRW loads a BMP image from a seekable SDL data stream (memory or file).
func SDL_LoadBMP_RW(src *SDL_RWops, freeSrc SDL_bool) *SDL_Surface {
	cSurface := C.SDL_LoadBMP_RW(src.cSDL_RWops(), cInt(freeSrc))
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

// LoadBMP loads a surface from a BMP file.
func LoadBMP(file string) *SDL_Surface {
	return SDL_LoadBMP_RW(SDL_RWFromFile(file, "rb"), SDL_TRUE)
}

// SaveBMPRW save the surface to a seekable SDL data stream (memory or file) in BMP format.
func SDL_SaveBMP_RW(surface *SDL_Surface, dst *SDL_RWops, freeDst SDL_bool) int {
	cRet := C.SDL_SaveBMP_RW(surface.cSDL_Surface(), dst.cSDL_RWops(), cInt(freeDst))
	return int(cRet)
}

// SaveBMP saves the surface to a BMP file.
func (surface *SDL_Surface) SaveBMP(file string) int {
	return SDL_SaveBMP_RW(surface, SDL_RWFromFile(file, "wb"), SDL_TRUE)
}

// SetRLE sets the RLE acceleration hint for the surface.
func SDL_SetSurfaceRLE(surface *SDL_Surface, flag SDL_bool) int {
	cRet := C.SDL_SetSurfaceRLE(surface.cSDL_Surface(), cInt(flag))
	return int(cRet)
}

// HasRLE returns whether the surface is RLE enabled.
func SDL_HasSurfaceRLE(surface *SDL_Surface) bool {
	cBool := C.SDL_HasSurfaceRLE(surface.cSDL_Surface())
	if int(cBool) == 1 {
		return true
	}
	return false
}

// SetColorKey sets the color key (transparent pixel) in the surface.
func SDL_SetColorKey(surface *SDL_Surface, flag SDL_bool, key uint32) int {
	cRet := C.SDL_SetColorKey(surface.cSDL_Surface(), cInt(flag), cUint32(key))
	return int(cRet)
}

// HasColorKey returns the color key (transparent pixel) for the surface.
func (surface *SDL_Surface) HasColorKey() bool {
	cBool := C.SDL_HasColorKey(surface.cSDL_Surface())
	if int(cBool) == 1 {
		return true
	}
	return false
}

// GetColorKey returns the color key (transparent pixel) for the surface.
func SDL_GetColorKey(surface *SDL_Surface, key *uint32) int {
	cKey := (*cUint32)(unsafe.Pointer(key))
	cRet := C.SDL_GetColorKey(surface.cSDL_Surface(), cKey)
	return int(cRet)
}

// SetColorMod sets an additional color value multiplied into blit operations.
func SDL_SetSurfaceColorMod(surface *SDL_Surface, r, g, b uint8) int {
	cRet := C.SDL_SetSurfaceColorMod(surface.cSDL_Surface(), cUint8(r), cUint8(g), cUint8(b))
	return int(cRet)
}

// GetColorMod returns the additional color value multiplied into blit operations.
func SDL_GetSurfaceColorMod(surface *SDL_Surface, r, g, b *uint8) int {
	cR := (*cUint8)(unsafe.Pointer(r))
	cG := (*cUint8)(unsafe.Pointer(g))
	cB := (*cUint8)(unsafe.Pointer(b))
	cRet := C.SDL_GetSurfaceColorMod(surface.cSDL_Surface(), cR, cG, cB)
	return int(cRet)
}

// SetAlphaMod sets an additional alpha value used in blit operations.
func SDL_SetSurfaceAlphaMod(surface *SDL_Surface, alpha uint8) int {
	cRet := C.SDL_SetSurfaceAlphaMod(surface.cSDL_Surface(), cUint8(alpha))
	return int(cRet)
}

// GetAlphaMod returns the additional alpha value used in blit operations.
func SDL_GetSurfaceAlphaMod(surface *SDL_Surface, alpha *uint8) int {
	cAlpha := (*cUint8)(unsafe.Pointer(alpha))
	cRet := C.SDL_GetSurfaceAlphaMod(surface.cSDL_Surface(), cAlpha)
	return int(cRet)
}

// SetBlendMode sets the blend mode used for blit operations.
func SDL_SetSurfaceBlendMode(surface *SDL_Surface, bm SDL_BlendMode) int {
	cRet := C.SDL_SetSurfaceBlendMode(surface.cSDL_Surface(), cSDL_BlendMode(bm))
	return int(cRet)
}

// GetBlendMode returns the blend mode used for blit operations.
func SDL_GetSurfaceBlendMode(surface *SDL_Surface, bm *SDL_BlendMode) int {
	cBM := (*cSDL_BlendMode)(unsafe.Pointer(bm))
	cRet := C.SDL_GetSurfaceBlendMode(surface.cSDL_Surface(), cBM)
	return int(cRet)
}

// SetClipRect sets the clipping rectangle for the surface
func SDL_SetClipRect(surface *SDL_Surface, rect *SDL_Rect) bool {
	cBool := C.SDL_SetClipRect(surface.cSDL_Surface(), rect.cSDL_Rect())
	if int(cBool) > 0 {
		return true
	}
	return false
}

// GetClipRect returns the clipping rectangle for a surface.
func SDL_GetClipRect(surface *SDL_Surface, rect *SDL_Rect) {
	C.SDL_GetClipRect(surface.cSDL_Surface(), rect.cSDL_Rect())
}

// Convert copies the existing surface into a new one that is optimized for blitting to a surface of a specified pixel format.
func SDL_ConvertSurface(surface *SDL_Surface, fmt *SDL_PixelFormat, flags uint32) *SDL_Surface {
	cSurface := C.SDL_ConvertSurface(surface.cSDL_Surface(), fmt.cSDL_PixelFormat(), cUint32(flags))
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

// ConvertFormat copies the existing surface to a new surface of the specified format.
func SDL_ConvertSurfaceFormat(surface *SDL_Surface, pixelFormat uint32, flags uint32) *SDL_Surface {
	cSurface := C.SDL_ConvertSurfaceFormat(surface.cSDL_Surface(), cUint32(pixelFormat), cUint32(flags))
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

// ConvertPixels copies a block of pixels of one format to another format.
func SDL_ConvertPixels(width, height int32, srcFormat uint32, src unsafe.Pointer, srcPitch int,
	dstFormat uint32, dst unsafe.Pointer, dstPitch int) int {
	cRet := C.SDL_ConvertPixels(cInt(width), cInt(height), cUint32(srcFormat), src, cInt(srcPitch), cUint32(dstFormat), dst, cInt(dstPitch))
	return int(cRet)
}

// FillRect performs a fast fill of a rectangle with a specific color.
func SDL_FillRect(surface *SDL_Surface, rect *SDL_Rect, color uint32) int {
	cRet := C.SDL_FillRect(surface.cSDL_Surface(), rect.cSDL_Rect(), cUint32(color))
	return int(cRet)
}

// FillRects performs a fast fill of a set of rectangles with a specific color.
func SDL_FillRects(surface *SDL_Surface, rects []SDL_Rect, color uint32) int {
	cRet := C.SDL_FillRects(surface.cSDL_Surface(), rects[0].cSDL_Rect(), cInt(len(rects)), cUint32(color))
	return int(cRet)
}

// Blit performs a fast surface copy to a destination surface.
func SDL_BlitSurface(surface *SDL_Surface, srcRect *SDL_Rect, dst *SDL_Surface, dstRect *SDL_Rect) int {
	cRet := C.SDL_BlitSurface(surface.cSDL_Surface(), srcRect.cSDL_Rect(), dst.cSDL_Surface(), dstRect.cSDL_Rect())
	return int(cRet)
}

// BlitScaled performs a scaled surface copy to a destination surface.
func SDL_BlitScaled(surface *SDL_Surface, srcRect *SDL_Rect, dst *SDL_Surface, dstRect *SDL_Rect) int {
	cRet := C.SDL_BlitScaled(surface.cSDL_Surface(), srcRect.cSDL_Rect(), dst.cSDL_Surface(), dstRect.cSDL_Rect())
	return int(cRet)
}

// UpperBlit has been replaced by Blit().
func SDL_UpperBlit(surface *SDL_Surface, srcRect *SDL_Rect, dst *SDL_Surface, dstRect *SDL_Rect) int {
	cRet := C.SDL_UpperBlit(surface.cSDL_Surface(), srcRect.cSDL_Rect(), dst.cSDL_Surface(), dstRect.cSDL_Rect())
	return int(cRet)
}

// LowerBlit performs low-level surface blitting only.
func SDL_LowerBlit(surface *SDL_Surface, srcRect *SDL_Rect, dst *SDL_Surface, dstRect *SDL_Rect) int {
	cRet := C.SDL_LowerBlit(surface.cSDL_Surface(), srcRect.cSDL_Rect(), dst.cSDL_Surface(), dstRect.cSDL_Rect())
	return int(cRet)
}

// SoftStretch has been replaced by BlitScaled().
func SDL_SoftStretch(surface *SDL_Surface, srcRect *SDL_Rect, dst *SDL_Surface, dstRect *SDL_Rect) int {
	cRet := C.SDL_SoftStretch(surface.cSDL_Surface(), srcRect.cSDL_Rect(), dst.cSDL_Surface(), dstRect.cSDL_Rect())
	return int(cRet)
}

// UpperBlitScaled has been replaced by BlitScaled().
func SDL_UpperBlitScaled(surface *SDL_Surface, srcRect *SDL_Rect, dst *SDL_Surface, dstRect *SDL_Rect) int {
	cRet := C.SDL_UpperBlitScaled(surface.cSDL_Surface(), srcRect.cSDL_Rect(), dst.cSDL_Surface(), dstRect.cSDL_Rect())
	return int(cRet)
}

// LowerBlitScaled performs low-level surface scaled blitting only.
func SDL_LowerBlitScaled(surface *SDL_Surface, srcRect *SDL_Rect, dst *SDL_Surface, dstRect *SDL_Rect) int {
	cRet := C.SDL_LowerBlitScaled(surface.cSDL_Surface(), srcRect.cSDL_Rect(), dst.cSDL_Surface(), dstRect.cSDL_Rect())
	return int(cRet)
}

// PixelNum returns the number of pixels stored in the surface.
func (surface *SDL_Surface) PixelNum() int {
	return int(surface.W * surface.H)
}

// BytesPerPixel return the number of significant bits in a pixel values of the surface.
func (surface *SDL_Surface) BytesPerPixel() int {
	return int(surface.Format.BytesPerPixel)
}

// Data returns the actual pixel data of the surface.
func (surface *SDL_Surface) Data() []byte {
	var b []byte
	length := int(surface.H) * int(surface.Pitch)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(surface.Pixels)
	return b
}

// Duplicate creates a new surface identical to the existing surface
func SDL_DuplicateSurface(surface *SDL_Surface) (newSurface *SDL_Surface) {
	cNewSurface := C.SDL_DuplicateSurface(surface.cSDL_Surface())
	newSurface = (*SDL_Surface)(unsafe.Pointer(cNewSurface))
	return
}

// ColorModel returns the color model used by this Surface.
func (surface *SDL_Surface) ColorModel() color.Model {
	switch surface.Format.Format {
	case SDL_PIXELFORMAT_ARGB8888:
		return ARGB8888Model
	case SDL_PIXELFORMAT_ABGR8888:
		return ABGR8888Model
	case SDL_PIXELFORMAT_RGB444:
		return RGB444Model
	case SDL_PIXELFORMAT_RGB332:
		return RGB332Model
	case SDL_PIXELFORMAT_RGB555:
		return RGB555Model
	case SDL_PIXELFORMAT_RGB565:
		return RGB565Model
	case SDL_PIXELFORMAT_BGR555:
		return BGR555Model
	case SDL_PIXELFORMAT_BGR565:
		return BGR565Model
	case SDL_PIXELFORMAT_RGB888:
		return RGB888Model
	case SDL_PIXELFORMAT_BGR888:
		return BGR888Model
	case SDL_PIXELFORMAT_ARGB4444:
		return ARGB4444Model
	case SDL_PIXELFORMAT_ABGR4444:
		return ABGR4444Model
	case SDL_PIXELFORMAT_RGBA4444:
		return RGBA4444Model
	case SDL_PIXELFORMAT_BGRA4444:
		return BGRA4444Model
	case SDL_PIXELFORMAT_ARGB1555:
		return ARGB1555Model
	case SDL_PIXELFORMAT_RGBA5551:
		return RGBA5551Model
	case SDL_PIXELFORMAT_ABGR1555:
		return ABGR1555Model
	case SDL_PIXELFORMAT_BGRA5551:
		return BGRA5551Model
	case SDL_PIXELFORMAT_RGBA8888:
		return RGBA8888Model
	case SDL_PIXELFORMAT_BGRA8888:
		return BGRA8888Model
	default:
		panic("Not implemented yet")
	}
}

// Bounds return the bounds of this surface. Currently, it always starts at
// (0,0), but this is not guaranteed in the future so don't rely on it.
func (surface *SDL_Surface) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(surface.W), int(surface.H))
}

// At returns the pixel color at (x, y)
func (surface *SDL_Surface) At(x, y int) color.Color {
	pix := surface.Data()
	i := int32(y)*surface.Pitch + int32(x)*int32(surface.Format.BytesPerPixel)
	var r, g, b, a uint8
	SDL_GetRGBA(*((*uint32)(unsafe.Pointer(&pix[i]))), surface.Format, &r, &g, &b, &a)
	return color.RGBA{r, g, b, a}
}

// Set the color of the pixel at (x, y) using this surface's color model to
// convert c to the appropriate color. This method is required for the
// draw.Image interface. The surface may require locking before calling Set.
func (surface *SDL_Surface) Set(x, y int, c color.Color) {
	pix := surface.Data()
	i := int32(y)*surface.Pitch + int32(x)*int32(surface.Format.BytesPerPixel)
	switch surface.Format.Format {
	case SDL_PIXELFORMAT_ARGB8888:
		col := surface.ColorModel().Convert(c).(ARGB8888)
		pix[i+3] = col.A
		pix[i+2] = col.R
		pix[i+1] = col.G
		pix[i+0] = col.B
	case SDL_PIXELFORMAT_ABGR8888:
		col := surface.ColorModel().Convert(c).(ABGR8888)
		pix[i+3] = col.A
		pix[i+2] = col.B
		pix[i+1] = col.G
		pix[i+0] = col.R
	case SDL_PIXELFORMAT_RGB24, SDL_PIXELFORMAT_RGB888:
		col := surface.ColorModel().Convert(c).(RGB888)
		pix[i+2] = col.R
		pix[i+1] = col.G
		pix[i+0] = col.B
	case SDL_PIXELFORMAT_BGR24, SDL_PIXELFORMAT_BGR888:
		col := surface.ColorModel().Convert(c).(BGR888)
		pix[i+2] = col.B
		pix[i+1] = col.G
		pix[i+0] = col.R
	case SDL_PIXELFORMAT_RGB444:
		col := surface.ColorModel().Convert(c).(RGB444)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 4 & 0x0F
		g := uint32(col.G) >> 4 & 0x0F
		b := uint32(col.B) >> 4 & 0x0F
		*buf = r<<8 | g<<4 | b
	case SDL_PIXELFORMAT_RGB332:
		col := surface.ColorModel().Convert(c).(RGB332)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 5 & 0x0F
		g := uint32(col.G) >> 5 & 0x0F
		b := uint32(col.B) >> 6 & 0x0F
		*buf = r<<5 | g<<2 | b
	case SDL_PIXELFORMAT_RGB565:
		col := surface.ColorModel().Convert(c).(RGB565)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 3 & 0xFF
		g := uint32(col.G) >> 2 & 0xFF
		b := uint32(col.B) >> 3 & 0xFF
		*buf = r<<11 | g<<5 | b
	case SDL_PIXELFORMAT_RGB555:
		col := surface.ColorModel().Convert(c).(RGB555)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 3 & 0xFF
		g := uint32(col.G) >> 3 & 0xFF
		b := uint32(col.B) >> 3 & 0xFF
		*buf = r<<10 | g<<5 | b
	case SDL_PIXELFORMAT_BGR565:
		col := surface.ColorModel().Convert(c).(BGR565)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 3 & 0xFF
		g := uint32(col.G) >> 2 & 0xFF
		b := uint32(col.B) >> 3 & 0xFF
		*buf = b<<11 | g<<5 | r
	case SDL_PIXELFORMAT_BGR555:
		col := surface.ColorModel().Convert(c).(BGR555)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 3 & 0xFF
		g := uint32(col.G) >> 3 & 0xFF
		b := uint32(col.B) >> 3 & 0xFF
		*buf = b<<10 | g<<5 | r
	case SDL_PIXELFORMAT_ARGB4444:
		col := surface.ColorModel().Convert(c).(ARGB4444)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		a := uint32(col.A) >> 4 & 0x0F
		r := uint32(col.R) >> 4 & 0x0F
		g := uint32(col.G) >> 4 & 0x0F
		b := uint32(col.B) >> 4 & 0x0F
		*buf = a<<12 | r<<8 | g<<4 | b
	case SDL_PIXELFORMAT_ABGR4444:
		col := surface.ColorModel().Convert(c).(ABGR4444)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		a := uint32(col.A) >> 4 & 0x0F
		r := uint32(col.R) >> 4 & 0x0F
		g := uint32(col.G) >> 4 & 0x0F
		b := uint32(col.B) >> 4 & 0x0F
		*buf = a<<12 | b<<8 | g<<4 | r
	case SDL_PIXELFORMAT_RGBA4444:
		col := surface.ColorModel().Convert(c).(RGBA4444)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 4 & 0x0F
		g := uint32(col.G) >> 4 & 0x0F
		b := uint32(col.B) >> 4 & 0x0F
		a := uint32(col.A) >> 4 & 0x0F
		*buf = r<<12 | g<<8 | b<<4 | a
	case SDL_PIXELFORMAT_BGRA4444:
		col := surface.ColorModel().Convert(c).(BGRA4444)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 4 & 0x0F
		g := uint32(col.G) >> 4 & 0x0F
		b := uint32(col.B) >> 4 & 0x0F
		a := uint32(col.A) >> 4 & 0x0F
		*buf = b<<12 | g<<8 | r<<4 | a
	case SDL_PIXELFORMAT_ARGB1555:
		col := surface.ColorModel().Convert(c).(ARGB1555)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 3 & 0xFF
		g := uint32(col.G) >> 3 & 0xFF
		b := uint32(col.B) >> 3 & 0xFF
		a := uint32(0)
		if col.A > 0 {
			a = 1
		}
		*buf = a<<15 | r<<10 | g<<5 | b
	case SDL_PIXELFORMAT_RGBA5551:
		col := surface.ColorModel().Convert(c).(RGBA5551)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 3 & 0xFF
		g := uint32(col.G) >> 3 & 0xFF
		b := uint32(col.B) >> 3 & 0xFF
		a := uint32(0)
		if col.A > 0 {
			a = 1
		}
		*buf = r<<11 | g<<6 | b<<1 | a
	case SDL_PIXELFORMAT_ABGR1555:
		col := surface.ColorModel().Convert(c).(ABGR1555)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 3 & 0xFF
		g := uint32(col.G) >> 3 & 0xFF
		b := uint32(col.B) >> 3 & 0xFF
		a := uint32(0)
		if col.A > 0 {
			a = 1
		}
		*buf = a<<15 | b<<10 | g<<5 | r
	case SDL_PIXELFORMAT_BGRA5551:
		col := surface.ColorModel().Convert(c).(BGRA5551)
		buf := (*uint32)(unsafe.Pointer(&pix[i]))
		r := uint32(col.R) >> 3 & 0xFF
		g := uint32(col.G) >> 3 & 0xFF
		b := uint32(col.B) >> 3 & 0xFF
		a := uint32(0)
		if col.A > 0 {
			a = 1
		}
		*buf = b<<11 | g<<6 | r<<1 | a
	case SDL_PIXELFORMAT_RGBA8888:
		col := surface.ColorModel().Convert(c).(RGBA8888)
		pix[i+3] = col.R
		pix[i+2] = col.G
		pix[i+1] = col.B
		pix[i+0] = col.A
	case SDL_PIXELFORMAT_BGRA8888:
		col := surface.ColorModel().Convert(c).(BGRA8888)
		pix[i+3] = col.B
		pix[i+2] = col.G
		pix[i+1] = col.R
		pix[i+0] = col.A
	default:
		panic("Unknown pixel format!")
	}
}

// SoftStretchLinear performs bilinear scaling between two surfaces of the same format, 32BPP.
func SDL_SoftStretchLinear(surface *SDL_Surface, srcRect *SDL_Rect, dst *SDL_Surface, dstRect *SDL_Rect) int {
	cRet := C.SDL_SoftStretchLinear(surface.cSDL_Surface(), srcRect.cSDL_Rect(), dst.cSDL_Surface(), dstRect.cSDL_Rect())
	return int(cRet)
}

// PremultiplyAlpha premultiplies the alpha on a block of pixels.
// This is safe to use with src == dst, but not for other overlapping areas.
// This function is currently only implemented for SDL_PIXELFORMAT_ARGB8888.
func SDL_PremultiplyAlpha(width, height int, srcFormat uint32, src []byte, srcPitch int, dstFormat uint32, dst []byte, dstPitch int) int {
	cWidth := cInt(width)
	cHeight := cInt(height)
	cSrcFormat := cUint32(srcFormat)
	cSrc := unsafe.Pointer(&src[0])
	cSrcPitch := cInt(srcPitch)
	cDstFormat := cUint32(dstFormat)
	cDst := unsafe.Pointer(&dst[0])
	cDstPitch := cInt(dstPitch)
	cRet := C.SDL_PremultiplyAlpha(cWidth, cHeight, cSrcFormat, cSrc, cSrcPitch, cDstFormat, cDst, cDstPitch)
	return int(cRet)
}
