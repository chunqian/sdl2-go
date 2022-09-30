package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

func SDL_GetPowerInfo(secs, pct *int) SDL_PowerState {
	cSecs := (*cInt)(unsafe.Pointer(secs))
	cPct := (*cInt)(unsafe.Pointer(pct))
	cState := C.SDL_GetPowerInfo(cSecs, cPct)
	return SDL_PowerState(cState)
}
