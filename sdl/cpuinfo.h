#ifndef _GO_SDL_CPUINFO_H // if {
#define _GO_SDL_CPUINFO_H

#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2, 24, 0)) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_HasLSX is not supported before SDL 2.24.0")
#pragma message("SDL_HasLASX is not supported before SDL 2.24.0")
#endif // }

static inline SDL_bool SDL_HasLSX() { return SDL_FALSE; }
static inline SDL_bool SDL_HasLASX() { return SDL_FALSE; }
#endif // }
#endif // }
