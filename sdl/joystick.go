package sdl

/*
#include "joystick.h"
*/
import "C"
import "unsafe"

// define
const (
	SDL_JOYSTICK_AXIS_MAX = 32767
	SDL_JOYSTICK_AXIS_MIN = -32768

	SDL_HAT_CENTERED  = 0x00
	SDL_HAT_UP        = 0x01
	SDL_HAT_RIGHT     = 0x02
	SDL_HAT_DOWN      = 0x04
	SDL_HAT_LEFT      = 0x08
	SDL_HAT_RIGHTUP   = SDL_HAT_RIGHT | SDL_HAT_UP
	SDL_HAT_RIGHTDOWN = SDL_HAT_RIGHT | SDL_HAT_DOWN
	SDL_HAT_LEFTUP    = SDL_HAT_LEFT | SDL_HAT_UP
	SDL_HAT_LEFTDOWN  = SDL_HAT_LEFT | SDL_HAT_DOWN
)

// typedef
type SDL_GUID = SDL_JoystickGUID
type SDL_JoystickID = int32
type SDL_JoystickType = int32
type SDL_JoystickPowerLevel = int32

// enum
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

const (
	SDL_JOYSTICK_POWER_UNKNOWN SDL_JoystickPowerLevel = -1
	SDL_JOYSTICK_POWER_EMPTY   SDL_JoystickPowerLevel = 0
	SDL_JOYSTICK_POWER_LOW     SDL_JoystickPowerLevel = 1
	SDL_JOYSTICK_POWER_MEDIUM  SDL_JoystickPowerLevel = 2
	SDL_JOYSTICK_POWER_FULL    SDL_JoystickPowerLevel = 3
	SDL_JOYSTICK_POWER_WIRED   SDL_JoystickPowerLevel = 4
	SDL_JOYSTICK_POWER_MAX     SDL_JoystickPowerLevel = 5
)

// struct
type SDL_JoystickGUID struct {
	Data [16]uint8
}

func cJoystick(js *SDL_Joystick) *C.SDL_Joystick {
	return (*C.SDL_Joystick)(unsafe.Pointer(js))
}

func cJoystickGUID(guid *SDL_JoystickGUID) *C.SDL_JoystickGUID {
	return (*C.SDL_JoystickGUID)(unsafe.Pointer(guid))
}

func SDL_NumJoysticks() int {
	return int(C.SDL_NumJoysticks())
}

func SDL_JoystickNameForIndex(index int) string {
	cStr := C.SDL_JoystickNameForIndex(cInt(index))
	return createGoString(cStr)
}

func SDL_JoystickPathForIndex(index int) string {
	cStr := C.SDL_JoystickPathForIndex(cInt(index))
	return createGoString(cStr)
}

func SDL_JoystickGetDevicePlayerIndex(index int) int {
	cRet := C.SDL_JoystickGetDevicePlayerIndex(cInt(index))
	return int(cRet)
}

func SDL_JoystickGetDeviceGUID(index int) SDL_JoystickGUID {
	cJoyGuid := C.SDL_JoystickGetDeviceGUID(cInt(index))
	return *(*SDL_JoystickGUID)(unsafe.Pointer(&cJoyGuid))
}

func SDL_JoystickGetDeviceVendor(index int) int {
	cRet := C.SDL_JoystickGetDeviceVendor(cInt(index))
	return int(cRet)
}

func SDL_JoystickGetDeviceProduct(index int) int {
	cRet := C.SDL_JoystickGetDeviceProduct(cInt(index))
	return int(cRet)
}

func SDL_JoystickGetDeviceProductVersion(index int) int {
	cRet := C.SDL_JoystickGetDeviceProductVersion(cInt(index))
	return int(cRet)
}

func SDL_JoystickGetDeviceType(index int) SDL_JoystickType {
	cJoy := C.SDL_JoystickGetDeviceType(cInt(index))
	return SDL_JoystickType(cJoy)
}

func SDL_JoystickGetDeviceInstanceID(index int) SDL_JoystickID {
	cJoyId := C.SDL_JoystickGetDeviceInstanceID(cInt(index))
	return SDL_JoystickID(cJoyId)
}

func SDL_JoystickGetGUIDString(guid SDL_JoystickGUID, pszGUID []byte, cbGUID int) string {
	cGuid := cJoystickGUID(&guid)
	cPszGUID := (*cChar)(unsafe.Pointer(&pszGUID[0]))
	C.SDL_JoystickGetGUIDString(*cGuid, cPszGUID, cInt(cbGUID))
	return createGoString(cPszGUID)
}

func SDL_JoystickGetGUIDFromString(pchGUID string) SDL_JoystickGUID {
	cPchGUID := createCString(SDL_GetMemoryPool(), pchGUID)
	defer destroyCString(SDL_GetMemoryPool(), cPchGUID)

	cJoyGuid := C.SDL_JoystickGetGUIDFromString(cPchGUID)
	return *(*SDL_JoystickGUID)(unsafe.Pointer(&cJoyGuid))
}

