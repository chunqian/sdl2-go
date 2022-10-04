package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

// typedef
type SDL_PowerState = int32

// enum
const (
	SDL_POWERSTATE_UNKNOWN    SDL_PowerState = 0
	SDL_POWERSTATE_ON_BATTERY SDL_PowerState = 1
	SDL_POWERSTATE_NO_BATTERY SDL_PowerState = 2
	SDL_POWERSTATE_CHARGING   SDL_PowerState = 3
	SDL_POWERSTATE_CHARGED    SDL_PowerState = 4
)

func SDL_GetPowerInfo(secs, pct *int) SDL_PowerState {
	cSecs := (*cInt)(unsafe.Pointer(secs))
	cPct := (*cInt)(unsafe.Pointer(pct))
	cState := C.SDL_GetPowerInfo(cSecs, cPct)
	return SDL_PowerState(cState)
}
