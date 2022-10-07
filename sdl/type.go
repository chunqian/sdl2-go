package sdl

/*
#include "sdl_wrapper.h"
#include "gamecontroller.h"
*/
import "C"
import (
	"reflect"
	"unsafe"

	. "github.com/chunqian/memory"
)

// c basic
type (
	cChar   = C.char   // byte
	cInt    = C.int    // int32
	cUint   = C.uint   // uint32
	cFloat  = C.float  // float32
	cDouble = C.double // float64

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

// c enum
type (
	cBool                 = C.SDL_bool
	cErrorcode            = C.SDL_errorcode
	cBlendMode            = C.SDL_BlendMode
	cBlendOperation       = C.SDL_BlendOperation
	cBlendFactor          = C.SDL_BlendFactor
	cScaleMode            = C.SDL_ScaleMode
	cRendererFlip         = C.SDL_RendererFlip
	cEventType            = C.SDL_EventType
	cEventaction          = C.SDL_eventaction
	cSystemCursor         = C.SDL_SystemCursor
	cTouchID              = C.SDL_TouchID
	cGestureID            = C.SDL_GestureID
	cGameControllerAxis   = C.SDL_GameControllerAxis
	cGameControllerButton = C.SDL_GameControllerButton
	cJoystickID           = C.SDL_JoystickID
	cJoystickPowerLevel   = C.SDL_JoystickPowerLevel
	cKeymod               = C.SDL_Keymod
	cScancode             = C.SDL_Scancode
	cKeycode              = C.SDL_Keycode
	cSensorID             = C.SDL_SensorID
	cHintPriority         = C.SDL_HintPriority
	cLogPriority          = C.SDL_LogPriority
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

func IsLittleEndian() bool {
	// 占4byte 转换成16进制 0x00 00 00 01
	var value int32 = 1

	// 大端(16进制): 00 00 00 01
	// 小端(16进制): 01 00 00 00
	pointer := unsafe.Pointer(&value)
	pb := (*byte)(pointer)
	if *pb != 1 {
		return false
	}
	return true
}

type SDL_MetalView = unsafe.Pointer

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

func createCString(mp *PX_memorypool, s string) *cChar {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	n := PX_uint(sh.Len)
	p := MP_Malloc(mp, n+1)
	PX_memcpy(p, unsafe.Pointer(sh.Data), PX_int(n))
	*(*cChar)(unsafe.Pointer(uintptr(p) + uintptr(n))) = '\x00'
	return (*cChar)(p)
}

func destroyCString(mp *PX_memorypool, cStr *cChar) {
	MP_Free(mp, unsafe.Pointer(cStr))
}

func createGoString(cStr *cChar) string {
	len := PX_strlen((*PX_char)(unsafe.Pointer(cStr)))

	var slice []uint8
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sh.Data = uintptr(unsafe.Pointer(cStr))
	sh.Len = int(len)
	sh.Cap = int(len)
	src := *(*[]byte)(unsafe.Pointer(sh))

	return string(src)
}
