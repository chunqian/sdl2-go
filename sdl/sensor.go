package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"

// define
const (
	SDL_STANDARD_GRAVITY = 9.80665
)

// typedef
type SDL_SensorID = int32
type SDL_SensorType = int32

// enum
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

func cSensor(ser *SDL_Sensor) *C.SDL_Sensor {
	return (*C.SDL_Sensor)(unsafe.Pointer(ser))
}

func SDL_LockSensors() {
	C.SDL_LockSensors()
}

func SDL_UnlockSensors() {
	C.SDL_UnlockSensors()
}

func SDL_NumSensors() int {
	return int(C.SDL_NumSensors())
}

func SDL_SensorGetDeviceName(deviceIndex int) string {
	cName := C.SDL_SensorGetDeviceName(cInt(deviceIndex))
	return createGoString(cName)
}

func SDL_SensorGetDeviceType(deviceIndex int) SDL_SensorType {
	cType := C.SDL_SensorGetDeviceType(cInt(deviceIndex))
	return SDL_SensorType(cType)
}

func SDL_SensorGetDeviceNonPortableType(deviceIndex int) int {
	cType := C.SDL_SensorGetDeviceNonPortableType(cInt(deviceIndex))
	return int(cType)
}

func SDL_SensorGetDeviceInstanceID(deviceIndex int) SDL_SensorID {
	cId := C.SDL_SensorGetDeviceInstanceID(cInt(deviceIndex))
	return SDL_SensorID(cId)
}

func SDL_SensorOpen(deviceIndex int) *SDL_Sensor {
	cSensor := C.SDL_SensorOpen(cInt(deviceIndex))
	return (*SDL_Sensor)(unsafe.Pointer(cSensor))
}

func SDL_SensorFromInstanceID(id SDL_SensorID) *SDL_Sensor {
	cSensor := C.SDL_SensorFromInstanceID(cSensorID(id))
	return (*SDL_Sensor)(unsafe.Pointer(cSensor))
}

func SDL_SensorGetName(sensor *SDL_Sensor) string {
	cName := C.SDL_SensorGetName(cSensor(sensor))
	return createGoString(cName)
}

func SDL_SensorGetType(sensor *SDL_Sensor) SDL_SensorType {
	cType := C.SDL_SensorGetType(cSensor(sensor))
	return SDL_SensorType(cType)
}

func SDL_SensorGetNonPortableType(sensor *SDL_Sensor) int {
	cRet := C.SDL_SensorGetNonPortableType(cSensor(sensor))
	return int(cRet)
}

func SDL_SensorGetInstanceID(sensor *SDL_Sensor) SDL_SensorID {
	cId := C.SDL_SensorGetInstanceID(cSensor(sensor))
	return SDL_SensorID(cId)
}

func SDL_SensorGetData(sensor *SDL_Sensor, data []float32, numValues int) int {
	if data == nil {
		return -1
	}
	cData := (*cFloat)(unsafe.Pointer(&data[0]))
	// cNumValues := cInt(len(data))
	cNumValues := cInt(numValues)

	cRet := int(C.SDL_SensorGetData(cSensor(sensor), cData, cNumValues))
	return int(cRet)
}

func SDL_SensorClose(sensor *SDL_Sensor) {
	C.SDL_SensorClose(cSensor(sensor))
}

func SDL_SensorUpdate() {
	C.SDL_SensorUpdate()
}
