package sdl

/*
#include "type.h"
*/
import "C"
import (
	"encoding/binary"
	"image/color"
	"unsafe"
)

// c basic
type (
	cBool   = C.char   // int8
	cChar   = C.char   // byte
	cInt    = C.int    // int32
	cUint   = C.uint   // uint32
	cFloat  = C.float  // float32
	cDouble = C.double // float64
	cSize_t = C.size_t // uint

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

	cLong      = C.long      // int32
	cUlong     = C.ulong     // int32
	cLonglong  = C.longlong  // int64
	cUlonglong = C.ulonglong // uint64
)

// c enum
type (
	cSDL_bool           = C.SDL_bool
	cSDL_errorcode      = C.SDL_errorcode
	cSDL_BlendMode      = C.SDL_BlendMode
	cSDL_BlendOperation = C.SDL_BlendOperation
	cSDL_BlendFactor    = C.SDL_BlendFactor
	cSDL_ScaleMode      = C.SDL_ScaleMode
	cSDL_RendererFlip   = C.SDL_RendererFlip
	cSDL_EventType      = C.SDL_EventType
	cSDL_eventaction    = C.SDL_eventaction
)

// c struct
type (
	SDL_BlitMap        C.SDL_BlitMap
	SDL_Cursor         C.SDL_Cursor
	SDL_Renderer       C.SDL_Renderer
	SDL_SysWMmsg       C.SDL_SysWMmsg
	SDL_Texture        C.SDL_Texture
	SDL_Thread         C.SDL_Thread
	SDL_Window         C.SDL_Window
	SDL_cond           C.SDL_cond
	SDL_hid_device     C.SDL_hid_device
	SDL_mutex          C.SDL_mutex
	SDL_semaphore      C.SDL_semaphore
	SDL_AudioStream    C.SDL_AudioStream
	SDL_GameController C.SDL_GameController
	SDL_Haptic         C.SDL_Haptic
	SDL_Joystick       C.SDL_Joystick
	SDL_Sensor         C.SDL_Sensor
	SDL_RWops          C.SDL_RWops
)

// c #define
const (
	SDL_MAJOR_VERSION = 2
	SDL_MINOR_VERSION = 0
	SDL_PATCHLEVEL    = 22

	SDL_WINDOWPOS_UNDEFINED_MASK int32 = 0x1FFF0000
	SDL_WINDOWPOS_UNDEFINED      int32 = SDL_WINDOWPOS_UNDEFINED_MASK | 0
	SDL_WINDOWPOS_CENTERED_MASK  int32 = 0x2FFF0000
	SDL_WINDOWPOS_CENTERED       int32 = SDL_WINDOWPOS_CENTERED_MASK | 0

	SDL_INIT_TIMER          uint32 = 0x00000001
	SDL_INIT_AUDIO          uint32 = 0x00000010
	SDL_INIT_VIDEO          uint32 = 0x00000020
	SDL_INIT_JOYSTICK       uint32 = 0x00000200
	SDL_INIT_HAPTIC         uint32 = 0x00001000
	SDL_INIT_GAMECONTROLLER uint32 = 0x00002000
	SDL_INIT_EVENTS         uint32 = 0x00004000
	SDL_INIT_NOPARACHUTE    uint32 = 0x00008000
	SDL_INIT_SENSOR         uint32 = 0x00100000
	SDL_INIT_EVERYTHING     uint32 = SDL_INIT_TIMER | SDL_INIT_AUDIO | SDL_INIT_VIDEO | SDL_INIT_EVENTS | SDL_INIT_JOYSTICK | SDL_INIT_HAPTIC | SDL_INIT_GAMECONTROLLER | SDL_INIT_SENSOR

	SDL_SWSURFACE    = 0
	SDL_PREALLOC     = 0x00000001
	SDL_RLEACCEL     = 0x00000002
	SDL_DONTFREE     = 0x00000004
	SDL_SIMD_ALIGNED = 0x00000008
)

const (
	SDL_FALSE SDL_bool = 0
	SDL_TRUE  SDL_bool = 1
)

type SDL_bool = int32

const (
	SDL_ENOMEM      SDL_errorcode = 0
	SDL_EFREAD      SDL_errorcode = 1
	SDL_EFWRITE     SDL_errorcode = 2
	SDL_EFSEEK      SDL_errorcode = 3
	SDL_UNSUPPORTED SDL_errorcode = 4
	SDL_LASTERROR   SDL_errorcode = 5
)

type SDL_errorcode = int32

const (
	RW_SEEK_SET = 0
	RW_SEEK_CUR = 1
	RW_SEEK_END = 2
)

type SDL_AudioFormat = uint16

// type SDL_AudioCallback = func(unsafe.Pointer, *uint8, int32)
type SDL_AudioSpec struct {
	Freq     int32
	Format   uint16
	Channels uint8
	Silence  uint8
	Samples  uint16
	Padding  uint16
	Size     uint32
	Callback func(unsafe.Pointer, *uint8, int32)
	Userdata unsafe.Pointer
}

type SDL_AudioFilter = func(*SDL_AudioCVT, uint16)

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

type SDL_AudioDeviceID = uint32

const (
	SDL_AUDIO_STOPPED SDL_AudioStatus = 0
	SDL_AUDIO_PLAYING SDL_AudioStatus = 1
	SDL_AUDIO_PAUSED  SDL_AudioStatus = 2
)

type SDL_AudioStatus = int32

const (
	SDL_PIXELTYPE_UNKNOWN  SDL_PixelType = 0
	SDL_PIXELTYPE_INDEX1   SDL_PixelType = 1
	SDL_PIXELTYPE_INDEX4   SDL_PixelType = 2
	SDL_PIXELTYPE_INDEX8   SDL_PixelType = 3
	SDL_PIXELTYPE_PACKED8  SDL_PixelType = 4
	SDL_PIXELTYPE_PACKED16 SDL_PixelType = 5
	SDL_PIXELTYPE_PACKED32 SDL_PixelType = 6
	SDL_PIXELTYPE_ARRAYU8  SDL_PixelType = 7
	SDL_PIXELTYPE_ARRAYU16 SDL_PixelType = 8
	SDL_PIXELTYPE_ARRAYU32 SDL_PixelType = 9
	SDL_PIXELTYPE_ARRAYF16 SDL_PixelType = 10
	SDL_PIXELTYPE_ARRAYF32 SDL_PixelType = 11
)

type SDL_PixelType = int32

const (
	SDL_BITMAPORDER_NONE SDL_BitmapOrder = 0
	SDL_BITMAPORDER_4321 SDL_BitmapOrder = 1
	SDL_BITMAPORDER_1234 SDL_BitmapOrder = 2
)

type SDL_BitmapOrder = int32

const (
	SDL_PACKEDORDER_NONE SDL_PackedOrder = 0
	SDL_PACKEDORDER_XRGB SDL_PackedOrder = 1
	SDL_PACKEDORDER_RGBX SDL_PackedOrder = 2
	SDL_PACKEDORDER_ARGB SDL_PackedOrder = 3
	SDL_PACKEDORDER_RGBA SDL_PackedOrder = 4
	SDL_PACKEDORDER_XBGR SDL_PackedOrder = 5
	SDL_PACKEDORDER_BGRX SDL_PackedOrder = 6
	SDL_PACKEDORDER_ABGR SDL_PackedOrder = 7
	SDL_PACKEDORDER_BGRA SDL_PackedOrder = 8
)

type SDL_PackedOrder = int32

const (
	SDL_ARRAYORDER_NONE SDL_ArrayOrder = 0
	SDL_ARRAYORDER_RGB  SDL_ArrayOrder = 1
	SDL_ARRAYORDER_RGBA SDL_ArrayOrder = 2
	SDL_ARRAYORDER_ARGB SDL_ArrayOrder = 3
	SDL_ARRAYORDER_BGR  SDL_ArrayOrder = 4
	SDL_ARRAYORDER_BGRA SDL_ArrayOrder = 5
	SDL_ARRAYORDER_ABGR SDL_ArrayOrder = 6
)

type SDL_ArrayOrder = int32

const (
	SDL_PACKEDLAYOUT_NONE    SDL_PackedLayout = 0
	SDL_PACKEDLAYOUT_332     SDL_PackedLayout = 1
	SDL_PACKEDLAYOUT_4444    SDL_PackedLayout = 2
	SDL_PACKEDLAYOUT_1555    SDL_PackedLayout = 3
	SDL_PACKEDLAYOUT_5551    SDL_PackedLayout = 4
	SDL_PACKEDLAYOUT_565     SDL_PackedLayout = 5
	SDL_PACKEDLAYOUT_8888    SDL_PackedLayout = 6
	SDL_PACKEDLAYOUT_2101010 SDL_PackedLayout = 7
	SDL_PACKEDLAYOUT_1010102 SDL_PackedLayout = 8
)

type SDL_PackedLayout = int32

func SDL_DEFINE_PIXELFORMAT(type_ int32, order_ int32, layout_ int32, bits_ int32, bytes_ int32) uint32 {
	var ret = (1 << 28) | ((type_) << 24) | ((order_) << 20) | ((layout_) << 16) | ((bits_) << 8) | ((bytes_) << 0)
	return uint32(ret)
}

func IsLittleEndian() bool {
	var value int32 = 1 // 占4byte 转换成16进制 0x00 00 00 01

	// 大端(16进制): 00 00 00 01
	// 小端(16进制): 01 00 00 00
	pointer := unsafe.Pointer(&value)
	pb := (*byte)(pointer)
	if *pb != 1 {
		return false
	}
	return true
}

func SDL_DEFINE_PIXELFOURCC(a byte, b byte, c byte, d byte) uint32 {
	var endian binary.ByteOrder
	if IsLittleEndian() {
		endian = binary.LittleEndian
	} else {
		endian = binary.BigEndian
	}
	arr := []byte{a, b, c, d}
	return endian.Uint32(arr)
}

func SDL_PIXELFORMATENUM_INIT() {
	if !IsLittleEndian() {
		SDL_PIXELFORMAT_RGBA32 = SDL_PIXELFORMAT_RGBA8888
		SDL_PIXELFORMAT_ARGB32 = SDL_PIXELFORMAT_ARGB8888
		SDL_PIXELFORMAT_BGRA32 = SDL_PIXELFORMAT_BGRA8888
		SDL_PIXELFORMAT_ABGR32 = SDL_PIXELFORMAT_ABGR8888
	}
}

