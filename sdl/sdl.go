package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "runtime"

// #define
const (
	SDL_INIT_TIMER          uint32 = 0x00000001
	SDL_INIT_AUDIO          uint32 = 0x00000010
	SDL_INIT_VIDEO          uint32 = 0x00000020
	SDL_INIT_JOYSTICK       uint32 = 0x00000200
	SDL_INIT_HAPTIC         uint32 = 0x00001000
	SDL_INIT_GAMECONTROLLER uint32 = 0x00002000
	SDL_INIT_EVENTS         uint32 = 0x00004000
	SDL_INIT_NOPARACHUTE    uint32 = 0x00008000
	SDL_INIT_SENSOR         uint32 = 0x00100000
	SDL_INIT_EVERYTHING     uint32 = SDL_INIT_TIMER | SDL_INIT_AUDIO | SDL_INIT_VIDEO | SDL_INIT_EVENTS | SDL_INIT_JOYSTICK | SDL_INIT_HAPTIC | SDL_INIT_GAMECONTROLLER | SDL_INIT_SENSOR
)

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
