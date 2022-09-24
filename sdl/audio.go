package sdl

/*
#include "audio.h"
*/
import "C"
import (
	"reflect"
	"unsafe"

	. "github.com/chunqian/memory"
)

func (as *SDL_AudioSpec) cSDL_AudioSpec() *C.SDL_AudioSpec {
	return (*C.SDL_AudioSpec)(unsafe.Pointer(as))
}

func (cvt *SDL_AudioCVT) cSDL_AudioCVT() *C.SDL_AudioCVT {
	return (*C.SDL_AudioCVT)(unsafe.Pointer(cvt))
}

func (stream *SDL_AudioStream) cSDL_AudioStream() *C.SDL_AudioStream {
	return (*C.SDL_AudioStream)(unsafe.Pointer(stream))
}

// AllocBuf allocates the requested memory for AudioCVT buffer.
func (cvt *SDL_AudioCVT) AllocBuf(size uintptr) {
	cvt.Buf = (*uint8)(MP_Malloc(SDL_GetMemoryPool(), uint32(size)))
}

// FreeBuf deallocates the memory previously allocated from AudioCVT buffer.
func (cvt *SDL_AudioCVT) FreeBuf() {
	MP_Free(SDL_GetMemoryPool(), unsafe.Pointer(cvt.Buf))
}

// BufAsSlice returns AudioCVT.buf as byte slice.
// NOTE: Must be used after ConvertAudio() because it uses LenCVT as slice length.
func (cvt SDL_AudioCVT) BufAsSlice() []byte {
	var b []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Len = int(cvt.LenCVT)
	sliceHeader.Cap = int(cvt.Len * cvt.LenMult)
	sliceHeader.Data = uintptr(unsafe.Pointer(cvt.Buf))
	return b
}

// GetNumAudioDrivers returns the number of built-in audio drivers.
func SDL_GetNumAudioDrivers() int {
	return int(C.SDL_GetNumAudioDrivers())
}

// GetAudioDriver returns the name of a built in audio driver.
func SDL_GetAudioDriver(index int) string {
	cstr := C.SDL_GetAudioDriver(cInt(index))
	return SDL_GoString(cstr)
}

// AudioInit initializes a particular audio driver.
func SDL_AudioInit(driverName string) int {
	cDriverName := SDL_CreateCString(SDL_GetMemoryPool(), driverName)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cDriverName) // memory free

	cRet := C.SDL_AudioInit(cDriverName)
	return int(cRet)
}

// AudioQuit shuts down audio if you initialized it with AudioInit().
func SDL_AudioQuit() {
	C.SDL_AudioQuit()
}

// GetCurrentAudioDriver returns the name of the current audio driver.
func SDL_GetCurrentAudioDriver() string {
	cstr := C.SDL_GetCurrentAudioDriver()
	return SDL_GoString(cstr)
}

// OpenAudio opens the audio device. New programs might want to use OpenAudioDevice() instead.
func SDL_OpenAudio(desired, obtained *SDL_AudioSpec) int {
	cRet := C.SDL_OpenAudio(desired.cSDL_AudioSpec(), obtained.cSDL_AudioSpec())
	return int(cRet)
}

// GetNumAudioDevices returns the number of built-in audio devices.
func SDL_GetNumAudioDevices(isCapture int32) int {
	cRet := C.SDL_GetNumAudioDevices(cInt(isCapture))
	return int(cRet)
}

// GetAudioDeviceName returns the name of a specific audio device.
func SDL_GetAudioDeviceName(index int32, isCapture int32) string {
	cstr := C.SDL_GetAudioDeviceName(cInt(index), cInt(isCapture))
	return SDL_GoString(cstr)
}

