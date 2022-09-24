#ifndef _GO_SDL_RWOPS_H // if {
#define _GO_SDL_RWOPS_H

#include "sdl_wrapper.h"
static Sint64 RWsize(SDL_RWops* ctx) { return ctx->size(ctx); }

static Sint64 RWseek(SDL_RWops* ctx, Sint64 offset, int whence) { return ctx->seek(ctx, offset, whence); }

static size_t RWread(SDL_RWops* ctx, void* ptr, size_t size, size_t maxnum) {
    return ctx->read(ctx, ptr, size, maxnum);
}

static size_t RWwrite(SDL_RWops* ctx, void* ptr, size_t size, size_t num) { return ctx->write(ctx, ptr, size, num); }

static int RWclose(SDL_RWops* ctx) { return ctx->close(ctx); }

#if !(SDL_VERSION_ATLEAST(2, 0, 6)) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_LoadFile_RW is not supported before SDL 2.0.6")
#endif // }

static void* SDL_LoadFile_RW(SDL_RWops* src, size_t* datasize, int freesrc) { return 0; }
#endif // }
#endif // }