var (
	SDL_PIXELFORMAT_UNKNOWN      SDL_PixelFormatEnum = 0
	SDL_PIXELFORMAT_INDEX1LSB    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX1, SDL_BITMAPORDER_4321, 0, 1, 0)
	SDL_PIXELFORMAT_INDEX1MSB    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX1, SDL_BITMAPORDER_1234, 0, 1, 0)
	SDL_PIXELFORMAT_INDEX4LSB    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX4, SDL_BITMAPORDER_4321, 0, 4, 0)
	SDL_PIXELFORMAT_INDEX4MSB    SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX4, SDL_BITMAPORDER_1234, 0, 4, 0)
	SDL_PIXELFORMAT_INDEX8       SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX8, 0, 0, 8, 1)
	SDL_PIXELFORMAT_RGB332       SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED8, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_332, 8, 1)
	SDL_PIXELFORMAT_XRGB4444     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_4444, 12, 2)
	SDL_PIXELFORMAT_RGB444       SDL_PixelFormatEnum = SDL_PIXELFORMAT_XRGB4444
	SDL_PIXELFORMAT_XBGR4444     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_4444, 12, 2)
	SDL_PIXELFORMAT_BGR444       SDL_PixelFormatEnum = SDL_PIXELFORMAT_XBGR4444
	SDL_PIXELFORMAT_XRGB1555     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_1555, 15, 2)
	SDL_PIXELFORMAT_RGB555       SDL_PixelFormatEnum = SDL_PIXELFORMAT_XRGB1555
	SDL_PIXELFORMAT_XBGR1555     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_1555, 15, 2)
	SDL_PIXELFORMAT_BGR555       SDL_PixelFormatEnum = SDL_PIXELFORMAT_XBGR1555
	SDL_PIXELFORMAT_ARGB4444     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_4444, 16, 2)
	SDL_PIXELFORMAT_RGBA4444     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_RGBA, SDL_PACKEDLAYOUT_4444, 16, 2)
	SDL_PIXELFORMAT_ABGR4444     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ABGR, SDL_PACKEDLAYOUT_4444, 16, 2)
	SDL_PIXELFORMAT_BGRA4444     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_BGRA, SDL_PACKEDLAYOUT_4444, 16, 2)
	SDL_PIXELFORMAT_ARGB1555     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_1555, 16, 2)
	SDL_PIXELFORMAT_RGBA5551     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_RGBA, SDL_PACKEDLAYOUT_5551, 16, 2)
	SDL_PIXELFORMAT_ABGR1555     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ABGR, SDL_PACKEDLAYOUT_1555, 16, 2)
	SDL_PIXELFORMAT_BGRA5551     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_BGRA, SDL_PACKEDLAYOUT_5551, 16, 2)
	SDL_PIXELFORMAT_RGB565       SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_565, 16, 2)
	SDL_PIXELFORMAT_BGR565       SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_565, 16, 2)
	SDL_PIXELFORMAT_RGB24        SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU8, SDL_ARRAYORDER_RGB, 0, 24, 3)
	SDL_PIXELFORMAT_BGR24        SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU8, SDL_ARRAYORDER_BGR, 0, 24, 3)
	SDL_PIXELFORMAT_XRGB8888     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_XRGB, SDL_PACKEDLAYOUT_8888, 24, 4)
	SDL_PIXELFORMAT_RGB888       SDL_PixelFormatEnum = SDL_PIXELFORMAT_XRGB8888
	SDL_PIXELFORMAT_RGBX8888     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_RGBX, SDL_PACKEDLAYOUT_8888, 24, 4)
	SDL_PIXELFORMAT_XBGR8888     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_XBGR, SDL_PACKEDLAYOUT_8888, 24, 4)
	SDL_PIXELFORMAT_BGR888       SDL_PixelFormatEnum = SDL_PIXELFORMAT_XBGR8888
	SDL_PIXELFORMAT_BGRX8888     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_BGRX, SDL_PACKEDLAYOUT_8888, 24, 4)
	SDL_PIXELFORMAT_ARGB8888     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_8888, 32, 4)
	SDL_PIXELFORMAT_RGBA8888     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_RGBA, SDL_PACKEDLAYOUT_8888, 32, 4)
	SDL_PIXELFORMAT_ABGR8888     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ABGR, SDL_PACKEDLAYOUT_8888, 32, 4)
	SDL_PIXELFORMAT_BGRA8888     SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_BGRA, SDL_PACKEDLAYOUT_8888, 32, 4)
	SDL_PIXELFORMAT_ARGB2101010  SDL_PixelFormatEnum = SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ARGB, SDL_PACKEDLAYOUT_2101010, 32, 4)
	SDL_PIXELFORMAT_RGBA32       SDL_PixelFormatEnum = SDL_PIXELFORMAT_ABGR8888
	SDL_PIXELFORMAT_ARGB32       SDL_PixelFormatEnum = SDL_PIXELFORMAT_BGRA8888
	SDL_PIXELFORMAT_BGRA32       SDL_PixelFormatEnum = SDL_PIXELFORMAT_ARGB8888
	SDL_PIXELFORMAT_ABGR32       SDL_PixelFormatEnum = SDL_PIXELFORMAT_RGBA8888
	SDL_PIXELFORMAT_YV12         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('Y', 'V', '1', '2')
	SDL_PIXELFORMAT_IYUV         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('I', 'Y', 'U', 'V')
	SDL_PIXELFORMAT_YUY2         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('Y', 'U', 'Y', '2')
	SDL_PIXELFORMAT_UYVY         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('U', 'Y', 'V', 'Y')
	SDL_PIXELFORMAT_YVYU         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('Y', 'V', 'Y', 'U')
	SDL_PIXELFORMAT_NV12         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('N', 'V', '1', '2')
	SDL_PIXELFORMAT_NV21         SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('N', 'V', '2', '1')
	SDL_PIXELFORMAT_EXTERNAL_OES SDL_PixelFormatEnum = SDL_DEFINE_PIXELFOURCC('O', 'E', 'S', ' ')
)

type SDL_PixelFormatEnum = uint32

// type SDL_Color struct {
// 	R uint8
// 	G uint8
// 	B uint8
// 	A uint8
// }

// Color represents a color. This implements image/color.Color interface.
type SDL_Color color.RGBA

type SDL_Palette struct {
	Ncolors  int32
	Colors   *SDL_Color
	Version  uint32
	Refcount int32
}

type SDL_PixelFormat struct {
	Format        uint32
	Palette       *SDL_Palette
	BitsPerPixel  uint8
	BytesPerPixel uint8
	Padding       [2]uint8
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	Refcount      int32
	Next          *SDL_PixelFormat
}

type SDL_Point struct {
	X int32
	Y int32
}

type SDL_FPoint struct {
	X float32
	Y float32
}

type SDL_Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

type SDL_FRect struct {
	X float32
	Y float32
	W float32
	H float32
}

const (
	SDL_BLENDMODE_NONE    SDL_BlendMode = 0
	SDL_BLENDMODE_BLEND   SDL_BlendMode = 1
	SDL_BLENDMODE_ADD     SDL_BlendMode = 2
	SDL_BLENDMODE_MOD     SDL_BlendMode = 4
	SDL_BLENDMODE_MUL     SDL_BlendMode = 8
	SDL_BLENDMODE_INVALID SDL_BlendMode = 0x7FFFFFFF
)

type SDL_BlendMode = int32

const (
	SDL_BLENDOPERATION_ADD          SDL_BlendOperation = 1
	SDL_BLENDOPERATION_SUBTRACT     SDL_BlendOperation = 2
	SDL_BLENDOPERATION_REV_SUBTRACT SDL_BlendOperation = 3
	SDL_BLENDOPERATION_MINIMUM      SDL_BlendOperation = 4
	SDL_BLENDOPERATION_MAXIMUM      SDL_BlendOperation = 5
)

type SDL_BlendOperation = int32

const (
	SDL_BLENDFACTOR_ZERO                SDL_BlendFactor = 1
	SDL_BLENDFACTOR_ONE                 SDL_BlendFactor = 2
	SDL_BLENDFACTOR_SRC_COLOR           SDL_BlendFactor = 3
	SDL_BLENDFACTOR_ONE_MINUS_SRC_COLOR SDL_BlendFactor = 4
	SDL_BLENDFACTOR_SRC_ALPHA           SDL_BlendFactor = 5
	SDL_BLENDFACTOR_ONE_MINUS_SRC_ALPHA SDL_BlendFactor = 6
	SDL_BLENDFACTOR_DST_COLOR           SDL_BlendFactor = 7
	SDL_BLENDFACTOR_ONE_MINUS_DST_COLOR SDL_BlendFactor = 8
	SDL_BLENDFACTOR_DST_ALPHA           SDL_BlendFactor = 9
	SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA SDL_BlendFactor = 10
)

type SDL_BlendFactor = int32

type SDL_Surface struct {
	Flags       uint32
	Format      *SDL_PixelFormat
	W           int32
	H           int32
	Pitch       int32
	Pixels      unsafe.Pointer
	Userdata    unsafe.Pointer
	Locked      int32
	ListBlitmap unsafe.Pointer
	ClipRect    SDL_Rect
	Map         *SDL_BlitMap
	Refcount    int32
}

type SDL_blit = func(*SDL_Surface, *SDL_Rect, *SDL_Surface, *SDL_Rect) int32

const (
	SDL_YUV_CONVERSION_JPEG      SDL_YUV_CONVERSION_MODE = 0
	SDL_YUV_CONVERSION_BT601     SDL_YUV_CONVERSION_MODE = 1
	SDL_YUV_CONVERSION_BT709     SDL_YUV_CONVERSION_MODE = 2
	SDL_YUV_CONVERSION_AUTOMATIC SDL_YUV_CONVERSION_MODE = 3
)

type SDL_YUV_CONVERSION_MODE = int32
type SDL_DisplayMode struct {
	Format      uint32
	W           int32
	H           int32
	RefreshRate int32
	Driverdata  unsafe.Pointer
}

const (
	SDL_WINDOW_FULLSCREEN         SDL_WindowFlags = 0x00000001
	SDL_WINDOW_OPENGL             SDL_WindowFlags = 0x00000002
	SDL_WINDOW_SHOWN              SDL_WindowFlags = 0x00000004
	SDL_WINDOW_HIDDEN             SDL_WindowFlags = 0x00000008
	SDL_WINDOW_BORDERLESS         SDL_WindowFlags = 0x00000010
	SDL_WINDOW_RESIZABLE          SDL_WindowFlags = 0x00000020
	SDL_WINDOW_MINIMIZED          SDL_WindowFlags = 0x00000040
	SDL_WINDOW_MAXIMIZED          SDL_WindowFlags = 0x00000080
	SDL_WINDOW_MOUSE_GRABBED      SDL_WindowFlags = 0x00000100
	SDL_WINDOW_INPUT_FOCUS        SDL_WindowFlags = 0x00000200
	SDL_WINDOW_MOUSE_FOCUS        SDL_WindowFlags = 0x00000400
	SDL_WINDOW_FULLSCREEN_DESKTOP SDL_WindowFlags = SDL_WINDOW_FULLSCREEN | 0x00001000
	SDL_WINDOW_FOREIGN            SDL_WindowFlags = 0x00000800
	SDL_WINDOW_ALLOW_HIGHDPI      SDL_WindowFlags = 0x00002000
	SDL_WINDOW_MOUSE_CAPTURE      SDL_WindowFlags = 0x00004000
	SDL_WINDOW_ALWAYS_ON_TOP      SDL_WindowFlags = 0x00008000
	SDL_WINDOW_SKIP_TASKBAR       SDL_WindowFlags = 0x00010000
	SDL_WINDOW_UTILITY            SDL_WindowFlags = 0x00020000
	SDL_WINDOW_TOOLTIP            SDL_WindowFlags = 0x00040000
	SDL_WINDOW_POPUP_MENU         SDL_WindowFlags = 0x00080000
	SDL_WINDOW_KEYBOARD_GRABBED   SDL_WindowFlags = 0x00100000
	SDL_WINDOW_VULKAN             SDL_WindowFlags = 0x10000000
	SDL_WINDOW_METAL              SDL_WindowFlags = 0x20000000
	SDL_WINDOW_INPUT_GRABBED      SDL_WindowFlags = SDL_WINDOW_MOUSE_GRABBED
)

type SDL_WindowFlags = uint32

const (
	SDL_WINDOWEVENT_NONE            SDL_WindowEventID = 0
	SDL_WINDOWEVENT_SHOWN           SDL_WindowEventID = 1
	SDL_WINDOWEVENT_HIDDEN          SDL_WindowEventID = 2
	SDL_WINDOWEVENT_EXPOSED         SDL_WindowEventID = 3
	SDL_WINDOWEVENT_MOVED           SDL_WindowEventID = 4
	SDL_WINDOWEVENT_RESIZED         SDL_WindowEventID = 5
	SDL_WINDOWEVENT_SIZE_CHANGED    SDL_WindowEventID = 6
	SDL_WINDOWEVENT_MINIMIZED       SDL_WindowEventID = 7
	SDL_WINDOWEVENT_MAXIMIZED       SDL_WindowEventID = 8
	SDL_WINDOWEVENT_RESTORED        SDL_WindowEventID = 9
	SDL_WINDOWEVENT_ENTER           SDL_WindowEventID = 10
	SDL_WINDOWEVENT_LEAVE           SDL_WindowEventID = 11
	SDL_WINDOWEVENT_FOCUS_GAINED    SDL_WindowEventID = 12
	SDL_WINDOWEVENT_FOCUS_LOST      SDL_WindowEventID = 13
	SDL_WINDOWEVENT_CLOSE           SDL_WindowEventID = 14
	SDL_WINDOWEVENT_TAKE_FOCUS      SDL_WindowEventID = 15
	SDL_WINDOWEVENT_HIT_TEST        SDL_WindowEventID = 16
	SDL_WINDOWEVENT_ICCPROF_CHANGED SDL_WindowEventID = 17
	SDL_WINDOWEVENT_DISPLAY_CHANGED SDL_WindowEventID = 18
)

