package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"

// GetTicks returns the number of milliseconds since the SDL library initialization.
func SDL_GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

// GetPerformanceCounter returns the current value of the high resolution counter.
func SDL_GetPerformanceCounter() uint64 {
	return uint64(C.SDL_GetPerformanceCounter())
}

// GetPerformanceFrequency returns the count per second of the high resolution counter.
func SDL_GetPerformanceFrequency() uint64 {
	return uint64(C.SDL_GetPerformanceFrequency())
}

// Delay waits a specified number of milliseconds before returning.
func SDL_Delay(ms uint32) {
	C.SDL_Delay(cUint32(ms))
}
