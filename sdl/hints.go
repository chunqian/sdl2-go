package sdl

/*
#include "hints.h"
*/
import "C"

type HintCallback func(data interface{}, name, oldValue, newValue string)

type HintCallbackAndData struct {
	callback HintCallback // the function to call when the hint value changes
	data     interface{}  // data to pass to the callback function
}

var hintCallbacks = make(map[string]HintCallbackAndData)

func SDL_SetHintWithPriority(name, value string, hp SDL_HintPriority) bool {
	cName := SDL_CreateCString(SDL_GetMemoryPool(), name)
	cValue := SDL_CreateCString(SDL_GetMemoryPool(), value)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cName)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cValue)

	cRet := C.SDL_SetHintWithPriority(cName.(*cChar), cValue.(*cChar), cHintPriority(hp))
	return int(cRet) > 0
}

func SDL_SetHint(name, value string) bool {
	cName := SDL_CreateCString(SDL_GetMemoryPool(), name)
	cValue := SDL_CreateCString(SDL_GetMemoryPool(), value)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cName)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cValue)

	cRet := C.SDL_SetHint(cName.(*cChar), cValue.(*cChar))
	return int(cRet) > 0
}

func SDL_ResetHint(name string) bool {
	cName := SDL_CreateCString(SDL_GetMemoryPool(), name)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cName)

	cRet := C.SDL_ResetHint(cName.(*cChar))
	return int(cRet) == 1
}

func SDL_GetHint(name string) string {
	cName := SDL_CreateCString(SDL_GetMemoryPool(), name)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cName)

	return SDL_GoString(C.SDL_GetHint(cName.(*cChar)))
}

func SDL_ClearHints() {
	C.SDL_ClearHints()
}

func SDL_AddHintCallback(name string, fn HintCallback, data interface{}) {
	cName := SDL_CreateCString(SDL_GetMemoryPool(), name)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cName)

	hintCallbacks[name] = HintCallbackAndData{
		callback: fn,
		data:     data,
	}
	C.addHintCallback(cName.(*cChar))
}

func SDL_DelHintCallback(name string) {
	cName := SDL_CreateCString(SDL_GetMemoryPool(), name)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cName)

	delete(hintCallbacks, name)
	C.delHintCallback(cName.(*cChar))
}

//export goHintCallback
func goHintCallback(cName, cOldValue, cNewValue *cChar) {
	name := SDL_GoString(cName)
	oldValue := SDL_GoString(cOldValue)
	newValue := SDL_GoString(cNewValue)
	if cb, ok := hintCallbacks[name]; ok {
		cb.callback(cb.data, name, oldValue, newValue)
	}
}