type SDL_WindowEventID = int32

const (
	SDL_DISPLAYEVENT_NONE         SDL_DisplayEventID = 0
	SDL_DISPLAYEVENT_ORIENTATION  SDL_DisplayEventID = 1
	SDL_DISPLAYEVENT_CONNECTED    SDL_DisplayEventID = 2
	SDL_DISPLAYEVENT_DISCONNECTED SDL_DisplayEventID = 3
)

type SDL_DisplayEventID = int32

const (
	SDL_ORIENTATION_UNKNOWN           SDL_DisplayOrientation = 0
	SDL_ORIENTATION_LANDSCAPE         SDL_DisplayOrientation = 1
	SDL_ORIENTATION_LANDSCAPE_FLIPPED SDL_DisplayOrientation = 2
	SDL_ORIENTATION_PORTRAIT          SDL_DisplayOrientation = 3
	SDL_ORIENTATION_PORTRAIT_FLIPPED  SDL_DisplayOrientation = 4
)

type SDL_DisplayOrientation = int32

const (
	SDL_FLASH_CANCEL        SDL_FlashOperation = 0
	SDL_FLASH_BRIEFLY       SDL_FlashOperation = 1
	SDL_FLASH_UNTIL_FOCUSED SDL_FlashOperation = 2
)

type SDL_FlashOperation = int32

type SDL_GLContext = unsafe.Pointer

const (
	SDL_GL_RED_SIZE                   SDL_GLattr = 0
	SDL_GL_GREEN_SIZE                 SDL_GLattr = 1
	SDL_GL_BLUE_SIZE                  SDL_GLattr = 2
	SDL_GL_ALPHA_SIZE                 SDL_GLattr = 3
	SDL_GL_BUFFER_SIZE                SDL_GLattr = 4
	SDL_GL_DOUBLEBUFFER               SDL_GLattr = 5
	SDL_GL_DEPTH_SIZE                 SDL_GLattr = 6
	SDL_GL_STENCIL_SIZE               SDL_GLattr = 7
	SDL_GL_ACCUM_RED_SIZE             SDL_GLattr = 8
	SDL_GL_ACCUM_GREEN_SIZE           SDL_GLattr = 9
	SDL_GL_ACCUM_BLUE_SIZE            SDL_GLattr = 10
	SDL_GL_ACCUM_ALPHA_SIZE           SDL_GLattr = 11
	SDL_GL_STEREO                     SDL_GLattr = 12
	SDL_GL_MULTISAMPLEBUFFERS         SDL_GLattr = 13
	SDL_GL_MULTISAMPLESAMPLES         SDL_GLattr = 14
	SDL_GL_ACCELERATED_VISUAL         SDL_GLattr = 15
	SDL_GL_RETAINED_BACKING           SDL_GLattr = 16
	SDL_GL_CONTEXT_MAJOR_VERSION      SDL_GLattr = 17
	SDL_GL_CONTEXT_MINOR_VERSION      SDL_GLattr = 18
	SDL_GL_CONTEXT_EGL                SDL_GLattr = 19
	SDL_GL_CONTEXT_FLAGS              SDL_GLattr = 20
	SDL_GL_CONTEXT_PROFILE_MASK       SDL_GLattr = 21
	SDL_GL_SHARE_WITH_CURRENT_CONTEXT SDL_GLattr = 22
	SDL_GL_FRAMEBUFFER_SRGB_CAPABLE   SDL_GLattr = 23
	SDL_GL_CONTEXT_RELEASE_BEHAVIOR   SDL_GLattr = 24
	SDL_GL_CONTEXT_RESET_NOTIFICATION SDL_GLattr = 25
	SDL_GL_CONTEXT_NO_ERROR           SDL_GLattr = 26
	SDL_GL_FLOATBUFFERS               SDL_GLattr = 27
)

type SDL_GLattr = int32

const (
	SDL_GL_CONTEXT_PROFILE_CORE          SDL_GLprofile = 1
	SDL_GL_CONTEXT_PROFILE_COMPATIBILITY SDL_GLprofile = 2
	SDL_GL_CONTEXT_PROFILE_ES            SDL_GLprofile = 4
)

type SDL_GLprofile = int32

const (
	SDL_GL_CONTEXT_DEBUG_FLAG              SDL_GLcontextFlag = 1
	SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG SDL_GLcontextFlag = 2
	SDL_GL_CONTEXT_ROBUST_ACCESS_FLAG      SDL_GLcontextFlag = 4
	SDL_GL_CONTEXT_RESET_ISOLATION_FLAG    SDL_GLcontextFlag = 8
)

type SDL_GLcontextFlag = int32

const (
	SDL_GL_CONTEXT_RELEASE_BEHAVIOR_NONE  SDL_GLcontextReleaseFlag = 0
	SDL_GL_CONTEXT_RELEASE_BEHAVIOR_FLUSH SDL_GLcontextReleaseFlag = 1
)

type SDL_GLcontextReleaseFlag = int32

const (
	SDL_GL_CONTEXT_RESET_NO_NOTIFICATION SDL_GLContextResetNotification = 0
	SDL_GL_CONTEXT_RESET_LOSE_CONTEXT    SDL_GLContextResetNotification = 1
)

type SDL_GLContextResetNotification = int32

const (
	SDL_HITTEST_NORMAL             SDL_HitTestResult = 0
	SDL_HITTEST_DRAGGABLE          SDL_HitTestResult = 1
	SDL_HITTEST_RESIZE_TOPLEFT     SDL_HitTestResult = 2
	SDL_HITTEST_RESIZE_TOP         SDL_HitTestResult = 3
	SDL_HITTEST_RESIZE_TOPRIGHT    SDL_HitTestResult = 4
	SDL_HITTEST_RESIZE_RIGHT       SDL_HitTestResult = 5
	SDL_HITTEST_RESIZE_BOTTOMRIGHT SDL_HitTestResult = 6
	SDL_HITTEST_RESIZE_BOTTOM      SDL_HitTestResult = 7
	SDL_HITTEST_RESIZE_BOTTOMLEFT  SDL_HitTestResult = 8
	SDL_HITTEST_RESIZE_LEFT        SDL_HitTestResult = 9
)

type SDL_HitTestResult = int32
type SDL_HitTest = func(*SDL_Window, *SDL_Point, unsafe.Pointer) int32