// OpenAudioDevice opens a specific audio device.
func SDL_OpenAudioDevice(device string, isCapture int32, desired, obtained *SDL_AudioSpec, allowedChanges int32) SDL_AudioDeviceID {
	cDevice := SDL_CreateCString(SDL_GetMemoryPool(), device)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cDevice) // memory free

	if device == "" {
		cDevice = nil
	}
	cId := C.SDL_OpenAudioDevice(cDevice, cInt(isCapture), desired.cSDL_AudioSpec(), obtained.cSDL_AudioSpec(), cInt(allowedChanges))
	return SDL_AudioDeviceID(cId)
}

// GetAudioStatus returns the current audio state of the audio device. New programs might want to use GetAudioDeviceStatus() instead.
func SDL_GetAudioStatus() SDL_AudioStatus {
	cRet := C.SDL_GetAudioStatus()
	return SDL_AudioStatus(cRet)
}

// GetAudioDeviceStatus returns the current audio state of an audio device.
func SDL_GetAudioDeviceStatus(dev SDL_AudioDeviceID) SDL_AudioStatus {
	cRet := C.SDL_GetAudioDeviceStatus(cUint32(dev))
	return SDL_AudioStatus(cRet)
}

// PauseAudio pauses and unpauses the audio device. New programs might want to use SDL_PauseAudioDevice() instead.
func SDL_PauseAudio(pauseOn int32) {
	C.SDL_PauseAudio(cInt(pauseOn))
}

// PauseAudioDevice pauses and unpauses audio playback on a specified device.
func SDL_PauseAudioDevice(dev SDL_AudioDeviceID, pauseOn int32) {
	C.SDL_PauseAudioDevice(cUint32(dev), cInt(pauseOn))
}

// LoadWAVRW loads a WAVE from the data source, automatically freeing that source if freeSrc is true.
func SDL_LoadWAV_RW(src *SDL_RWops, freeSrc int32, spec *SDL_AudioSpec, audioBuf *[]uint8, audioLen *uint32) *SDL_AudioSpec {
	cAudioBuf := (**cUint8)(unsafe.Pointer(audioBuf))
	cAudioLen := (*cUint32)(unsafe.Pointer(audioLen))
	cAudioSpec := C.SDL_LoadWAV_RW(src.cSDL_RWops(), cInt(freeSrc), spec.cSDL_AudioSpec(), cAudioBuf, cAudioLen)
	audioSpec := (*SDL_AudioSpec)(unsafe.Pointer(cAudioSpec))

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(audioBuf))
	sliceHeader.Len = int(*audioLen)
	sliceHeader.Cap = int(*audioLen)
	sliceHeader.Data = uintptr(unsafe.Pointer(audioBuf))
	return audioSpec
}

// LoadWAV loads a WAVE from a file.
func SDL_LoadWAV(file string, spec *SDL_AudioSpec, audioBuf *[]uint8, audioLen *uint32) *SDL_AudioSpec {
	cFile := SDL_CreateCString(SDL_GetMemoryPool(), file)
	cRB := SDL_CreateCString(SDL_GetMemoryPool(), "rb")
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cFile) // memory free
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cRB)   // memory free

	cAudioBuf := (**cUint8)(unsafe.Pointer(audioBuf))
	cAudioLen := (*cUint32)(unsafe.Pointer(audioLen))
	cAudioSpec := C.SDL_LoadWAV_RW(C.SDL_RWFromFile(cFile, cRB), 1, spec.cSDL_AudioSpec(), cAudioBuf, cAudioLen)
	audioSpec := (*SDL_AudioSpec)(unsafe.Pointer(cAudioSpec))

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(audioBuf))
	sliceHeader.Len = int(*audioLen)
	sliceHeader.Cap = int(*audioLen)
	sliceHeader.Data = uintptr(unsafe.Pointer(audioBuf))
	return audioSpec
}

// FreeWAV frees data previously allocated with LoadWAV() or LoadWAVRW().
func SDL_FreeWAV(audioBuf []uint8) {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&audioBuf))
	cAudioBuf := (*cUint8)(unsafe.Pointer(sliceHeader.Data))
	C.SDL_FreeWAV(cAudioBuf)
}

