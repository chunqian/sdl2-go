package sdl

/*
#include "filesystem.h"
*/
import "C"
import "unsafe"

// GetBasePath returns the directory where the application was run from. This is where the application data directory is.
func SDL_GetBasePath() string {
	cVal := C.SDL_GetBasePath()
	defer C.free(unsafe.Pointer(cVal))

	return SDL_GoString(cVal)
}

// GetPrefPath returns the "pref dir".
// This is meant to be where the application can write personal files (Preferences and save games, etc.) that are specific to the application.
// This directory is unique per user and per application.
func SDL_GetPrefPath(org, app string) string {
	cOrg := SDL_CreateCString(SDL_GetMemoryPool(), org)
	cApp := SDL_CreateCString(SDL_GetMemoryPool(), app)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cOrg)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cApp)

	cVal := C.SDL_GetPrefPath(cOrg, cApp)
	defer C.free(unsafe.Pointer(cVal))

	return SDL_GoString(cVal)
}