const (
	SDL_SCANCODE_UNKNOWN            SDL_Scancode = 0
	SDL_SCANCODE_A                  SDL_Scancode = 4
	SDL_SCANCODE_B                  SDL_Scancode = 5
	SDL_SCANCODE_C                  SDL_Scancode = 6
	SDL_SCANCODE_D                  SDL_Scancode = 7
	SDL_SCANCODE_E                  SDL_Scancode = 8
	SDL_SCANCODE_F                  SDL_Scancode = 9
	SDL_SCANCODE_G                  SDL_Scancode = 10
	SDL_SCANCODE_H                  SDL_Scancode = 11
	SDL_SCANCODE_I                  SDL_Scancode = 12
	SDL_SCANCODE_J                  SDL_Scancode = 13
	SDL_SCANCODE_K                  SDL_Scancode = 14
	SDL_SCANCODE_L                  SDL_Scancode = 15
	SDL_SCANCODE_M                  SDL_Scancode = 16
	SDL_SCANCODE_N                  SDL_Scancode = 17
	SDL_SCANCODE_O                  SDL_Scancode = 18
	SDL_SCANCODE_P                  SDL_Scancode = 19
	SDL_SCANCODE_Q                  SDL_Scancode = 20
	SDL_SCANCODE_R                  SDL_Scancode = 21
	SDL_SCANCODE_S                  SDL_Scancode = 22
	SDL_SCANCODE_T                  SDL_Scancode = 23
	SDL_SCANCODE_U                  SDL_Scancode = 24
	SDL_SCANCODE_V                  SDL_Scancode = 25
	SDL_SCANCODE_W                  SDL_Scancode = 26
	SDL_SCANCODE_X                  SDL_Scancode = 27
	SDL_SCANCODE_Y                  SDL_Scancode = 28
	SDL_SCANCODE_Z                  SDL_Scancode = 29
	SDL_SCANCODE_1                  SDL_Scancode = 30
	SDL_SCANCODE_2                  SDL_Scancode = 31
	SDL_SCANCODE_3                  SDL_Scancode = 32
	SDL_SCANCODE_4                  SDL_Scancode = 33
	SDL_SCANCODE_5                  SDL_Scancode = 34
	SDL_SCANCODE_6                  SDL_Scancode = 35
	SDL_SCANCODE_7                  SDL_Scancode = 36
	SDL_SCANCODE_8                  SDL_Scancode = 37
	SDL_SCANCODE_9                  SDL_Scancode = 38
	SDL_SCANCODE_0                  SDL_Scancode = 39
	SDL_SCANCODE_RETURN             SDL_Scancode = 40
	SDL_SCANCODE_ESCAPE             SDL_Scancode = 41
	SDL_SCANCODE_BACKSPACE          SDL_Scancode = 42
	SDL_SCANCODE_TAB                SDL_Scancode = 43
	SDL_SCANCODE_SPACE              SDL_Scancode = 44
	SDL_SCANCODE_MINUS              SDL_Scancode = 45
	SDL_SCANCODE_EQUALS             SDL_Scancode = 46
	SDL_SCANCODE_LEFTBRACKET        SDL_Scancode = 47
	SDL_SCANCODE_RIGHTBRACKET       SDL_Scancode = 48
	SDL_SCANCODE_BACKSLASH          SDL_Scancode = 49
	SDL_SCANCODE_NONUSHASH          SDL_Scancode = 50
	SDL_SCANCODE_SEMICOLON          SDL_Scancode = 51
	SDL_SCANCODE_APOSTROPHE         SDL_Scancode = 52
	SDL_SCANCODE_GRAVE              SDL_Scancode = 53
	SDL_SCANCODE_COMMA              SDL_Scancode = 54
	SDL_SCANCODE_PERIOD             SDL_Scancode = 55
	SDL_SCANCODE_SLASH              SDL_Scancode = 56
	SDL_SCANCODE_CAPSLOCK           SDL_Scancode = 57
	SDL_SCANCODE_F1                 SDL_Scancode = 58
	SDL_SCANCODE_F2                 SDL_Scancode = 59
	SDL_SCANCODE_F3                 SDL_Scancode = 60
	SDL_SCANCODE_F4                 SDL_Scancode = 61
	SDL_SCANCODE_F5                 SDL_Scancode = 62
	SDL_SCANCODE_F6                 SDL_Scancode = 63
	SDL_SCANCODE_F7                 SDL_Scancode = 64
	SDL_SCANCODE_F8                 SDL_Scancode = 65
	SDL_SCANCODE_F9                 SDL_Scancode = 66
	SDL_SCANCODE_F10                SDL_Scancode = 67
	SDL_SCANCODE_F11                SDL_Scancode = 68
	SDL_SCANCODE_F12                SDL_Scancode = 69
	SDL_SCANCODE_PRINTSCREEN        SDL_Scancode = 70
	SDL_SCANCODE_SCROLLLOCK         SDL_Scancode = 71
	SDL_SCANCODE_PAUSE              SDL_Scancode = 72
	SDL_SCANCODE_INSERT             SDL_Scancode = 73
	SDL_SCANCODE_HOME               SDL_Scancode = 74
	SDL_SCANCODE_PAGEUP             SDL_Scancode = 75
	SDL_SCANCODE_DELETE             SDL_Scancode = 76
	SDL_SCANCODE_END                SDL_Scancode = 77
	SDL_SCANCODE_PAGEDOWN           SDL_Scancode = 78
	SDL_SCANCODE_RIGHT              SDL_Scancode = 79
	SDL_SCANCODE_LEFT               SDL_Scancode = 80
	SDL_SCANCODE_DOWN               SDL_Scancode = 81
	SDL_SCANCODE_UP                 SDL_Scancode = 82
	SDL_SCANCODE_NUMLOCKCLEAR       SDL_Scancode = 83
	SDL_SCANCODE_KP_DIVIDE          SDL_Scancode = 84
	SDL_SCANCODE_KP_MULTIPLY        SDL_Scancode = 85
	SDL_SCANCODE_KP_MINUS           SDL_Scancode = 86
	SDL_SCANCODE_KP_PLUS            SDL_Scancode = 87
	SDL_SCANCODE_KP_ENTER           SDL_Scancode = 88
	SDL_SCANCODE_KP_1               SDL_Scancode = 89
	SDL_SCANCODE_KP_2               SDL_Scancode = 90
	SDL_SCANCODE_KP_3               SDL_Scancode = 91
	SDL_SCANCODE_KP_4               SDL_Scancode = 92
	SDL_SCANCODE_KP_5               SDL_Scancode = 93
	SDL_SCANCODE_KP_6               SDL_Scancode = 94
	SDL_SCANCODE_KP_7               SDL_Scancode = 95
	SDL_SCANCODE_KP_8               SDL_Scancode = 96
	SDL_SCANCODE_KP_9               SDL_Scancode = 97
	SDL_SCANCODE_KP_0               SDL_Scancode = 98
	SDL_SCANCODE_KP_PERIOD          SDL_Scancode = 99
	SDL_SCANCODE_NONUSBACKSLASH     SDL_Scancode = 100
	SDL_SCANCODE_APPLICATION        SDL_Scancode = 101
	SDL_SCANCODE_POWER              SDL_Scancode = 102
	SDL_SCANCODE_KP_EQUALS          SDL_Scancode = 103
	SDL_SCANCODE_F13                SDL_Scancode = 104
	SDL_SCANCODE_F14                SDL_Scancode = 105
	SDL_SCANCODE_F15                SDL_Scancode = 106
	SDL_SCANCODE_F16                SDL_Scancode = 107
	SDL_SCANCODE_F17                SDL_Scancode = 108
	SDL_SCANCODE_F18                SDL_Scancode = 109
	SDL_SCANCODE_F19                SDL_Scancode = 110
	SDL_SCANCODE_F20                SDL_Scancode = 111
	SDL_SCANCODE_F21                SDL_Scancode = 112
	SDL_SCANCODE_F22                SDL_Scancode = 113
	SDL_SCANCODE_F23                SDL_Scancode = 114
	SDL_SCANCODE_F24                SDL_Scancode = 115
	SDL_SCANCODE_EXECUTE            SDL_Scancode = 116
	SDL_SCANCODE_HELP               SDL_Scancode = 117
	SDL_SCANCODE_MENU               SDL_Scancode = 118
	SDL_SCANCODE_SELECT             SDL_Scancode = 119
	SDL_SCANCODE_STOP               SDL_Scancode = 120
	SDL_SCANCODE_AGAIN              SDL_Scancode = 121
	SDL_SCANCODE_UNDO               SDL_Scancode = 122
	SDL_SCANCODE_CUT                SDL_Scancode = 123
	SDL_SCANCODE_COPY               SDL_Scancode = 124
	SDL_SCANCODE_PASTE              SDL_Scancode = 125
	SDL_SCANCODE_FIND               SDL_Scancode = 126
	SDL_SCANCODE_MUTE               SDL_Scancode = 127
	SDL_SCANCODE_VOLUMEUP           SDL_Scancode = 128
	SDL_SCANCODE_VOLUMEDOWN         SDL_Scancode = 129
	SDL_SCANCODE_KP_COMMA           SDL_Scancode = 133
	SDL_SCANCODE_KP_EQUALSAS400     SDL_Scancode = 134
	SDL_SCANCODE_INTERNATIONAL1     SDL_Scancode = 135
	SDL_SCANCODE_INTERNATIONAL2     SDL_Scancode = 136
	SDL_SCANCODE_INTERNATIONAL3     SDL_Scancode = 137
	SDL_SCANCODE_INTERNATIONAL4     SDL_Scancode = 138
	SDL_SCANCODE_INTERNATIONAL5     SDL_Scancode = 139
	SDL_SCANCODE_INTERNATIONAL6     SDL_Scancode = 140
	SDL_SCANCODE_INTERNATIONAL7     SDL_Scancode = 141
	SDL_SCANCODE_INTERNATIONAL8     SDL_Scancode = 142
	SDL_SCANCODE_INTERNATIONAL9     SDL_Scancode = 143
	SDL_SCANCODE_LANG1              SDL_Scancode = 144
	SDL_SCANCODE_LANG2              SDL_Scancode = 145
	SDL_SCANCODE_LANG3              SDL_Scancode = 146
	SDL_SCANCODE_LANG4              SDL_Scancode = 147
	SDL_SCANCODE_LANG5              SDL_Scancode = 148
	SDL_SCANCODE_LANG6              SDL_Scancode = 149
	SDL_SCANCODE_LANG7              SDL_Scancode = 150
	SDL_SCANCODE_LANG8              SDL_Scancode = 151
	SDL_SCANCODE_LANG9              SDL_Scancode = 152
	SDL_SCANCODE_ALTERASE           SDL_Scancode = 153
	SDL_SCANCODE_SYSREQ             SDL_Scancode = 154
	SDL_SCANCODE_CANCEL             SDL_Scancode = 155
	SDL_SCANCODE_CLEAR              SDL_Scancode = 156
	SDL_SCANCODE_PRIOR              SDL_Scancode = 157
	SDL_SCANCODE_RETURN2            SDL_Scancode = 158
	SDL_SCANCODE_SEPARATOR          SDL_Scancode = 159
	SDL_SCANCODE_OUT                SDL_Scancode = 160
	SDL_SCANCODE_OPER               SDL_Scancode = 161
	SDL_SCANCODE_CLEARAGAIN         SDL_Scancode = 162
	SDL_SCANCODE_CRSEL              SDL_Scancode = 163
	SDL_SCANCODE_EXSEL              SDL_Scancode = 164
	SDL_SCANCODE_KP_00              SDL_Scancode = 176
	SDL_SCANCODE_KP_000             SDL_Scancode = 177
	SDL_SCANCODE_THOUSANDSSEPARATOR SDL_Scancode = 178
	SDL_SCANCODE_DECIMALSEPARATOR   SDL_Scancode = 179
	SDL_SCANCODE_CURRENCYUNIT       SDL_Scancode = 180
	SDL_SCANCODE_CURRENCYSUBUNIT    SDL_Scancode = 181
	SDL_SCANCODE_KP_LEFTPAREN       SDL_Scancode = 182
	SDL_SCANCODE_KP_RIGHTPAREN      SDL_Scancode = 183
	SDL_SCANCODE_KP_LEFTBRACE       SDL_Scancode = 184
	SDL_SCANCODE_KP_RIGHTBRACE      SDL_Scancode = 185
	SDL_SCANCODE_KP_TAB             SDL_Scancode = 186
	SDL_SCANCODE_KP_BACKSPACE       SDL_Scancode = 187
	SDL_SCANCODE_KP_A               SDL_Scancode = 188
	SDL_SCANCODE_KP_B               SDL_Scancode = 189
	SDL_SCANCODE_KP_C               SDL_Scancode = 190
	SDL_SCANCODE_KP_D               SDL_Scancode = 191
	SDL_SCANCODE_KP_E               SDL_Scancode = 192
	SDL_SCANCODE_KP_F               SDL_Scancode = 193
	SDL_SCANCODE_KP_XOR             SDL_Scancode = 194
	SDL_SCANCODE_KP_POWER           SDL_Scancode = 195
	SDL_SCANCODE_KP_PERCENT         SDL_Scancode = 196
	SDL_SCANCODE_KP_LESS            SDL_Scancode = 197
	SDL_SCANCODE_KP_GREATER         SDL_Scancode = 198
	SDL_SCANCODE_KP_AMPERSAND       SDL_Scancode = 199
	SDL_SCANCODE_KP_DBLAMPERSAND    SDL_Scancode = 200
	SDL_SCANCODE_KP_VERTICALBAR     SDL_Scancode = 201
	SDL_SCANCODE_KP_DBLVERTICALBAR  SDL_Scancode = 202
	SDL_SCANCODE_KP_COLON           SDL_Scancode = 203
	SDL_SCANCODE_KP_HASH            SDL_Scancode = 204
	SDL_SCANCODE_KP_SPACE           SDL_Scancode = 205
	SDL_SCANCODE_KP_AT              SDL_Scancode = 206
	SDL_SCANCODE_KP_EXCLAM          SDL_Scancode = 207
	SDL_SCANCODE_KP_MEMSTORE        SDL_Scancode = 208
	SDL_SCANCODE_KP_MEMRECALL       SDL_Scancode = 209
	SDL_SCANCODE_KP_MEMCLEAR        SDL_Scancode = 210
	SDL_SCANCODE_KP_MEMADD          SDL_Scancode = 211
	SDL_SCANCODE_KP_MEMSUBTRACT     SDL_Scancode = 212
	SDL_SCANCODE_KP_MEMMULTIPLY     SDL_Scancode = 213
	SDL_SCANCODE_KP_MEMDIVIDE       SDL_Scancode = 214
	SDL_SCANCODE_KP_PLUSMINUS       SDL_Scancode = 215
	SDL_SCANCODE_KP_CLEAR           SDL_Scancode = 216
	SDL_SCANCODE_KP_CLEARENTRY      SDL_Scancode = 217
	SDL_SCANCODE_KP_BINARY          SDL_Scancode = 218
	SDL_SCANCODE_KP_OCTAL           SDL_Scancode = 219
	SDL_SCANCODE_KP_DECIMAL         SDL_Scancode = 220
	SDL_SCANCODE_KP_HEXADECIMAL     SDL_Scancode = 221
	SDL_SCANCODE_LCTRL              SDL_Scancode = 224
	SDL_SCANCODE_LSHIFT             SDL_Scancode = 225
	SDL_SCANCODE_LALT               SDL_Scancode = 226
	SDL_SCANCODE_LGUI               SDL_Scancode = 227
	SDL_SCANCODE_RCTRL              SDL_Scancode = 228
	SDL_SCANCODE_RSHIFT             SDL_Scancode = 229
	SDL_SCANCODE_RALT               SDL_Scancode = 230
	SDL_SCANCODE_RGUI               SDL_Scancode = 231
	SDL_SCANCODE_MODE               SDL_Scancode = 257
	SDL_SCANCODE_AUDIONEXT          SDL_Scancode = 258
	SDL_SCANCODE_AUDIOPREV          SDL_Scancode = 259
	SDL_SCANCODE_AUDIOSTOP          SDL_Scancode = 260
	SDL_SCANCODE_AUDIOPLAY          SDL_Scancode = 261
	SDL_SCANCODE_AUDIOMUTE          SDL_Scancode = 262
	SDL_SCANCODE_MEDIASELECT        SDL_Scancode = 263
	SDL_SCANCODE_WWW                SDL_Scancode = 264
	SDL_SCANCODE_MAIL               SDL_Scancode = 265
	SDL_SCANCODE_CALCULATOR         SDL_Scancode = 266
	SDL_SCANCODE_COMPUTER           SDL_Scancode = 267
	SDL_SCANCODE_AC_SEARCH          SDL_Scancode = 268
	SDL_SCANCODE_AC_HOME            SDL_Scancode = 269
	SDL_SCANCODE_AC_BACK            SDL_Scancode = 270
	SDL_SCANCODE_AC_FORWARD         SDL_Scancode = 271
	SDL_SCANCODE_AC_STOP            SDL_Scancode = 272
	SDL_SCANCODE_AC_REFRESH         SDL_Scancode = 273
	SDL_SCANCODE_AC_BOOKMARKS       SDL_Scancode = 274
	SDL_SCANCODE_BRIGHTNESSDOWN     SDL_Scancode = 275
	SDL_SCANCODE_BRIGHTNESSUP       SDL_Scancode = 276
	SDL_SCANCODE_DISPLAYSWITCH      SDL_Scancode = 277
	SDL_SCANCODE_KBDILLUMTOGGLE     SDL_Scancode = 278
	SDL_SCANCODE_KBDILLUMDOWN       SDL_Scancode = 279
	SDL_SCANCODE_KBDILLUMUP         SDL_Scancode = 280
	SDL_SCANCODE_EJECT              SDL_Scancode = 281
	SDL_SCANCODE_SLEEP              SDL_Scancode = 282
	SDL_SCANCODE_APP1               SDL_Scancode = 283
	SDL_SCANCODE_APP2               SDL_Scancode = 284
	SDL_SCANCODE_AUDIOREWIND        SDL_Scancode = 285
	SDL_SCANCODE_AUDIOFASTFORWARD   SDL_Scancode = 286
	SDL_SCANCODE_SOFTLEFT           SDL_Scancode = 287
	SDL_SCANCODE_SOFTRIGHT          SDL_Scancode = 288
	SDL_SCANCODE_CALL               SDL_Scancode = 289
	SDL_SCANCODE_ENDCALL            SDL_Scancode = 290
	SDL_NUM_SCANCODES               SDL_Scancode = 512
)