// BuildAudioCVT initializes an AudioCVT structure for conversion.
func SDL_BuildAudioCVT(cvt *SDL_AudioCVT, srcFormat SDL_AudioFormat, srcChannels uint8, srcRate int, dstFormat SDL_AudioFormat, dstChannels uint8, dstRate int) (converted int) {
	cRet := C.SDL_BuildAudioCVT(cvt.cSDL_AudioCVT(), cUint16(srcFormat), cUint8(srcChannels), cInt(srcRate), cUint16(dstFormat), cUint8(dstChannels), cInt(dstRate))
	return int(cRet)
}

// ConvertAudio converts audio data to a desired audio format.
func SDL_ConvertAudio(cvt *SDL_AudioCVT) int {
	cRet := C.SDL_ConvertAudio(cvt.cSDL_AudioCVT())
	return int(cRet)
}

// QueueAudio queues more audio on non-callback devices.
func SDL_QueueAudio(dev SDL_AudioDeviceID, data []byte, len uint32) int {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	cData := unsafe.Pointer(sliceHeader.Data)
	// cLen := cUint32(sliceHeader.Len)
	cRet := C.SDL_QueueAudio(cUint32(dev), cData, cUint32(len))
	return int(cRet)
}

// DequeueAudio dequeues more audio on non-callback devices.
func SDL_DequeueAudio(dev SDL_AudioDeviceID, data []byte, len uint32) uint32 {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	cData := unsafe.Pointer(sliceHeader.Data)
	// cLen := cUint32(sliceHeader.Len)
	cRet := C.SDL_DequeueAudio(cUint32(dev), cData, cUint32(len))
	return uint32(cRet)
}

// GetQueuedAudioSize returns the number of bytes of still-queued audio.
func SDL_GetQueuedAudioSize(dev SDL_AudioDeviceID) uint32 {
	cRet := C.SDL_GetQueuedAudioSize(cUint32(dev))
	return uint32(cRet)
}

// ClearQueuedAudio drops any queued audio data waiting to be sent to the hardware.
func SDL_ClearQueuedAudio(dev SDL_AudioDeviceID) {
	C.SDL_ClearQueuedAudio(cUint32(dev))
}

// MixAudio mixes audio data. New programs might want to use MixAudioFormat() instead.
func SDL_MixAudio(dst, src *uint8, len uint32, volume int) {
	cDst := (*cUint8)(unsafe.Pointer(dst))
	cSrc := (*cUint8)(unsafe.Pointer(src))
	C.SDL_MixAudio(cDst, cSrc, cUint32(len), cInt(volume))
}

// MixAudioFormat mixes audio data in a specified format.
func SDL_MixAudioFormat(dst, src *uint8, format SDL_AudioFormat, len uint32, volume int) {
	cDst := (*cUint8)(unsafe.Pointer(dst))
	cSrc := (*cUint8)(unsafe.Pointer(src))
	C.SDL_MixAudioFormat(cDst, cSrc, cUint16(format), cUint32(len), cInt(volume))
}

// LockAudio locks the audio device. New programs might want to use LockAudioDevice() instead.
func SDL_LockAudio() {
	C.SDL_LockAudio()
}

// LockAudioDevice locks out the audio callback function for a specified device.
func SDL_LockAudioDevice(dev SDL_AudioDeviceID) {
	C.SDL_LockAudioDevice(cUint32(dev))
}

// UnlockAudio unlocks the audio device. New programs might want to use UnlockAudioDevice() instead.
func SDL_UnlockAudio() {
	C.SDL_UnlockAudio()
}

// UnlockAudioDevice unlocks the audio callback function for a specified device.
func SDL_UnlockAudioDevice(dev SDL_AudioDeviceID) {
	C.SDL_UnlockAudioDevice(cUint32(dev))
}

// CloseAudio closes the audio device. New programs might want to use CloseAudioDevice() instead.
func SDL_CloseAudio() {
	C.SDL_CloseAudio()
}

