#ifndef _GO_SDL_RENDER_H // if {
#define _GO_SDL_RENDER_H

#include "sdl_wrapper.h"

#if SDL_COMPILEDVERSION == SDL_VERSIONNUM(2, 0, 18) // if {
static inline int RenderGeometryRaw(SDL_Renderer* renderer, SDL_Texture* texture, const float* xy, int xy_stride,
                                    const SDL_Color* color, int color_stride, const float* uv, int uv_stride,
                                    int num_vertices, const void* indices, int num_indices, int size_indices) {
    return SDL_RenderGeometryRaw(renderer, texture, xy, xy_stride, (int*)color, color_stride, uv, uv_stride,
                                 num_vertices, indices, num_indices, size_indices);
}
#else // } else {
static inline int RenderGeometryRaw(SDL_Renderer *renderer, SDL_Texture *texture, const float *xy, int xy_stride,
                                    const SDL_Color *color, int color_stride, const float *uv, int uv_stride,
                                    int num_vertices, const void *indices, int num_indices, int size_indices) {
    return SDL_RenderGeometryRaw(renderer, texture, xy, xy_stride, color, color_stride, uv, uv_stride, num_vertices,
                                 indices, num_indices, size_indices);
}
#endif // }

// WORKAROUND: This prevents audio from seemingly going corrupt when drawing outside the screen bounding box?
// It does that by allocating SDL_Rect in the C context instead of Go context.
static inline int RenderCopy(SDL_Renderer* renderer, SDL_Texture* texture, SDL_Rect* src, int dst_x, int dst_y,
                             int dst_w, int dst_h) {
    SDL_Rect dst = {dst_x, dst_y, dst_w, dst_h};
    return SDL_RenderCopy(renderer, texture, src, &dst);
}
#endif // }
