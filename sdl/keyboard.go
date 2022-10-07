package sdl

/*
#include "keyboard.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// struct
type SDL_Keysym struct {
	Scancode int32
	Sym      int32
	Mod      uint16
	Unused   uint32
}

func SDL_GetKeyboardFocus() *SDL_Window {
	cWnd := C.SDL_GetKeyboardFocus()
	return (*SDL_Window)(unsafe.Pointer(cWnd))
}

func SDL_GetKeyboardState(numkeys *int) []uint8 {
	// var numkeys cInt
	cNumkeys := (*cInt)(unsafe.Pointer(numkeys))
	start := C.SDL_GetKeyboardState(cNumkeys)

	var slice []uint8
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sh.Len = int(*numkeys)
	sh.Cap = int(*numkeys)
	sh.Data = uintptr(unsafe.Pointer(start))
	return *(*[]uint8)(unsafe.Pointer(&sh))
}

func SDL_GetModState() SDL_Keymod {
	cKeymod := C.SDL_GetModState()
	return SDL_Keymod(cKeymod)
}

func SDL_SetModState(mod SDL_Keymod) {
	C.SDL_SetModState(cKeymod(mod))
}

func SDL_GetKeyFromScancode(code SDL_Scancode) SDL_Keycode {
	cKeycode := C.SDL_GetKeyFromScancode(cScancode(code))
	return SDL_Keycode(cKeycode)
}

func SDL_GetScancodeFromKey(code SDL_Keycode) SDL_Scancode {
	cKeycode := C.SDL_GetScancodeFromKey(cKeycode(code))
	return SDL_Scancode(cKeycode)
}

func SDL_GetScancodeName(code SDL_Scancode) string {
	cName := C.SDL_GetScancodeName(cScancode(code))
	return createGoString(cName)
}

func SDL_GetScancodeFromName(name string) SDL_Scancode {
	cName := createCString(SDL_GetMemoryPool(), name)
	defer destroyCString(SDL_GetMemoryPool(), cName)

	cScancode := C.SDL_GetScancodeFromName(cName)
	return SDL_Scancode(cScancode)
}

func SDL_GetKeyName(code SDL_Keycode) string {
	cKeyname := C.SDL_GetKeyName(cKeycode(code))
	return createGoString(cKeyname)
}

func SDL_GetKeyFromName(name string) SDL_Keycode {
	cName := createCString(SDL_GetMemoryPool(), name)
	defer destroyCString(SDL_GetMemoryPool(), cName)

	cKeycode := C.SDL_GetKeyFromName(cName)
	return SDL_Keycode(cKeycode)
}

func SDL_StartTextInput() {
	C.SDL_StartTextInput()
}

func SDL_IsTextInputActive() bool {
	cRet := C.SDL_IsTextInputActive()
	return int(cRet) > 0
}

func SDL_StopTextInput() {
	C.SDL_StopTextInput()
}

func SDL_SetTextInputRect(rect *SDL_Rect) {
	C.SDL_SetTextInputRect(cRect(rect))
}

func SDL_HasScreenKeyboardSupport() bool {
	cRet := C.SDL_HasScreenKeyboardSupport()
	return int(cRet) > 0
}

func SDL_IsScreenKeyboardShown(window *SDL_Window) bool {
	cRet := C.SDL_IsScreenKeyboardShown(cWindow(window))
	return int(cRet) > 0
}

func SDL_IsTextInputShown() bool {
	cRet := C.SDL_IsTextInputShown()
	return int(cRet) > 0
}

func SDL_ClearComposition() {
	C.SDL_ClearComposition()
}
