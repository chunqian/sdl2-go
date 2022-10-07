package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

func SDL_SetClipboardText(text string) int {
	cText := createCString(SDL_GetMemoryPool(), text)
	defer destroyCString(SDL_GetMemoryPool(), cText) // memory free

	cRet := C.SDL_SetClipboardText(cText)
	return int(cRet)
}

func SDL_GetClipboardText() string {
	cText := C.SDL_GetClipboardText()
	if cText == nil {
		return ""
	}
	defer C.free(unsafe.Pointer(cText))

	text := createGoString(cText)
	return text
}

func SDL_HasClipboardText() bool {
	cRet := C.SDL_HasClipboardText()
	if int(cRet) > 0 {
		return true
	}
	return false
}
