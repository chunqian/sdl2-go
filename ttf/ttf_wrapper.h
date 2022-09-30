#ifndef _GO_SDL_TTF_WRAPPER_H // if {
#define _GO_SDL_TTF_WRAPPER_H

#if defined(__WIN32) // if {
#include <SDL2/SDL_ttf.h>
#include <stdlib.h>
#else // } else {
#include <SDL_ttf.h>
#endif // }

typedef struct _TTF_Font TTF_Font;

static inline void Do_TTF_SetError(const char *str) { TTF_SetError("%s", str); }

#endif // }
