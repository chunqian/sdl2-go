package sdl

/*
#include "gamecontroller.h"
*/
import "C"
import (
	"encoding/binary"
	"unsafe"
)

func cGameController(ctrl *SDL_GameController) *C.SDL_GameController {
	return (*C.SDL_GameController)(unsafe.Pointer(ctrl))
}

func SDL_GameControllerAddMapping(mappingString string) int {
	cMappingString := SDL_CreateCString(SDL_GetMemoryPool(), mappingString)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cMappingString)

	cRet := C.SDL_GameControllerAddMapping(cMappingString.(*cChar))
	return int(cRet)
}

func SDL_GameControllerNumMappings() int {
	cRet := C.SDL_GameControllerNumMappings()
	return int(cRet)
}

func SDL_GameControllerMappingForIndex(index int) string {
	cMappingString := C.SDL_GameControllerMappingForIndex(cInt(index))
	defer C.free(unsafe.Pointer(cMappingString))

	return SDL_GoString(cMappingString)
}

func SDL_GameControllerMappingForGUID(guid SDL_JoystickGUID) string {
	cGuid := cJoystickGUID(&guid)
	cMappingString := C.SDL_GameControllerMappingForGUID(*cGuid)
	defer C.free(unsafe.Pointer(cMappingString))

	return SDL_GoString(cMappingString)
}

func SDL_IsGameController(index int) bool {
	cRet := C.SDL_IsGameController(cInt(index))
	return int(cRet) == 1
}

func SDL_GameControllerNameForIndex(index int) string {
	cRet := C.SDL_GameControllerNameForIndex(cInt(index))
	return SDL_GoString(cRet)
}

func SDL_GameControllerPathForIndex(index int) string {
	cRet := C.SDL_GameControllerPathForIndex(cInt(index))
	return SDL_GoString(cRet)
}

func SDL_GameControllerTypeForIndex(index int) SDL_GameControllerType {
	cRet := C.SDL_GameControllerTypeForIndex(cInt(index))
	return SDL_GameControllerType(cRet)
}

func SDL_GameControllerMappingForDeviceIndex(index int) string {
	cMappingString := C.SDL_GameControllerMappingForDeviceIndex(cInt(index))
	defer C.free(unsafe.Pointer(cMappingString))

	return SDL_GoString(cMappingString)
}

func SDL_GameControllerOpen(index int) *SDL_GameController {
	cCtrl := C.SDL_GameControllerOpen(cInt(index))
	return (*SDL_GameController)(unsafe.Pointer(cCtrl))
}

func SDL_GameControllerFromInstanceID(joyid SDL_JoystickID) *SDL_GameController {
	cCtrl := C.SDL_GameControllerFromInstanceID(cJoystickID(joyid))
	return (*SDL_GameController)(unsafe.Pointer(cCtrl))
}

func SDL_GameControllerFromPlayerIndex(playerIndex int) *SDL_GameController {
	cCtrl := C.SDL_GameControllerFromPlayerIndex(cInt(playerIndex))
	return (*SDL_GameController)(unsafe.Pointer(cCtrl))
}

func SDL_GameControllerName(ctrl *SDL_GameController) string {
	cName := C.SDL_GameControllerName(cGameController(ctrl))
	return SDL_GoString(cName)
}

func SDL_GameControllerPath(ctrl *SDL_GameController) string {
	cPath := C.SDL_GameControllerPath(cGameController(ctrl))
	return SDL_GoString(cPath)
}

func SDL_GameControllerGetType(ctrl *SDL_GameController) SDL_GameControllerType {
	cType := C.SDL_GameControllerGetType(cGameController(ctrl))
	return SDL_GameControllerType(cType)
}

func SDL_GameControllerGetPlayerIndex(ctrl *SDL_GameController) int {
	cRet := C.SDL_GameControllerGetPlayerIndex(cGameController(ctrl))
	return int(cRet)
}

func SDL_GameControllerSetPlayerIndex(ctrl *SDL_GameController, playerIndex int) {
	C.SDL_GameControllerSetPlayerIndex(cGameController(ctrl), cInt(playerIndex))
}

func SDL_GameControllerGetVendor(ctrl *SDL_GameController) int {
	cRet := C.SDL_GameControllerGetVendor(cGameController(ctrl))
	return int(cRet)
}

func SDL_GameControllerGetProduct(ctrl *SDL_GameController) int {
	cRet := C.SDL_GameControllerGetProduct(cGameController(ctrl))
	return int(cRet)
}

func SDL_GameControllerGetProductVersion(ctrl *SDL_GameController) int {
	cRet := C.SDL_GameControllerGetProductVersion(cGameController(ctrl))
	return int(cRet)
}

func SDL_GameControllerGetFirmwareVersion(ctrl *SDL_GameController) uint16 {
	cRet := C.SDL_GameControllerGetFirmwareVersion(cGameController(ctrl))
	return uint16(cRet)
}

func SDL_GameControllerGetSerial(ctrl *SDL_GameController) string {
	cSerial := C.SDL_GameControllerGetSerial(cGameController(ctrl))
	return SDL_GoString(cSerial)
}

