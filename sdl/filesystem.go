package sdl

/*
#include "filesystem.h"
*/
import "C"
import "unsafe"

func SDL_GetBasePath() string {
	cVal := C.SDL_GetBasePath()
	defer C.free(unsafe.Pointer(cVal))

	return SDL_GoString(cVal)
}

func SDL_GetPrefPath(org, app string) string {
	cOrg := SDL_CreateCString(SDL_GetMemoryPool(), org)
	cApp := SDL_CreateCString(SDL_GetMemoryPool(), app)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cOrg)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cApp)

	cVal := C.SDL_GetPrefPath(cOrg.(*cChar), cApp.(*cChar))
	defer C.free(unsafe.Pointer(cVal))

	return SDL_GoString(cVal)
}
