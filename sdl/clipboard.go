package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

// SetClipboardText puts UTF-8 text into the clipboard.
func SDL_SetClipboardText(text string) int {
	cText := SDL_CreateCString(SDL_GetMemoryPool(), text)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cText) // memory free

	cRet := C.SDL_SetClipboardText(cText)
	return int(cRet)
}

// GetClipboardText returns UTF-8 text from the clipboard.
func SDL_GetClipboardText() string {
	cText := C.SDL_GetClipboardText()
	if cText == nil {
		return ""
	}
	defer C.SDL_free(unsafe.Pointer(cText))

	text := SDL_GoString(cText)
	return text
}

// HasClipboardText reports whether the clipboard exists and contains a text string that is non-empty.
func SDL_HasClipboardText() bool {
	cRet := C.SDL_HasClipboardText()
	if int(cRet) > 0 {
		return true
	}
	return false
}
