#ifndef _GO_SDL_IMAGE_WRAPPER_H // if {
#define _GO_SDL_IMAGE_WRAPPER_H

#if defined(__WIN32) // if {
#include <SDL2/SDL_image.h>
#include <stdlib.h>
#else // } else {
#include <SDL_image.h>
#endif // }
#endif // }
