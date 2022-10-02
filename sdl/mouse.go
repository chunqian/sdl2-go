package sdl

/*
#include "mouse.h"
*/
import "C"
import "unsafe"

func cCursor(c *SDL_Cursor) *C.SDL_Cursor {
	return (*C.SDL_Cursor)(unsafe.Pointer(c))
}

func SDL_GetMouseFocus() *SDL_Window {
	return (*SDL_Window)(unsafe.Pointer(C.SDL_GetMouseFocus()))
}

func SDL_GetGlobalMouseState(x, y *int32) uint32 {
	cX := (*cInt)(unsafe.Pointer(x))
	cY := (*cInt)(unsafe.Pointer(y))

	cRet := C.SDL_GetGlobalMouseState(cX, cY)
	return uint32(cRet)
}

func SDL_GetMouseState(x, y *int32) uint32 {
	cX := (*cInt)(unsafe.Pointer(x))
	cY := (*cInt)(unsafe.Pointer(y))

	cRet := C.SDL_GetMouseState(cX, cY)
	return uint32(cRet)
}

func SDL_GetRelativeMouseState(x, y *int32) uint32 {
	cX := (*cInt)(unsafe.Pointer(x))
	cY := (*cInt)(unsafe.Pointer(y))

	cRet := C.SDL_GetRelativeMouseState(cX, cY)
	return uint32(cRet)
}

func SDL_WarpMouseInWindow(window *SDL_Window, x, y int32) {
	C.SDL_WarpMouseInWindow(cWindow(window), cInt(x), cInt(y))
}

func SDL_SetRelativeMouseMode(enabled SDL_bool) int {
	cRet := C.SDL_SetRelativeMouseMode(cBool(enabled))
	return int(cRet)
}

func SDL_GetRelativeMouseMode() bool {
	return C.SDL_GetRelativeMouseMode() > 0
}

func SDL_CreateCursor(data, mask *uint8, w, h, hotX, hotY int32) *SDL_Cursor {
	cData := (*C.Uint8)(unsafe.Pointer(data))
	cMask := (*C.Uint8)(unsafe.Pointer(mask))

	cCursor := C.SDL_CreateCursor(cData, cMask, cInt(w), cInt(h), cInt(hotX), cInt(hotY))
	return (*SDL_Cursor)(unsafe.Pointer(cCursor))
}

func SDL_CreateColorCursor(surface *SDL_Surface, hotX, hotY int32) *SDL_Cursor {
	cCursor := C.SDL_CreateColorCursor(cSurface(surface), cInt(hotX), cInt(hotY))
	return (*SDL_Cursor)(unsafe.Pointer(cCursor))
}

func SDL_CreateSystemCursor(id SDL_SystemCursor) *SDL_Cursor {
	cId := cSystemCursor(id)

	cCursor := C.SDL_CreateSystemCursor(cId)
	return (*SDL_Cursor)(unsafe.Pointer(cCursor))
}

func SDL_SetCursor(cursor *SDL_Cursor) {
	C.SDL_SetCursor(cCursor(cursor))
}

func SDL_GetCursor() *SDL_Cursor {
	cCursor := C.SDL_GetCursor()
	return (*SDL_Cursor)(unsafe.Pointer(cCursor))
}

func SDL_GetDefaultCursor() *SDL_Cursor {
	cCursor := C.SDL_GetDefaultCursor()
	return (*SDL_Cursor)(unsafe.Pointer(cCursor))
}

func SDL_FreeCursor(cursor *SDL_Cursor) {
	C.SDL_FreeCursor(cCursor(cursor))
}

func SDL_ShowCursor(toggle int) int {
	cRet := C.SDL_ShowCursor(cInt(toggle))
	return int(cRet)
}

func SDL_CaptureMouse(toggle SDL_bool) int {
	cRet := C.SDL_CaptureMouse(cBool(toggle))
	return int(cRet)
}

func SDL_WarpMouseGlobal(x, y int32) int {
	cRet := C.SDL_WarpMouseGlobal(cInt(x), cInt(y))
	return int(cRet)
}
