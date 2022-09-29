package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

func SDL_SetClipboardText(text string) int {
	cText := SDL_CreateCString(SDL_GetMemoryPool(), text)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cText) // memory free

	cRet := C.SDL_SetClipboardText(cText)
	return int(cRet)
}

func SDL_GetClipboardText() string {
	cText := C.SDL_GetClipboardText()
	if cText == nil {
		return ""
	}
	defer C.SDL_free(unsafe.Pointer(cText))

	text := SDL_GoString(cText)
	return text
}

func SDL_HasClipboardText() bool {
	cRet := C.SDL_HasClipboardText()
	if int(cRet) > 0 {
		return true
	}
	return false
}