func SDL_GameControllerGetAttached(ctrl *SDL_GameController) bool {
	cRet := C.SDL_GameControllerGetAttached(cGameController(ctrl))
	return int(cRet) == 1
}

func SDL_GameControllerMapping(ctrl *SDL_GameController) string {
	cMappingString := C.SDL_GameControllerMapping(cGameController(ctrl))
	defer C.free(unsafe.Pointer(cMappingString))

	return SDL_GoString(cMappingString)
}

func SDL_GameControllerGetJoystick(ctrl *SDL_GameController) *SDL_Joystick {
	cJs := C.SDL_GameControllerGetJoystick(cGameController(ctrl))
	return (*SDL_Joystick)(unsafe.Pointer(cJs))
}

func SDL_GameControllerEventState(state int) int {
	cRet := C.SDL_GameControllerEventState(cInt(state))
	return int(cRet)
}

func SDL_GameControllerUpdate() {
	C.SDL_GameControllerUpdate()
}

func SDL_GameControllerGetAxisFromString(pchString string) SDL_GameControllerAxis {
	cPchString := SDL_CreateCString(SDL_GetMemoryPool(), pchString)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cPchString)

	cRet := C.SDL_GameControllerGetAxisFromString(cPchString.(*cChar))
	return SDL_GameControllerAxis(cRet)
}

func SDL_GameControllerGetStringForAxis(axis SDL_GameControllerAxis) string {
	cRet := C.SDL_GameControllerGetStringForAxis(cGameControllerAxis(axis))
	return SDL_GoString(cRet)
}

func SDL_GameControllerGetBindForAxis(ctrl *SDL_GameController, axis SDL_GameControllerAxis) SDL_GameControllerButtonBind {
	cRet := C.SDL_GameControllerGetBindForAxis(cGameController(ctrl), cGameControllerAxis(axis))
	return *(*SDL_GameControllerButtonBind)(unsafe.Pointer(&cRet))
}

func SDL_GameControllerHasAxis(ctrl *SDL_GameController, axis SDL_GameControllerAxis) bool {
	cRet := C.SDL_GameControllerHasAxis(cGameController(ctrl), cGameControllerAxis(axis))
	return int(cRet) == 1
}

func SDL_GameControllerGetAxis(ctrl *SDL_GameController, axis SDL_GameControllerAxis) int16 {
	cRet := C.SDL_GameControllerGetAxis(cGameController(ctrl), cGameControllerAxis(axis))
	return int16(cRet)
}

func SDL_GameControllerGetButtonFromString(pchString string) SDL_GameControllerButton {
	cPchString := SDL_CreateCString(SDL_GetMemoryPool(), pchString)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cPchString)

	cRet := C.SDL_GameControllerGetButtonFromString(cPchString.(*cChar))
	return SDL_GameControllerButton(cRet)
}

func SDL_GameControllerGetStringForButton(btn SDL_GameControllerButton) string {
	cRet := C.SDL_GameControllerGetStringForButton(cGameControllerButton(btn))
	return SDL_GoString(cRet)
}

func SDL_GameControllerGetBindForButton(ctrl *SDL_GameController, btn SDL_GameControllerButton) SDL_GameControllerButtonBind {
	cBind := C.SDL_GameControllerGetBindForButton(cGameController(ctrl), cGameControllerButton(btn))
	return *(*SDL_GameControllerButtonBind)(unsafe.Pointer(&cBind))
}

func SDL_GameControllerHasButton(ctrl *SDL_GameController, btn SDL_GameControllerButton) bool {
	cRet := C.SDL_GameControllerHasButton(cGameController(ctrl), cGameControllerButton(btn))
	return int(cRet) == 1
}

func SDL_GameControllerRumble(ctrl *SDL_GameController, lowFrequencyRumble, highFrequencyRumble uint16, durationMS uint32) int {
	cRet := C.SDL_GameControllerRumble(cGameController(ctrl), cUint16(lowFrequencyRumble), cUint16(highFrequencyRumble), cUint32(durationMS))
	return int(cRet)
}

func SDL_GameControllerRumbleTriggers(ctrl *SDL_GameController, leftRumble, rightRumble uint16, durationMS uint32) int {
	cRet := C.SDL_GameControllerRumbleTriggers(cGameController(ctrl), cUint16(leftRumble), cUint16(rightRumble), cUint32(durationMS))
	return int(cRet)
}

func SDL_GameControllerHasLED(ctrl *SDL_GameController) bool {
	cRet := C.SDL_GameControllerHasLED(cGameController(ctrl))
	return int(cRet) == 1
}

func SDL_GameControllerSetLED(ctrl *SDL_GameController, red, green, blue uint8) int {
	cRet := C.SDL_GameControllerSetLED(cGameController(ctrl), cUint8(red), cUint8(blue), cUint8(green))
	return int(cRet)
}

func SDL_GameControllerGetButton(ctrl *SDL_GameController, btn SDL_GameControllerButton) byte {
	cRet := C.SDL_GameControllerGetButton(cGameController(ctrl), cGameControllerButton(btn))
	return byte(cRet)
}

