package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"

// typedef
type SDL_BlendMode = int32
type SDL_BlendOperation = int32
type SDL_BlendFactor = int32

// enum
const (
	SDL_BLENDMODE_NONE    SDL_BlendMode = 0
	SDL_BLENDMODE_BLEND   SDL_BlendMode = 1
	SDL_BLENDMODE_ADD     SDL_BlendMode = 2
	SDL_BLENDMODE_MOD     SDL_BlendMode = 4
	SDL_BLENDMODE_MUL     SDL_BlendMode = 8
	SDL_BLENDMODE_INVALID SDL_BlendMode = 0x7FFFFFFF
)

const (
	SDL_BLENDOPERATION_ADD          SDL_BlendOperation = 1
	SDL_BLENDOPERATION_SUBTRACT     SDL_BlendOperation = 2
	SDL_BLENDOPERATION_REV_SUBTRACT SDL_BlendOperation = 3
	SDL_BLENDOPERATION_MINIMUM      SDL_BlendOperation = 4
	SDL_BLENDOPERATION_MAXIMUM      SDL_BlendOperation = 5
)

const (
	SDL_BLENDFACTOR_ZERO                SDL_BlendFactor = 1
	SDL_BLENDFACTOR_ONE                 SDL_BlendFactor = 2
	SDL_BLENDFACTOR_SRC_COLOR           SDL_BlendFactor = 3
	SDL_BLENDFACTOR_ONE_MINUS_SRC_COLOR SDL_BlendFactor = 4
	SDL_BLENDFACTOR_SRC_ALPHA           SDL_BlendFactor = 5
	SDL_BLENDFACTOR_ONE_MINUS_SRC_ALPHA SDL_BlendFactor = 6
	SDL_BLENDFACTOR_DST_COLOR           SDL_BlendFactor = 7
	SDL_BLENDFACTOR_ONE_MINUS_DST_COLOR SDL_BlendFactor = 8
	SDL_BLENDFACTOR_DST_ALPHA           SDL_BlendFactor = 9
	SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA SDL_BlendFactor = 10
)

func SDL_ComposeCustomBlendMode(srcColorFactor, dstColorFactor SDL_BlendFactor, colorOperation SDL_BlendOperation,
	srcAlphaFactor, dstAlphaFactor SDL_BlendFactor, alphaOperation SDL_BlendOperation) SDL_BlendMode {
	cSrcColorFactor := cBlendFactor(srcColorFactor)
	cDstColorFactor := cBlendFactor(dstColorFactor)
	cColorOperation := cBlendOperation(colorOperation)
	cSrcAlphaFactor := cBlendFactor(srcAlphaFactor)
	cDstAlphaFactor := cBlendFactor(dstAlphaFactor)
	cAlphaOperation := cBlendOperation(alphaOperation)
	return SDL_BlendMode(C.SDL_ComposeCustomBlendMode(cSrcColorFactor, cDstColorFactor, cColorOperation,
		cSrcAlphaFactor, cDstAlphaFactor, cAlphaOperation))
}
