package sdl

/*
#include "blendmode.h"
*/
import "C"

// ComposeCustomBlendMode creates a custom blend mode, which may or may not be supported by a given renderer
// The result of the blend mode operation will be:
//
// dstRGB = colorOperation(srcRGB * srcColorFactor, dstRGB * dstColorFactor);
// dstA = alphaOperation(srcA * srcAlphaFactor, dstA * dstAlphaFactor);
func SDL_ComposeCustomBlendMode(srcColorFactor, dstColorFactor SDL_BlendFactor, colorOperation SDL_BlendOperation,
	srcAlphaFactor, dstAlphaFactor SDL_BlendFactor, alphaOperation SDL_BlendOperation) SDL_BlendMode {
	cSrcColorFactor := cSDL_BlendFactor(srcColorFactor)
	cDstColorFactor := cSDL_BlendFactor(dstColorFactor)
	cColorOperation := cSDL_BlendOperation(colorOperation)
	cSrcAlphaFactor := cSDL_BlendFactor(srcAlphaFactor)
	cDstAlphaFactor := cSDL_BlendFactor(dstAlphaFactor)
	cAlphaOperation := cSDL_BlendOperation(alphaOperation)
	return SDL_BlendMode(C.SDL_ComposeCustomBlendMode(cSrcColorFactor, cDstColorFactor, cColorOperation, cSrcAlphaFactor, cDstAlphaFactor, cAlphaOperation))
}
