package sdl

/*
#include "audio.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func cSDL_AudioSpec(as *SDL_AudioSpec) *C.SDL_AudioSpec {
	return (*C.SDL_AudioSpec)(unsafe.Pointer(as))
}

func cSDL_AudioCVT(cvt *SDL_AudioCVT) *C.SDL_AudioCVT {
	return (*C.SDL_AudioCVT)(unsafe.Pointer(cvt))
}

func cSDL_AudioStream(stream *SDL_AudioStream) *C.SDL_AudioStream {
	return (*C.SDL_AudioStream)(unsafe.Pointer(stream))
}

func (cvt *SDL_AudioCVT) BufAlloc(size uintptr) {
	cvt.Buf = (*uint8)(SDL_malloc(SDL_GetMemoryPool(), int(size)))
}

func (cvt *SDL_AudioCVT) BufFree() {
	SDL_free(SDL_GetMemoryPool(), unsafe.Pointer(cvt.Buf))
}

// BufAsSlice方法返回的切片, 禁止使用append扩容, 会导致内存泄露
func (cvt SDL_AudioCVT) BufAsSlice() []byte {
	var b []byte
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh.Len = int(cvt.LenCVT)
	sh.Cap = int(cvt.Len * cvt.LenMult)
	sh.Data = uintptr(unsafe.Pointer(cvt.Buf))
	return b
}

func SDL_GetNumAudioDrivers() int {
	return int(C.SDL_GetNumAudioDrivers())
}

func SDL_GetAudioDriver(index int) string {
	cStr := C.SDL_GetAudioDriver(cInt(index))
	return SDL_GoString(cStr)
}

func SDL_AudioInit(driverName string) int {
	cDriverName := SDL_CreateCString(SDL_GetMemoryPool(), driverName)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cDriverName) // memory free

	cRet := C.SDL_AudioInit(cDriverName.(*cChar))
	return int(cRet)
}

func SDL_AudioQuit() {
	C.SDL_AudioQuit()
}

func SDL_GetCurrentAudioDriver() string {
	cStr := C.SDL_GetCurrentAudioDriver()
	return SDL_GoString(cStr)
}

func SDL_OpenAudio(desired, obtained *SDL_AudioSpec) int {
	cRet := C.SDL_OpenAudio(cSDL_AudioSpec(desired), cSDL_AudioSpec(obtained))
	return int(cRet)
}

func SDL_GetNumAudioDevices(isCapture int32) int {
	cRet := C.SDL_GetNumAudioDevices(cInt(isCapture))
	return int(cRet)
}

func SDL_GetAudioDeviceName(index int32, isCapture int32) string {
	cStr := C.SDL_GetAudioDeviceName(cInt(index), cInt(isCapture))
	return SDL_GoString(cStr)
}

func SDL_OpenAudioDevice(device string, isCapture int32, desired, obtained *SDL_AudioSpec, allowedChanges int32) SDL_AudioDeviceID {
	cDevice := SDL_CreateCString(SDL_GetMemoryPool(), device)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cDevice) // memory free

	if device == "" {
		cDevice = nil
	}
	cId := C.SDL_OpenAudioDevice(cDevice.(*cChar), cInt(isCapture), cSDL_AudioSpec(desired), cSDL_AudioSpec(obtained), cInt(allowedChanges))
	return SDL_AudioDeviceID(cId)
}

func SDL_GetAudioStatus() SDL_AudioStatus {
	cRet := C.SDL_GetAudioStatus()
	return SDL_AudioStatus(cRet)
}

func SDL_GetAudioDeviceStatus(dev SDL_AudioDeviceID) SDL_AudioStatus {
	cRet := C.SDL_GetAudioDeviceStatus(cUint32(dev))
	return SDL_AudioStatus(cRet)
}

func SDL_PauseAudio(pauseOn int32) {
	C.SDL_PauseAudio(cInt(pauseOn))
}

func SDL_PauseAudioDevice(dev SDL_AudioDeviceID, pauseOn int32) {
	C.SDL_PauseAudioDevice(cUint32(dev), cInt(pauseOn))
}

func SDL_LoadWAV_RW(src *SDL_RWops, freeSrc int32, spec *SDL_AudioSpec, audioBuf *[]uint8, audioLen *uint32) *SDL_AudioSpec {
	cAudioBuf := (**cUint8)(unsafe.Pointer(audioBuf))
	cAudioLen := (*cUint32)(unsafe.Pointer(audioLen))
	cAudioSpec := C.SDL_LoadWAV_RW(cSDL_RWops(src), cInt(freeSrc), cSDL_AudioSpec(spec), cAudioBuf, cAudioLen)
	audioSpec := (*SDL_AudioSpec)(unsafe.Pointer(cAudioSpec))

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(audioBuf))
	sliceHeader.Len = int(*audioLen)
	sliceHeader.Cap = int(*audioLen)
	sliceHeader.Data = uintptr(unsafe.Pointer(audioBuf))
	return audioSpec
}

func SDL_LoadWAV(file string, spec *SDL_AudioSpec, audioBuf *[]uint8, audioLen *uint32) *SDL_AudioSpec {
	cFile := SDL_CreateCString(SDL_GetMemoryPool(), file)
	cRB := SDL_CreateCString(SDL_GetMemoryPool(), "rb")
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cFile) // memory free
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cRB)   // memory free

	cAudioBuf := (**cUint8)(unsafe.Pointer(audioBuf))
	cAudioLen := (*cUint32)(unsafe.Pointer(audioLen))
	cAudioSpec := C.SDL_LoadWAV_RW(C.SDL_RWFromFile(cFile.(*cChar), cRB.(*cChar)), 1, cSDL_AudioSpec(spec), cAudioBuf, cAudioLen)
	audioSpec := (*SDL_AudioSpec)(unsafe.Pointer(cAudioSpec))

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(audioBuf))
	sliceHeader.Len = int(*audioLen)
	sliceHeader.Cap = int(*audioLen)
	sliceHeader.Data = uintptr(unsafe.Pointer(audioBuf))
	return audioSpec
}

func SDL_FreeWAV(audioBuf []uint8) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&audioBuf))
	cAudioBuf := (*cUint8)(unsafe.Pointer(sliceHeader.Data))
	C.SDL_FreeWAV(cAudioBuf)
}

func SDL_BuildAudioCVT(cvt *SDL_AudioCVT, srcFormat SDL_AudioFormat, srcChannels uint8, srcRate int, dstFormat SDL_AudioFormat, dstChannels uint8, dstRate int) (converted int) {
	cRet := C.SDL_BuildAudioCVT(cSDL_AudioCVT(cvt), cUint16(srcFormat), cUint8(srcChannels), cInt(srcRate), cUint16(dstFormat), cUint8(dstChannels), cInt(dstRate))
	return int(cRet)
}

func SDL_ConvertAudio(cvt *SDL_AudioCVT) int {
	cRet := C.SDL_ConvertAudio(cSDL_AudioCVT(cvt))
	return int(cRet)
}

func SDL_QueueAudio(dev SDL_AudioDeviceID, data []byte, len uint32) int {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	cData := unsafe.Pointer(sliceHeader.Data)
	// cLen := cUint32(sliceHeader.Len)
	cRet := C.SDL_QueueAudio(cUint32(dev), cData, cUint32(len))
	return int(cRet)
}

func SDL_DequeueAudio(dev SDL_AudioDeviceID, data []byte, len uint32) uint32 {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	cData := unsafe.Pointer(sliceHeader.Data)
	// cLen := cUint32(sliceHeader.Len)
	cRet := C.SDL_DequeueAudio(cUint32(dev), cData, cUint32(len))
	return uint32(cRet)
}

func SDL_GetQueuedAudioSize(dev SDL_AudioDeviceID) uint32 {
	cRet := C.SDL_GetQueuedAudioSize(cUint32(dev))
	return uint32(cRet)
}

func SDL_ClearQueuedAudio(dev SDL_AudioDeviceID) {
	C.SDL_ClearQueuedAudio(cUint32(dev))
}

func SDL_MixAudio(dst, src *uint8, len uint32, volume int) {
	cDst := (*cUint8)(unsafe.Pointer(dst))
	cSrc := (*cUint8)(unsafe.Pointer(src))
	C.SDL_MixAudio(cDst, cSrc, cUint32(len), cInt(volume))
}

func SDL_MixAudioFormat(dst, src *uint8, format SDL_AudioFormat, len uint32, volume int) {
	cDst := (*cUint8)(unsafe.Pointer(dst))
	cSrc := (*cUint8)(unsafe.Pointer(src))
	C.SDL_MixAudioFormat(cDst, cSrc, cUint16(format), cUint32(len), cInt(volume))
}

func SDL_LockAudio() {
	C.SDL_LockAudio()
}

func SDL_LockAudioDevice(dev SDL_AudioDeviceID) {
	C.SDL_LockAudioDevice(cUint32(dev))
}

func SDL_UnlockAudio() {
	C.SDL_UnlockAudio()
}

func SDL_UnlockAudioDevice(dev SDL_AudioDeviceID) {
	C.SDL_UnlockAudioDevice(cUint32(dev))
}

func SDL_CloseAudio() {
	C.SDL_CloseAudio()
}

func SDL_CloseAudioDevice(dev SDL_AudioDeviceID) {
	C.SDL_CloseAudioDevice(cUint32(dev))
}

func SDL_NewAudioStream(srcFormat SDL_AudioFormat, srcChannels uint8, srcRate int, dstFormat SDL_AudioFormat, dstChannels uint8, dstRate int) (stream *SDL_AudioStream) {
	cSrcFormat := C.SDL_AudioFormat(srcFormat)
	cSrcChannels := cUint8(srcChannels)
	cSrcRate := cInt(srcRate)
	cDstFormat := C.SDL_AudioFormat(dstFormat)
	cDstChannels := cUint8(dstChannels)
	cDstRate := cInt(dstRate)

	cStream := C.SDL_NewAudioStream(cSrcFormat, cSrcChannels, cSrcRate, cDstFormat, cDstChannels, cDstRate)
	stream = (*SDL_AudioStream)(unsafe.Pointer(cStream))
	return
}

func SDL_AudioStreamPut(stream *SDL_AudioStream, buf []byte, bufLen int) int {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cBuf := unsafe.Pointer(sliceHeader.Data)
	// cLen := cInt(len(buf))
	cRet := C.SDL_AudioStreamPut(cSDL_AudioStream(stream), cBuf, cInt(bufLen))
	return int(cRet)
}

func SDL_AudioStreamGet(stream *SDL_AudioStream, buf []byte, bufLen int) int {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cBuf := unsafe.Pointer(sliceHeader.Data)
	// cLen := cInt(len(buf))
	cRet := C.SDL_AudioStreamGet(cSDL_AudioStream(stream), cBuf, cInt(bufLen))
	return int(cRet)
}

func SDL_AudioStreamAvailable(stream *SDL_AudioStream) int {
	cRet := C.SDL_AudioStreamAvailable(cSDL_AudioStream(stream))
	return int(cRet)
}

func SDL_AudioStreamFlush(stream *SDL_AudioStream) int {
	cRet := C.SDL_AudioStreamFlush(cSDL_AudioStream(stream))
	return int(cRet)
}

func SDL_AudioStreamClear(stream *SDL_AudioStream) {
	C.SDL_AudioStreamClear(cSDL_AudioStream(stream))
}

func SDL_AudoiStreamFree(stream *SDL_AudioStream) {
	C.SDL_FreeAudioStream(cSDL_AudioStream(stream))
}

//export SDL_AudioCallbackWrapper
func SDL_AudioCallbackWrapper(_ unsafe.Pointer, stream *cUint8, streamLen cInt) {
	if SDL_AudioOK.callback != nil {
		SDL_AudioOK.callback(SDL_AudioOK.userdata, (*uint8)(unsafe.Pointer(stream)), int(streamLen))
	}
}

type SDL_AudioWatcher struct {
	callback SDL_AudioCallback
	userdata any
}

var SDL_AudioOK SDL_AudioWatcher

type SDL_AudioCallback func(userdata any, stream *uint8, streamLen int)

func SDL_SetAudioCallback(audioSpec *SDL_AudioSpec, callback SDL_AudioCallback, userdata any) {
	SDL_AudioOK.callback = callback
	SDL_AudioOK.userdata = userdata

	cAudioSpec := cSDL_AudioSpec(audioSpec)
	cAudioSpec.callback = C.SDL_AudioCallback(C.SDL_AudioCallbackWrapper)
}

func SDL_GetAudioDeviceSpec(index int, isCapture int, spec *SDL_AudioSpec) int {
	cRet := C.SDL_GetAudioDeviceSpec(cInt(index), cInt(isCapture), cSDL_AudioSpec(spec))
	return int(cRet)
}

func SDL_GetDefaultAudioInfo(name *[]byte, spec *SDL_AudioSpec, isCapture int) int {
	cName := (**cChar)(unsafe.Pointer(name))
	cRet := C.SDL_GetDefaultAudioInfo(cName, cSDL_AudioSpec(spec), cInt(isCapture))
	return int(cRet)
}
