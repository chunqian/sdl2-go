package gfx

/*
#include "gfx_wrapper.h"
*/
import "C"
import (
	"unsafe"

	. "github.com/chunqian/sdl2-go/sdl"
)

func (manager *FPSmanager) cFPSmanager() *C.FPSmanager {
	return (*C.FPSmanager)(unsafe.Pointer(manager))
}
