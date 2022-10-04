package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

// define
var (
	SDL_HAPTIC              = func(x uint) uint { return 1 << x }
	SDL_HAPTIC_CONSTANT     = SDL_HAPTIC(0)
	SDL_HAPTIC_SINE         = SDL_HAPTIC(1)
	SDL_HAPTIC_LEFTRIGHT    = SDL_HAPTIC(2)
	SDL_HAPTIC_TRIANGLE     = SDL_HAPTIC(3)
	SDL_HAPTIC_SAWTOOTHUP   = SDL_HAPTIC(4)
	SDL_HAPTIC_SAWTOOTHDOWN = SDL_HAPTIC(5)
	SDL_HAPTIC_RAMP         = SDL_HAPTIC(6)
	SDL_HAPTIC_SPRING       = SDL_HAPTIC(7)
	SDL_HAPTIC_DAMPER       = SDL_HAPTIC(8)
	SDL_HAPTIC_INERTIA      = SDL_HAPTIC(9)
	SDL_HAPTIC_FRICTION     = SDL_HAPTIC(10)
	SDL_HAPTIC_CUSTOM       = SDL_HAPTIC(11)
	SDL_HAPTIC_GAIN         = SDL_HAPTIC(12)
	SDL_HAPTIC_AUTOCENTER   = SDL_HAPTIC(13)
	SDL_HAPTIC_STATUS       = SDL_HAPTIC(14)
	SDL_HAPTIC_PAUSE        = SDL_HAPTIC(15)

	SDL_HAPTIC_POLAR         = 0
	SDL_HAPTIC_CARTESIAN     = 1
	SDL_HAPTIC_SPHERICAL     = 2
	SDL_HAPTIC_STEERING_AXIS = 3
	SDL_HAPTIC_INFINITY      = 0xFFFFFFFF
)

// struct
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

func cHapticEffect(he *SDL_HapticEffect) *C.SDL_HapticEffect {
	return (*C.SDL_HapticEffect)(unsafe.Pointer(he))
}

func cHaptic(h *SDL_Haptic) *C.SDL_Haptic {
	return (*C.SDL_Haptic)(unsafe.Pointer(h))
}

func SDL_NumHaptics() int {
	cRet := C.SDL_NumHaptics()
	return int(cRet)
}

func SDL_HapticName(index int) string {
	name := C.SDL_HapticName(cInt(index))
	if name == nil {
		return ""
	}
	return SDL_GoString(name)
}

func SDL_HapticOpen(index int) *SDL_Haptic {
	cHaptic := C.SDL_HapticOpen(cInt(index))
	haptic := (*SDL_Haptic)(unsafe.Pointer(cHaptic))
	if haptic == nil {
		return nil
	}
	return haptic
}

func SDL_HapticOpened(index int) bool {
	cRet := C.SDL_HapticOpened(cInt(index))
	return int(cRet) == 1
}

func SDL_HapticIndex(h *SDL_Haptic) int {
	cRet := C.SDL_HapticIndex(cHaptic(h))
	return int(cRet)
}

func SDL_MouseIsHaptic() bool {
	cRet := C.SDL_MouseIsHaptic()
	return int(cRet) == 1
}

func SDL_HapticOpenFromMouse() *SDL_Haptic {
	cHaptic := C.SDL_HapticOpenFromMouse()
	haptic := (*SDL_Haptic)(unsafe.Pointer(cHaptic))
	if haptic == nil {
		return nil
	}
	return haptic
}

func SDL_JoystickIsHaptic(joy *SDL_Joystick) bool {
	cRet := C.SDL_JoystickIsHaptic(cJoystick(joy))
	return int(cRet) == 1
}

func SDL_HapticOpenFromJoystick(joy *SDL_Joystick) *SDL_Haptic {
	cHaptic := C.SDL_HapticOpenFromJoystick(cJoystick(joy))
	haptic := (*SDL_Haptic)(unsafe.Pointer(cHaptic))
	if haptic == nil {
		return nil
	}
	return haptic
}

func SDL_HapticClose(h *SDL_Haptic) {
	C.SDL_HapticClose(cHaptic(h))
}

