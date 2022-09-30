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

func Mix_OpenAudio(frequency int, format uint16, channels, chunksize int) error {
	cFrequency := cInt(frequency)
	cFormat := cUint16(format)
	cChannels := cInt(channels)
	cChunksize := cInt(chunksize)
	if C.Mix_OpenAudio(cFrequency, cFormat, cChannels, cChunksize) < 0 {
		return SDL_GetError()
	}
	return nil
}

func Mix_OpenAudioDevice(frequency int, format uint16, channels, chunksize int, device string, allowedChanges int) error {
	cFrequency := cInt(frequency)
	cFormat := cUint16(format)
	cChannels := cInt(channels)
	cChunksize := cInt(chunksize)
	cAllowedChanges := cInt(allowedChanges)
	cDevice := C.CString(device)
	defer C.free(unsafe.Pointer(cDevice))
	if device == "" {
		// Passing in a device name of NULL requests the most reasonable default
		// (and is equivalent to what SDL_OpenAudio() does to choose a device)
		cDevice = nil
	}
	if C.Mix_OpenAudioDevice(cFrequency, cFormat, cChannels, cChunksize, cDevice, cAllowedChanges) < 0 {
		return SDL_GetError()
	}
	return nil
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

func Mix_LoadWAV_RW(src *SDL_RWops, freesrc SDL_bool) (chunk *Mix_Chunk, err error) {
	cSrc := cSDL_RWops(src)
	cFreesrc := cInt(freesrc)

	cChunk := C.Mix_LoadWAV_RW(cSrc, cFreesrc)
	chunk = (*Mix_Chunk)(unsafe.Pointer(cChunk))
	if chunk == nil {
		err = SDL_GetError()
	}
	return
}

func Mix_LoadWAV(file string) (chunk *Mix_Chunk, err error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	cRB := C.CString("rb")
	defer C.free(unsafe.Pointer(cRB))
	// why doesn't this call Mix_LoadWAV ?
	cChunk := C.Mix_LoadWAV_RW(C.SDL_RWFromFile(cFile, cRB), 1)
	chunk = (*Mix_Chunk)(unsafe.Pointer(cChunk))
	if chunk == nil {
		err = SDL_GetError()
	}
	return
}

func Mix_LoadMUS(file string) (mus *Mix_Music, err error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))

	cMus := C.Mix_LoadMUS(cFile)
	mus = (*Mix_Music)(unsafe.Pointer(cMus))
	if mus == nil {
		err = SDL_GetError()
	}
	return
}

func Mix_LoadMUS_RW(src *SDL_RWops, freesrc int) (mus *Mix_Music, err error) {
	cSrc := cSDL_RWops(src)
	cFreesrc := cInt(freesrc)

	cMus := C.Mix_LoadMUS_RW(cSrc, cFreesrc)
	mus = (*Mix_Music)(unsafe.Pointer(cMus))
	if mus == nil {
		err = SDL_GetError()
	}
	return
}

func Mix_LoadMUSType_RW(src *SDL_RWops, type_ Mix_MusicType, freesrc int) (mus *Mix_Music, err error) {
	cSrc := cSDL_RWops(src)
	cType := (C.Mix_MusicType)(type_)
	cFreesrc := cInt(freesrc)

	cMus := C.Mix_LoadMUSType_RW(cSrc, cType, cFreesrc)
	mus = (*Mix_Music)(unsafe.Pointer(cMus))
	if mus == nil {
		err = SDL_GetError()
	}
	return
}

func Mix_QuickLoad_WAV(mem []byte) (chunk *Mix_Chunk, err error) {
	cMem := (*cUint8)(&mem[0])

	cChunk := C.Mix_QuickLoad_WAV(cMem)
	chunk = (*Mix_Chunk)(unsafe.Pointer(cChunk))
	if chunk == nil {
		err = SDL_GetError()
	}
	return
}

func Mix_QuickLoad_RAW(mem *uint8, len_ uint32) (chunk *Mix_Chunk, err error) {
	cMem := (*cUint8)(mem)
	cLen := cUint32(len_)

	cChunk := C.Mix_QuickLoad_RAW(cMem, cLen)
	chunk = (*Mix_Chunk)(unsafe.Pointer(cChunk))
	if chunk == nil {
		err = SDL_GetError()
	}
	return
}

func (chunk *Mix_Chunk) Mix_FreeChunk() {
	cChunk := cMix_Chunk(chunk)
	C.Mix_FreeChunk(cChunk)
}

func (music *Mix_Music) Mix_FreeMusic() {
	cMusic := cMix_Music(music)
	C.Mix_FreeMusic(cMusic)
}

func (music *Mix_Music) Mix_GetMusicType() Mix_MusicType {
	cMusic := cMix_Music(music)
	return (Mix_MusicType)(C.Mix_GetMusicType(cMusic))
}

func Mix_SetPanning(channel int, left, right uint8) error {
	cChannel := cInt(channel)
	cLeft := cUint8(left)
	cRight := cUint8(right)
	if C.Mix_SetPanning(cChannel, cLeft, cRight) == 0 {
		return SDL_GetError()
	}
	return nil
}

func Mix_SetPosition(channel int, angle int16, distance uint8) error {
	cChannel := cInt(channel)
	cAngle := cInt16(angle)
	cDistance := cUint8(distance)
	if (C.Mix_SetPosition(cChannel, cAngle, cDistance)) == 0 {
		return SDL_GetError()
	}
	return nil
}

