package mix

/*
#include "mixer_wrapper.h"
*/
import "C"
import (
	"reflect"
	"unsafe"

	. "github.com/chunqian/sdl2-go/sdl"
)

func cSDL_RWops(rw *SDL_RWops) *C.SDL_RWops {
	return (*C.SDL_RWops)(unsafe.Pointer(rw))
}

func cMix_Music(m *Mix_Music) *C.Mix_Music {
	return (*C.Mix_Music)(unsafe.Pointer(m))
}

func cMix_Chunk(c *Mix_Chunk) *C.Mix_Chunk {
	return (*C.Mix_Chunk)(unsafe.Pointer(c))
}

func Mix_Init(flags MIX_InitFlags) int {
	cRet := C.Mix_Init(cInt(flags))
	return int(cRet)
}

func Mix_Quit() {
	C.Mix_Quit()
}

func Mix_OpenAudio(frequency int, format uint16, channels, chunksize int) int {
	cFrequency := cInt(frequency)
	cFormat := cUint16(format)
	cChannels := cInt(channels)
	cChunksize := cInt(chunksize)

	cRet := C.Mix_OpenAudio(cFrequency, cFormat, cChannels, cChunksize)
	return int(cRet)
}

func Mix_OpenAudioDevice(frequency int, format uint16, channels, chunksize int, device string, allowedChanges int) int {
	cFrequency := cInt(frequency)
	cFormat := cUint16(format)
	cChannels := cInt(channels)
	cChunksize := cInt(chunksize)
	cAllowedChanges := cInt(allowedChanges)

	cDevice := SDL_CreateCString(SDL_GetMemoryPool(), device)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cDevice) // memory free

	if device == "" {
		// Passing in a device name of NULL requests the most reasonable default
		// (and is equivalent to what SDL_OpenAudio() does to choose a device)
		cDevice = nil
	}

	cRet := C.Mix_OpenAudioDevice(cFrequency, cFormat, cChannels, cChunksize, cDevice.(*cChar), cAllowedChanges)
	return int(cRet)
}

func Mix_AllocateChannels(numchans int) int {
	cNumchans := cInt(numchans)
	return int(C.Mix_AllocateChannels(cNumchans))
}

func Mix_QuerySpec(frequency *int, format *uint16, channels *int) int {
	cFrequency := (*cInt)(unsafe.Pointer(frequency))
	cFormat := (*cUint16)(unsafe.Pointer(format))
	cChannels := (*cInt)(unsafe.Pointer(channels))

	cRet := C.Mix_QuerySpec(cFrequency, cFormat, cChannels)
	return int(cRet)
}

func Mix_LoadWAV_RW(src *SDL_RWops, freesrc SDL_bool) (chunk *Mix_Chunk) {
	cSrc := cSDL_RWops(src)
	cFreesrc := cInt(freesrc)

	cChunk := C.Mix_LoadWAV_RW(cSrc, cFreesrc)
	chunk = (*Mix_Chunk)(unsafe.Pointer(cChunk))
	return
}

func Mix_LoadWAV(file string) (chunk *Mix_Chunk) {
	cFile := SDL_CreateCString(SDL_GetMemoryPool(), file)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cFile) // memory free

	cRB := SDL_CreateCString(SDL_GetMemoryPool(), "rb")
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cRB) // memory free

	// why doesn't this call Mix_LoadWAV ?
	cChunk := C.Mix_LoadWAV_RW(C.SDL_RWFromFile(cFile.(*cChar), cRB.(*cChar)), 1)
	chunk = (*Mix_Chunk)(unsafe.Pointer(cChunk))
	return
}

func Mix_LoadMUS(file string) (mus *Mix_Music) {
	cFile := SDL_CreateCString(SDL_GetMemoryPool(), file)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cFile) // memory free

	cMus := C.Mix_LoadMUS(cFile.(*cChar))
	mus = (*Mix_Music)(unsafe.Pointer(cMus))
	return
}

func Mix_LoadMUS_RW(src *SDL_RWops, freesrc int) (mus *Mix_Music) {
	cSrc := cSDL_RWops(src)
	cFreesrc := cInt(freesrc)

	cMus := C.Mix_LoadMUS_RW(cSrc, cFreesrc)
	mus = (*Mix_Music)(unsafe.Pointer(cMus))
	return
}

func Mix_LoadMUSType_RW(src *SDL_RWops, type_ Mix_MusicType, freesrc int) (mus *Mix_Music) {
	cSrc := cSDL_RWops(src)
	cType := (C.Mix_MusicType)(type_)
	cFreesrc := cInt(freesrc)

	cMus := C.Mix_LoadMUSType_RW(cSrc, cType, cFreesrc)
	mus = (*Mix_Music)(unsafe.Pointer(cMus))
	return
}

