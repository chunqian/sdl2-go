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

// define
var (
	SDL_AUDIO_MASK_BITSIZE   = 0xFF
	SDL_AUDIO_MASK_DATATYPE  = 1 << 8
	SDL_AUDIO_MASK_ENDIAN    = 1 << 12
	SDL_AUDIO_MASK_SIGNED    = 1 << 15
	SDL_AUDIO_BITSIZE        = func(x int) int { return x & SDL_AUDIO_MASK_BITSIZE }
	SDL_AUDIO_ISFLOAT        = func(x int) bool { return (x & SDL_AUDIO_MASK_DATATYPE) != 0 }
	SDL_AUDIO_ISBIGENDIAN    = func(x int) bool { return (x & SDL_AUDIO_MASK_ENDIAN) != 0 }
	SDL_AUDIO_ISSIGNED       = func(x int) bool { return (x & SDL_AUDIO_MASK_SIGNED) != 0 }
	SDL_AUDIO_ISINT          = func(x int) bool { return !SDL_AUDIO_ISFLOAT(x) }
	SDL_AUDIO_ISLITTLEENDIAN = func(x int) bool { return !SDL_AUDIO_ISBIGENDIAN(x) }
	SDL_AUDIO_ISUNSIGNED     = func(x int) bool { return !SDL_AUDIO_ISSIGNED(x) }

	SDL_MIX_MAXVOLUME = 128
)

var (
	AUDIO_U8     = 0x0008
	AUDIO_S8     = 0x8008
	AUDIO_U16LSB = 0x0010
	AUDIO_S16LSB = 0x8010
	AUDIO_U16MSB = 0x1010
	AUDIO_S16MSB = 0x9010
	AUDIO_U16    = AUDIO_U16LSB
	AUDIO_S16    = AUDIO_S16LSB

	AUDIO_S32LSB = 0x8020
	AUDIO_S32MSB = 0x9020
	AUDIO_S32    = AUDIO_S32LSB

	AUDIO_F32LSB = 0x8120
	AUDIO_F32MSB = 0x9120
	AUDIO_F32    = AUDIO_F32LSB

	AUDIO_U16SYS = func() int {
		if !IsLittleEndian() {
			return AUDIO_U16LSB
		} else {
			return AUDIO_U16MSB
		}
	}()
	AUDIO_S16SYS = func() int {
		if !IsLittleEndian() {
			return AUDIO_S16LSB
		} else {
			return AUDIO_S16MSB
		}
	}()
	AUDIO_S32SYS = func() int {
		if !IsLittleEndian() {
			return AUDIO_S32LSB
		} else {
			return AUDIO_S32MSB
		}
	}()
	AUDIO_F32SYS = func() int {
		if !IsLittleEndian() {
			return AUDIO_F32LSB
		} else {
			return AUDIO_F32MSB
		}
	}()
)

const (
	SDL_AUDIO_ALLOW_FREQUENCY_CHANGE = 0x00000001
	SDL_AUDIO_ALLOW_FORMAT_CHANGE    = 0x00000002
	SDL_AUDIO_ALLOW_CHANNELS_CHANGE  = 0x00000004
	SDL_AUDIO_ALLOW_SAMPLES_CHANGE   = 0x00000008
	SDL_AUDIO_ALLOW_ANY_CHANGE       = SDL_AUDIO_ALLOW_FREQUENCY_CHANGE | SDL_AUDIO_ALLOW_FORMAT_CHANGE | SDL_AUDIO_ALLOW_CHANNELS_CHANGE | SDL_AUDIO_ALLOW_SAMPLES_CHANGE
)

// typedef
type SDL_AudioFormat = uint16
type SDL_AudioFilter = func(*SDL_AudioCVT, uint16)
type SDL_AudioDeviceID = uint32
type SDL_AudioStatus = int32

// enum
const (
	SDL_AUDIO_STOPPED SDL_AudioStatus = 0
	SDL_AUDIO_PLAYING SDL_AudioStatus = 1
	SDL_AUDIO_PAUSED  SDL_AudioStatus = 2
)

