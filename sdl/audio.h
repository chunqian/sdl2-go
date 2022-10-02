#ifndef _GO_SDL_AUDIO_H // if {
#define _GO_SDL_AUDIO_H

#include "sdl_wrapper.h"

typedef void (*SDL_AudioCallback)(void *userdata, Uint8 *stream, int len);

void callAudio(void *userdata, Uint8 *stream, int len);

#if !(SDL_VERSION_ATLEAST(2, 24, 0)) // if {
#if defined(WARN_OUTDATED)
#pragma message("SDL_GetDefaultAudioInfo is not supported before SDL 2.24.0")
#endif // }

static inline int SDL_GetDefaultAudioInfo(char **name, SDL_AudioSpec *spec, int iscapture) { return -1; }

#endif // }
#endif // }
