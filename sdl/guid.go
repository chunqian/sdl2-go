package sdl

/*
#include "guid.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func cGUID(g *SDL_GUID) *C.SDL_GUID {
	return (*C.SDL_GUID)(unsafe.Pointer(g))
}

func SDL_GUIDToString(guid SDL_GUID, pszGUID []byte, cbGUID int) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&pszGUID))
	sh.Len = int(cbGUID)
	sh.Cap = int(cbGUID)
	cPszGUID := (*cChar)(unsafe.Pointer(sh.Data))
	cGuid := cGUID(&guid)
	C.SDL_GUIDToString(*cGuid, cPszGUID, cInt(cbGUID))
}

func SDL_GUIDFromString(pchGUID string) SDL_GUID {
	cPchGUID := createCString(SDL_GetMemoryPool(), pchGUID)
	defer destroyCString(SDL_GetMemoryPool(), cPchGUID)

	cGuid := C.SDL_GUIDFromString(cPchGUID)
	return *(*SDL_GUID)(unsafe.Pointer(&cGuid))
}
