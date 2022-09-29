package ttf

/*
#include "ttf_wrapper.h"
*/
import "C"
import (
	"unsafe"

	. "github.com/chunqian/sdl2-go/sdl"
)

func cSDL_Color(c *SDL_Color) *C.SDL_Color {
	return (*C.SDL_Color)(unsafe.Pointer(c))
}

func cTTF_Font(font *TTF_Font) *C.TTF_Font {
	return (*C.TTF_Font)(unsafe.Pointer(font))
}

func TTF_Init() int {
	cRet := C.TTF_Init()
	return int(cRet)
}

func TTF_WasInit() bool {
	return int(C.TTF_WasInit()) != 0
}

func TTF_Quit() {
	C.TTF_Quit()
}

func TTF_ByteSwappedUNICODE(swap SDL_bool) {
	C.TTF_ByteSwappedUNICODE(cSDL_bool(swap))
}

func TTF_OpenFont(file string, size int) *TTF_Font {
	cFile := SDL_CreateCString(SDL_GetMemoryPool(), file)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cFile) // memory free

	cSize := cInt(size)
	cFont := C.TTF_OpenFont(cFile.(*cChar), cSize)
	return (*TTF_Font)(unsafe.Pointer(cFont))
}

func TTF_OpenFontIndex(file string, size int, index int) *TTF_Font {
	cFile := SDL_CreateCString(SDL_GetMemoryPool(), file)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cFile) // memory free

	cSize := cInt(size)
	cIndex := cLong(index)
	cFont := C.TTF_OpenFontIndex(cFile.(*cChar), cSize, cIndex)
	return (*TTF_Font)(unsafe.Pointer(cFont))
}

func OpenFontRW(src *SDL_RWops, freesrc, size int) *TTF_Font {
	return TTF_OpenFontIndexRW(src, freesrc, size, 0)
}

func TTF_OpenFontIndexRW(src *SDL_RWops, freesrc, size, index int) *TTF_Font {
	cSrc := (*C.SDL_RWops)(unsafe.Pointer(src))
	cFreesrc := cInt(freesrc)
	cSize := cInt(size)
	cIndex := cLong(index)
	cFont := C.TTF_OpenFontIndexRW(cSrc, cFreesrc, cSize, cIndex)
	return (*TTF_Font)(unsafe.Pointer(cFont))
}

func TTF_RenderUTF8_Solid(font *TTF_Font, text string, color SDL_Color) *SDL_Surface {
	cText := SDL_CreateCString(SDL_GetMemoryPool(), text)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cText) // memory free

	cColor := cSDL_Color(&color)
	cSurface := C.TTF_RenderUTF8_Solid(cTTF_Font(font), cText.(*cChar), *cColor)
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

func TTF_RenderUTF8_Shaded(font *TTF_Font, text string, fg, bg SDL_Color) *SDL_Surface {
	cText := SDL_CreateCString(SDL_GetMemoryPool(), text)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cText) // memory free

	cFg := cSDL_Color(&fg)
	cBg := cSDL_Color(&bg)
	cSurface := C.TTF_RenderUTF8_Shaded(cTTF_Font(font), cText.(*cChar), *cFg, *cBg)
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

func TTF_RenderUTF8_Blended(font *TTF_Font, text string, color SDL_Color) *SDL_Surface {
	cText := SDL_CreateCString(SDL_GetMemoryPool(), text)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cText) // memory free

	cColor := cSDL_Color(&color)
	cSurface := C.TTF_RenderUTF8_Blended(cTTF_Font(font), cText.(*cChar), *cColor)
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

func TTF_RenderUTF8_Blended_Wrapped(font *TTF_Font, text string, fg SDL_Color, wrapLength int) *SDL_Surface {
	cText := SDL_CreateCString(SDL_GetMemoryPool(), text)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cText) // memory free

	cFg := cSDL_Color(&fg)
	cSurface := C.TTF_RenderUTF8_Blended_Wrapped(cTTF_Font(font), cText.(*cChar), *cFg, cUint32(wrapLength))
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

func TTF_SizeUTF8(font *TTF_Font, text string) (int, int) {
	cText := SDL_CreateCString(SDL_GetMemoryPool(), text)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cText) // memory free

	var w cInt
	var h cInt
	result := C.TTF_SizeUTF8(cTTF_Font(font), cText.(*cChar), &w, &h)
	if result == 0 {
		return int(w), int(h)
	}
	return int(w), int(h)
}

