package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "runtime"

func init() {
	// Make sure the main goroutine is bound to the main thread.
	runtime.LockOSThread()
}

func SDL_Init(flags uint32) int {
	cRet := C.SDL_Init(cUint32(flags))
	return int(cRet)
}

func SDL_Quit() {
	C.SDL_Quit()
}

func SDL_InitSubSystem(flags uint32) int {
	cRet := C.SDL_InitSubSystem(cUint32(flags))
	return int(cRet)
}

func SDL_QuitSubSystem(flags uint32) {
	C.SDL_QuitSubSystem(cUint32(flags))
}

func SDL_WasInit(flags uint32) uint32 {
	return uint32(C.SDL_WasInit(cUint32(flags)))
}

func SDL_GetPlatform() string {
	return SDL_GoString(C.SDL_GetPlatform())
}
