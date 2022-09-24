package sdl

/*
#include "cpuinfo.h"
*/
import "C"
import "unsafe"

// GetCPUCount returns the number of CPU cores available.
func SDL_GetCPUCount() int {
	return int(C.SDL_GetCPUCount())
}

// GetCPUCacheLineSize returns the L1 cache line size of the CPU.
func SDL_GetCPUCacheLineSize() int {
	return int(C.SDL_GetCPUCacheLineSize())
}

// HasRDTSC reports whether the CPU has the RDTSC instruction.
func SDL_HasRDTSC() bool {
	cRet := C.SDL_HasRDTSC()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasAltiVec reports whether the CPU has AltiVec features.
func SDL_HasAltiVec() bool {
	cRet := C.SDL_HasAltiVec()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasMMX reports whether the CPU has MMX features.
func SDL_HasMMX() bool {
	cRet := C.SDL_HasMMX()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// Has3DNow reports whether the CPU has 3DNow! features.
func SDL_Has3DNow() bool {
	cRet := C.SDL_Has3DNow()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasSSE reports whether the CPU has SSE features.
func SDL_HasSSE() bool {
	cRet := C.SDL_HasSSE()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasSSE2 reports whether the CPU has SSE2 features.
func SDL_HasSSE2() bool {
	cRet := C.SDL_HasSSE2()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasSSE3 reports whether the CPU has SSE3 features.
func SDL_HasSSE3() bool {
	cRet := C.SDL_HasSSE3()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasSSE41 reports whether the CPU has SSE4.1 features.
func SDL_HasSSE41() bool {
	cRet := C.SDL_HasSSE41()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasSSE42 reports whether the CPU has SSE4.2 features.
func SDL_HasSSE42() bool {
	cRet := C.SDL_HasSSE42()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// GetSystemRAM returns the amount of RAM configured in the system.
func SDL_GetSystemRAM() int {
	return int(C.SDL_GetSystemRAM())
}

// HasAVX reports whether the CPU has AVX features.
func SDL_HasAVX() bool {
	cRet := C.SDL_HasAVX()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasAVX512F reports whether the CPU has AVX-512F (foundation) features.
func SDL_HasAVX512F() bool {
	cRet := C.SDL_HasAVX512F()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasAVX2 reports whether the CPU has AVX2 features.
func SDL_HasAVX2() bool {
	cRet := C.SDL_HasAVX2()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasARMSIMD reports whether the CPU has ARM SIMD (ARMv6) features.
func SDL_HasARMSIMD() bool {
	cRet := C.SDL_HasARMSIMD()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// HasNEON reports whether the CPU has NEON features.
func SDL_HasNEON() bool {
	cRet := C.SDL_HasNEON()
	if int(cRet) > 0 {
		return true
	}
	return false
}

// SIMDGetAlignment reports the alignment this system needs for SIMD allocations.
func SDL_SIMDGetAlignment() int {
	return int(C.SDL_SIMDGetAlignment())
}

// SIMDAlloc allocates memory in a SIMD-friendly way.
func SDL_SIMDAlloc(memLen int) unsafe.Pointer {
	return C.SDL_SIMDAlloc(cSize_t(memLen))
}

// SIMDRealloc reallocates memory obtained from SDL_SIMDAlloc.
func SDL_SIMDRealloc(mem unsafe.Pointer, memLen int) unsafe.Pointer {
	return C.SDL_SIMDRealloc(mem, cSize_t(memLen))
}

// SIMDFree deallocates memory obtained from SDL_SIMDAlloc.
func SDL_SIMDFree(p unsafe.Pointer) {
	C.SDL_SIMDFree(p)
}
