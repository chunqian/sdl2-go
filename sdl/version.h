#ifndef _GO_SDL_VERSION_H // if {
#define _GO_SDL_VERSION_H

#include "sdl_wrapper.h"

#if SDL_VERSION_ATLEAST(2, 0, 16) // if {

static inline int GetRevisionNumber(void) { return 0; }
#else // } else {
static inline int GetRevisionNumber(void) { return SDL_GetRevisionNumber(); }

#endif // }
#endif // }