type SDL_Scancode = int32
type SDL_Keycode = int32

const SDLK_SCANCODE_MASK = (1 << 30)

func SDL_SCANCODE_TO_KEYCODE(x int32) int32 {
	return x | SDLK_SCANCODE_MASK
}

var (
	SDLK_UNKNOWN = 0

	SDLK_RETURN     SDL_KeyCode = '\r'
	SDLK_ESCAPE     SDL_KeyCode = '\x1B'
	SDLK_BACKSPACE  SDL_KeyCode = '\b'
	SDLK_TAB        SDL_KeyCode = '\t'
	SDLK_SPACE      SDL_KeyCode = ' '
	SDLK_EXCLAIM    SDL_KeyCode = '!'
	SDLK_QUOTEDBL   SDL_KeyCode = '"'
	SDLK_HASH       SDL_KeyCode = '#'
	SDLK_PERCENT    SDL_KeyCode = '%'
	SDLK_DOLLAR     SDL_KeyCode = '$'
	SDLK_AMPERSAND  SDL_KeyCode = '&'
	SDLK_QUOTE      SDL_KeyCode = '\''
	SDLK_LEFTPAREN  SDL_KeyCode = '('
	SDLK_RIGHTPAREN SDL_KeyCode = ')'
	SDLK_ASTERISK   SDL_KeyCode = '*'
	SDLK_PLUS       SDL_KeyCode = '+'
	SDLK_COMMA      SDL_KeyCode = ','
	SDLK_MINUS      SDL_KeyCode = '-'
	SDLK_PERIOD     SDL_KeyCode = '.'
	SDLK_SLASH      SDL_KeyCode = '/'
	SDLK_0          SDL_KeyCode = '0'
	SDLK_1          SDL_KeyCode = '1'
	SDLK_2          SDL_KeyCode = '2'
	SDLK_3          SDL_KeyCode = '3'
	SDLK_4          SDL_KeyCode = '4'
	SDLK_5          SDL_KeyCode = '5'
	SDLK_6          SDL_KeyCode = '6'
	SDLK_7          SDL_KeyCode = '7'
	SDLK_8          SDL_KeyCode = '8'
	SDLK_9          SDL_KeyCode = '9'
	SDLK_COLON      SDL_KeyCode = ':'
	SDLK_SEMICOLON  SDL_KeyCode = ';'
	SDLK_LESS       SDL_KeyCode = '<'
	SDLK_EQUALS     SDL_KeyCode = '='
	SDLK_GREATER    SDL_KeyCode = '>'
	SDLK_QUESTION   SDL_KeyCode = '?'
	SDLK_AT         SDL_KeyCode = '@'

	// Skip uppercase letters
	SDLK_LEFTBRACKET  SDL_KeyCode = '['
	SDLK_BACKSLASH    SDL_KeyCode = '\\'
	SDLK_RIGHTBRACKET SDL_KeyCode = ']'
	SDLK_CARET        SDL_KeyCode = '^'
	SDLK_UNDERSCORE   SDL_KeyCode = '_'
	SDLK_BACKQUOTE    SDL_KeyCode = '`'
	SDLK_a            SDL_KeyCode = 'a'
	SDLK_b            SDL_KeyCode = 'b'
	SDLK_c            SDL_KeyCode = 'c'
	SDLK_d            SDL_KeyCode = 'd'
	SDLK_e            SDL_KeyCode = 'e'
	SDLK_f            SDL_KeyCode = 'f'
	SDLK_g            SDL_KeyCode = 'g'
	SDLK_h            SDL_KeyCode = 'h'
	SDLK_i            SDL_KeyCode = 'i'
	SDLK_j            SDL_KeyCode = 'j'
	SDLK_k            SDL_KeyCode = 'k'
	SDLK_l            SDL_KeyCode = 'l'
	SDLK_m            SDL_KeyCode = 'm'
	SDLK_n            SDL_KeyCode = 'n'
	SDLK_o            SDL_KeyCode = 'o'
	SDLK_p            SDL_KeyCode = 'p'
	SDLK_q            SDL_KeyCode = 'q'
	SDLK_r            SDL_KeyCode = 'r'
	SDLK_s            SDL_KeyCode = 's'
	SDLK_t            SDL_KeyCode = 't'
	SDLK_u            SDL_KeyCode = 'u'
	SDLK_v            SDL_KeyCode = 'v'
	SDLK_w            SDL_KeyCode = 'w'
	SDLK_x            SDL_KeyCode = 'x'
	SDLK_y            SDL_KeyCode = 'y'
	SDLK_z            SDL_KeyCode = 'z'

	SDLK_CAPSLOCK SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CAPSLOCK)

	SDLK_F1  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F1)
	SDLK_F2  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F2)
	SDLK_F3  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F3)
	SDLK_F4  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F4)
	SDLK_F5  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F5)
	SDLK_F6  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F6)
	SDLK_F7  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F7)
	SDLK_F8  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F8)
	SDLK_F9  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F9)
	SDLK_F10 SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F10)
	SDLK_F11 SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F11)
	SDLK_F12 SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F12)

	SDLK_PRINTSCREEN SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PRINTSCREEN)
	SDLK_SCROLLLOCK  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SCROLLLOCK)
	SDLK_PAUSE       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAUSE)
	SDLK_INSERT      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_INSERT)
	SDLK_HOME        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_HOME)
	SDLK_PAGEUP      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAGEUP)
	SDLK_DELETE      SDL_KeyCode = '\x7F'
	SDLK_END         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_END)
	SDLK_PAGEDOWN    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAGEDOWN)
	SDLK_RIGHT       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RIGHT)
	SDLK_LEFT        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LEFT)
	SDLK_DOWN        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DOWN)
	SDLK_UP          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_UP)

	SDLK_NUMLOCKCLEAR SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_NUMLOCKCLEAR)
	SDLK_KP_DIVIDE    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DIVIDE)
	SDLK_KP_MULTIPLY  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MULTIPLY)
	SDLK_KP_MINUS     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MINUS)
	SDLK_KP_PLUS      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PLUS)
	SDLK_KP_ENTER     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_ENTER)
	SDLK_KP_1         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_1)
	SDLK_KP_2         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_2)
	SDLK_KP_3         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_3)
	SDLK_KP_4         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_4)
	SDLK_KP_5         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_5)
	SDLK_KP_6         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_6)
	SDLK_KP_7         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_7)
	SDLK_KP_8         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_8)
	SDLK_KP_9         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_9)
	SDLK_KP_0         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_0)
	SDLK_KP_PERIOD    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PERIOD)

	SDLK_APPLICATION    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_APPLICATION)
	SDLK_POWER          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_POWER)
	SDLK_KP_EQUALS      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EQUALS)
	SDLK_F13            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F13)
	SDLK_F14            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F14)
	SDLK_F15            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F15)
	SDLK_F16            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F16)
	SDLK_F17            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F17)
	SDLK_F18            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F18)
	SDLK_F19            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F19)
	SDLK_F20            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F20)
	SDLK_F21            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F21)
	SDLK_F22            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F22)
	SDLK_F23            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F23)
	SDLK_F24            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F24)
	SDLK_EXECUTE        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EXECUTE)
	SDLK_HELP           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_HELP)
	SDLK_MENU           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MENU)
	SDLK_SELECT         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SELECT)
	SDLK_STOP           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_STOP)
	SDLK_AGAIN          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AGAIN)
	SDLK_UNDO           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_UNDO)
	SDLK_CUT            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CUT)
	SDLK_COPY           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_COPY)
	SDLK_PASTE          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PASTE)
	SDLK_FIND           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_FIND)
	SDLK_MUTE           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MUTE)
	SDLK_VOLUMEUP       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_VOLUMEUP)
	SDLK_VOLUMEDOWN     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_VOLUMEDOWN)
	SDLK_KP_COMMA       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_COMMA)
	SDLK_KP_EQUALSAS400 SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EQUALSAS400)

	SDLK_ALTERASE   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_ALTERASE)
	SDLK_SYSREQ     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SYSREQ)
	SDLK_CANCEL     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CANCEL)
	SDLK_CLEAR      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CLEAR)
	SDLK_PRIOR      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PRIOR)
	SDLK_RETURN2    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RETURN2)
	SDLK_SEPARATOR  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SEPARATOR)
	SDLK_OUT        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_OUT)
	SDLK_OPER       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_OPER)
	SDLK_CLEARAGAIN SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CLEARAGAIN)
	SDLK_CRSEL      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CRSEL)
	SDLK_EXSEL      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EXSEL)

	SDLK_KP_00              SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_00)
	SDLK_KP_000             SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_000)
	SDLK_THOUSANDSSEPARATOR SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_THOUSANDSSEPARATOR)
	SDLK_DECIMALSEPARATOR   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DECIMALSEPARATOR)
	SDLK_CURRENCYUNIT       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CURRENCYUNIT)
	SDLK_CURRENCYSUBUNIT    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CURRENCYSUBUNIT)
	SDLK_KP_LEFTPAREN       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LEFTPAREN)
	SDLK_KP_RIGHTPAREN      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_RIGHTPAREN)
	SDLK_KP_LEFTBRACE       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LEFTBRACE)
	SDLK_KP_RIGHTBRACE      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_RIGHTBRACE)
	SDLK_KP_TAB             SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_TAB)
	SDLK_KP_BACKSPACE       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_BACKSPACE)
	SDLK_KP_A               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_A)
	SDLK_KP_B               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_B)
	SDLK_KP_C               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_C)
	SDLK_KP_D               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_D)
	SDLK_KP_E               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_E)
	SDLK_KP_F               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_F)
	SDLK_KP_XOR             SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_XOR)
	SDLK_KP_POWER           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_POWER)
	SDLK_KP_PERCENT         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PERCENT)
	SDLK_KP_LESS            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LESS)
	SDLK_KP_GREATER         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_GREATER)
	SDLK_KP_AMPERSAND       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_AMPERSAND)
	SDLK_KP_DBLAMPERSAND    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DBLAMPERSAND)
	SDLK_KP_VERTICALBAR     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_VERTICALBAR)
	SDLK_KP_DBLVERTICALBAR  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DBLVERTICALBAR)
	SDLK_KP_COLON           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_COLON)
	SDLK_KP_HASH            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_HASH)
	SDLK_KP_SPACE           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_SPACE)
	SDLK_KP_AT              SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_AT)
	SDLK_KP_EXCLAM          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EXCLAM)
	SDLK_KP_MEMSTORE        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMSTORE)
	SDLK_KP_MEMRECALL       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMRECALL)
	SDLK_KP_MEMCLEAR        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMCLEAR)
	SDLK_KP_MEMADD          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMADD)
	SDLK_KP_MEMSUBTRACT     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMSUBTRACT)
	SDLK_KP_MEMMULTIPLY     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMMULTIPLY)
	SDLK_KP_MEMDIVIDE       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMDIVIDE)
	SDLK_KP_PLUSMINUS       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PLUSMINUS)
	SDLK_KP_CLEAR           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_CLEAR)
	SDLK_KP_CLEARENTRY      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_CLEARENTRY)
	SDLK_KP_BINARY          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_BINARY)
	SDLK_KP_OCTAL           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_OCTAL)
	SDLK_KP_DECIMAL         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DECIMAL)
	SDLK_KP_HEXADECIMAL     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_HEXADECIMAL)

	SDLK_LCTRL  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LCTRL)
	SDLK_LSHIFT SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LSHIFT)
	SDLK_LALT   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LALT)
	SDLK_LGUI   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LGUI)
	SDLK_RCTRL  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RCTRL)
	SDLK_RSHIFT SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RSHIFT)
	SDLK_RALT   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RALT)
	SDLK_RGUI   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RGUI)

	SDLK_MODE SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MODE)

	SDLK_AUDIONEXT    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIONEXT)
	SDLK_AUDIOPREV    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOPREV)
	SDLK_AUDIOSTOP    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOSTOP)
	SDLK_AUDIOPLAY    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOPLAY)
	SDLK_AUDIOMUTE    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOMUTE)
	SDLK_MEDIASELECT  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIASELECT)
	SDLK_WWW          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_WWW)
	SDLK_MAIL         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MAIL)
	SDLK_CALCULATOR   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CALCULATOR)
	SDLK_COMPUTER     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_COMPUTER)
	SDLK_AC_SEARCH    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_SEARCH)
	SDLK_AC_HOME      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_HOME)
	SDLK_AC_BACK      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_BACK)
	SDLK_AC_FORWARD   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_FORWARD)
	SDLK_AC_STOP      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_STOP)
	SDLK_AC_REFRESH   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_REFRESH)
	SDLK_AC_BOOKMARKS SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_BOOKMARKS)

	SDLK_BRIGHTNESSDOWN SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_BRIGHTNESSDOWN)
	SDLK_BRIGHTNESSUP   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_BRIGHTNESSUP)
	SDLK_DISPLAYSWITCH  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DISPLAYSWITCH)
	SDLK_KBDILLUMTOGGLE SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KBDILLUMTOGGLE)
	SDLK_KBDILLUMDOWN   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KBDILLUMDOWN)
	SDLK_KBDILLUMUP     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KBDILLUMUP)
	SDLK_EJECT          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EJECT)
	SDLK_SLEEP          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SLEEP)
	SDLK_APP1           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_APP1)
	SDLK_APP2           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_APP2)

	SDLK_AUDIOREWIND      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOREWIND)
	SDLK_AUDIOFASTFORWARD SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOFASTFORWARD)
)

