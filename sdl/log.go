package sdl

/*
#include "log.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func SDL_LogSetAllPriority(p SDL_LogPriority) {
	C.SDL_LogSetAllPriority(cLogPriority(p))
}

func SDL_LogSetPriority(category SDL_LogCategory, p SDL_LogPriority) {
	C.SDL_LogSetPriority(cInt(category), cLogPriority(p))
}

func SDL_LogGetPriority(category SDL_LogCategory) SDL_LogPriority {
	cP := C.SDL_LogGetPriority(cInt(category))
	return SDL_LogPriority(cP)
}

func SDL_LogResetPriorities() {
	C.SDL_LogResetPriorities()
}

func Log(str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := SDL_CreateCString(SDL_GetMemoryPool(), str)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_Log(cStr.(*cChar))
}

func SDL_LogVerbose(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := SDL_CreateCString(SDL_GetMemoryPool(), str)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogVerbose(cInt(category), cStr.(*cChar))
}

func SDL_LogDebug(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := SDL_CreateCString(SDL_GetMemoryPool(), str)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogDebug(cInt(category), cStr.(*cChar))
}

func SDL_LogInfo(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := SDL_CreateCString(SDL_GetMemoryPool(), str)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogInfo(cInt(category), cStr.(*cChar))
}

func SDL_LogWarn(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := SDL_CreateCString(SDL_GetMemoryPool(), str)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogWarn(cInt(category), cStr.(*cChar))
}

func SDL_LogError(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := SDL_CreateCString(SDL_GetMemoryPool(), str)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogError(cInt(category), cStr.(*cChar))
}

func SDL_LogCritical(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := SDL_CreateCString(SDL_GetMemoryPool(), str)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogCritical(cInt(category), cStr.(*cChar))
}

func SDL_LogMessage(category SDL_LogCategory, pri SDL_LogPriority, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := SDL_CreateCString(SDL_GetMemoryPool(), str)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogMessage(cInt(category), cLogPriority(pri), cStr.(*cChar))
}

type LogOutputFunction func(data any, category SDL_LogCategory, pri SDL_LogPriority, message string)

type logOutputFunctionCtx struct {
	f LogOutputFunction
	d any
}

//export logOutputFunction
func logOutputFunction(data unsafe.Pointer, category cInt, pri cLogPriority, message *cChar) {
	ctx := (*logOutputFunctionCtx)(data)

	ctx.f(ctx.d, SDL_LogCategory(category), SDL_LogPriority(pri), SDL_GoString(message))
}

var (
	logOutputFunctionCache LogOutputFunction
	logOutputDataCache     any
)

func SDL_LogGetOutputFunction() (LogOutputFunction, any) {
	return logOutputFunctionCache, logOutputDataCache
}

func SDL_LogSetOutputFunction(f LogOutputFunction, data any) {
	ctx := &logOutputFunctionCtx{
		f: f,
		d: data,
	}

	C.LogSetOutputFunction(unsafe.Pointer(ctx))

	logOutputFunctionCache = f
	logOutputDataCache = data
}
