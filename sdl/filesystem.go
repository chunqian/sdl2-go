package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

func SDL_GetBasePath() string {
	cVal := C.SDL_GetBasePath()
	defer C.free(unsafe.Pointer(cVal))

	return createGoString(cVal)
}

func SDL_GetPrefPath(org, app string) string {
	cOrg := createCString(SDL_GetMemoryPool(), org)
	cApp := createCString(SDL_GetMemoryPool(), app)
	defer destroyCString(SDL_GetMemoryPool(), cOrg)
	defer destroyCString(SDL_GetMemoryPool(), cApp)

	cVal := C.SDL_GetPrefPath(cOrg, cApp)
	defer C.free(unsafe.Pointer(cVal))

	return createGoString(cVal)
}
