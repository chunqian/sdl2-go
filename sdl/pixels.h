#ifndef _GO_SDL_PIXELS_H // if {
#define _GO_SDL_PIXELS_H

#include "sdl_wrapper.h"

static int SDL_BytesPerPixel(Uint32 format) { return SDL_BYTESPERPIXEL(format); }
static int SDL_BitsPerPixel(Uint32 format) { return SDL_BITSPERPIXEL(format); }
#endif // }
