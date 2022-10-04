package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import (
	"math"
	"unsafe"
)

// define
const (
	SDL_TOUCH_MOUSEID = math.MaxUint32
	SDL_MOUSE_TOUCHID = int64(-1)
)

// typedef
type SDL_TouchID = int64
type SDL_FingerID = int64
type SDL_TouchDeviceType = int32

// enum
const (
	SDL_TOUCH_DEVICE_INVALID           SDL_TouchDeviceType = -1
	SDL_TOUCH_DEVICE_DIRECT            SDL_TouchDeviceType = 0
	SDL_TOUCH_DEVICE_INDIRECT_ABSOLUTE SDL_TouchDeviceType = 1
	SDL_TOUCH_DEVICE_INDIRECT_RELATIVE SDL_TouchDeviceType = 2
)

// struct
type SDL_Finger struct {
	Id       int64
	X        float32
	Y        float32
	Pressure float32
}

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