func SDL_HapticNumAxes(h *SDL_Haptic) int {
	cRet := C.SDL_HapticNumAxes(cHaptic(h))
	return int(cRet)
}

func SDL_HapticNumEffects(h *SDL_Haptic) int {
	cRet := C.SDL_HapticNumEffects(cHaptic(h))
	return int(cRet)
}

func SDL_HapticNumEffectsPlaying(h *SDL_Haptic) int {
	cRet := C.SDL_HapticNumEffectsPlaying(cHaptic(h))
	return int(cRet)
}

func SDL_HapticQuery(h *SDL_Haptic) uint32 {
	cRet := C.SDL_HapticQuery(cHaptic(h))
	return uint32(cRet)
}

func SDL_HapticEffectSupported(h *SDL_Haptic, he SDL_HapticEffect) bool {
	cHe := cHapticEffect(&he)
	cRet := C.SDL_HapticEffectSupported(cHaptic(h), cHe)
	return int(cRet) == 1
}

func SDL_HapticNewEffect(h *SDL_Haptic, he SDL_HapticEffect) int {
	cHe := cHapticEffect(&he)
	cRet := C.SDL_HapticNewEffect(cHaptic(h), cHe)
	return int(cRet)
}

func SDL_HapticUpdateEffect(h *SDL_Haptic, effect int, data SDL_HapticEffect) int {
	cData := cHapticEffect(&data)
	cRet := C.SDL_HapticUpdateEffect(cHaptic(h), cInt(effect), cData)
	return int(cRet)
}

func SDL_HapticRunEffect(h *SDL_Haptic, effect int, iterations uint32) int {
	cRet := C.SDL_HapticRunEffect(cHaptic(h), cInt(effect), cUint32(iterations))
	return int(cRet)
}

func SDL_HapticStopEffect(h *SDL_Haptic, effect int) int {
	cRet := C.SDL_HapticStopEffect(cHaptic(h), cInt(effect))
	return int(cRet)
}

func SDL_HapticDestroyEffect(h *SDL_Haptic, effect int) {
	C.SDL_HapticDestroyEffect(cHaptic(h), cInt(effect))
}

func SDL_HapticGetEffectStatus(h *SDL_Haptic, effect int) int {
	cRet := C.SDL_HapticGetEffectStatus(cHaptic(h), cInt(effect))
	return int(cRet)

}

func SDL_HapticSetGain(h *SDL_Haptic, gain int) int {
	cRet := C.SDL_HapticSetGain(cHaptic(h), cInt(gain))
	return int(cRet)
}

func SDL_HapticSetAutocenter(h *SDL_Haptic, autocenter int) int {
	cRet := C.SDL_HapticSetAutocenter(cHaptic(h), cInt(autocenter))
	return int(cRet)
}

func SDL_HapticPause(h *SDL_Haptic) int {
	cRet := C.SDL_HapticPause(cHaptic(h))
	return int(cRet)
}

func SDL_HapticUnpause(h *SDL_Haptic) int {
	cRet := C.SDL_HapticUnpause(cHaptic(h))
	return int(cRet)
}

func SDL_HapticStopAll(h *SDL_Haptic) int {
	cRet := C.SDL_HapticStopAll(cHaptic(h))
	return int(cRet)
}

func SDL_HapticRumbleSupported(h *SDL_Haptic) bool {
	cRet := C.SDL_HapticRumbleSupported(cHaptic(h))
	return int(cRet) == 1
}

func SDL_HapticRumbleInit(h *SDL_Haptic) int {
	cRet := C.SDL_HapticRumbleInit(cHaptic(h))
	return int(cRet)
}

func SDL_HapticRumblePlay(h *SDL_Haptic, strength float32, length uint32) int {
	cRet := C.SDL_HapticRumblePlay(cHaptic(h), cFloat(strength), cUint32(length))
	return int(cRet)
}

func SDL_HapticRumbleStop(h *SDL_Haptic) int {
	cRet := C.SDL_HapticRumbleStop(cHaptic(h))
	return int(cRet)
}