func Mix_QuickLoad_WAV(mem []byte) (chunk *Mix_Chunk) {
	cMem := (*cUint8)(&mem[0])

	cChunk := C.Mix_QuickLoad_WAV(cMem)
	chunk = (*Mix_Chunk)(unsafe.Pointer(cChunk))
	return
}

func Mix_QuickLoad_RAW(mem *uint8, len_ uint32) (chunk *Mix_Chunk) {
	cMem := (*cUint8)(mem)
	cLen := cUint32(len_)

	cChunk := C.Mix_QuickLoad_RAW(cMem, cLen)
	chunk = (*Mix_Chunk)(unsafe.Pointer(cChunk))
	return
}

func Mix_FreeChunk(chunk *Mix_Chunk) {
	cChunk := cMix_Chunk(chunk)
	C.Mix_FreeChunk(cChunk)
}

func Mix_FreeMusic(music *Mix_Music) {
	cMusic := cMix_Music(music)
	C.Mix_FreeMusic(cMusic)
}

func Mix_GetMusicType(music *Mix_Music) Mix_MusicType {
	cMusic := cMix_Music(music)
	return (Mix_MusicType)(C.Mix_GetMusicType(cMusic))
}

func Mix_SetPanning(channel int, left, right uint8) int {
	cChannel := cInt(channel)
	cLeft := cUint8(left)
	cRight := cUint8(right)

	cRet := C.Mix_SetPanning(cChannel, cLeft, cRight)
	return int(cRet)
}

func Mix_SetPosition(channel int, angle int16, distance uint8) int {
	cChannel := cInt(channel)
	cAngle := cInt16(angle)
	cDistance := cUint8(distance)

	cRet := C.Mix_SetPosition(cChannel, cAngle, cDistance)
	return int(cRet)
}

func Mix_SetDistance(channel int, distance uint8) int {
	cChannel := cInt(channel)
	cDistance := cUint8(distance)

	cRet := C.Mix_SetDistance(cChannel, cDistance)
	return int(cRet)
}

func Mix_SetReverseStereo(channel, flip int) int {
	cChannel := cInt(channel)
	cFlip := cInt(flip)

	cRet := C.Mix_SetReverseStereo(cChannel, cFlip)
	return int(cRet)
}

func Mix_ReserveChannels(num int) int {
	cNum := cInt(num)
	return int(C.Mix_ReserveChannels(cNum))
}

func Mix_GroupChannel(which, tag int) bool {
	cWhich := cInt(which)
	cTag := cInt(tag)
	return C.Mix_GroupChannel(cWhich, cTag) != 0
}

func Mix_GroupChannels(from, to, tag int) int {
	cFrom := cInt(from)
	cTo := cInt(to)
	cTag := cInt(tag)
	return int(C.Mix_GroupChannels(cFrom, cTo, cTag))
}

func Mix_GroupAvailable(tag int) int {
	cTag := cInt(tag)
	return int(C.Mix_GroupAvailable(cTag))
}

func Mix_GroupCount(tag int) int {
	cTag := cInt(tag)
	return int(C.Mix_GroupCount(cTag))
}

func Mix_GroupOldest(tag int) int {
	cTag := cInt(tag)
	return int(C.Mix_GroupOldest(cTag))
}

func Mix_GroupNewer(tag int) int {
	cTag := cInt(tag)
	return int(C.Mix_GroupNewer(cTag))
}

func Mix_PlayChannelTimed(channel int, chunk *Mix_Chunk, loops, ticks int) int {
	cChannel := cInt(channel)
	cChunk := cMix_Chunk(chunk)
	cLoops := cInt(loops)
	cTicks := cInt(ticks)

	cRet := C.Mix_PlayChannelTimed(cChannel, cChunk, cLoops, cTicks)
	return int(cRet)
}

func Mix_GetChunkTimeMilliseconds(chunk *Mix_Chunk) int {
	cChunk := cMix_Chunk(chunk)

	cRet := C.Mix_GetChunkTimeMilliseconds(cChunk)
	return int(cRet)
}

func Mix_PlayChannel(channel int, chunk *Mix_Chunk, loops int) int {
	cChannel := cInt(channel)
	cChunk := cMix_Chunk(chunk)
	cLoops := cInt(loops)

	cRet := C.Mix_PlayChannelTimed(cChannel, cChunk, cLoops, -1)
	return int(cRet)
}