func SDL_GameControllerGetNumTouchpads(ctrl *SDL_GameController) int {
	cRet := C.SDL_GameControllerGetNumTouchpads(cGameController(ctrl))
	return int(cRet)
}

func SDL_GameControllerGetNumTouchpadFingers(ctrl *SDL_GameController, touchpad int) int {
	cRet := C.SDL_GameControllerGetNumTouchpadFingers(cGameController(ctrl), cInt(touchpad))
	return int(cRet)
}

func SDL_GameControllerGetTouchpadFinger(ctrl *SDL_GameController, touchpad, finger int, state *uint8, x, y, pressure *float32) int {
	cState := (*cUint8)(unsafe.Pointer(state))
	cX := (*cFloat)(unsafe.Pointer(x))
	cY := (*cFloat)(unsafe.Pointer(y))
	cPressure := (*cFloat)(unsafe.Pointer(pressure))

	cRet := C.SDL_GameControllerGetTouchpadFinger(cGameController(ctrl), cInt(touchpad), cInt(finger), cState, cX, cY, cPressure)
	return int(cRet)
}

func SDL_GameControllerHasSensor(ctrl *SDL_GameController, typ SDL_SensorType) bool {
	cRet := C.SDL_GameControllerHasSensor(cGameController(ctrl), C.SDL_SensorType(typ))
	return int(cRet) == 1
}

func SDL_GameControllerSetSensorEnabled(ctrl *SDL_GameController, typ SDL_SensorType, enabled bool) int {
	cRet := C.SDL_GameControllerSetSensorEnabled(cGameController(ctrl), C.SDL_SensorType(typ), C.SDL_bool(Btoi(enabled)))
	return int(cRet)
}

func SDL_GameControllerIsSensorEnabled(ctrl *SDL_GameController, typ SDL_SensorType) bool {
	cRet := C.SDL_GameControllerIsSensorEnabled(cGameController(ctrl), C.SDL_SensorType(typ))
	return int(cRet) == 1
}

func SDL_GameControllerGetSensorData(ctrl *SDL_GameController, typ SDL_SensorType, data []float32, numValues int) int {
	cData := (*cFloat)(unsafe.Pointer(&data))
	cRet := C.SDL_GameControllerGetSensorData(cGameController(ctrl), C.SDL_SensorType(typ), cData, cInt(numValues))
	return int(cRet)
}

func SDL_GameControllerClose(ctrl *SDL_GameController) {
	C.SDL_GameControllerClose(cGameController(ctrl))
}

func (bind *SDL_GameControllerButtonBind) Type() int {
	return int(bind.BindType)
}

func (bind *SDL_GameControllerButtonBind) Button() int {
	val, _ := binary.Varint(bind.Value[:4])
	return int(val)
}

func (bind *SDL_GameControllerButtonBind) Axis() int {
	val, _ := binary.Varint(bind.Value[:4])
	return int(val)
}

func (bind *SDL_GameControllerButtonBind) Hat() int {
	val, _ := binary.Varint(bind.Value[:4])
	return int(val)
}

func (bind *SDL_GameControllerButtonBind) HatMask() int {
	val, _ := binary.Varint(bind.Value[4:8])
	return int(val)
}

func SDL_GameControllerSendEffect(ctrl *SDL_GameController, data []byte) int {
	cData := unsafe.Pointer(&data)
	cSize := cInt(len(data))

	cRet := C.SDL_GameControllerSendEffect(cGameController(ctrl), cData, cSize)
	return int(cRet)
}

func SDL_GameControllerGetSensorDataRate(ctrl *SDL_GameController, typ SDL_SensorType) (rate float32) {
	cRet := C.SDL_GameControllerGetSensorDataRate(cGameController(ctrl), C.SDL_SensorType(typ))
	return float32(cRet)
}

func SDL_GameControllerHasRumble(ctrl *SDL_GameController) bool {
	cRet := C.SDL_GameControllerHasRumble(cGameController(ctrl))
	return int(cRet) == 1
}

func SDL_GameControllerHasRumbleTriggers(ctrl *SDL_GameController) bool {
	cRet := C.SDL_GameControllerHasRumbleTriggers(cGameController(ctrl))
	return int(cRet) == 1
}

func SDL_GameControllerGetAppleSFSymbolsNameForButton(ctrl *SDL_GameController, button SDL_GameControllerButton) (sfSymbolsName string) {
	cButton := cGameControllerButton(button)
	cSfSymbolsName := C.SDL_GameControllerGetAppleSFSymbolsNameForButton(cGameController(ctrl), cButton)
	sfSymbolsName = SDL_GoString(cSfSymbolsName)
	return
}

func SDL_GameControllerGetAppleSFSymbolsNameForAxis(ctrl *SDL_GameController, axis SDL_GameControllerAxis) (sfSymbolsName string) {
	cAxis := cGameControllerAxis(axis)
	cSfSymbolsName := C.SDL_GameControllerGetAppleSFSymbolsNameForAxis(cGameController(ctrl), cAxis)
	sfSymbolsName = SDL_GoString(cSfSymbolsName)
	return
}
