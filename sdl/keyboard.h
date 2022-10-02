#ifndef _GO_SDL_KEYBOARD_H // if {
#define _GO_SDL_KEYBOARD_H

#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2, 24, 0)) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_ResetKeyboard is not supported before SDL 2.24.0")
#endif // }

static inline void SDL_ResetKeyboard(void) { return; }
#endif // }

#if !(SDL_VERSION_ATLEAST(2, 0, 22)) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_IsTextInputShown is not supported before SDL 2.0.22")
#pragma message("SDL_ClearComposition is not supported before SDL 2.0.22")
#endif // }

static inline SDL_bool SDL_IsTextInputShown(void) { return SDL_FALSE; }

static inline void SDL_ClearComposition(void) {}
#endif // }
#endif // }