func SDL_JoystickUpdate() {
	C.SDL_JoystickUpdate()
}

func SDL_JoystickEventState(state int) int {
	cRet := C.SDL_JoystickEventState(cInt(state))
	return int(cRet)
}

func SDL_JoystickOpen(index int) *SDL_Joystick {
	cJoy := C.SDL_JoystickOpen(cInt(index))
	return (*SDL_Joystick)(unsafe.Pointer(cJoy))
}

func SDL_JoystickFromInstanceID(joyid SDL_JoystickID) *SDL_Joystick {
	cJoy := C.SDL_JoystickFromInstanceID(cJoystickID(joyid))
	return (*SDL_Joystick)(unsafe.Pointer(cJoy))
}

func SDL_JoystickFromPlayerIndex(playerIndex int) *SDL_Joystick {
	cJoy := C.SDL_JoystickFromPlayerIndex(cInt(playerIndex))
	return (*SDL_Joystick)(unsafe.Pointer(cJoy))
}

func SDL_JoystickAttachVirtual(typ SDL_JoystickType, naxes, nbuttons, nhats int) (deviceIndex int) {
	cRet := C.SDL_JoystickAttachVirtual(C.SDL_JoystickType(typ), cInt(naxes), cInt(nbuttons), cInt(nhats))
	deviceIndex = int(cRet)
	return
}

func SDL_JoystickAttachVirtualEx(desc *C.SDL_VirtualJoystickDesc) (deviceIndex int) {
	cRet := C.SDL_JoystickAttachVirtualEx(desc)
	deviceIndex = int(cRet)
	return
}

func SDL_JoystickDetachVirtual(deviceIndex int) int {
	cRet := C.SDL_JoystickDetachVirtual(cInt(deviceIndex))
	return int(cRet)
}

func SDL_JoystickIsVirtual(deviceIndex int) bool {
	cRet := C.SDL_JoystickIsVirtual(cInt(deviceIndex))
	return int(cRet) == 1
}

func SDL_JoystickSetVirtualAxis(joy *SDL_Joystick, axis int, value int16) int {
	cRet := C.SDL_JoystickSetVirtualAxis(cJoystick(joy), cInt(axis), cInt16(value))
	return int(cRet)
}

func SDL_JoystickSetVirtualButton(joy *SDL_Joystick, button int, value uint8) int {
	cRet := C.SDL_JoystickSetVirtualButton(cJoystick(joy), cInt(button), cUint8(value))
	return int(cRet)
}

func SDL_JoystickSetVirtualHat(joy *SDL_Joystick, hat int, value uint8) int {
	cRet := C.SDL_JoystickSetVirtualHat(cJoystick(joy), cInt(hat), cUint8(value))
	return int(cRet)
}

func SDL_LockJoysticks() {
	C.SDL_LockJoysticks()
}

func SDL_UnlockJoysticks() {
	C.SDL_UnlockJoysticks()
}

func SDL_JoystickName(joy *SDL_Joystick) string {
	cName := C.SDL_JoystickName(cJoystick(joy))
	return createGoString(cName)
}

func SDL_JoystickPath(joy *SDL_Joystick) string {
	cPath := C.SDL_JoystickPath(cJoystick(joy))
	return createGoString(cPath)
}

func SDL_JoystickGetPlayerIndex(joy *SDL_Joystick) int {
	cRet := C.SDL_JoystickGetPlayerIndex(cJoystick(joy))
	return int(cRet)
}

func SDL_JoystickSetPlayerIndex(joy *SDL_Joystick, playerIndex int) {
	C.SDL_JoystickSetPlayerIndex(cJoystick(joy), cInt(playerIndex))
}

func SDL_JoystickGetGUID(joy *SDL_Joystick) SDL_JoystickGUID {
	cJoyId := C.SDL_JoystickGetGUID(cJoystick(joy))
	return *(*SDL_JoystickGUID)(unsafe.Pointer(&cJoyId))
}

func SDL_JoystickGetVendor(joy *SDL_Joystick) int {
	cRet := C.SDL_JoystickGetVendor(cJoystick(joy))
	return int(cRet)
}

func SDL_JoystickGetProduct(joy *SDL_Joystick) int {
	cRet := C.SDL_JoystickGetProduct(cJoystick(joy))
	return int(cRet)
}

func SDL_JoystickGetProductVersion(joy *SDL_Joystick) int {
	cRet := C.SDL_JoystickGetProductVersion(cJoystick(joy))
	return int(cRet)
}

func SDL_JoystickGetFirmwareVersion(joy *SDL_Joystick) uint16 {
	cRet := C.SDL_JoystickGetFirmwareVersion(cJoystick(joy))
	return uint16(cRet)
}