// CloseAudioDevice shuts down audio processing and closes the audio device.
func SDL_CloseAudioDevice(dev SDL_AudioDeviceID) {
	C.SDL_CloseAudioDevice(cUint32(dev))
}

// NewAudioStream creates a new audio stream
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

// Put adds data to be converted/resampled to the stream
func SDL_AudioStreamPut(stream *SDL_AudioStream, buf []byte, bufLen int) int {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cBuf := unsafe.Pointer(sliceHeader.Data)
	// cLen := cInt(len(buf))
	cRet := C.SDL_AudioStreamPut(stream.cSDL_AudioStream(), cBuf, cInt(bufLen))
	return int(cRet)
}

// Get gets converted/resampled data from the stream
func SDL_AudioStreamGet(stream *SDL_AudioStream, buf []byte, bufLen int) int {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cBuf := unsafe.Pointer(sliceHeader.Data)
	// cLen := cInt(len(buf))
	cRet := C.SDL_AudioStreamGet(stream.cSDL_AudioStream(), cBuf, cInt(bufLen))
	return int(cRet)
}

// Available gets the number of converted/resampled bytes available
func SDL_AudioStreamAvailable(stream *SDL_AudioStream) int {
	cRet := C.SDL_AudioStreamAvailable(stream.cSDL_AudioStream())
	return int(cRet)
}

// Flush tells the stream that you're done sending data, and anything being buffered
// should be converted/resampled and made available immediately.
func SDL_AudioStreamFlush(stream *SDL_AudioStream) int {
	cRet := C.SDL_AudioStreamFlush(stream.cSDL_AudioStream())
	return int(cRet)
}

// Clear clears any pending data in the stream without converting it
func SDL_AudioStreamClear(stream *SDL_AudioStream) {
	C.SDL_AudioStreamClear(stream.cSDL_AudioStream())
}

// Free frees the audio stream
func SDL_AudoiStreamFree(stream *SDL_AudioStream) {
	C.SDL_FreeAudioStream(stream.cSDL_AudioStream())
}

//export SDL_AudioCallbackWrapper
func SDL_AudioCallbackWrapper(_ unsafe.Pointer, stream *cUint8, streamLen cInt) {
	if SDL_AudioOK.callback != nil {
		SDL_AudioOK.callback(SDL_AudioOK.userdata, (*uint8)(unsafe.Pointer(stream)), int(streamLen))
	}
}

type SDL_AudioWatcher struct {
	callback SDL_AudioCallback
	userdata interface{}
}

var SDL_AudioOK SDL_AudioWatcher

// AudioCallback is a function to call when the audio device needs more data.`
type SDL_AudioCallback func(userdata interface{}, stream *uint8, streamLen int)

func SDL_SetAudioCallback(audioSpec *SDL_AudioSpec, callback SDL_AudioCallback, userdata interface{}) {
	SDL_AudioOK.callback = callback
	SDL_AudioOK.userdata = userdata

	cAudioSpec := audioSpec.cSDL_AudioSpec()
	cAudioSpec.callback = C.SDL_AudioCallback(C.SDL_AudioCallbackWrapper)
}

// GetAudioDeviceSpec returns the preferred audio format of a specific audio device.
func SDL_GetAudioDeviceSpec(index int, isCapture int, spec *SDL_AudioSpec) int {
	cRet := C.SDL_GetAudioDeviceSpec(cInt(index), cInt(isCapture), spec.cSDL_AudioSpec())
	return int(cRet)
}

// GetDefaultAudioInfo returns the name and preferred format of the default audio device.
func SDL_GetDefaultAudioInfo(name *[]byte, spec *SDL_AudioSpec, isCapture int) int {
	cName := (**cChar)(unsafe.Pointer(name))
	cRet := C.SDL_GetDefaultAudioInfo(cName, spec.cSDL_AudioSpec(), cInt(isCapture))
	return int(cRet)
}