// struct
type SDL_AudioSpec struct {
	Freq     int32
	Format   uint16
	Channels uint8
	Silence  uint8
	Samples  uint16
	Padding  uint16
	Size     uint32
	callback SDL_AudioCallback
	userdata unsafe.Pointer
}

type SDL_AudioCVT struct {
	Needed      int32
	SrcFormat   uint16
	DstFormat   uint16
	RateIncr    float64
	Buf         *uint8
	Len         int32
	LenCVT      int32
	LenMult     int32
	LenRatio    float64
	Filters     [10]func(*SDL_AudioCVT, uint16)
	FilterIndex int32
}

func cAudioSpec(as *SDL_AudioSpec) *C.SDL_AudioSpec {
	return (*C.SDL_AudioSpec)(unsafe.Pointer(as))
}

func cAudioCVT(cvt *SDL_AudioCVT) *C.SDL_AudioCVT {
	return (*C.SDL_AudioCVT)(unsafe.Pointer(cvt))
}

func cAudioStream(stream *SDL_AudioStream) *C.SDL_AudioStream {
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
	cRet := C.SDL_OpenAudio(cAudioSpec(desired), cAudioSpec(obtained))
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
	cId := C.SDL_OpenAudioDevice(cDevice.(*cChar), cInt(isCapture), cAudioSpec(desired), cAudioSpec(obtained), cInt(allowedChanges))
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
	var cAudioBuf *cUint8

	cAudioLen := (*cUint32)(unsafe.Pointer(audioLen))
	cSpec := C.SDL_LoadWAV_RW(cRWops(src), cInt(freeSrc), cAudioSpec(spec), &cAudioBuf, cAudioLen)
	dstSpec := (*SDL_AudioSpec)(unsafe.Pointer(cSpec))

	sh := (*reflect.SliceHeader)(unsafe.Pointer(audioBuf))
	sh.Len = int(*audioLen)
	sh.Cap = int(*audioLen)
	sh.Data = uintptr(unsafe.Pointer(cAudioBuf))
	return dstSpec
}

func SDL_LoadWAV(file string, spec *SDL_AudioSpec, audioBuf *[]uint8) *SDL_AudioSpec {
	var cAudioBuf *cUint8
	var audioLen uint32

	cFile := SDL_CreateCString(SDL_GetMemoryPool(), file)
	cRB := SDL_CreateCString(SDL_GetMemoryPool(), "rb")
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cFile) // memory free
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cRB)   // memory free

	cAudioLen := (*cUint32)(unsafe.Pointer(&audioLen))
	cSpec := C.SDL_LoadWAV_RW(C.SDL_RWFromFile(cFile.(*cChar), cRB.(*cChar)), 1, cAudioSpec(spec), &cAudioBuf, cAudioLen)
	dstSpec := (*SDL_AudioSpec)(unsafe.Pointer(cSpec))

	sh := (*reflect.SliceHeader)(unsafe.Pointer(audioBuf))
	sh.Len = int(audioLen)
	sh.Cap = int(audioLen)
	sh.Data = uintptr(unsafe.Pointer(cAudioBuf))
	return dstSpec
}

func SDL_FreeWAV(audioBuf []uint8) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&audioBuf))
	cAudioBuf := (*cUint8)(unsafe.Pointer(sh.Data))
	C.SDL_FreeWAV(cAudioBuf)
}

func SDL_BuildAudioCVT(cvt *SDL_AudioCVT, srcFormat SDL_AudioFormat, srcChannels uint8, srcRate int, dstFormat SDL_AudioFormat, dstChannels uint8, dstRate int) (converted int) {
	cRet := C.SDL_BuildAudioCVT(cAudioCVT(cvt), cUint16(srcFormat), cUint8(srcChannels), cInt(srcRate), cUint16(dstFormat), cUint8(dstChannels), cInt(dstRate))
	return int(cRet)
}

