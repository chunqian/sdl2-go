#ifndef _GO_SDL_MESSAGEBOX_H // if {
#define _GO_SDL_MESSAGEBOX_H

#include "sdl_wrapper.h"

static inline Sint32 ShowMessageBox(SDL_MessageBoxData data) {
    Sint32 buttonid;
    SDL_ShowMessageBox(&data, &buttonid);
    return buttonid;
}
#endif // }