type SDL_KeyCode = int32

const (
	KMOD_NONE     SDL_Keymod = 0x0000
	KMOD_LSHIFT   SDL_Keymod = 0x0001
	KMOD_RSHIFT   SDL_Keymod = 0x0002
	KMOD_LCTRL    SDL_Keymod = 0x0040
	KMOD_RCTRL    SDL_Keymod = 0x0080
	KMOD_LALT     SDL_Keymod = 0x0100
	KMOD_RALT     SDL_Keymod = 0x0200
	KMOD_LGUI     SDL_Keymod = 0x0400
	KMOD_RGUI     SDL_Keymod = 0x0800
	KMOD_NUM      SDL_Keymod = 0x1000
	KMOD_CAPS     SDL_Keymod = 0x2000
	KMOD_MODE     SDL_Keymod = 0x4000
	KMOD_SCROLL   SDL_Keymod = 0x8000
	KMOD_CTRL     SDL_Keymod = KMOD_LCTRL | KMOD_RCTRL
	KMOD_SHIFT    SDL_Keymod = KMOD_LSHIFT | KMOD_RSHIFT
	KMOD_ALT      SDL_Keymod = KMOD_LALT | KMOD_RALT
	KMOD_GUI      SDL_Keymod = KMOD_LGUI | KMOD_RGUI
	KMOD_RESERVED SDL_Keymod = KMOD_SCROLL
)

type SDL_Keymod = uint32

type SDL_Keysym struct {
	Scancode int32
	Sym      int32
	Mod      uint16
	Unused   uint32
}

const (
	SDL_SYSTEM_CURSOR_ARROW     SDL_SystemCursor = 0
	SDL_SYSTEM_CURSOR_IBEAM     SDL_SystemCursor = 1
	SDL_SYSTEM_CURSOR_WAIT      SDL_SystemCursor = 2
	SDL_SYSTEM_CURSOR_CROSSHAIR SDL_SystemCursor = 3
	SDL_SYSTEM_CURSOR_WAITARROW SDL_SystemCursor = 4
	SDL_SYSTEM_CURSOR_SIZENWSE  SDL_SystemCursor = 5
	SDL_SYSTEM_CURSOR_SIZENESW  SDL_SystemCursor = 6
	SDL_SYSTEM_CURSOR_SIZEWE    SDL_SystemCursor = 7
	SDL_SYSTEM_CURSOR_SIZENS    SDL_SystemCursor = 8
	SDL_SYSTEM_CURSOR_SIZEALL   SDL_SystemCursor = 9
	SDL_SYSTEM_CURSOR_NO        SDL_SystemCursor = 10
	SDL_SYSTEM_CURSOR_HAND      SDL_SystemCursor = 11
	SDL_NUM_SYSTEM_CURSORS      SDL_SystemCursor = 12
)

type SDL_SystemCursor = int32

const (
	SDL_MOUSEWHEEL_NORMAL  SDL_MouseWheelDirection = 0
	SDL_MOUSEWHEEL_FLIPPED SDL_MouseWheelDirection = 1
)

type SDL_MouseWheelDirection = int32

type SDL_GUID struct {
	Data [16]uint8
}

type SDL_JoystickGUID = SDL_GUID
type SDL_JoystickID = int32

const (
	SDL_JOYSTICK_TYPE_UNKNOWN        SDL_JoystickType = 0
	SDL_JOYSTICK_TYPE_GAMECONTROLLER SDL_JoystickType = 1
	SDL_JOYSTICK_TYPE_WHEEL          SDL_JoystickType = 2
	SDL_JOYSTICK_TYPE_ARCADE_STICK   SDL_JoystickType = 3
	SDL_JOYSTICK_TYPE_FLIGHT_STICK   SDL_JoystickType = 4
	SDL_JOYSTICK_TYPE_DANCE_PAD      SDL_JoystickType = 5
	SDL_JOYSTICK_TYPE_GUITAR         SDL_JoystickType = 6
	SDL_JOYSTICK_TYPE_DRUM_KIT       SDL_JoystickType = 7
	SDL_JOYSTICK_TYPE_ARCADE_PAD     SDL_JoystickType = 8
	SDL_JOYSTICK_TYPE_THROTTLE       SDL_JoystickType = 9
)

type SDL_JoystickType = int32

const (
	SDL_JOYSTICK_POWER_UNKNOWN SDL_JoystickPowerLevel = -1
	SDL_JOYSTICK_POWER_EMPTY   SDL_JoystickPowerLevel = 0
	SDL_JOYSTICK_POWER_LOW     SDL_JoystickPowerLevel = 1
	SDL_JOYSTICK_POWER_MEDIUM  SDL_JoystickPowerLevel = 2
	SDL_JOYSTICK_POWER_FULL    SDL_JoystickPowerLevel = 3
	SDL_JOYSTICK_POWER_WIRED   SDL_JoystickPowerLevel = 4
	SDL_JOYSTICK_POWER_MAX     SDL_JoystickPowerLevel = 5
)

type SDL_JoystickPowerLevel = int32

type SDL_SensorID = int32

const (
	SDL_SENSOR_INVALID SDL_SensorType = -1
	SDL_SENSOR_UNKNOWN SDL_SensorType = 0
	SDL_SENSOR_ACCEL   SDL_SensorType = 1
	SDL_SENSOR_GYRO    SDL_SensorType = 2
	// SDL_SENSOR_ACCEL_L SDL_SensorType = 3
	// SDL_SENSOR_GYRO_L  SDL_SensorType = 4
	// SDL_SENSOR_ACCEL_R SDL_SensorType = 5
	// SDL_SENSOR_GYRO_R  SDL_SensorType = 6
)

type SDL_SensorType = int32

const (
	SDL_CONTROLLER_TYPE_UNKNOWN             SDL_GameControllerType = 0
	SDL_CONTROLLER_TYPE_XBOX360             SDL_GameControllerType = 1
	SDL_CONTROLLER_TYPE_XBOXONE             SDL_GameControllerType = 2
	SDL_CONTROLLER_TYPE_PS3                 SDL_GameControllerType = 3
	SDL_CONTROLLER_TYPE_PS4                 SDL_GameControllerType = 4
	SDL_CONTROLLER_TYPE_NINTENDO_SWITCH_PRO SDL_GameControllerType = 5
	SDL_CONTROLLER_TYPE_VIRTUAL             SDL_GameControllerType = 6
	SDL_CONTROLLER_TYPE_PS5                 SDL_GameControllerType = 7
	SDL_CONTROLLER_TYPE_AMAZON_LUNA         SDL_GameControllerType = 8
	SDL_CONTROLLER_TYPE_GOOGLE_STADIA       SDL_GameControllerType = 9
	// SDL_CONTROLLER_TYPE_NVIDIA_SHIELD                SDL_GameControllerType = 10
	// SDL_CONTROLLER_TYPE_NINTENDO_SWITCH_JOYCON_LEFT  SDL_GameControllerType = 11
	// SDL_CONTROLLER_TYPE_NINTENDO_SWITCH_JOYCON_RIGHT SDL_GameControllerType = 12
	// SDL_CONTROLLER_TYPE_NINTENDO_SWITCH_JOYCON_PAIR  SDL_GameControllerType = 13
)

type SDL_GameControllerType = int32

const (
	SDL_CONTROLLER_BINDTYPE_NONE   SDL_GameControllerBindType = 0
	SDL_CONTROLLER_BINDTYPE_BUTTON SDL_GameControllerBindType = 1
	SDL_CONTROLLER_BINDTYPE_AXIS   SDL_GameControllerBindType = 2
	SDL_CONTROLLER_BINDTYPE_HAT    SDL_GameControllerBindType = 3
)

type SDL_GameControllerBindType = int32

type SDL_GameControllerButtonBindButton struct {
	Button int32
}

type SDL_GameControllerButtonBindAxis struct {
	Axis int32
}

type SDL_GameControllerButtonBindHat struct {
	Hat     int32
	HatMask int32
}

// button or axis or hat
type SDL_GameControllerButtonBindPadding struct {
	Padding1 int32
	Padding2 int32
}

type SDL_GameControllerButtonBind struct {
	BindType int32
	Value    SDL_GameControllerButtonBindPadding
}

