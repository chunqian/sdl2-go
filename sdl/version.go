package sdl

/*
#include "version.h"
*/
import "C"
import "unsafe"

// #define
const (
	SDL_MAJOR_VERSION = 2
	SDL_MINOR_VERSION = 0
	SDL_PATCHLEVEL    = 22
)

func cVersion(v *SDL_version) *C.SDL_version {
	return (*C.SDL_version)(unsafe.Pointer(v))
}

func SDL_VERSION(v *SDL_version) {
	v.Major = SDL_MAJOR_VERSION
	v.Minor = SDL_MINOR_VERSION
	v.Patch = SDL_PATCHLEVEL
}

func SDL_VERSIONNUM(x, y, z int) int {
	return (x*1000 + y*100 + z)
}

func SDL_COMPILEDVERSION() int {
	return SDL_VERSIONNUM(SDL_MAJOR_VERSION, SDL_MINOR_VERSION, SDL_PATCHLEVEL)
}

func SDL_VERSION_ATLEAST(x, y, z int) bool {
	return SDL_COMPILEDVERSION() >= SDL_VERSIONNUM(x, y, z)
}

func SDL_GetVersion(v *SDL_version) {
	C.SDL_GetVersion(cVersion(v))
}

func SDL_GetRevision() string {
	return SDL_GoString(C.SDL_GetRevision())
}

func GetRevisionNumber() int {
	return (int)(C.GetRevisionNumber())
}
