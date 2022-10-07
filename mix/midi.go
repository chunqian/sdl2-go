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

type Mix_EachSoundFontCallback = func(path string, userdata any) int

type Mix_SoundFontWatcher struct {
	callback Mix_EachSoundFontCallback
	data     any
}

var Mix_SoundFontOK Mix_SoundFontWatcher

//export callEachSoundFont
func callEachSoundFont(cStr *cChar, _ unsafe.Pointer) cInt {
	var ret int
	if Mix_SoundFontOK.callback != nil {
		ret = Mix_SoundFontOK.callback(createGoString(cStr), Mix_SoundFontOK.data)
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
	cPaths := createCString(SDL_GetMemoryPool(), paths)
	defer destroyCString(SDL_GetMemoryPool(), cPaths) // memory free

	cRet := C.Mix_SetSoundFonts(cPaths)
	return int(cRet) == 0
}

func Mix_GetSoundFonts() string {
	cStr := C.Mix_GetSoundFonts()
	return createGoString(cStr)
}