func SDL_JoystickGetSerial(joy *SDL_Joystick) string {
	cSerial := C.SDL_JoystickGetSerial(cJoystick(joy))
	return createGoString(cSerial)
}

func SDL_JoystickGetType(joy *SDL_Joystick) SDL_JoystickType {
	cType := C.SDL_JoystickGetType(cJoystick(joy))
	return SDL_JoystickType(cType)
}

func SDL_JoystickGetAttached(joy *SDL_Joystick) bool {
	cRet := C.SDL_JoystickGetAttached(cJoystick(joy))
	return int(cRet) == 1
}

func SDL_JoystickInstanceID(joy *SDL_Joystick) SDL_JoystickID {
	cJoyId := C.SDL_JoystickInstanceID(cJoystick(joy))
	return SDL_JoystickID(cJoyId)
}

func SDL_JoystickNumAxes(joy *SDL_Joystick) int {
	cRet := C.SDL_JoystickNumAxes(cJoystick(joy))
	return int(cRet)
}

func SDL_JoystickNumBalls(joy *SDL_Joystick) int {
	cRet := C.SDL_JoystickNumBalls(cJoystick(joy))
	return int(cRet)
}

func SDL_JoystickNumHats(joy *SDL_Joystick) int {
	cRet := C.SDL_JoystickNumHats(cJoystick(joy))
	return int(cRet)
}

func SDL_JoystickNumButtons(joy *SDL_Joystick) int {
	cRet := C.SDL_JoystickNumButtons(cJoystick(joy))
	return int(cRet)
}

func SDL_JoystickGetAxis(joy *SDL_Joystick, axis int) int16 {
	cRet := C.SDL_JoystickGetAxis(cJoystick(joy), cInt(axis))
	return int16(cRet)
}

func SDL_JoystickGetAxisInitialState(joy *SDL_Joystick, axis int, state *int16) (ok bool) {
	cState := (*cInt16)(unsafe.Pointer(state))
	cOk := C.SDL_JoystickGetAxisInitialState(cJoystick(joy), cInt(axis), cState)
	return int(cOk) == 1
}

func SDL_JoystickGetHat(joy *SDL_Joystick, hat int) byte {
	return byte(C.SDL_JoystickGetHat(cJoystick(joy), cInt(hat)))
}

func SDL_JoystickGetBall(joy *SDL_Joystick, ball int, dx, dy *int32) int {
	cDx := (*cInt)(unsafe.Pointer(dx))
	cDy := (*cInt)(unsafe.Pointer(dy))

	cRet := C.SDL_JoystickGetBall(cJoystick(joy), cInt(ball), cDx, cDy)
	return int(cRet)
}

func SDL_JoystickGetButton(joy *SDL_Joystick, button int) byte {
	return byte(C.SDL_JoystickGetButton(cJoystick(joy), cInt(button)))
}

func SDL_JoystickRumble(joy *SDL_Joystick, lowFrequencyRumble, highFrequencyRumble uint16, durationMS uint32) int {
	cRet := C.SDL_JoystickRumble(cJoystick(joy), cUint16(lowFrequencyRumble), cUint16(highFrequencyRumble), cUint32(durationMS))
	return int(cRet)
}

func SDL_JoystickRumbleTriggers(joy *SDL_Joystick, leftRumble, rightRumble uint16, durationMS uint32) int {
	cRet := C.SDL_JoystickRumbleTriggers(cJoystick(joy), cUint16(leftRumble), cUint16(rightRumble), cUint32(durationMS))
	return int(cRet)
}

func SDL_JoystickHasLED(joy *SDL_Joystick) bool {
	cRet := C.SDL_JoystickHasLED(cJoystick(joy))
	return int(cRet) == 1
}

func SDL_JoystickSetLED(joy *SDL_Joystick, red, green, blue uint8) int {
	cRet := C.SDL_JoystickSetLED(cJoystick(joy), cUint8(red), cUint8(green), cUint8(blue))
	return int(cRet)
}

func SDL_JoystickClose(joy *SDL_Joystick) {
	C.SDL_JoystickClose(cJoystick(joy))
}

func SDL_JoystickCurrentPowerLevel(joy *SDL_Joystick) SDL_JoystickPowerLevel {
	cJoyLevel := C.SDL_JoystickCurrentPowerLevel(cJoystick(joy))
	return SDL_JoystickPowerLevel(cJoyLevel)
}

func SDL_JoystickHasRumble(joy *SDL_Joystick) bool {
	cRet := C.SDL_JoystickHasRumble(cJoystick(joy))
	return int(cRet) == 1
}

func SDL_JoystickHasRumbleTriggers(joy *SDL_Joystick) bool {
	cRet := C.SDL_JoystickHasRumbleTriggers(cJoystick(joy))
	return int(cRet) == 1
}
