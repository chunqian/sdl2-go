package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"

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