const (
	SDL_CONTROLLER_AXIS_INVALID      SDL_GameControllerAxis = -1
	SDL_CONTROLLER_AXIS_LEFTX        SDL_GameControllerAxis = 0
	SDL_CONTROLLER_AXIS_LEFTY        SDL_GameControllerAxis = 1
	SDL_CONTROLLER_AXIS_RIGHTX       SDL_GameControllerAxis = 2
	SDL_CONTROLLER_AXIS_RIGHTY       SDL_GameControllerAxis = 3
	SDL_CONTROLLER_AXIS_TRIGGERLEFT  SDL_GameControllerAxis = 4
	SDL_CONTROLLER_AXIS_TRIGGERRIGHT SDL_GameControllerAxis = 5
	SDL_CONTROLLER_AXIS_MAX          SDL_GameControllerAxis = 6
)

type SDL_GameControllerAxis = int32

const (
	SDL_CONTROLLER_BUTTON_INVALID       SDL_GameControllerButton = -1
	SDL_CONTROLLER_BUTTON_A             SDL_GameControllerButton = 0
	SDL_CONTROLLER_BUTTON_B             SDL_GameControllerButton = 1
	SDL_CONTROLLER_BUTTON_X             SDL_GameControllerButton = 2
	SDL_CONTROLLER_BUTTON_Y             SDL_GameControllerButton = 3
	SDL_CONTROLLER_BUTTON_BACK          SDL_GameControllerButton = 4
	SDL_CONTROLLER_BUTTON_GUIDE         SDL_GameControllerButton = 5
	SDL_CONTROLLER_BUTTON_START         SDL_GameControllerButton = 6
	SDL_CONTROLLER_BUTTON_LEFTSTICK     SDL_GameControllerButton = 7
	SDL_CONTROLLER_BUTTON_RIGHTSTICK    SDL_GameControllerButton = 8
	SDL_CONTROLLER_BUTTON_LEFTSHOULDER  SDL_GameControllerButton = 9
	SDL_CONTROLLER_BUTTON_RIGHTSHOULDER SDL_GameControllerButton = 10
	SDL_CONTROLLER_BUTTON_DPAD_UP       SDL_GameControllerButton = 11
	SDL_CONTROLLER_BUTTON_DPAD_DOWN     SDL_GameControllerButton = 12
	SDL_CONTROLLER_BUTTON_DPAD_LEFT     SDL_GameControllerButton = 13
	SDL_CONTROLLER_BUTTON_DPAD_RIGHT    SDL_GameControllerButton = 14
	SDL_CONTROLLER_BUTTON_MISC1         SDL_GameControllerButton = 15
	SDL_CONTROLLER_BUTTON_PADDLE1       SDL_GameControllerButton = 16
	SDL_CONTROLLER_BUTTON_PADDLE2       SDL_GameControllerButton = 17
	SDL_CONTROLLER_BUTTON_PADDLE3       SDL_GameControllerButton = 18
	SDL_CONTROLLER_BUTTON_PADDLE4       SDL_GameControllerButton = 19
	SDL_CONTROLLER_BUTTON_TOUCHPAD      SDL_GameControllerButton = 20
	SDL_CONTROLLER_BUTTON_MAX           SDL_GameControllerButton = 21
)

type SDL_GameControllerButton = int32

type SDL_TouchID = int64
type SDL_FingerID = int64

const (
	SDL_TOUCH_DEVICE_INVALID           SDL_TouchDeviceType = -1
	SDL_TOUCH_DEVICE_DIRECT            SDL_TouchDeviceType = 0
	SDL_TOUCH_DEVICE_INDIRECT_ABSOLUTE SDL_TouchDeviceType = 1
	SDL_TOUCH_DEVICE_INDIRECT_RELATIVE SDL_TouchDeviceType = 2
)

type SDL_TouchDeviceType = int32

type SDL_Finger struct {
	Id       int64
	X        float32
	Y        float32
	Pressure float32
}
type SDL_GestureID = int64

const (
	SDL_FIRSTEVENT              SDL_EventType = 0
	SDL_QUIT                    SDL_EventType = 0x100
	SDL_APP_TERMINATING         SDL_EventType = 0x101
	SDL_APP_LOWMEMORY           SDL_EventType = 0x102
	SDL_APP_WILLENTERBACKGROUND SDL_EventType = 0x103
	SDL_APP_DIDENTERBACKGROUND  SDL_EventType = 0x104
	SDL_APP_WILLENTERFOREGROUND SDL_EventType = 0x105
	SDL_APP_DIDENTERFOREGROUND  SDL_EventType = 0x106
	SDL_LOCALECHANGED           SDL_EventType = 0x107

	SDL_DISPLAYEVENT SDL_EventType = 0x150

	SDL_WINDOWEVENT SDL_EventType = 0x200
	SDL_SYSWMEVENT  SDL_EventType = 0x201

	SDL_KEYDOWN         SDL_EventType = 0x300
	SDL_KEYUP           SDL_EventType = 0x301
	SDL_TEXTEDITING     SDL_EventType = 0x302
	SDL_TEXTINPUT       SDL_EventType = 0x303
	SDL_KEYMAPCHANGED   SDL_EventType = 0x304
	SDL_TEXTEDITING_EXT SDL_EventType = 0x305

	SDL_MOUSEMOTION     SDL_EventType = 0x400
	SDL_MOUSEBUTTONDOWN SDL_EventType = 0x401
	SDL_MOUSEBUTTONUP   SDL_EventType = 0x402
	SDL_MOUSEWHEEL      SDL_EventType = 0x403

	SDL_JOYAXISMOTION     SDL_EventType = 0x600
	SDL_JOYBALLMOTION     SDL_EventType = 0x601
	SDL_JOYHATMOTION      SDL_EventType = 0x602
	SDL_JOYBUTTONDOWN     SDL_EventType = 0x603
	SDL_JOYBUTTONUP       SDL_EventType = 0x604
	SDL_JOYDEVICEADDED    SDL_EventType = 0x605
	SDL_JOYDEVICEREMOVED  SDL_EventType = 0x606
	SDL_JOYBATTERYUPDATED SDL_EventType = 0x607

	SDL_CONTROLLERAXISMOTION     SDL_EventType = 0x650
	SDL_CONTROLLERBUTTONDOWN     SDL_EventType = 0x651
	SDL_CONTROLLERBUTTONUP       SDL_EventType = 0x652
	SDL_CONTROLLERDEVICEADDED    SDL_EventType = 0x653
	SDL_CONTROLLERDEVICEREMOVED  SDL_EventType = 0x654
	SDL_CONTROLLERDEVICEREMAPPED SDL_EventType = 0x655
	SDL_CONTROLLERTOUCHPADDOWN   SDL_EventType = 0x656
	SDL_CONTROLLERTOUCHPADMOTION SDL_EventType = 0x657
	SDL_CONTROLLERTOUCHPADUP     SDL_EventType = 0x658
	SDL_CONTROLLERSENSORUPDATE   SDL_EventType = 0x659

	SDL_FINGERDOWN   SDL_EventType = 0x700
	SDL_FINGERUP     SDL_EventType = 0x701
	SDL_FINGERMOTION SDL_EventType = 0x702

	SDL_DOLLARGESTURE SDL_EventType = 0x800
	SDL_DOLLARRECORD  SDL_EventType = 0x801
	SDL_MULTIGESTURE  SDL_EventType = 0x802

	SDL_CLIPBOARDUPDATE SDL_EventType = 0x900

	SDL_DROPFILE     SDL_EventType = 0x1000
	SDL_DROPTEXT     SDL_EventType = 0x1001
	SDL_DROPBEGIN    SDL_EventType = 0x1002
	SDL_DROPCOMPLETE SDL_EventType = 0x1003

	SDL_AUDIODEVICEADDED   SDL_EventType = 0x1100
	SDL_AUDIODEVICEREMOVED SDL_EventType = 0x1101

	SDL_SENSORUPDATE SDL_EventType = 0x1200

	SDL_RENDER_TARGETS_RESET SDL_EventType = 0x2000
	SDL_RENDER_DEVICE_RESET  SDL_EventType = 0x2001

	SDL_POLLSENTINEL SDL_EventType = 0x7F00

	SDL_USEREVENT SDL_EventType = 0x8000

	SDL_LASTEVENT SDL_EventType = 0xFFFF
)

type SDL_EventType = uint32

type SDL_CommonEvent struct {
	Type      uint32
	Timestamp uint32
}

type SDL_DisplayEvent struct {
	Type      uint32
	Timestamp uint32
	Display   uint32
	Event     uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Data1     int32
}

type SDL_WindowEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Event     uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Data1     int32
	Data2     int32
}

type SDL_KeyboardEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	State     uint8
	Repeat    uint8
	Padding2  uint8
	Padding3  uint8
	Keysym    SDL_Keysym
}

type SDL_TextEditingEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Text      [32]int8
	Start     int32
	Length    int32
}

type SDL_TextEditingExtEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Text      *int8
	Start     int32
	Length    int32
}

type SDL_TextInputEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Text      [32]int8
}

type SDL_MouseMotionEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	State     uint32
	X         int32
	Y         int32
	XRel      int32
	YRel      int32
}

type SDL_MouseButtonEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	Button    uint8
	State     uint8
	Clicks    uint8
	Padding1  uint8
	X         int32
	Y         int32
}

type SDL_MouseWheelEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	X         int32
	Y         int32
	Direction uint32
	PreciseX  float32
	PreciseY  float32
}

type SDL_JoyAxisEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Axis      uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Value     int16
	Padding4  uint16
}

type SDL_JoyBallEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Ball      uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	XRel      int16
	YRel      int16
}

type SDL_JoyHatEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Hat       uint8
	Value     uint8
	Padding1  uint8
	Padding2  uint8
}

type SDL_JoyButtonEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Button    uint8
	State     uint8
	Padding1  uint8
	Padding2  uint8
}

type SDL_JoyDeviceEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
}

type SDL_JoyBatteryEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Level     int32
}

type SDL_ControllerAxisEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Axis      uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Value     int16
	Padding4  uint16
}

type SDL_ControllerButtonEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Button    uint8
	State     uint8
	Padding1  uint8
	Padding2  uint8
}

type SDL_ControllerDeviceEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
}

type SDL_ControllerTouchpadEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Touchpad  int32
	Finger    int32
	X         float32
	Y         float32
	Pressure  float32
}

type SDL_ControllerSensorEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Sensor    int32
	Data      [3]float32
}

type SDL_AudioDeviceEvent struct {
	Type      uint32
	Timestamp uint32
	Which     uint32
	Iscapture uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
}

type SDL_TouchFingerEvent struct {
	Type      uint32
	Timestamp uint32
	TouchId   int64
	FingerId  int64
	X         float32
	Y         float32
	Dx        float32
	Dy        float32
	Pressure  float32
	WindowID  uint32
}

type SDL_MultiGestureEvent struct {
	Type       uint32
	Timestamp  uint32
	TouchId    int64
	DTheta     float32
	DDist      float32
	X          float32
	Y          float32
	NumFingers uint16
	Padding    uint16
}

type SDL_DollarGestureEvent struct {
	Type       uint32
	Timestamp  uint32
	TouchId    int64
	GestureId  int64
	NumFingers uint32
	Error      float32
	X          float32
	Y          float32
}

type SDL_DropEvent struct {
	Type      uint32
	Timestamp uint32
	File      *int8
	WindowID  uint32
}

type SDL_SensorEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Data      [6]float32
}

type SDL_QuitEvent struct {
	Type      uint32
	Timestamp uint32
}

type SDL_OSEvent struct {
	Type      uint32
	Timestamp uint32
}

type SDL_UserEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Code      int32
	Data1     unsafe.Pointer
	Data2     unsafe.Pointer
}

type SDL_SysWMEvent struct {
	Type      uint32
	Timestamp uint32
	Msg       *SDL_SysWMmsg
}

type SDL_Event struct {
	Type    uint32
	Padding [56]uint8 // padding
}

const (
	SDL_ADDEVENT  SDL_eventaction = 0
	SDL_PEEKEVENT SDL_eventaction = 1
	SDL_GETEVENT  SDL_eventaction = 2
)

type SDL_eventaction = int32

// type SDL_EventFilter = func(unsafe.Pointer, *SDL_Event) int32

