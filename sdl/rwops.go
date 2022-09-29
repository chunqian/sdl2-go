package sdl

/*
#include "rwops.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func (rwops *SDL_RWops) cSDL_RWops() *C.SDL_RWops {
	return (*C.SDL_RWops)(unsafe.Pointer(rwops))
}

func SDL_RWFromFile(file, mode string) *SDL_RWops {
	cFile := SDL_CreateCString(SDL_GetMemoryPool(), file)
	cMode := SDL_CreateCString(SDL_GetMemoryPool(), mode)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cFile) // memory free
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cMode) // memory free

	return (*SDL_RWops)(unsafe.Pointer(C.SDL_RWFromFile(cFile, cMode)))
}

func SDL_RWFromMem(mem []byte) *SDL_RWops {
	if mem == nil {
		return nil
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	cRwops := C.SDL_RWFromMem(unsafe.Pointer(header.Data), cInt(len(mem)))
	rwops := (*SDL_RWops)(unsafe.Pointer(cRwops))
	return rwops
}

func SDL_AllocRW() *SDL_RWops {
	return (*SDL_RWops)(unsafe.Pointer(C.SDL_AllocRW()))
}

func SDL_FreeRW(rwops *SDL_RWops) {
	C.SDL_FreeRW(rwops.cSDL_RWops())
}

func SDL_RWsize(rwops *SDL_RWops) int64 {
	cN := C.RWsize(rwops.cSDL_RWops())
	return int64(cN)
}

func SDL_RWseek(rwops *SDL_RWops, offset int64, whence int32) int64 {
	if rwops == nil {
		return -1
	}
	cRet := C.RWseek(rwops.cSDL_RWops(), cInt64(offset), cInt(whence))
	return int64(cRet)
}

func SDL_RWread(rwops *SDL_RWops, buf []byte) int {
	return rwops.reading(buf, 1, uint(len(buf)))
}

func (rwops *SDL_RWops) reading(buf []byte, size, maxnum uint) int {
	if rwops == nil || buf == nil {
		return 0
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cN := C.RWread(rwops.cSDL_RWops(), unsafe.Pointer(header.Data), cSize_t(size), cSize_t(maxnum))
	return int(cN)
}

func SDL_RWtell(rwops *SDL_RWops) int64 {
	if rwops == nil {
		return 0
	}
	cRet := C.RWseek(rwops.cSDL_RWops(), 0, cInt(RW_SEEK_CUR))
	return int64(cRet)
}

func SDL_RWwrite(rwops *SDL_RWops, buf []byte) int {
	return rwops.writing(buf, 1, uint(len(buf)))
}

func (rwops *SDL_RWops) writing(buf []byte, size, num uint) int {
	if rwops == nil || buf == nil {
		return 0
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cN := C.RWwrite(rwops.cSDL_RWops(), unsafe.Pointer(header.Data), cSize_t(size), cSize_t(num))
	return int(cN)
}

func SDL_RWclose(rwops *SDL_RWops) int {
	cRet := C.RWclose(rwops.cSDL_RWops())
	if rwops != nil && int(cRet) != 0 {
		return -1
	}
	return 0
}

func SDL_ReadU8(rwops *SDL_RWops) uint8 {
	if rwops == nil {
		return 0
	}
	return uint8(C.SDL_ReadU8(rwops.cSDL_RWops()))
}

func SDL_ReadLE16(rwops *SDL_RWops) uint16 {
	if rwops == nil {
		return 0
	}
	return uint16(C.SDL_ReadLE16(rwops.cSDL_RWops()))
}

func SDL_ReadBE16(rwops *SDL_RWops) uint16 {
	if rwops == nil {
		return 0
	}
	return uint16(C.SDL_ReadBE16(rwops.cSDL_RWops()))
}

func SDL_ReadLE32(rwops *SDL_RWops) uint32 {
	if rwops == nil {
		return 0
	}
	return uint32(C.SDL_ReadLE32(rwops.cSDL_RWops()))
}

func SDL_ReadBE32(rwops *SDL_RWops) uint32 {
	if rwops == nil {
		return 0
	}
	return uint32(C.SDL_ReadBE32(rwops.cSDL_RWops()))
}

func SDL_ReadLE64(rwops *SDL_RWops) uint64 {
	if rwops == nil {
		return 0
	}
	return uint64(C.SDL_ReadLE64(rwops.cSDL_RWops()))
}

func SDL_ReadBE64(rwops *SDL_RWops) uint64 {
	if rwops == nil {
		return 0
	}
	return uint64(C.SDL_ReadBE64(rwops.cSDL_RWops()))
}

func SDL_LoadFile_RW(rwops *SDL_RWops, freesrc SDL_bool) (data []byte, size int) {
	var cSize cSize_t
	var cFreesrc cInt = 0

	if freesrc == SDL_TRUE {
		cFreesrc = 1
	}

	cData := C.SDL_LoadFile_RW(rwops.cSDL_RWops(), &cSize, cFreesrc)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Cap = int(cSize)
	sliceHeader.Len = int(cSize)
	sliceHeader.Data = uintptr(cData)
	size = int(cSize)
	return
}

func SDL_LoadFile(file string) (data []byte, size int) {
	rwops := SDL_RWFromFile(file, "rb")
	return SDL_LoadFile_RW(rwops, SDL_TRUE)
}

func SDL_WriteU8(rwops *SDL_RWops, value uint8) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteU8(rwops.cSDL_RWops(), cUint8(value)))
}

func SDL_WriteLE16(rwops *SDL_RWops, value uint16) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteLE16(rwops.cSDL_RWops(), cUint16(value)))
}

func SDL_WriteBE16(rwops *SDL_RWops, value uint16) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteBE16(rwops.cSDL_RWops(), cUint16(value)))
}

func SDL_WriteLE32(rwops *SDL_RWops, value uint32) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteLE32(rwops.cSDL_RWops(), cUint32(value)))
}

func SDL_WriteBE32(rwops *SDL_RWops, value uint32) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteBE32(rwops.cSDL_RWops(), cUint32(value)))
}

func SDL_WriteLE64(rwops *SDL_RWops, value uint64) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteLE64(rwops.cSDL_RWops(), cUint64(value)))
}

func SDL_WriteBE64(rwops *SDL_RWops, value uint64) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteBE64(rwops.cSDL_RWops(), cUint64(value)))
}
