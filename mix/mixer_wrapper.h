#ifndef _GO_SDL_MIXER_WRAPPER_H // if {
#define _GO_SDL_MIXER_WRAPPER_H

#if defined(__WIN32) // if {
#include <SDL2/SDL_mixer.h>
#else // } else {
#include <SDL_mixer.h>
#endif // }

// The internal format for a music chunk interpreted via mikmod
typedef struct _Mix_Music Mix_Music;

#if defined(SDL_SUPPORTED_MIDI_BACKENDS) // if {

typedef int (*Mix_EachSoundFontCallback)(const char *, void *);

int callEachSoundFont(char *str, void *udata);
#endif // }

typedef void (*Mix_SetPostMixCallback)(void *udata, Uint8 *stream, int len);
typedef void (*Mix_HookMusicCallback)(void *udata, Uint8 *stream, int len);
typedef void (*Mix_HookMusicFinishedCallback)(void);
typedef void (*Mix_ChannelFinishedCallback)(int channel);
typedef void (*Mix_RegisterEffectCallback)(int chan, void *stream, int len, void *udata);
typedef void (*Mix_RegisterEffectDoneCallback)(int chan, void *udata);

void callPostMixFunction(void *udata, Uint8 *stream, int length);
void callHookMusic(void *udata, Uint8 *stream, int length);
void callHookMusicFinished();
void callChannelFinished(int channel);
void callEffectFunc(int channel, void *stream, int len, void *udata);
void callEffectDone(int channel, void *udata);

static inline Uint32 Mix_GetChunkTimeMilliseconds(Mix_Chunk *chunk) {
    Uint32 points = 0;
    Uint32 frames = 0;
    int freq = 0;
    Uint16 fmt = 0;
    int chans = 0;
    if (!Mix_QuerySpec(&freq, &fmt, &chans)) return 0;
    points = (chunk->alen / ((fmt & 0xFF) / 8));
    frames = (points / chans);
    return (frames * 1000) / freq;
}

#if (SDL_MIXER_MAJOR_VERSION == 2) && (SDL_MIXER_MINOR_VERSION == 0) && (SDL_MIXER_PATCHLEVEL < 2) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("Mix_OpenAudioDevice is not supported before SDL_mixer 2.0.2")
#endif // }

static inline int Mix_OpenAudioDevice(int frequency, Uint16 format, int channels, int chunksize, const char *device,
                                      int allowed_changes) {
    return -1;
}
#endif // }
#endif // }
