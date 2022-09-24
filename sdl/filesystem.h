#ifndef _GO_SDL_FILESYSTEM_H // if {
#define _GO_SDL_FILESYSTEM_H

#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2, 0, 1)) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_GetBasePath is not supported before SDL 2.0.1")
#endif // }

static inline char* SDL_GetBasePath() { return NULL; }

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_GetPrefPath is not supported before SDL 2.0.1")
#endif // }

static inline char* SDL_GetPrefPath(const char* org, const char* app) { return NULL; }
#endif // }
#endif // }
