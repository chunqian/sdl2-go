package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

func SDL_GetNumTouchDevices() int {
	return int(C.SDL_GetNumTouchDevices())
}

func SDL_GetTouchDevice(index int) SDL_TouchID {
	return SDL_TouchID(C.SDL_GetTouchDevice(cInt(index)))
}

func SDL_GetTouchDeviceType(id SDL_TouchID) SDL_TouchDeviceType {
	return SDL_TouchDeviceType(C.SDL_GetTouchDeviceType(C.SDL_TouchID(id)))
}

func SDL_GetNumTouchFingers(id SDL_TouchID) int {
	cId := cTouchID(id)

	return int(C.SDL_GetNumTouchFingers(cId))
}

func SDL_GetTouchFinger(id SDL_TouchID, index int) *SDL_Finger {
	cId := cTouchID(id)

	cFinger := C.SDL_GetTouchFinger(cId, cInt(index))
	return (*SDL_Finger)(unsafe.Pointer(cFinger))
}
