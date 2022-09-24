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

// RWFromFile creates a new RWops structure for reading from and/or writing to a named file.
func SDL_RWFromFile(file, mode string) *SDL_RWops {
	cFile := SDL_CreateCString(SDL_GetMemoryPool(), file)
	cMode := SDL_CreateCString(SDL_GetMemoryPool(), mode)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cFile) // memory free
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cMode) // memory free

	return (*SDL_RWops)(unsafe.Pointer(C.SDL_RWFromFile(cFile, cMode)))
}

// RWFromMem prepares a read-write memory buffer for use with RWops.
func SDL_RWFromMem(mem []byte) *SDL_RWops {
	if mem == nil {
		return nil
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	cRwops := C.SDL_RWFromMem(unsafe.Pointer(header.Data), cInt(len(mem)))
	rwops := (*SDL_RWops)(unsafe.Pointer(cRwops))
	return rwops
}

// AllocRW allocates an empty, unpopulated RWops structure.
func SDL_AllocRW() *SDL_RWops {
	return (*SDL_RWops)(unsafe.Pointer(C.SDL_AllocRW()))
}

// Free frees the RWops structure allocated by AllocRW().
func SDL_FreeRW(rwops *SDL_RWops) {
	C.SDL_FreeRW(rwops.cSDL_RWops())
}

// Size returns the size of the data stream in the RWops.
func SDL_RWsize(rwops *SDL_RWops) int64 {
	cN := C.RWsize(rwops.cSDL_RWops())
	return int64(cN)
}

// Seek seeks within the RWops data stream.
func SDL_RWseek(rwops *SDL_RWops, offset int64, whence int32) int64 {
	if rwops == nil {
		return -1
	}
	cRet := C.RWseek(rwops.cSDL_RWops(), cInt64(offset), cInt(whence))
	return int64(cRet)
}

// Read reads from a data source.
func SDL_RWread(rwops *SDL_RWops, buf []byte) int {
	return rwops.read(buf, 1, uint(len(buf)))
}

// Read reads from a data source (native).
func (rwops *SDL_RWops) read(buf []byte, size, maxnum uint) int {
	if rwops == nil || buf == nil {
		return 0
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cN := C.RWread(rwops.cSDL_RWops(), unsafe.Pointer(header.Data), cSize_t(size), cSize_t(maxnum))
	return int(cN)
}

// Tell returns the current read/write offset in the RWops data stream.
func SDL_RWtell(rwops *SDL_RWops) int64 {
	if rwops == nil {
		return 0
	}
	cRet := C.RWseek(rwops.cSDL_RWops(), 0, cInt(RW_SEEK_CUR))
	return int64(cRet)
}

// Write writes to the RWops data stream.
func SDL_RWwrite(rwops *SDL_RWops, buf []byte) int {
	return rwops.write(buf, 1, uint(len(buf)))
}

// Write writes to the RWops data stream (native).
func (rwops *SDL_RWops) write(buf []byte, size, num uint) int {
	if rwops == nil || buf == nil {
		return 0
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cN := C.RWwrite(rwops.cSDL_RWops(), unsafe.Pointer(header.Data), cSize_t(size), cSize_t(num))
	return int(cN)
}

// Close closes and frees the allocated RWops structure.
func SDL_RWclose(rwops *SDL_RWops) int {
	cRet := C.RWclose(rwops.cSDL_RWops())
	if rwops != nil && int(cRet) != 0 {
		return -1
	}
	return 0
}

// ReadU8 reads a byte from the RWops.
func SDL_ReadU8(rwops *SDL_RWops) uint8 {
	if rwops == nil {
		return 0
	}
	return uint8(C.SDL_ReadU8(rwops.cSDL_RWops()))
}

// ReadLE16 reads 16 bits of little-endian data from the RWops and returns in native format.
func SDL_ReadLE16(rwops *SDL_RWops) uint16 {
	if rwops == nil {
		return 0
	}
	return uint16(C.SDL_ReadLE16(rwops.cSDL_RWops()))
}

// ReadBE16 read 16 bits of big-endian data from the RWops and returns in native format.
func SDL_ReadBE16(rwops *SDL_RWops) uint16 {
	if rwops == nil {
		return 0
	}
	return uint16(C.SDL_ReadBE16(rwops.cSDL_RWops()))
}

// ReadLE32 reads 32 bits of little-endian data from the RWops and returns in native format.
func SDL_ReadLE32(rwops *SDL_RWops) uint32 {
	if rwops == nil {
		return 0
	}
	return uint32(C.SDL_ReadLE32(rwops.cSDL_RWops()))
}

// ReadBE32 reads 32 bits of big-endian data from the RWops and returns in native format.
func SDL_ReadBE32(rwops *SDL_RWops) uint32 {
	if rwops == nil {
		return 0
	}
	return uint32(C.SDL_ReadBE32(rwops.cSDL_RWops()))
}

// ReadLE64 reads 64 bits of little-endian data from the RWops and returns in native format.
func SDL_ReadLE64(rwops *SDL_RWops) uint64 {
	if rwops == nil {
		return 0
	}
	return uint64(C.SDL_ReadLE64(rwops.cSDL_RWops()))
}

// ReadBE64 reads 64 bits of big-endian data from the RWops and returns in native format.
func SDL_ReadBE64(rwops *SDL_RWops) uint64 {
	if rwops == nil {
		return 0
	}
	return uint64(C.SDL_ReadBE64(rwops.cSDL_RWops()))
}

// LoadFile_RW loads all the data from an SDL data stream.
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

// LoadFile loads an entire file
func SDL_LoadFile(file string) (data []byte, size int) {
	rwops := SDL_RWFromFile(file, "rb")
	return SDL_LoadFile_RW(rwops, SDL_TRUE)
}

// WriteU8 writes a byte to the RWops.
func SDL_WriteU8(rwops *SDL_RWops, value uint8) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteU8(rwops.cSDL_RWops(), cUint8(value)))
}

// WriteLE16 writes 16 bits in native format to the RWops as little-endian data.
func SDL_WriteLE16(rwops *SDL_RWops, value uint16) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteLE16(rwops.cSDL_RWops(), cUint16(value)))
}

// WriteBE16 writes 16 bits in native format to the RWops as big-endian data.
func SDL_WriteBE16(rwops *SDL_RWops, value uint16) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteBE16(rwops.cSDL_RWops(), cUint16(value)))
}

// WriteLE32 writes 32 bits in native format to the RWops as little-endian data.
func SDL_WriteLE32(rwops *SDL_RWops, value uint32) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteLE32(rwops.cSDL_RWops(), cUint32(value)))
}

// WriteBE32 writes 32 bits in native format to the RWops as big-endian data.
func SDL_WriteBE32(rwops *SDL_RWops, value uint32) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteBE32(rwops.cSDL_RWops(), cUint32(value)))
}

// WriteLE64 writes 64 bits in native format to the RWops as little-endian data.
func SDL_WriteLE64(rwops *SDL_RWops, value uint64) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteLE64(rwops.cSDL_RWops(), cUint64(value)))
}

// WriteBE64 writes 64 bits in native format to the RWops as big-endian data.
func SDL_WriteBE64(rwops *SDL_RWops, value uint64) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteBE64(rwops.cSDL_RWops(), cUint64(value)))
}
