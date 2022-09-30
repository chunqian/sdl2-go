package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"

func SDL_RecordGesture(t SDL_TouchID) int {
	return int(C.SDL_RecordGesture(cSDL_TouchID(t)))
}

func SDL_SaveAllDollarTemplates(src *SDL_RWops) int {
	cRet := C.SDL_SaveAllDollarTemplates(cSDL_RWops(src))
	return int(cRet)
}

func SDL_SaveDollarTemplate(g SDL_GestureID, src *SDL_RWops) int {
	cRet := C.SDL_SaveDollarTemplate(ccSDL_GestureID(g), cSDL_RWops(src))
	return int(cRet)
}

func SDL_LoadDollarTemplates(t SDL_TouchID, src *SDL_RWops) int {
	cRet := C.SDL_LoadDollarTemplates(cSDL_TouchID(t), cSDL_RWops(src))
	return int(cRet)
}