func Mix_SetDistance(channel int, distance uint8) error {
	cChannel := cInt(channel)
	cDistance := cUint8(distance)
	if (C.Mix_SetDistance(cChannel, cDistance)) == 0 {
		return SDL_GetError()
	}
	return nil
}

func Mix_SetReverseStereo(channel, flip int) error {
	cChannel := cInt(channel)
	cFlip := cInt(flip)
	if (C.Mix_SetReverseStereo(cChannel, cFlip)) == 0 {
		return SDL_GetError()
	}
	return nil
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

func (chunk *Mix_Chunk) Mix_PlayChannelTimed(channel, loops, ticks int) (channel_ int, err error) {
	cChannel := cInt(channel)
	cChunk := cMix_Chunk(chunk)
	cLoops := cInt(loops)
	cTicks := cInt(ticks)
	channel_ = int(C.Mix_PlayChannelTimed(cChannel, cChunk, cLoops, cTicks))
	if channel_ == -1 {
		err = SDL_GetError()
	}
	return
}

func (chunk *Mix_Chunk) getChunkTimeMilliseconds() int {
	cChunk := cMix_Chunk(chunk)
	return int(C.getChunkTimeMilliseconds(cChunk))
}

func (chunk *Mix_Chunk) Mix_PlayChannel(channel, loops int) (channel_ int, err error) {
	cChannel := cInt(channel)
	cChunk := cMix_Chunk(chunk)
	cLoops := cInt(loops)
	channel_ = int(C.Mix_PlayChannelTimed(cChannel, cChunk, cLoops, -1))
	if channel_ == -1 {
		err = SDL_GetError()
	}
	return
}

func (music *Mix_Music) Mix_PlayMusic(loops int) error {
	cMusic := cMix_Music(music)
	cLoops := cInt(loops)
	if C.Mix_PlayMusic(cMusic, cLoops) == -1 {
		return SDL_GetError()
	}
	return nil
}

func (music *Mix_Music) Mix_FadeInMusic(loops, ms int) error {
	cMusic := cMix_Music(music)
	cLoops := cInt(loops)
	cMs := cInt(ms)
	if C.Mix_FadeInMusic(cMusic, cLoops, cMs) == -1 {
		return SDL_GetError()
	}
	return nil
}

func (music *Mix_Music) Mix_FadeInMusicPos(loops, ms int, position float64) error {
	cMusic := cMix_Music(music)
	cLoops := cInt(loops)
	cMs := cInt(ms)
	cPosition := cDouble(position)
	if C.Mix_FadeInMusicPos(cMusic, cLoops, cMs, cPosition) == -1 {
		return SDL_GetError()
	}
	return nil
}

func (chunk *Mix_Chunk) FadeIn(channel, loops, ms int) (channel_ int, err error) {
	return chunk.Mix_FadeInChannelTimed(channel, loops, ms, -1)
}

func (chunk *Mix_Chunk) Mix_FadeInChannelTimed(channel, loops, ms, ticks int) (channel_ int, err error) {
	cChannel := cInt(channel)
	cChunk := cMix_Chunk(chunk)
	cLoops := cInt(loops)
	cMs := cInt(ms)
	cTicks := cInt(ticks)
	channel_ = int(C.Mix_FadeInChannelTimed(cChannel, cChunk, cLoops, cMs, cTicks))
	if channel_ == -1 {
		err = SDL_GetError()
	}
	return
}

func Mix_Volume(channel, volume int) int {
	cChannel := cInt(channel)
	cVolume := cInt(volume)
	return int(C.Mix_Volume(cChannel, cVolume))
}

func (chunk *Mix_Chunk) Mix_VolumeChunk(volume int) int {
	cChunk := cMix_Chunk(chunk)
	cVolume := cInt(volume)
	return int(C.Mix_VolumeChunk(cChunk, cVolume))
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

func Mix_SetMusicPosition(position int64) error {
	cPosition := cDouble(position)
	if C.Mix_SetMusicPosition(cPosition) == -1 {
		return SDL_GetError()
	}
	return nil
}

func Mix_Playing(channel int) int {
	cChannel := cInt(channel)
	return int(C.Mix_Playing(cChannel))
}

func Mix_PlayingMusic() bool {
	return int(C.Mix_PlayingMusic()) > 0
}

func Mix_SetMusicCMD(command string) error {
	cCommand := C.CString(command)
	defer C.free(unsafe.Pointer(cCommand))
	if C.Mix_SetMusicCMD(cCommand) == -1 {
		return SDL_GetError()
	}
	return nil
}

func Mix_SetSynchroValue(value int) bool {
	_value := cInt(value)
	return int(C.Mix_SetSynchroValue(_value)) == 0
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

func Mix_RegisterEffect(channel int, f EffectFuncT, d EffectDoneT) error {
	// the user data pointer is not required, because go has proper closures
	index := len(allEffectFunc)
	// since go functions can't be cast to C function pointers, we have a workaround here.
	allEffectFunc = append(allEffectFunc, f)
	allEffectDone = append(allEffectDone, d)
	if C.Mix_RegisterEffect(cInt(channel), (*[0]byte)(C.callEffectFunc), (*[0]byte)(C.callEffectDone), unsafe.Pointer(uintptr(index))) == 0 {
		return SDL_GetError()
	}
	return nil
}

func Mix_UnregisterAllEffects(channel int) error {
	// release all effect functions
	allEffectFunc = nil
	allEffectDone = nil
	if C.Mix_UnregisterAllEffects(cInt(channel)) == 0 {
		return SDL_GetError()
	}
	return nil
}
