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

type Mix_EachSoundFontCallback = func(string, any) int

type Mix_SoundFontWatcher struct {
	callback Mix_EachSoundFontCallback
	data     any
}

var Mix_SoundFontOK Mix_SoundFontWatcher

//export callEachSoundFont
func callEachSoundFont(cstr *cChar, _ unsafe.Pointer) cInt {
	var ret int
	if Mix_SoundFontOK.callback != nil {
		ret = Mix_SoundFontOK.callback(SDL_GoString(cstr), Mix_SoundFontOK.data)
	}
	return cInt(ret)
}

func Mix_EachSoundFont(callback Mix_EachSoundFontCallback, data any) int {
	Mix_SoundFontOK.callback = callback
	Mix_SoundFontOK.data = data
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
