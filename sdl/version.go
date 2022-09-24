package sdl

/*
#include "version.h"
*/
import "C"
import "unsafe"

func (v *SDL_version) cSDL_version() *C.SDL_version {
	return (*C.SDL_version)(unsafe.Pointer(v))
}

// VERSION fills the selected struct with the version of SDL in use.
func SDL_VERSION(v *SDL_version) {
	v.Major = SDL_MAJOR_VERSION
	v.Minor = SDL_MINOR_VERSION
	v.Patch = SDL_PATCHLEVEL
}

// VERSIONNUM converts separate version components into a single numeric value.
func SDL_VERSIONNUM(x, y, z int) int {
	return (x*1000 + y*100 + z)
}

// COMPILEDVERSION returns the SDL version number that you compiled against.
func SDL_COMPILEDVERSION() int {
	return SDL_VERSIONNUM(SDL_MAJOR_VERSION, SDL_MINOR_VERSION, SDL_PATCHLEVEL)
}

// VERSION_ATLEAST reports whether the SDL version compiled against is at least as new as the specified version.
func SDL_VERSION_ATLEAST(x, y, z int) bool {
	return SDL_COMPILEDVERSION() >= SDL_VERSIONNUM(x, y, z)
}

// GetVersion returns the version of SDL that is linked against your program.
func SDL_GetVersion(v *SDL_version) {
	C.SDL_GetVersion(v.cSDL_version())
}

// GetRevision returns the code revision of SDL that is linked against your program.
func SDL_GetRevision() string {
	return SDL_GoString(C.SDL_GetRevision())
}

// Deprecated: GetRevisionNumber is deprecated in SDL2 2.0.16 and will return 0. Users should use GetRevision instead.
func GetRevisionNumber() int {
	return (int)(C.GetRevisionNumber())
}
