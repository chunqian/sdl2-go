package sdl

/*
#include "error.h"
*/
import "C"
import "errors"

var ErrArrayIndexOutOfBounds = errors.New("Array Index Out Of Bounds Exception.")

func SDL_GetError() error {
	if cErrstr := C.SDL_GetError(); cErrstr != nil {
		errstr := SDL_GoString(cErrstr)
		// SDL_GetError returns "an empty string if there hasn't been an error message"
		if len(errstr) > 0 {
			return errors.New(errstr)
		}
	}
	return nil
}

func SDL_GetErrorMsg(maxlen int) error {
	cBuf := SDL_malloc(SDL_GetMemoryPool(), maxlen)
	defer SDL_free(SDL_GetMemoryPool(), cBuf)

	if cErrstr := C.SDL_GetErrorMsg((*cChar)(cBuf), cInt(maxlen)); cErrstr != nil {
		errstr := SDL_GoString(cErrstr)
		// SDL_GetErrorMsg returns "an empty string if there hasn't been an error message"
		if len(errstr) > 0 {
			return errors.New(errstr)
		}
	}
	return nil
}

func SDL_SetError(err error) {
	cStr := SDL_CreateCString(SDL_GetMemoryPool(), err.Error())
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cStr)

	C.SDL_SetErrorWrapper(cStr.(*cChar))
}

func SDL_ClearError() {
	C.SDL_ClearError()
}

func SDL_Error(code SDL_errorcode) {
	C.SDL_Error(cSDL_errorcode(code))
}

func SDL_OutOfMemory() {
	SDL_Error(SDL_ENOMEM)
}

func SDL_Unsupported() {
	SDL_Error(SDL_UNSUPPORTED)
}

func SDL_InvalidParamError(param string) {
	SDL_SetError(errors.New(param))
}