type SDL_HapticDirection struct {
	Type uint8
	Dir  [3]int32
}

type SDL_HapticConstant struct {
	Type         uint16
	Direction    SDL_HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Level        int16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type SDL_HapticPeriodic struct {
	Type         uint16
	Direction    SDL_HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Period       uint16
	Magnitude    int16
	Offset       int16
	Phase        uint16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type SDL_HapticCondition struct {
	Type       uint16
	Direction  SDL_HapticDirection
	Length     uint32
	Delay      uint16
	Button     uint16
	Interval   uint16
	RightSat   [3]uint16
	LeftSat    [3]uint16
	RightCoeff [3]int16
	LeftCoeff  [3]int16
	Deadband   [3]uint16
	Center     [3]int16
}

type SDL_HapticRamp struct {
	Type         uint16
	Direction    SDL_HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Start        int16
	End          int16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type SDL_HapticLeftRight struct {
	Type           uint16
	Length         uint32
	LargeMagnitude uint16
	SmallMagnitude uint16
}

type SDL_HapticCustom struct {
	Type         uint16
	Direction    SDL_HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Channels     uint8
	Period       uint16
	Samples      uint16
	Data         *uint16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type SDL_HapticEffect struct {
	condition SDL_HapticCondition
}

type SDL_hid_device_info struct {
	Path               *int8
	VendorId           uint16
	ProductId          uint16
	SerialNumber       *int32
	ReleaseNumber      uint16
	ManufacturerString *int32
	ProductString      *int32
	UsagePage          uint16
	Usage              uint16
	InterfaceNumber    int32
	InterfaceClass     int32
	InterfaceSubclass  int32
	InterfaceProtocol  int32
	Next               *SDL_hid_device_info
}

const (
	SDL_HINT_DEFAULT  SDL_HintPriority = 0
	SDL_HINT_NORMAL   SDL_HintPriority = 1
	SDL_HINT_OVERRIDE SDL_HintPriority = 2
)

type SDL_HintPriority = int32

type SDL_HintCallback = func(unsafe.Pointer, *int8, *int8, *int8)

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

type SDL_LogCategory = int32

const (
	SDL_LOG_PRIORITY_VERBOSE  SDL_LogPriority = 1
	SDL_LOG_PRIORITY_DEBUG    SDL_LogPriority = 2
	SDL_LOG_PRIORITY_INFO     SDL_LogPriority = 3
	SDL_LOG_PRIORITY_WARN     SDL_LogPriority = 4
	SDL_LOG_PRIORITY_ERROR    SDL_LogPriority = 5
	SDL_LOG_PRIORITY_CRITICAL SDL_LogPriority = 6
	SDL_NUM_LOG_PRIORITIES    SDL_LogPriority = 7
)

type SDL_LogPriority = int32

type SDL_LogOutputFunction = func(unsafe.Pointer, int32, int32, *int8)

const (
	SDL_MESSAGEBOX_ERROR                 SDL_MessageBoxFlags = 16
	SDL_MESSAGEBOX_WARNING               SDL_MessageBoxFlags = 32
	SDL_MESSAGEBOX_INFORMATION           SDL_MessageBoxFlags = 64
	SDL_MESSAGEBOX_BUTTONS_LEFT_TO_RIGHT SDL_MessageBoxFlags = 128
	SDL_MESSAGEBOX_BUTTONS_RIGHT_TO_LEFT SDL_MessageBoxFlags = 256
)

type SDL_MessageBoxFlags = int32

const (
	SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT SDL_MessageBoxButtonFlags = 1
	SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT SDL_MessageBoxButtonFlags = 2
)

type SDL_MessageBoxButtonFlags = int32

type SDL_MessageBoxButtonData struct {
	Flags    uint32
	Buttonid int32
	Text     *int8
}

type SDL_MessageBoxColor struct {
	R uint8
	G uint8
	B uint8
}

const (
	SDL_MESSAGEBOX_COLOR_BACKGROUND        SDL_MessageBoxColorType = 0
	SDL_MESSAGEBOX_COLOR_TEXT              SDL_MessageBoxColorType = 1
	SDL_MESSAGEBOX_COLOR_BUTTON_BORDER     SDL_MessageBoxColorType = 2
	SDL_MESSAGEBOX_COLOR_BUTTON_BACKGROUND SDL_MessageBoxColorType = 3
	SDL_MESSAGEBOX_COLOR_BUTTON_SELECTED   SDL_MessageBoxColorType = 4
	SDL_MESSAGEBOX_COLOR_MAX               SDL_MessageBoxColorType = 5
)

type SDL_MessageBoxColorType = int32

type SDL_MessageBoxColorScheme struct {
	Colors [5]SDL_MessageBoxColor
}

type SDL_MessageBoxData struct {
	Flags       uint32
	Window      *SDL_Window
	Title       *int8
	Message     *int8
	Numbuttons  int32
	Buttons     *SDL_MessageBoxButtonData
	ColorScheme *SDL_MessageBoxColorScheme
}

type SDL_MetalView = unsafe.Pointer

const (
	SDL_POWERSTATE_UNKNOWN    SDL_PowerState = 0
	SDL_POWERSTATE_ON_BATTERY SDL_PowerState = 1
	SDL_POWERSTATE_NO_BATTERY SDL_PowerState = 2
	SDL_POWERSTATE_CHARGING   SDL_PowerState = 3
	SDL_POWERSTATE_CHARGED    SDL_PowerState = 4
)

type SDL_PowerState = int32

const (
	SDL_RENDERER_SOFTWARE      SDL_RendererFlags = 1
	SDL_RENDERER_ACCELERATED   SDL_RendererFlags = 2
	SDL_RENDERER_PRESENTVSYNC  SDL_RendererFlags = 4
	SDL_RENDERER_TARGETTEXTURE SDL_RendererFlags = 8
)

type SDL_RendererFlags = int32

type SDL_RendererInfo struct {
	Name              *int8
	Flags             uint32
	NumTextureFormats uint32
	TextureFormats    [16]uint32
	MaxTextureWidth   int32
	MaxTextureHeight  int32
}

type SDL_Vertex struct {
	Position SDL_FPoint
	Color    SDL_Color
	TexCoord SDL_FPoint
}

const (
	SDL_ScaleModeNearest SDL_ScaleMode = 0
	SDL_ScaleModeLinear  SDL_ScaleMode = 1
	SDL_ScaleModeBest    SDL_ScaleMode = 2
)

type SDL_ScaleMode = int32

const (
	SDL_TEXTUREACCESS_STATIC    SDL_TextureAccess = 0
	SDL_TEXTUREACCESS_STREAMING SDL_TextureAccess = 1
	SDL_TEXTUREACCESS_TARGET    SDL_TextureAccess = 2
)

type SDL_TextureAccess = int32

const (
	SDL_TEXTUREMODULATE_NONE  SDL_TextureModulate = 0
	SDL_TEXTUREMODULATE_COLOR SDL_TextureModulate = 1
	SDL_TEXTUREMODULATE_ALPHA SDL_TextureModulate = 2
)

type SDL_TextureModulate = int32

const (
	SDL_FLIP_NONE       SDL_RendererFlip = 0
	SDL_FLIP_HORIZONTAL SDL_RendererFlip = 1
	SDL_FLIP_VERTICAL   SDL_RendererFlip = 2
)

type SDL_RendererFlip = int32

const (
	ShapeModeDefault              WindowShapeMode = 0
	ShapeModeBinarizeAlpha        WindowShapeMode = 1
	ShapeModeReverseBinarizeAlpha WindowShapeMode = 2
	ShapeModeColorKey             WindowShapeMode = 3
)

type WindowShapeMode = int32

type SDL_WindowShapeParams struct {
	ColorKey SDL_Color
}

type SDL_WindowShapeMode struct {
	Mode       int32
	Parameters SDL_WindowShapeParams
}

type SDL_TimerCallback = func(uint32, unsafe.Pointer) uint32
type SDL_TimerID = int32

type SDL_version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

type SDL_Locale struct {
	Language *int8
	Country  *int8
}

// 空接口
type emptyInterface struct {
	Typ  unsafe.Pointer
	Word unsafe.Pointer
}

type SDL_PixelsArray interface {
	[]byte | []int8 | []int16 | []uint16 | []int32 | []uint32
}

type SDL_DataStructures interface {
	SDL_DisplayMode |
		SDL_WindowShapeParams |
		SDL_WindowShapeMode |
		SDL_Window |
		SDL_Texture |
		SDL_Renderer |
		SDL_Surface |
		SDL_RWops |
		SDL_RendererInfo |
		SDL_Vertex |
		SDL_Point |
		SDL_Rect |
		SDL_FPoint |
		SDL_FRect |
		SDL_PixelFormat |
		SDL_Palette |
		SDL_CommonEvent |
		SDL_DisplayEvent |
		SDL_WindowEvent |
		SDL_KeyboardEvent |
		SDL_TextEditingEvent |
		SDL_TextEditingExtEvent |
		SDL_TextInputEvent |
		SDL_MouseMotionEvent |
		SDL_MouseButtonEvent |
		SDL_MouseWheelEvent |
		SDL_JoyAxisEvent |
		SDL_JoyBallEvent |
		SDL_JoyHatEvent |
		SDL_JoyButtonEvent |
		SDL_JoyDeviceEvent |
		SDL_ControllerAxisEvent |
		SDL_ControllerButtonEvent |
		SDL_ControllerDeviceEvent |
		SDL_ControllerTouchpadEvent |
		SDL_ControllerSensorEvent |
		SDL_AudioDeviceEvent |
		SDL_TouchFingerEvent |
		SDL_MultiGestureEvent |
		SDL_DollarGestureEvent |
		SDL_DropEvent |
		SDL_SensorEvent |
		SDL_QuitEvent |
		SDL_OSEvent |
		SDL_UserEvent |
		SDL_SysWMEvent |
		SDL_Event |
		SDL_EventWatcher |
		SDL_AudioSpec |
		SDL_AudioCVT |
		SDL_AudioStream |
		SDL_AudioWatcher
}

type SDL_EventStructures interface {
	SDL_CommonEvent |
		SDL_DisplayEvent |
		SDL_WindowEvent |
		SDL_KeyboardEvent |
		SDL_TextEditingEvent |
		SDL_TextEditingExtEvent |
		SDL_TextInputEvent |
		SDL_MouseMotionEvent |
		SDL_MouseButtonEvent |
		SDL_MouseWheelEvent |
		SDL_JoyAxisEvent |
		SDL_JoyBallEvent |
		SDL_JoyHatEvent |
		SDL_JoyButtonEvent |
		SDL_JoyDeviceEvent |
		SDL_ControllerAxisEvent |
		SDL_ControllerButtonEvent |
		SDL_ControllerDeviceEvent |
		SDL_ControllerTouchpadEvent |
		SDL_ControllerSensorEvent |
		SDL_AudioDeviceEvent |
		SDL_TouchFingerEvent |
		SDL_MultiGestureEvent |
		SDL_DollarGestureEvent |
		SDL_DropEvent |
		SDL_SensorEvent |
		SDL_QuitEvent |
		SDL_OSEvent |
		SDL_UserEvent |
		SDL_SysWMEvent |
		SDL_Event
}

func SDL_EventConvert[T SDL_EventStructures, T2 SDL_EventStructures](e *T, _ T2) *T2 {
	return (*T2)(unsafe.Pointer(e))
}

type SDL_ConvertStructures interface {
	SDL_GameControllerButtonBindPadding |
		SDL_GameControllerButtonBindButton |
		SDL_GameControllerButtonBindAxis |
		SDL_GameControllerButtonBindHat
}

func SDL_Convert[T SDL_ConvertStructures, T2 SDL_ConvertStructures](s *T, _ T2) *T2 {
	return (*T2)(unsafe.Pointer(s))
}

func init() {
	// 大小端, 初始化处理
	SDL_PIXELFORMATENUM_INIT()
}
