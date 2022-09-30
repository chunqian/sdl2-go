#ifndef _GO_SDL_GUID_H // if {
#define _GO_SDL_GUID_H

#include "sdl_wrapper.h"

#if !SDL_VERSION_ATLEAST(2, 24, 0) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_GUIDToString is not supported before SDL 2.24.0")
#pragma message("SDL_GUIDFromString is not supported before SDL 2.24.0")
#endif // }

typedef struct {
    Uint8 data[16];
} SDL_GUID;

static inline void SDL_GUIDToString(SDL_GUID guid, char *pszGUID, int cbGUID) { return; }

static inline SDL_GUID SDL_GUIDFromString(const char *pchGUID) {
    SDL_GUID guid;
    return guid;
}
#endif // }
#endif // }