func TTF_RenderGlyph_Solid(font *TTF_Font, ch rune, fg SDL_Color) *SDL_Surface {
	cCh := cUint16(ch)
	cFg := cSDL_Color(&fg)
	cSurface := C.TTF_RenderGlyph_Solid(cTTF_Font(font), cCh, *cFg)
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

func TTF_RenderGlyph_Shaded(font *TTF_Font, ch rune, fg, bg SDL_Color) *SDL_Surface {
	cCh := cUint16(ch)
	cFg := cSDL_Color(&fg)
	cBg := cSDL_Color(&bg)
	cSurface := C.TTF_RenderGlyph_Shaded(cTTF_Font(font), cCh, *cFg, *cBg)
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

func TTF_RenderGlyph_Blended(font *TTF_Font, ch rune, fg SDL_Color) *SDL_Surface {
	cCh := cUint16(ch)
	cFg := cSDL_Color(&fg)
	cSurface := C.TTF_RenderGlyph_Blended(cTTF_Font(font), cCh, *cFg)
	surface := (*SDL_Surface)(unsafe.Pointer(cSurface))
	return surface
}

func TTF_CloseFont(font *TTF_Font) {
	C.TTF_CloseFont(cTTF_Font(font))
	font = nil
}

func TTF_FontHeight(font *TTF_Font) int {
	cRet := C.TTF_FontHeight(cTTF_Font(font))
	return int(cRet)
}

func TTF_FontAscent(font *TTF_Font) int {
	cRet := C.TTF_FontAscent(cTTF_Font(font))
	return int(cRet)
}

func TTF_FontDescent(font *TTF_Font) int {
	cRet := C.TTF_FontDescent(cTTF_Font(font))
	return int(cRet)
}

func TTF_FontLineSkip(font *TTF_Font) int {
	cRet := C.TTF_FontLineSkip(cTTF_Font(font))
	return int(cRet)
}

func TTF_FontFaces(font *TTF_Font) int {
	cRet := C.TTF_FontFaces(cTTF_Font(font))
	return int(cRet)
}

func TTF_GetFontStyle(font *TTF_Font) int {
	cRet := C.TTF_GetFontStyle(cTTF_Font(font))
	return int(cRet)
}

func TTF_SetFontStyle(font *TTF_Font, style int32) {
	C.TTF_SetFontStyle(cTTF_Font(font), cInt(style))
}

func TTF_GetFontHinting(font *TTF_Font) int {
	cRet := C.TTF_GetFontHinting(cTTF_Font(font))
	return int(cRet)
}

func TTF_SetFontHinting(font *TTF_Font, hinting int32) {
	C.TTF_SetFontHinting(cTTF_Font(font), cInt(hinting))
}

func TTF_GetFontKerning(font *TTF_Font) bool {
	cRet := C.TTF_GetFontKerning(cTTF_Font(font))
	return int(cRet) == 1
}

func TTF_SetFontKerning(font *TTF_Font, allowed SDL_bool) {
	C.TTF_SetFontKerning(cTTF_Font(font), cInt(allowed))
}

func TTF_GetFontOutline(font *TTF_Font) int {
	cRet := C.TTF_GetFontOutline(cTTF_Font(font))
	return int(cRet)
}

func TTF_SetFontOutline(font *TTF_Font, outline int) {
	C.TTF_SetFontOutline(cTTF_Font(font), cInt(outline))
}

func TTF_FontFaceIsFixedWidth(font *TTF_Font) bool {
	cRet := C.TTF_FontFaceIsFixedWidth(cTTF_Font(font))
	return int(cRet) != 0
}

func TTF_FontFaceFamilyName(font *TTF_Font) string {
	cFname := C.TTF_FontFaceFamilyName(cTTF_Font(font))
	fname := SDL_GoString(cFname)
	return fname
}

type GlyphMetrics struct {
	MinX, MaxX int
	MinY, MaxY int
	Advance    int
}

func TTF_GlyphMetrics(font *TTF_Font, ch rune, gm *GlyphMetrics) int {
	cCh := cUint16(ch)
	cMinX := (*cInt)(unsafe.Pointer(&gm.MinX))
	cMaxX := (*cInt)(unsafe.Pointer(&gm.MaxX))
	cMinY := (*cInt)(unsafe.Pointer(&gm.MinY))
	cMaxY := (*cInt)(unsafe.Pointer(&gm.MaxY))
	cAdvance := (*cInt)(unsafe.Pointer(&gm.Advance))
	cRet := C.TTF_GlyphMetrics(cTTF_Font(font), cCh, cMinX, cMaxX, cMinY, cMaxY, cAdvance)
	return int(cRet)
}