func Mix_PlayMusic(music *Mix_Music, loops int) int {
	cMusic := cMix_Music(music)
	cLoops := cInt(loops)

	cRet := C.Mix_PlayMusic(cMusic, cLoops)
	return int(cRet)
}

func Mix_FadeInMusic(music *Mix_Music, loops, ms int) int {
	cMusic := cMix_Music(music)
	cLoops := cInt(loops)
	cMs := cInt(ms)

	cRet := C.Mix_FadeInMusic(cMusic, cLoops, cMs)
	return int(cRet)
}

func Mix_FadeInMusicPos(music *Mix_Music, loops, ms int, position float64) int {
	cMusic := cMix_Music(music)
	cLoops := cInt(loops)
	cMs := cInt(ms)
	cPosition := cDouble(position)

	cRet := C.Mix_FadeInMusicPos(cMusic, cLoops, cMs, cPosition)
	return int(cRet)
}

func Mix_FadeInChannel(channel int, chunk *Mix_Chunk, loops, ms int) int {
	return Mix_FadeInChannelTimed(channel, chunk, loops, ms, -1)
}

func Mix_FadeInChannelTimed(channel int, chunk *Mix_Chunk, loops, ms, ticks int) int {
	cChannel := cInt(channel)
	cChunk := cMix_Chunk(chunk)
	cLoops := cInt(loops)
	cMs := cInt(ms)
	cTicks := cInt(ticks)

	cRet := C.Mix_FadeInChannelTimed(cChannel, cChunk, cLoops, cMs, cTicks)
	return int(cRet)
}

func Mix_Volume(channel, volume int) int {
	cChannel := cInt(channel)
	cVolume := cInt(volume)
	return int(C.Mix_Volume(cChannel, cVolume))
}

func Mix_VolumeChunk(chunk *Mix_Chunk, volume int) int {
	cChunk := cMix_Chunk(chunk)
	cVolume := cInt(volume)

	cRet := C.Mix_VolumeChunk(cChunk, cVolume)
	return int(cRet)
}

func Mix_VolumeMusic(volume int) int {
	cVolume := cInt(volume)
	return int(C.Mix_VolumeMusic(cVolume))
}

func Mix_HaltChannel(channel int) {
	cChannel := cInt(channel)
	C.Mix_HaltChannel(cChannel)
}

func Mix_HaltGroup(tag int) {
	cTag := cInt(tag)
	C.Mix_HaltGroup(cTag)
}

func Mix_HaltMusic() {
	C.Mix_HaltMusic()
}

func Mix_ExpireChannel(channel, ticks int) int {
	cChannel := cInt(channel)
	cTicks := cInt(ticks)
	return int(C.Mix_ExpireChannel(cChannel, cTicks))
}

func Mix_FadeOutChannel(which, ms int) int {
	cWhich := cInt(which)
	cMs := cInt(ms)
	return int(C.Mix_FadeOutChannel(cWhich, cMs))
}

func Mix_FadeOutGroup(tag, ms int) int {
	cTag := cInt(tag)
	cMs := cInt(ms)
	return int(C.Mix_FadeOutGroup(cTag, cMs))
}

func Mix_FadeOutMusic(ms int) bool {
	cMs := cInt(ms)
	return int(C.Mix_FadeOutMusic(cMs)) == 0
}

func Mix_FadingMusic() Mix_Fading {
	return (Mix_Fading)(C.Mix_FadingMusic())
}

func Mix_FadingChannel(which int) Mix_Fading {
	cWhich := cInt(which)
	return (Mix_Fading)(C.Mix_FadingChannel(cWhich))
}

func Mix_Pause(channel int) {
	cChannel := cInt(channel)
	C.Mix_Pause(cChannel)
}

func Mix_Resume(channel int) {
	cChannel := cInt(channel)
	C.Mix_Resume(cChannel)
}

func Mix_Paused(channel int) int {
	cChannel := cInt(channel)
	return int(C.Mix_Paused(cChannel))
}

func Mix_PauseMusic() {
	C.Mix_PauseMusic()
}

func Mix_ResumeMusic() {
	C.Mix_ResumeMusic()
}

func Mix_RewindMusic() {
	C.Mix_RewindMusic()
}

func Mix_PausedMusic() bool {
	return int(C.Mix_PausedMusic()) > 0
}

func Mix_SetMusicPosition(position int64) int {
	cPosition := cDouble(position)

	cRet := C.Mix_SetMusicPosition(cPosition)
	return int(cRet)
}

func Mix_Playing(channel int) int {
	cChannel := cInt(channel)
	return int(C.Mix_Playing(cChannel))
}

func Mix_PlayingMusic() bool {
	return int(C.Mix_PlayingMusic()) > 0
}

