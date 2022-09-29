#ifndef _GO_SDL_GFX_WRAPPER_H // if {
#define _GO_SDL_GFX_WRAPPER_H

#if defined(__WIN32) // if {
#include <SDL2/SDL2_framerate.h>
#include <SDL2/SDL2_gfxPrimitives.h>
#include <SDL2/SDL2_imageFilter.h>
#include <SDL2/SDL2_rotozoom.h>
#include <stdlib.h>
#else // } else {
#include <SDL2_framerate.h>
#include <SDL2_gfxPrimitives.h>
#include <SDL2_imageFilter.h>
#include <SDL2_rotozoom.h>
#endif // }
#endif // }
