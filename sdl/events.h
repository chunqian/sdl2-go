#ifndef _GO_SDL_EVENTS_H // if {
#define _GO_SDL_EVENTS_H

#if defined(_WIN32) // if {
#include <SDL2/SDL_events.h>
#else // } else {
#include <SDL_events.h>
#endif // }

typedef int (*SDL_EventFilterCallback)(void* userdata, SDL_Event* event);

int callEventFilter(void* userdata, SDL_Event* event);
int callEventWatch(void* userdata, SDL_Event* event);
int callFilterEvents(void* userdata, SDL_Event* event);

#include "sdl_wrapper.h"

#endif // }
