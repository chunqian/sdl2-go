package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

// typedef
type SharedObject unsafe.Pointer

func SDL_LoadObject(sofile string) SharedObject {
	cSofile := createCString(SDL_GetMemoryPool(), sofile)
	defer destroyCString(SDL_GetMemoryPool(), cSofile)

	sharedObject := C.SDL_LoadObject(cSofile)
	return (SharedObject)(unsafe.Pointer(sharedObject))
}

func SDL_LoadFunction(handle SharedObject, name string) unsafe.Pointer {
	cName := createCString(SDL_GetMemoryPool(), name)
	defer destroyCString(SDL_GetMemoryPool(), cName)

	return C.SDL_LoadFunction(unsafe.Pointer(handle), cName)
}

func SDL_UnloadObject(handle SharedObject) {
	C.SDL_UnloadObject(unsafe.Pointer(handle))
}
