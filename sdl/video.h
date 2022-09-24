#ifndef _GO_SDL_VIDEO_H // if {
#define _GO_SDL_VIDEO_H

#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2, 24, 0)) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_GetPointDisplayIndex is not supported before SDL 2.24.0")
#pragma message("SDL_GetRectDisplayIndex is not supported before SDL 2.24.0")
#endif // }

#define SDL_GL_FLOATBUFFERS (27)

static inline int SDL_GetPointDisplayIndex(const SDL_Point *point) { return -1; }

static inline int SDL_GetRectDisplayIndex(const SDL_Rect *rect) { return -1; }

#endif // }
#endif // }
