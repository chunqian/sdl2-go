package sdl

import "C"
import (
	"reflect"
	unsafe "unsafe"

	. "github.com/chunqian/memory"
)

const memorySize = 1024 * 1024 * 20 // 20M
var calcBuffer [memorySize]PX_char
var memoryPool PX_memorypool

func init() {
	// 内存池初始化
	calcBuffer = [memorySize]PX_char{PX_char(0)}
	memoryPool = MP_Create(unsafe.Pointer(&calcBuffer), PX_uint(memorySize))
}

func SDL_GetMemoryPool() *PX_memorypool {
	return &memoryPool
}

func SDL_malloc(mp *PX_memorypool, maxlen int) unsafe.Pointer {
	n := PX_uint(maxlen)
	p := MP_Malloc(mp, n)
	PX_memset(p, PX_byte(0), PX_int(n))
	return p
}

func SDL_free(mp *PX_memorypool, paddr unsafe.Pointer) {
	MP_Free(mp, paddr)
}

func SDL_CreateDataStructures[T SDL_DataStructures](mp *PX_memorypool, d T) *T {
	n := PX_uint(unsafe.Sizeof(d))
	p := MP_Malloc(mp, n)
	PX_memcpy(p, unsafe.Pointer(&d), PX_int(n))
	return (*T)(p)
}

func SDL_DestroyDataStructures[T SDL_DataStructures](mp *PX_memorypool, d *T) {
	MP_Free(mp, unsafe.Pointer(d))
}

func SDL_CreateArrayDataStructures[T SDL_DataStructures](mp *PX_memorypool, d []T) *[]T {
	dh := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	n := PX_uint(dh.Len)

	p := MP_Malloc(mp, n)
	PX_memcpy(p, unsafe.Pointer(dh.Data), PX_int(n))

	var slice []uint8
	length := len(d)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sh.Data = uintptr(unsafe.Pointer(p))
	sh.Len = int(length)
	sh.Cap = int(length)
	return (*[]T)(unsafe.Pointer(sh))
}

func SDL_DestroyArrayDataStructures[T SDL_DataStructures](mp *PX_memorypool, d *[]T) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&d))
	MP_Free(mp, unsafe.Pointer(sh.Data))
}