func Mix_SetMusicCMD(command string) int {
	cCommand := SDL_CreateCString(SDL_GetMemoryPool(), command)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cCommand) // memory free

	cRet := C.Mix_SetMusicCMD(cCommand.(*cChar))
	return int(cRet)
}

func Mix_SetSynchroValue(value int) bool {
	cValue := cInt(value)

	cRet := C.Mix_SetSynchroValue(cValue)
	return int(cRet) == 0
}

func Mix_GetSynchroValue() int {
	return int(C.Mix_GetSynchroValue())
}

func Mix_GetChunk(channel int) *Mix_Chunk {
	cChannel := cInt(channel)
	return (*Mix_Chunk)(unsafe.Pointer(C.Mix_GetChunk(cChannel)))
}

func Mix_CloseAudio() {
	C.Mix_CloseAudio()
}

func Mix_GetNumChunkDecoders() int {
	return int(C.Mix_GetNumChunkDecoders())
}

func Mix_GetChunkDecoder(index int) string {
	return C.GoString(C.Mix_GetChunkDecoder(cInt(index)))
}

func Mix_GetNumMusicDecoders() int {
	return int(C.Mix_GetNumMusicDecoders())
}

func Mix_GetMusicDecoder(index int) string {
	return C.GoString(C.Mix_GetMusicDecoder(cInt(index)))
}

//export callPostMixFunction
func callPostMixFunction(udata unsafe.Pointer, stream *C.Uint8, length C.int) {
	var slice []uint8
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Data = uintptr(unsafe.Pointer(stream))
	header.Len = int(length)
	header.Cap = int(length)
	postMixFunc(slice)
}

var postMixFunc func([]uint8)

func Mix_SetPostMix(mixFunc func([]uint8)) {
	postMixFunc = mixFunc
	C.Mix_SetPostMix((*[0]byte)(C.callPostMixFunction), nil)
}

//export callHookMusic
func callHookMusic(udata unsafe.Pointer, stream *C.Uint8, length C.int) {
	var slice []uint8
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Data = uintptr(unsafe.Pointer(stream))
	header.Len = int(length)
	header.Cap = int(length)
	hookMusicFunc(slice)
}

var hookMusicFunc func([]uint8)

func Mix_HookMusic(musicFunc func([]uint8)) {
	hookMusicFunc = musicFunc
	C.Mix_HookMusic((*[0]byte)(C.callHookMusic), nil)
}

//export callHookMusicFinished
func callHookMusicFinished() {
	hookMusicFinishedFunc()
}

var hookMusicFinishedFunc func()

func Mix_HookMusicFinished(musicFinished func()) {
	hookMusicFinishedFunc = musicFinished
	C.Mix_HookMusicFinished((*[0]byte)(C.callHookMusicFinished))
}

//export callChannelFinished
func callChannelFinished(channel C.int) {
	channelFinishedFunc(int(channel))
}

var channelFinishedFunc func(int)

func Mix_ChannelFinished(channelFinished func(int)) {
	channelFinishedFunc = channelFinished
	C.Mix_ChannelFinished((*[0]byte)(C.callChannelFinished))
}

type EffectFuncT func(channel int, stream []byte)

type EffectDoneT func(channel int)

var allEffectFunc []EffectFuncT
var allEffectDone []EffectDoneT

//export callEffectFunc
func callEffectFunc(channel C.int, stream unsafe.Pointer, length C.int, udata unsafe.Pointer) {
	index := int(uintptr(udata))
	var slice []byte
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Data = uintptr(stream)
	header.Len = int(length)
	header.Cap = int(length)
	allEffectFunc[index](int(channel), slice)
}

//export callEffectDone
func callEffectDone(channel C.int, udata unsafe.Pointer) {
	index := int(uintptr(udata))
	allEffectDone[index](int(channel))
}

func Mix_RegisterEffect(channel int, f EffectFuncT, d EffectDoneT) int {
	// the user data pointer is not required, because go has proper closures
	index := len(allEffectFunc)
	// since go functions can't be cast to C function pointers, we have a workaround here.
	allEffectFunc = append(allEffectFunc, f)
	allEffectDone = append(allEffectDone, d)

	cRet := C.Mix_RegisterEffect(cInt(channel), (*[0]byte)(C.callEffectFunc), (*[0]byte)(C.callEffectDone), unsafe.Pointer(uintptr(index)))
	return int(cRet)
}

func Mix_UnregisterAllEffects(channel int) int {
	// release all effect functions
	allEffectFunc = nil
	allEffectDone = nil

	cRet := C.Mix_UnregisterAllEffects(cInt(channel))
	return int(cRet)
}
