package mix

/*
#include "mixer_wrapper.h"
*/
import "C"

// c basic
type (
	cChar   = C.char    // byte
	cInt    = C.int     // int32
	cUint   = C.uint    // uint32
	cFloat  = C.float   // float32
	cDouble = C.double  // float64

	cInt8   = C.int8_t   // int8
	cInt16  = C.int16_t  // int16
	cInt32  = C.int32_t  // int32
	cInt64  = C.int64_t  // int64
	cIntPtr = C.intptr_t // uintptr

	cUint8   = C.uint8_t   // uint8
	cUint16  = C.uint16_t  // uint16
	cUint32  = C.uint32_t  // uint32
	cUint64  = C.uint64_t  // uint64
	cUintPtr = C.uintptr_t // uintptr

	cSize  = C.size_t  // uint
	cWchar = C.wchar_t // windows uint16, unix int32

	cSchar  = C.schar  // int8
	cUchar  = C.uchar  // uint8
	cShort  = C.short  // int16
	cUshort = C.ushort // uint16

	cLong      = C.long      // int32
	cUlong     = C.ulong     // int32
	cLonglong  = C.longlong  // int64
	cUlonglong = C.ulonglong // uint64
)

// c #define
const (
	SDL_MIXER_MAJOR_VERSION = 2
	SDL_MIXER_MINOR_VERSION = 0
	SDL_MIXER_PATCHLEVEL    = 4

	MIX_MAX_VOLUME = 128 // SDL_MIX_MAXVOLUME

	MIX_CHANNELS = 8

	MIX_CHANNEL_POST = -2

	MIX_EFFECTSMAXSPEED = "MIX_EFFECTSMAXSPEED"
)

// c struct
type Mix_Music C.Mix_Music

const (
	MIX_INIT_FLAC MIX_InitFlags = 1
	MIX_INIT_MOD  MIX_InitFlags = 2
	MIX_INIT_MP3  MIX_InitFlags = 8
	MIX_INIT_OGG  MIX_InitFlags = 16
	MIX_INIT_MID  MIX_InitFlags = 32
	MIX_INIT_OPUS MIX_InitFlags = 64
)

type MIX_InitFlags = int32

type Mix_Chunk struct {
	allocated int32
	abuf      *uint8
	alen      uint32
	volume    uint8
}

const (
	MIX_NO_FADING  Mix_Fading = 0
	MIX_FADING_OUT Mix_Fading = 1
	MIX_FADING_IN  Mix_Fading = 2
)

type Mix_Fading = int32

const (
	MUS_NONE           Mix_MusicType = 0
	MUS_CMD            Mix_MusicType = 1
	MUS_WAV            Mix_MusicType = 2
	MUS_MOD            Mix_MusicType = 3
	MUS_MID            Mix_MusicType = 4
	MUS_OGG            Mix_MusicType = 5
	MUS_MP3            Mix_MusicType = 6
	MUS_MP3_MAD_UNUSED Mix_MusicType = 7
	MUS_FLAC           Mix_MusicType = 8
	MUS_MODPLUG_UNUSED Mix_MusicType = 9
	MUS_OPUS           Mix_MusicType = 10
)

type Mix_MusicType = int32
