package sdl

/*
#include "error.h"
*/
import "C"
import "errors"

// typedef
type SDL_errorcode = int32

// enum
const (
	SDL_ENOMEM      SDL_errorcode = 0
	SDL_EFREAD      SDL_errorcode = 1
	SDL_EFWRITE     SDL_errorcode = 2
	SDL_EFSEEK      SDL_errorcode = 3
	SDL_UNSUPPORTED SDL_errorcode = 4
	SDL_LASTERROR   SDL_errorcode = 5
)

func SDL_GetError() error {
	if cErrstr := C.SDL_GetError(); cErrstr != nil {
		errstr := createGoString(cErrstr)
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
		errstr := createGoString(cErrstr)
		// SDL_GetErrorMsg returns "an empty string if there hasn't been an error message"
		if len(errstr) > 0 {
			return errors.New(errstr)
		}
	}
	return nil
}

func SDL_SetError(err error) {
	cStr := createCString(SDL_GetMemoryPool(), err.Error())
	defer destroyCString(SDL_GetMemoryPool(), cStr)

	C.SDL_SetErrorWrapper(cStr)
}

func SDL_ClearError() {
	C.SDL_ClearError()
}

func SDL_Error(code SDL_errorcode) {
	C.SDL_Error(cErrorcode(code))
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
