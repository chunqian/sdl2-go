package sdl

/*
#include "rwops.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func cRWops(rw *SDL_RWops) *C.SDL_RWops {
	return (*C.SDL_RWops)(unsafe.Pointer(rw))
}

func SDL_RWFromFile(file, mode string) *SDL_RWops {
	cFile := SDL_CreateCString(SDL_GetMemoryPool(), file)
	cMode := SDL_CreateCString(SDL_GetMemoryPool(), mode)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cFile) // memory free
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cMode) // memory free

	cRW := C.SDL_RWFromFile(cFile.(*cChar), cMode.(*cChar))
	return (*SDL_RWops)(unsafe.Pointer(cRW))
}

func SDL_RWFromMem(mem []byte) *SDL_RWops {
	if mem == nil {
		return nil
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	cRW := C.SDL_RWFromMem(unsafe.Pointer(header.Data), cInt(len(mem)))
	rw := (*SDL_RWops)(unsafe.Pointer(cRW))
	return rw
}

func SDL_AllocRW() *SDL_RWops {
	return (*SDL_RWops)(unsafe.Pointer(C.SDL_AllocRW()))
}

func SDL_FreeRW(rw *SDL_RWops) {
	C.SDL_FreeRW(cRWops(rw))
}

func SDL_RWsize(rw *SDL_RWops) int64 {
	cN := C.RWsize(cRWops(rw))
	return int64(cN)
}

func SDL_RWseek(rw *SDL_RWops, offset int64, whence int32) int64 {
	if rw == nil {
		return -1
	}
	cRet := C.RWseek(cRWops(rw), cInt64(offset), cInt(whence))
	return int64(cRet)
}

func SDL_RWread(rw *SDL_RWops, buf []byte) int {
	return rw.reading(buf, 1, uint(len(buf)))
}

func (rw *SDL_RWops) reading(buf []byte, size, maxnum uint) int {
	if rw == nil || buf == nil {
		return 0
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cN := C.RWread(cRWops(rw), unsafe.Pointer(header.Data), cSize_t(size), cSize_t(maxnum))
	return int(cN)
}

func SDL_RWtell(rw *SDL_RWops) int64 {
	if rw == nil {
		return 0
	}
	cRet := C.RWseek(cRWops(rw), 0, cInt(RW_SEEK_CUR))
	return int64(cRet)
}

func SDL_RWwrite(rw *SDL_RWops, buf []byte) int {
	return rw.writing(buf, 1, uint(len(buf)))
}

func (rw *SDL_RWops) writing(buf []byte, size, num uint) int {
	if rw == nil || buf == nil {
		return 0
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cN := C.RWwrite(cRWops(rw), unsafe.Pointer(header.Data), cSize_t(size), cSize_t(num))
	return int(cN)
}

func SDL_RWclose(rw *SDL_RWops) int {
	cRet := C.RWclose(cRWops(rw))
	if rw != nil && int(cRet) != 0 {
		return -1
	}
	return 0
}

func SDL_ReadU8(rw *SDL_RWops) uint8 {
	if rw == nil {
		return 0
	}
	return uint8(C.SDL_ReadU8(cRWops(rw)))
}

func SDL_ReadLE16(rw *SDL_RWops) uint16 {
	if rw == nil {
		return 0
	}
	return uint16(C.SDL_ReadLE16(cRWops(rw)))
}

func SDL_ReadBE16(rw *SDL_RWops) uint16 {
	if rw == nil {
		return 0
	}
	return uint16(C.SDL_ReadBE16(cRWops(rw)))
}

func SDL_ReadLE32(rw *SDL_RWops) uint32 {
	if rw == nil {
		return 0
	}
	return uint32(C.SDL_ReadLE32(cRWops(rw)))
}

func SDL_ReadBE32(rw *SDL_RWops) uint32 {
	if rw == nil {
		return 0
	}
	return uint32(C.SDL_ReadBE32(cRWops(rw)))
}

func SDL_ReadLE64(rw *SDL_RWops) uint64 {
	if rw == nil {
		return 0
	}
	return uint64(C.SDL_ReadLE64(cRWops(rw)))
}

func SDL_ReadBE64(rw *SDL_RWops) uint64 {
	if rw == nil {
		return 0
	}
	return uint64(C.SDL_ReadBE64(cRWops(rw)))
}

func SDL_LoadFile_RW(rw *SDL_RWops, freesrc SDL_bool) (data []byte, size int) {
	var cSize cSize_t
	var cFreesrc cInt = 0

	if freesrc == SDL_TRUE {
		cFreesrc = 1
	}

	cData := C.SDL_LoadFile_RW(cRWops(rw), &cSize, cFreesrc)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Cap = int(cSize)
	sliceHeader.Len = int(cSize)
	sliceHeader.Data = uintptr(cData)
	size = int(cSize)
	return
}

func SDL_LoadFile(file string) (data []byte, size int) {
	rw := SDL_RWFromFile(file, "rb")
	return SDL_LoadFile_RW(rw, SDL_TRUE)
}

func SDL_WriteU8(rw *SDL_RWops, value uint8) uint {
	if rw == nil {
		return 0
	}
	return uint(C.SDL_WriteU8(cRWops(rw), cUint8(value)))
}

func SDL_WriteLE16(rw *SDL_RWops, value uint16) uint {
	if rw == nil {
		return 0
	}
	return uint(C.SDL_WriteLE16(cRWops(rw), cUint16(value)))
}

func SDL_WriteBE16(rw *SDL_RWops, value uint16) uint {
	if rw == nil {
		return 0
	}
	return uint(C.SDL_WriteBE16(cRWops(rw), cUint16(value)))
}

func SDL_WriteLE32(rw *SDL_RWops, value uint32) uint {
	if rw == nil {
		return 0
	}
	return uint(C.SDL_WriteLE32(cRWops(rw), cUint32(value)))
}

func SDL_WriteBE32(rw *SDL_RWops, value uint32) uint {
	if rw == nil {
		return 0
	}
	return uint(C.SDL_WriteBE32(cRWops(rw), cUint32(value)))
}

func SDL_WriteLE64(rw *SDL_RWops, value uint64) uint {
	if rw == nil {
		return 0
	}
	return uint(C.SDL_WriteLE64(cRWops(rw), cUint64(value)))
}

func SDL_WriteBE64(rw *SDL_RWops, value uint64) uint {
	if rw == nil {
		return 0
	}
	return uint(C.SDL_WriteBE64(cRWops(rw), cUint64(value)))
}
