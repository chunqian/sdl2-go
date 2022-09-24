#ifndef _GO_SDL_EVENTS_H // if {
#define _GO_SDL_EVENTS_H

#if defined(_WIN32) // if {
#include <SDL2/SDL_events.h>
#else // } else {
#include <SDL_events.h>
#endif // }

int SDL_EventFilterWrapper(void* userdata, SDL_Event* event);
int SDL_EventWatchWrapper(void* userdata, SDL_Event* event);
int SDL_FilterEventsWrapper(void* userdata, SDL_Event* event);

#include "sdl_wrapper.h"

#endif // }
