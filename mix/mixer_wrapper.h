#ifndef _GO_SDL_MIXER_WRAPPER_H // if {
#define _GO_SDL_MIXER_WRAPPER_H

#if defined(__WIN32) // if {
#include <SDL2/SDL_mixer.h>
#else // } else {
#include <SDL_mixer.h>
#endif // }

#if defined(SDL_SUPPORTED_MIDI_BACKENDS) // if {
extern int callEachSoundFont(char* str, void* udata);
#endif // }

extern void callPostMixFunction(void *udata, Uint8 *stream, int length);
extern void callHookMusic(void *udata, Uint8 *stream, int length);
extern void callHookMusicFinished();
extern void callChannelFinished(int channel);
extern void callEffectFunc(int channel, void *stream, int len, void *udata);
extern void callEffectDone(int channel, void *udata);

static inline Uint32 getChunkTimeMilliseconds(Mix_Chunk *chunk) {
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
