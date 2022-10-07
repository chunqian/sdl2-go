package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"

func SDL_OpenURL(url string) int {
	cUrl := createCString(SDL_GetMemoryPool(), url)
	defer destroyCString(SDL_GetMemoryPool(), cUrl)

	cRet := C.SDL_OpenURL(cUrl)
	return int(cRet)
}
