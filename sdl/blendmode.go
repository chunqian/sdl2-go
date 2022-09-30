package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"

func SDL_ComposeCustomBlendMode(srcColorFactor, dstColorFactor SDL_BlendFactor, colorOperation SDL_BlendOperation,
	srcAlphaFactor, dstAlphaFactor SDL_BlendFactor, alphaOperation SDL_BlendOperation) SDL_BlendMode {
	cSrcColorFactor := cSDL_BlendFactor(srcColorFactor)
	cDstColorFactor := cSDL_BlendFactor(dstColorFactor)
	cColorOperation := cSDL_BlendOperation(colorOperation)
	cSrcAlphaFactor := cSDL_BlendFactor(srcAlphaFactor)
	cDstAlphaFactor := cSDL_BlendFactor(dstAlphaFactor)
	cAlphaOperation := cSDL_BlendOperation(alphaOperation)
	return SDL_BlendMode(C.SDL_ComposeCustomBlendMode(cSrcColorFactor, cDstColorFactor, cColorOperation,
		cSrcAlphaFactor, cDstAlphaFactor, cAlphaOperation))
}
