package sdl

/*
#include "log.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// typedef
type SDL_LogCategory = int32
type SDL_LogPriority = int32
type SDL_LogOutputFunction = func(unsafe.Pointer, int32, int32, *int8)

// enum
const (
	SDL_LOG_CATEGORY_APPLICATION SDL_LogCategory = 0
	SDL_LOG_CATEGORY_ERROR       SDL_LogCategory = 1
	SDL_LOG_CATEGORY_ASSERT      SDL_LogCategory = 2
	SDL_LOG_CATEGORY_SYSTEM      SDL_LogCategory = 3
	SDL_LOG_CATEGORY_AUDIO       SDL_LogCategory = 4
	SDL_LOG_CATEGORY_VIDEO       SDL_LogCategory = 5
	SDL_LOG_CATEGORY_RENDER      SDL_LogCategory = 6
	SDL_LOG_CATEGORY_INPUT       SDL_LogCategory = 7
	SDL_LOG_CATEGORY_TEST        SDL_LogCategory = 8
	SDL_LOG_CATEGORY_RESERVED1   SDL_LogCategory = 9
	SDL_LOG_CATEGORY_RESERVED2   SDL_LogCategory = 10
	SDL_LOG_CATEGORY_RESERVED3   SDL_LogCategory = 11
	SDL_LOG_CATEGORY_RESERVED4   SDL_LogCategory = 12
	SDL_LOG_CATEGORY_RESERVED5   SDL_LogCategory = 13
	SDL_LOG_CATEGORY_RESERVED6   SDL_LogCategory = 14
	SDL_LOG_CATEGORY_RESERVED7   SDL_LogCategory = 15
	SDL_LOG_CATEGORY_RESERVED8   SDL_LogCategory = 16
	SDL_LOG_CATEGORY_RESERVED9   SDL_LogCategory = 17
	SDL_LOG_CATEGORY_RESERVED10  SDL_LogCategory = 18
	SDL_LOG_CATEGORY_CUSTOM      SDL_LogCategory = 19
)

const (
	SDL_LOG_PRIORITY_VERBOSE  SDL_LogPriority = 1
	SDL_LOG_PRIORITY_DEBUG    SDL_LogPriority = 2
	SDL_LOG_PRIORITY_INFO     SDL_LogPriority = 3
	SDL_LOG_PRIORITY_WARN     SDL_LogPriority = 4
	SDL_LOG_PRIORITY_ERROR    SDL_LogPriority = 5
	SDL_LOG_PRIORITY_CRITICAL SDL_LogPriority = 6
	SDL_NUM_LOG_PRIORITIES    SDL_LogPriority = 7
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

	cStr := createCString(SDL_GetMemoryPool(), str)
	defer destroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_Log(cStr)
}

func SDL_LogVerbose(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := createCString(SDL_GetMemoryPool(), str)
	defer destroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogVerbose(cInt(category), cStr)
}

func SDL_LogDebug(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := createCString(SDL_GetMemoryPool(), str)
	defer destroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogDebug(cInt(category), cStr)
}

func SDL_LogInfo(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := createCString(SDL_GetMemoryPool(), str)
	defer destroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogInfo(cInt(category), cStr)
}

func SDL_LogWarn(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := createCString(SDL_GetMemoryPool(), str)
	defer destroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogWarn(cInt(category), cStr)
}

func SDL_LogError(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := createCString(SDL_GetMemoryPool(), str)
	defer destroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogError(cInt(category), cStr)
}

func SDL_LogCritical(category SDL_LogCategory, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := createCString(SDL_GetMemoryPool(), str)
	defer destroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogCritical(cInt(category), cStr)
}

func SDL_LogMessage(category SDL_LogCategory, pri SDL_LogPriority, str string, args ...any) {
	str = fmt.Sprintf(str, args...)

	cStr := createCString(SDL_GetMemoryPool(), str)
	defer destroyCString(SDL_GetMemoryPool(), cStr)

	C._SDL_LogMessage(cInt(category), cLogPriority(pri), cStr)
}

type LogOutputFunction func(data any, category SDL_LogCategory, pri SDL_LogPriority, message string)

type logOutputFunctionCtx struct {
	f LogOutputFunction
	d any
}

//export logOutputFunction
func logOutputFunction(data unsafe.Pointer, category cInt, pri cLogPriority, message *cChar) {
	ctx := (*logOutputFunctionCtx)(data)

	ctx.f(ctx.d, SDL_LogCategory(category), SDL_LogPriority(pri), createGoString(message))
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
