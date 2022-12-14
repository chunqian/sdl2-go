package img

/*
#include "image_wrapper.h"
*/
import "C"
import (
	"unsafe"

	. "github.com/chunqian/sdl2-go/sdl"
)

// define
const (
	SDL_IMAGE_MAJOR_VERSION = 2
	SDL_IMAGE_MINOR_VERSION = 0
	SDL_IMAGE_PATCHLEVEL    = 5
)

// typedef
type IMG_InitFlags = int32

// enum
const (
	IMG_INIT_JPG  IMG_InitFlags = 0x00000001
	IMG_INIT_PNG  IMG_InitFlags = 0x00000002
	IMG_INIT_TIF  IMG_InitFlags = 0x00000004
	IMG_INIT_WEBP IMG_InitFlags = 0x00000008
)

func cRWops(rw *SDL_RWops) *C.SDL_RWops {
	return (*C.SDL_RWops)(unsafe.Pointer(rw))
}

func cSurface(surface *SDL_Surface) *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(surface))
}

func cRenderer(r *SDL_Renderer) *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
}

func IMG_Linked_Version() *SDL_version {
	return (*SDL_version)(unsafe.Pointer(C.IMG_Linked_Version()))
}

func IMG_Init(flags IMG_InitFlags) int {
	cFlags := cInt(flags)
	cRet := C.IMG_Init(cFlags)
	return int(cRet)
}

func IMG_Quit() {
	C.IMG_Quit()
}

func IMG_LoadTyped_RW(src *SDL_RWops, freesrc SDL_bool, type_ string) *SDL_Surface {
	cSrc := cRWops(src)
	cFreesrc := cInt(freesrc)

	cType := createCString(SDL_GetMemoryPool(), type_)
	defer destroyCString(SDL_GetMemoryPool(), cType) // memory free

	cSurface := C.IMG_LoadTyped_RW(cSrc, cFreesrc, cType)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_Load(file string) *SDL_Surface {
	cFile := createCString(SDL_GetMemoryPool(), file)
	defer destroyCString(SDL_GetMemoryPool(), cFile) // memory free

	cSurface := C.IMG_Load(cFile)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_Load_RW(src *SDL_RWops, freesrc SDL_bool) *SDL_Surface {
	cSrc := cRWops(src)
	cFreesrc := cInt(freesrc)
	cSurface := C.IMG_Load_RW(cSrc, cFreesrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadTexture(renderer *SDL_Renderer, file string) *SDL_Texture {
	cRenderer := cRenderer(renderer)

	cFile := createCString(SDL_GetMemoryPool(), file)
	defer destroyCString(SDL_GetMemoryPool(), cFile) // memory free

	cSurface := C.IMG_LoadTexture(cRenderer, cFile)
	return (*SDL_Texture)(unsafe.Pointer(cSurface))
}

func IMG_LoadTexture_RW(renderer *SDL_Renderer, src *SDL_RWops, freesrc SDL_bool) *SDL_Texture {
	cRenderer := cRenderer(renderer)
	cSrc := cRWops(src)
	cFreesrc := cInt(freesrc)
	cSurface := C.IMG_LoadTexture_RW(cRenderer, cSrc, cFreesrc)
	return (*SDL_Texture)(unsafe.Pointer(cSurface))
}

func IMG_isICO(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isICO(cSrc)) > 0
}

func IMG_isCUR(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isCUR(cSrc)) > 0
}

func IMG_isBMP(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isBMP(cSrc)) > 0
}

func IMG_isGIF(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isGIF(cSrc)) > 0
}

func IMG_isJPG(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isJPG(cSrc)) > 0
}

func IMG_isLBM(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isLBM(cSrc)) > 0
}

func IMG_isPCX(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isPCX(cSrc)) > 0
}

func IMG_isPNG(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isPNG(cSrc)) > 0
}

func IMG_isPNM(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isPNM(cSrc)) > 0
}

func IMG_isTIF(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isTIF(cSrc)) > 0
}

func IMG_isXCF(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isXCF(cSrc)) > 0
}

func IMG_isXPM(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isXPM(cSrc)) > 0
}

func IMG_isXV(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isXV(cSrc)) > 0
}

func IMG_isWEBP(src *SDL_RWops) bool {
	cSrc := cRWops(src)
	return int(C.IMG_isWEBP(cSrc)) > 0
}

func IMG_LoadICO_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadICO_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadCUR_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadCUR_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadBMP_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadBMP_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadGIF_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadGIF_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadJPG_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadJPG_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadLBM_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadLBM_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadPCX_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadPCX_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadPNG_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadPNG_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadPNM_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadPNM_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadTGA_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadTGA_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadTIF_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadTIF_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadXCF_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadXCF_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadXPM_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadXPM_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadXV_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadXV_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_LoadWEBP_RW(src *SDL_RWops) *SDL_Surface {
	cSrc := cRWops(src)
	cSurface := C.IMG_LoadWEBP_RW(cSrc)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_ReadXPMFromArray(xpm string) *SDL_Surface {
	cXpm := createCString(SDL_GetMemoryPool(), xpm)
	defer destroyCString(SDL_GetMemoryPool(), cXpm) // memory free

	cSurface := C.IMG_ReadXPMFromArray(&cXpm)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func IMG_SavePNG(surface *SDL_Surface, file string) int {
	cSurface := cSurface(surface)

	cFile := createCString(SDL_GetMemoryPool(), file)
	defer destroyCString(SDL_GetMemoryPool(), cFile) // memory free

	cRet := C.IMG_SavePNG(cSurface, cFile)
	return int(cRet)
}

func IMG_SavePNG_RW(surface *SDL_Surface, dst *SDL_RWops, freedst int) int {
	cSurface := cSurface(surface)
	cDst := cRWops(dst)
	cFreedst := cInt(freedst)
	cRet := C.IMG_SavePNG_RW(cSurface, cDst, cFreedst)
	return int(cRet)
}
