#ifndef _GO_SDL_GAMECONTROLLER_H // if {
#define _GO_SDL_GAMECONTROLLER_H

#include "sdl_wrapper.h"

#define SDL_GAMECONTROLLER_BUTTON_BIND_VALUE_SIZE sizeof(((SDL_GameControllerButtonBind*)0)->value)

#if !(SDL_VERSION_ATLEAST(2, 24, 0)) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_GameControllerPathForIndex is not supported before SDL 2.24.0")
#pragma message("SDL_GameControllerPath is not supported before SDL 2.24.0")
#pragma message("SDL_GameControllerGetFirmwareVersion is not supported before SDL 2.24.0")
#endif // }

static const char *SDL_GameControllerPathForIndex(int device_index) { return NULL; }

static const char *SDL_GameControllerPath(SDL_GameController *gamecontroller) { return NULL; }

static Uint16 SDL_GameControllerGetFirmwareVersion(SDL_GameController *gamecontroller) { return 0; }

#endif // }
#endif // }
