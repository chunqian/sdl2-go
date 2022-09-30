#ifndef _GO_SDL_ERROR_H // if {
#define _GO_SDL_ERROR_H

#include "sdl_wrapper.h"

static inline void SDL_SetErrorWrapper(const char *fmt) { SDL_SetError("%s", fmt); }
#endif // }
