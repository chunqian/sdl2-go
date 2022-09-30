package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"

func SDL_OpenURL(url string) int {
	cUrl := SDL_CreateCString(SDL_GetMemoryPool(), url)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cUrl)

	cRet := C.SDL_OpenURL(cUrl.(*cChar))
	return int(cRet)
}