func SDL_ConvertAudio(cvt *SDL_AudioCVT) int {
	cRet := C.SDL_ConvertAudio(cAudioCVT(cvt))
	return int(cRet)
}

func SDL_QueueAudio(dev SDL_AudioDeviceID, data []byte, len uint32) int {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	cData := unsafe.Pointer(sh.Data)
	// cLen := cUint32(sh.Len)
	cRet := C.SDL_QueueAudio(cUint32(dev), cData, cUint32(len))
	return int(cRet)
}

func SDL_DequeueAudio(dev SDL_AudioDeviceID, data []byte, len uint32) uint32 {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	cData := unsafe.Pointer(sh.Data)
	// cLen := cUint32(sh.Len)
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
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cBuf := unsafe.Pointer(sh.Data)
	// cLen := cInt(len(buf))
	cRet := C.SDL_AudioStreamPut(cAudioStream(stream), cBuf, cInt(bufLen))
	return int(cRet)
}

func SDL_AudioStreamGet(stream *SDL_AudioStream, buf []byte, bufLen int) int {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	cBuf := unsafe.Pointer(sh.Data)
	// cLen := cInt(len(buf))
	cRet := C.SDL_AudioStreamGet(cAudioStream(stream), cBuf, cInt(bufLen))
	return int(cRet)
}

func SDL_AudioStreamAvailable(stream *SDL_AudioStream) int {
	cRet := C.SDL_AudioStreamAvailable(cAudioStream(stream))
	return int(cRet)
}

func SDL_AudioStreamFlush(stream *SDL_AudioStream) int {
	cRet := C.SDL_AudioStreamFlush(cAudioStream(stream))
	return int(cRet)
}

func SDL_AudioStreamClear(stream *SDL_AudioStream) {
	C.SDL_AudioStreamClear(cAudioStream(stream))
}

func SDL_AudoiStreamFree(stream *SDL_AudioStream) {
	C.SDL_FreeAudioStream(cAudioStream(stream))
}

//export callAudio
func callAudio(_ unsafe.Pointer, stream *cUint8, streamLen cInt) {
	var slice []uint8
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sh.Data = uintptr(unsafe.Pointer(stream))
	sh.Len = int(streamLen)
	sh.Cap = int(streamLen)
	if SDL_AudioOK.callback != nil {
		SDL_AudioOK.callback(SDL_AudioOK.userdata, slice)
	}
}

type SDL_AudioWatcher struct {
	callback SDL_AudioCallback
	userdata any
}

var SDL_AudioOK SDL_AudioWatcher

type SDL_AudioCallback func(userdata any, stream []uint8)

func SDL_SetAudioCallback(audioSpec *SDL_AudioSpec, callback SDL_AudioCallback, userdata any) {
	SDL_AudioOK.callback = callback
	SDL_AudioOK.userdata = userdata

	cSpec := cAudioSpec(audioSpec)
	cSpec.callback = C.SDL_AudioCallback(C.callAudio)
}

func SDL_GetAudioDeviceSpec(index int, isCapture int, spec *SDL_AudioSpec) int {
	cRet := C.SDL_GetAudioDeviceSpec(cInt(index), cInt(isCapture), cAudioSpec(spec))
	return int(cRet)
}

func SDL_GetDefaultAudioInfo(name *[]byte, spec *SDL_AudioSpec, isCapture int) int {
	var cName *cChar
	cRet := C.SDL_GetDefaultAudioInfo(&cName, cAudioSpec(spec), cInt(isCapture))
	defer C.free(unsafe.Pointer(cName))

	// strcpy
	size := PX_strlen((*PX_char)(unsafe.Pointer(cName)))
	nameSlice := make([]byte, size)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&nameSlice))
	src := (*PX_char)(unsafe.Pointer(cName))
	dst := (*PX_char)(unsafe.Pointer(sh.Data))
	PX_strcpy(dst, src, size)
	name = &nameSlice
	return int(cRet)
}
