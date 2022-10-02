package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

type SharedObject unsafe.Pointer

func SDL_LoadObject(sofile string) SharedObject {
	cSofile := SDL_CreateCString(SDL_GetMemoryPool(), sofile)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cSofile)

	sharedObject := C.SDL_LoadObject(cSofile.(*cChar))
	return (SharedObject)(unsafe.Pointer(sharedObject))
}

func SDL_LoadFunction(handle SharedObject, name string) unsafe.Pointer {
	cName := SDL_CreateCString(SDL_GetMemoryPool(), name)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cName)

	return C.SDL_LoadFunction(unsafe.Pointer(handle), cName.(*cChar))
}

func SDL_UnloadObject(handle SharedObject) {
	C.SDL_UnloadObject(unsafe.Pointer(handle))
}
