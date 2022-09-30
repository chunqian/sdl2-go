//go:build !nomidi

package mix

/*
#define SDL_SUPPORTED_MIDI_BACKENDS
#include "mixer_wrapper.h"
*/
import "C"
import (
	"unsafe"

	. "github.com/chunqian/sdl2-go/sdl"
)

type SDL_SoundFontCallback = func(string, any) int

type SDL_SoundFontWatcher struct {
	callback SDL_SoundFontCallback
	data     any
}

var SDL_SoundFontOK SDL_SoundFontWatcher

//export callEachSoundFont
func callEachSoundFont(cstr *cChar, _ unsafe.Pointer) cInt {
	var ret int
	if SDL_SoundFontOK.callback != nil {
		ret = SDL_SoundFontOK.callback(SDL_GoString(cstr), SDL_SoundFontOK.data)
	}
	return cInt(ret)
}

func Mix_EachSoundFont(callback SDL_SoundFontCallback, data any) int {
	SDL_SoundFontOK.callback = callback
	SDL_SoundFontOK.data = data
	cRet := C.Mix_EachSoundFont(C.Mix_EachSoundFontCallback(C.callEachSoundFont), nil)
	return int(cRet)
}

func Mix_SetSoundFonts(paths string) bool {
	cPaths := SDL_CreateCString(SDL_GetMemoryPool(), paths)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cPaths) // memory free

	cRet := C.Mix_SetSoundFonts(cPaths.(*cChar))
	return int(cRet) == 0
}

func Mix_GetSoundFonts() string {
	cStr := C.Mix_GetSoundFonts()
	return SDL_GoString(cStr)
}
