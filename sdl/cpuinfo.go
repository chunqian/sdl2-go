package sdl

/*
#include "cpuinfo.h"
*/
import "C"
import "unsafe"

// define
const (
	SDL_CACHELINE_SIZE = 128
)

func SDL_GetCPUCount() int {
	return int(C.SDL_GetCPUCount())
}

func SDL_GetCPUCacheLineSize() int {
	return int(C.SDL_GetCPUCacheLineSize())
}

func SDL_HasRDTSC() bool {
	cRet := C.SDL_HasRDTSC()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasAltiVec() bool {
	cRet := C.SDL_HasAltiVec()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasMMX() bool {
	cRet := C.SDL_HasMMX()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_Has3DNow() bool {
	cRet := C.SDL_Has3DNow()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasSSE() bool {
	cRet := C.SDL_HasSSE()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasSSE2() bool {
	cRet := C.SDL_HasSSE2()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasSSE3() bool {
	cRet := C.SDL_HasSSE3()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasSSE41() bool {
	cRet := C.SDL_HasSSE41()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasSSE42() bool {
	cRet := C.SDL_HasSSE42()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_GetSystemRAM() int {
	return int(C.SDL_GetSystemRAM())
}

func SDL_HasAVX() bool {
	cRet := C.SDL_HasAVX()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasAVX512F() bool {
	cRet := C.SDL_HasAVX512F()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasAVX2() bool {
	cRet := C.SDL_HasAVX2()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasARMSIMD() bool {
	cRet := C.SDL_HasARMSIMD()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_HasNEON() bool {
	cRet := C.SDL_HasNEON()
	if int(cRet) > 0 {
		return true
	}
	return false
}

func SDL_SIMDGetAlignment() int {
	return int(C.SDL_SIMDGetAlignment())
}

func SDL_SIMDAlloc(memLen int) unsafe.Pointer {
	return C.SDL_SIMDAlloc(cSize(memLen))
}

func SDL_SIMDRealloc(mem unsafe.Pointer, memLen int) unsafe.Pointer {
	return C.SDL_SIMDRealloc(mem, cSize(memLen))
}

func SDL_SIMDFree(p unsafe.Pointer) {
	C.SDL_SIMDFree(p)
}
