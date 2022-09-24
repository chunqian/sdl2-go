package sdl

/*
#include "error.h"
*/
import "C"
import "errors"

// GetError returns the last error that occurred, or an empty string if there hasn't been an error message set since the last call to ClearError().
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

// GetErrorMsg returns the last error message that was set for the current thread.
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

// SetError set the SDL error message.
func SDL_SetError(err error) {
	cstr := SDL_CreateCString(SDL_GetMemoryPool(), err.Error())
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cstr)

	C.SDL_SetErrorWrapper(cstr)
}

// ClearError clears any previous error message.
func SDL_ClearError() {
	C.SDL_ClearError()
}

// Error sets the SDL error message to the specified error code.
func SDL_Error(code SDL_errorcode) {
	C.SDL_Error(cSDL_errorcode(code))
}

// OutOfMemory sets SDL error message to ENOMEM (out of memory).
func SDL_OutOfMemory() {
	SDL_Error(SDL_ENOMEM)
}

// Unsupported sets SDL error message to UNSUPPORTED (that operation is not supported).
func SDL_Unsupported() {
	SDL_Error(SDL_UNSUPPORTED)
}

func SDL_InvalidParamError(param string) {
	SDL_SetError(errors.New(param))
}
