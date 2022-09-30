#ifndef _GO_SDL_MOUSE_H // if {
#define _GO_SDL_MOUSE_H

#include "sdl_wrapper.h"

#if defined(__WIN32) // if {
#include <SDL2/SDL_syswm.h>
#else // } else {
#include <SDL_syswm.h>
#endif // }

#endif // }
