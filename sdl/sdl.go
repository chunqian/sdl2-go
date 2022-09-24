package sdl

/*
#include "sdl_go.h"
*/
import "C"
import "runtime"

func init() {
	// Make sure the main goroutine is bound to the main thread.
	runtime.LockOSThread()
}

// Init initialize the SDL library. This must be called before using most other SDL functions.
func SDL_Init(flags uint32) int {
	cRet := C.SDL_Init(cUint32(flags))
	return int(cRet)
}

// Quit cleans up all initialized subsystems. You should call it upon all exit conditions.
func SDL_Quit() {
	C.SDL_Quit()
}

// InitSubSystem initializes specific SDL subsystems.
func SDL_InitSubSystem(flags uint32) int {
	cRet := C.SDL_InitSubSystem(cUint32(flags))
	return int(cRet)
}

// QuitSubSystem shuts down specific SDL subsystems.
func SDL_QuitSubSystem(flags uint32) {
	C.SDL_QuitSubSystem(cUint32(flags))
}

// WasInit returns a mask of the specified subsystems which have previously been initialized.
func SDL_WasInit(flags uint32) uint32 {
	return uint32(C.SDL_WasInit(cUint32(flags)))
}

// GetPlatform returns the name of the platform.
func SDL_GetPlatform() string {
	return SDL_GoString(C.SDL_GetPlatform())
}
