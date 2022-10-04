package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

// typedef
type SDL_TimerCallback = func(uint32, unsafe.Pointer) uint32
type SDL_TimerID = int32

func SDL_GetTicks() uint32 {
	return uint32(C.SDL_GetTicks())
}

func SDL_GetPerformanceCounter() uint64 {
	return uint64(C.SDL_GetPerformanceCounter())
}

func SDL_GetPerformanceFrequency() uint64 {
	return uint64(C.SDL_GetPerformanceFrequency())
}

func SDL_Delay(ms uint32) {
	C.SDL_Delay(cUint32(ms))
}
