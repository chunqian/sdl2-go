package sdl

/*
#include "guid.h"
*/
import "C"
import (
	"unsafe"
)

func cSDL_GUID(g *SDL_GUID) *C.SDL_GUID {
	return (*C.SDL_GUID)(unsafe.Pointer(g))
}

func SDL_GUIDToString(guid SDL_GUID, pszGUID *byte, cbGUID int) {
	cPszGUID := (*cChar)(unsafe.Pointer(pszGUID))
	cGuid := cSDL_GUID(&guid)
	C.SDL_GUIDToString(*cGuid, cPszGUID, cInt(cbGUID))
}

func SDL_GUIDFromString(pchGUID string) SDL_GUID {
	cPchGUID := SDL_CreateCString(SDL_GetMemoryPool(), pchGUID)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cPchGUID)

	cGuid := C.SDL_GUIDFromString(cPchGUID.(*cChar))
	return *(*SDL_GUID)(unsafe.Pointer(&cGuid))
}
