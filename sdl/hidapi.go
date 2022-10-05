package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// struct
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

func cHidDevice(info *SDL_hid_device) *C.SDL_hid_device {
	return (*C.SDL_hid_device)(unsafe.Pointer(info))
}

func (info *SDL_hid_device_info) PathAsString() string {
	return SDL_GoString(info.Path)
}

func SDL_hid_init() int {
	cRet := C.SDL_hid_init()
	return int(cRet)
}

func SDL_hid_exit() int {
	cRet := C.SDL_hid_exit()
	return int(cRet)
}

func SDL_hid_device_change_count() (n uint32) {
	cRet := C.SDL_hid_device_change_count()
	return uint32(cRet)
}

func SDL_hid_enumerate(vendorID, productID uint16) *SDL_hid_device_info {
	cVendorID := cUint16(vendorID)
	cProductID := cUint16(productID)

	cInfo := C.SDL_hid_enumerate(cVendorID, cProductID)
	return (*SDL_hid_device_info)(unsafe.Pointer(cInfo))
}

func SDL_hid_free_enumeration(info *SDL_hid_device_info) {
	cInfo := (*C.SDL_hid_device_info)(unsafe.Pointer(info))
	C.SDL_hid_free_enumeration(cInfo)
}

func SDL_hid_open[T []uint16 | []int32](vendorID, productID uint16, serialNumber T) *SDL_hid_device {
	cVendorID := cUint16(vendorID)
	cProductID := cUint16(productID)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&serialNumber))
	cSerialNumber := (*cWchar)(unsafe.Pointer(sh.Data))

	cDevice := C.SDL_hid_open(cVendorID, cProductID, cSerialNumber)
	return (*SDL_hid_device)(unsafe.Pointer(cDevice))
}

func SDL_hid_open_path(path string, exclusive SDL_bool) *SDL_hid_device {
	cPath := SDL_CreateCString(SDL_GetMemoryPool(), path)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cPath)

	cDevice := C.SDL_hid_open_path(cPath.(*cChar), cInt(exclusive))
	return (*SDL_hid_device)(unsafe.Pointer(cDevice))
}

func SDL_hid_write(device *SDL_hid_device, data []byte) int {
	cLength := cSize(len(data))
	cData := (*cUchar)(unsafe.Pointer(&data[0]))
	cDevice := cHidDevice(device)

	cRet := C.SDL_hid_write(cDevice, cData, cLength)
	return int(cRet)
}

func SDL_hid_read_timeout(device *SDL_hid_device, data []byte, milliseconds int) int {
	cLength := cSize(len(data))
	cData := (*cUchar)(unsafe.Pointer(&data[0]))
	cDevice := cHidDevice(device)
	cMilliseconds := cInt(milliseconds)

	cRet := C.SDL_hid_read_timeout(cDevice, cData, cLength, cMilliseconds)
	return int(cRet)
}

func SDL_hid_read(device *SDL_hid_device, data []byte) int {
	cLength := cSize(len(data))
	cData := (*cUchar)(unsafe.Pointer(&data[0]))
	cDevice := cHidDevice(device)

	cRet := C.SDL_hid_read(cDevice, cData, cLength)
	return int(cRet)
}

func SDL_hid_set_nonblocking(device *SDL_hid_device, nonblock SDL_bool) int {
	cDevice := cHidDevice(device)
	cNonblock := cInt(nonblock)

	cRet := C.SDL_hid_set_nonblocking(cDevice, cNonblock)
	return int(cRet)
}

func SDL_hid_send_feature_report(device *SDL_hid_device, data []byte) int {
	cLength := cSize(len(data))
	cData := (*cUchar)(unsafe.Pointer(&data[0]))
	cDevice := cHidDevice(device)

	cRet := C.SDL_hid_send_feature_report(cDevice, cData, cLength)
	return int(cRet)
}

func SDL_hid_get_feature_report(device *SDL_hid_device, data []byte) int {
	cLength := cSize(len(data))
	cData := (*cUchar)(unsafe.Pointer(&data[0]))
	cDevice := cHidDevice(device)

	cRet := C.SDL_hid_get_feature_report(cDevice, cData, cLength)
	return int(cRet)
}

func SDL_hid_close(device *SDL_hid_device) {
	cDevice := cHidDevice(device)
	C.SDL_hid_close(cDevice)
}

func SDL_hid_get_manufacturer_string[T []uint16 | []int32](device *SDL_hid_device, str T, maxlen uint) int {
	cDevice := cHidDevice(device)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	cStr := (*cWchar)(unsafe.Pointer(sh.Data))
	cMaxlen := cSize(maxlen)

	cRet := C.SDL_hid_get_manufacturer_string(cDevice, cStr, cMaxlen)
	return int(cRet)
}

func SDL_hid_get_product_string[T []uint16 | []int32](device *SDL_hid_device, str T, maxlen uint) int {
	cDevice := cHidDevice(device)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	cStr := (*cWchar)(unsafe.Pointer(sh.Data))
	cMaxlen := cSize(maxlen)

	cRet := C.SDL_hid_get_product_string(cDevice, cStr, cMaxlen)
	return int(cRet)
}

func SDL_hid_get_serial_number_string[T []uint16 | []int32](device *SDL_hid_device, str T, maxlen uint) int {
	cDevice := cHidDevice(device)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	cStr := (*cWchar)(unsafe.Pointer(sh.Data))
	cMaxlen := cSize(maxlen)

	cRet := C.SDL_hid_get_serial_number_string(cDevice, cStr, cMaxlen)
	return int(cRet)
}

func SDL_hid_get_indexed_string[T []uint16 | []int32](device *SDL_hid_device, index int, str T, maxlen uint) int {
	cDevice := cHidDevice(device)
	cIndex := cInt(index)

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&str))
	cStr := (*cWchar)(unsafe.Pointer(sh.Data))
	cMaxlen := cSize(maxlen)

	cRet := C.SDL_hid_get_indexed_string(cDevice, cIndex, cStr, cMaxlen)
	return int(cRet)
}

func SDL_hid_ble_scan(device *SDL_hid_device, active bool) {
	cActive := C.SDL_bool(Btoi(active))
	C.SDL_hid_ble_scan(cActive)
}
