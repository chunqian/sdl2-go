#ifndef _GO_SDL_LOG_H // if {
#define _GO_SDL_LOG_H

#include "sdl_wrapper.h"

void logOutputFunction(void *userdata, int category, SDL_LogPriority priority, char *message);

static void LogSetOutputFunction(void *data) {
    SDL_LogSetOutputFunction((SDL_LogOutputFunction)logOutputFunction, data);
}

static inline void _SDL_Log(const char *fmt) { SDL_Log("%s", fmt); }

static inline void _SDL_LogVerbose(int category, const char *fmt) { SDL_LogVerbose(category, "%s", fmt); }

static inline void _SDL_LogDebug(int category, const char *fmt) { SDL_LogDebug(category, "%s", fmt); }

static inline void _SDL_LogInfo(int category, const char *fmt) { SDL_LogInfo(category, "%s", fmt); }

static inline void _SDL_LogWarn(int category, const char *fmt) { SDL_LogWarn(category, "%s", fmt); }

static inline void _SDL_LogError(int category, const char *fmt) { SDL_LogError(category, "%s", fmt); }

static inline void _SDL_LogCritical(int category, const char *fmt) { SDL_LogCritical(category, "%s", fmt); }

static inline void _SDL_LogMessage(int category, SDL_LogPriority priority, const char *fmt) {
    SDL_LogCritical(category, "%s", fmt);
}

#if !(SDL_VERSION_ATLEAST(2, 0, 12)) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_LogCategory is not supported before SDL 2.0.12")
#endif

typedef int SDL_LogCategory;

#endif
#endif // }
