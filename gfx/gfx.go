package gfx

/*
#include "gfx_wrapper.h"
*/
import "C"
import (
	"unsafe"

	. "github.com/chunqian/sdl2-go/sdl"
)

func (manager *FPSmanager) cFPSmanager() *C.FPSmanager {
	return (*C.FPSmanager)(unsafe.Pointer(manager))
}

func min(a ...int) int {
	b := a[0]
	for i := 1; i < len(a); i++ {
		if b > a[i] {
			b = a[i]
		}
	}
	return b
}

func gfxColor(color SDL_Color) uint32 {
	return (uint32(color.A) << 24) | (uint32(color.B) << 16) | (uint32(color.G) << 8) | uint32(color.R)
}

func SDL_initFramerate(manager *FPSmanager) {
	C.SDL_initFramerate(manager.cFPSmanager())
}

func SDL_setFramerate(manager *FPSmanager, rate uint32) bool {
	_rate := C.Uint32(rate)
	return C.SDL_setFramerate(manager.cFPSmanager(), _rate) == 0
}

func SDL_getFramerate(manager *FPSmanager) (int, bool) {
	fps := int(C.SDL_getFramerate(manager.cFPSmanager()))
	return fps, fps >= 0
}

func SDL_getFramecount(manager *FPSmanager) (int, bool) {
	count := int(C.SDL_getFramecount(manager.cFPSmanager()))
	return count, count >= 0
}

func SDL_framerateDelay(manager *FPSmanager) uint32 {
	return uint32(C.SDL_framerateDelay(manager.cFPSmanager()))
}

func PixelColor(renderer *SDL_Renderer, x, y int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_color := C.Uint32(gfxColor(color))
	return C.pixelColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _color) == 0
}

func PixelRGBA(renderer *SDL_Renderer, x, y int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.pixelRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _r, _g, _b, _a) == 0
}

func HlineColor(renderer *SDL_Renderer, x1, x2, y int32, color SDL_Color) bool {
	_x1 := C.Sint16(x1)
	_x2 := C.Sint16(x2)
	_y := C.Sint16(y)
	_color := C.Uint32(gfxColor(color))
	return C.hlineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _x2, _y, _color) == 0
}

func HlineRGBA(renderer *SDL_Renderer, x1, x2, y int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_x2 := C.Sint16(x2)
	_y := C.Sint16(y)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.hlineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _x2, _y, _r, _g, _b, _a) == 0
}

func VlineColor(renderer *SDL_Renderer, x, y1, y2 int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y1 := C.Sint16(y1)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.vlineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y1, _y2, _color) == 0
}

func VlineRGBA(renderer *SDL_Renderer, x, y1, y2 int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y1 := C.Sint16(y1)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.vlineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y1, _y2, _r, _g, _b, _a) == 0
}

func RectangleColor(renderer *SDL_Renderer, x1, y1, x2, y2 int32, color SDL_Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.rectangleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

func RectangleRGBA(renderer *SDL_Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.rectangleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

func RoundedRectangleColor(renderer *SDL_Renderer, x1, y1, x2, y2, rad int32, color SDL_Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.roundedRectangleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _rad, _color) == 0
}

func RoundedRectangleRGBA(renderer *SDL_Renderer, x1, y1, x2, y2, rad int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.roundedRectangleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _rad, _r, _g, _b, _a) == 0
}

func BoxColor(renderer *SDL_Renderer, x1, y1, x2, y2 int32, color SDL_Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.boxColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

func BoxRGBA(renderer *SDL_Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.boxRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

func RoundedBoxColor(renderer *SDL_Renderer, x1, y1, x2, y2, rad int32, color SDL_Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.roundedBoxColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _rad, _color) == 0
}

func RoundedBoxRGBA(renderer *SDL_Renderer, x1, y1, x2, y2, rad int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.roundedBoxRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _rad, _r, _g, _b, _a) == 0
}

func LineColor(renderer *SDL_Renderer, x1, y1, x2, y2 int32, color SDL_Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.lineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

func LineRGBA(renderer *SDL_Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.lineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

func AALineColor(renderer *SDL_Renderer, x1, y1, x2, y2 int32, color SDL_Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_color := C.Uint32(gfxColor(color))
	return C.aalineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _color) == 0
}

func AALineRGBA(renderer *SDL_Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aalineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _r, _g, _b, _a) == 0
}

func ThickLineColor(renderer *SDL_Renderer, x1, y1, x2, y2, width int32, color SDL_Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_width := C.Uint8(width)
	_color := C.Uint32(gfxColor(color))
	return C.thickLineColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _width, _color) == 0
}

func ThickLineRGBA(renderer *SDL_Renderer, x1, y1, x2, y2, width int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_width := C.Uint8(width)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.thickLineRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _width, _r, _g, _b, _a) == 0
}

func CircleColor(renderer *SDL_Renderer, x, y, rad int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.circleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

func CircleRGBA(renderer *SDL_Renderer, x, y, rad int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.circleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _r, _g, _b, _a) == 0
}

func ArcColor(renderer *SDL_Renderer, x, y, rad, start, end int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(gfxColor(color))
	return C.arcColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _color) == 0
}

func ArcRGBA(renderer *SDL_Renderer, x, y, rad, start, end int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.arcRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _r, _g, _b, _a) == 0
}

func AACircleColor(renderer *SDL_Renderer, x, y, rad int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.aacircleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

func AACircleRGBA(renderer *SDL_Renderer, x, y, rad int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aacircleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _r, _g, _b, _a) == 0
}

func FilledCircleColor(renderer *SDL_Renderer, x, y, rad int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_color := C.Uint32(gfxColor(color))
	return C.filledCircleColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _color) == 0
}

func FilledCircleRGBA(renderer *SDL_Renderer, x, y, rad int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledCircleRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _r, _g, _b, _a) == 0
}

func EllipseColor(renderer *SDL_Renderer, x, y, rx, ry int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.ellipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

func EllipseRGBA(renderer *SDL_Renderer, x, y, rx, ry int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.ellipseRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _r, _g, _b, _a) == 0
}

func AAEllipseColor(renderer *SDL_Renderer, x, y, rx, ry int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.aaellipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

func AAEllipseRGBA(renderer *SDL_Renderer, x, y, rx, ry int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aaellipseRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _r, _g, _b, _a) == 0
}

func FilledEllipseColor(renderer *SDL_Renderer, x, y, rx, ry int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_color := C.Uint32(gfxColor(color))
	return C.filledEllipseColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _color) == 0
}

func FilledEllipseRGBA(renderer *SDL_Renderer, x, y, rx, ry int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rx := C.Sint16(rx)
	_ry := C.Sint16(ry)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledEllipseRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rx, _ry, _r, _g, _b, _a) == 0
}

func PieColor(renderer *SDL_Renderer, x, y, rad, start, end int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(gfxColor(color))
	return C.pieColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _color) == 0
}

func PieRGBA(renderer *SDL_Renderer, x, y, rad, start, end int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.pieRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _r, _g, _b, _a) == 0
}

func FilledPieColor(renderer *SDL_Renderer, x, y, rad, start, end int32, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_color := C.Uint32(gfxColor(color))
	return C.filledPieColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _color) == 0
}

func FilledPieRGBA(renderer *SDL_Renderer, x, y, rad, start, end int32, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_rad := C.Sint16(rad)
	_start := C.Sint16(start)
	_end := C.Sint16(end)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledPieRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _rad, _start, _end, _r, _g, _b, _a) == 0
}

func TrigonColor(renderer *SDL_Renderer, x1, y1, x2, y2, x3, y3 int32, color SDL_Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_color := C.Uint32(gfxColor(color))
	return C.trigonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _x3, _y3, _color) == 0
}

func TrigonRGBA(renderer *SDL_Renderer, x1, y1, x2, y2, x3, y3 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.trigonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _x3, _y3, _r, _g, _b, _a) == 0
}

func FilledTrigonColor(renderer *SDL_Renderer, x1, y1, x2, y2, x3, y3 int32, color SDL_Color) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_color := C.Uint32(gfxColor(color))
	return C.filledTrigonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _x3, _y3, _color) == 0
}

func FilledTrigonRGBA(renderer *SDL_Renderer, x1, y1, x2, y2, x3, y3 int32, r, g, b, a uint8) bool {
	_x1 := C.Sint16(x1)
	_y1 := C.Sint16(y1)
	_x2 := C.Sint16(x2)
	_y2 := C.Sint16(y2)
	_x3 := C.Sint16(x3)
	_y3 := C.Sint16(y3)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledTrigonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x1, _y1, _x2, _y2, _x3, _y3, _r, _g, _b, _a) == 0
}

func PolygonColor(renderer *SDL_Renderer, vx, vy []int16, color SDL_Color) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_color := C.Uint32(gfxColor(color))
	return C.polygonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _color) == 0
}

func PolygonRGBA(renderer *SDL_Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.polygonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _r, _g, _b, _a) == 0
}

func AAPolygonColor(renderer *SDL_Renderer, vx, vy []int16, color SDL_Color) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_color := C.Uint32(gfxColor(color))
	return C.aapolygonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _color) == 0
}

func AAPolygonRGBA(renderer *SDL_Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.aapolygonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _r, _g, _b, _a) == 0
}

func FilledPolygonColor(renderer *SDL_Renderer, vx, vy []int16, color SDL_Color) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_color := C.Uint32(gfxColor(color))
	return C.filledPolygonColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _color) == 0
}

func FilledPolygonRGBA(renderer *SDL_Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.filledPolygonRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _r, _g, _b, _a) == 0
}

func TexturedPolygon(renderer *SDL_Renderer, vx, vy []int16, surface *SDL_Surface, textureDX, textureDY int) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	_textureDX := C.int(textureDX)
	_textureDY := C.int(textureDY)
	return C.texturedPolygon((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _surface, _textureDX, _textureDY) == 0
}

func BezierColor(renderer *SDL_Renderer, vx, vy []int16, s int, color SDL_Color) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_s := C.int(s)
	_color := C.Uint32(gfxColor(color))
	return C.bezierColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _s, _color) == 0
}

func BezierRGBA(renderer *SDL_Renderer, vx, vy []int16, s int, r, g, b, a uint8) bool {
	_len := C.int(min(len(vx), len(vy)))
	if _len == 0 {
		return true
	}
	_vx := (*C.Sint16)(unsafe.Pointer(&vx[0]))
	_vy := (*C.Sint16)(unsafe.Pointer(&vy[0]))
	_s := C.int(s)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.bezierRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _vx, _vy, _len, _s, _r, _g, _b, _a) == 0
}

func SetFont(fontdata []byte, cw, ch uint32) {
	_fontdata := unsafe.Pointer(nil)
	if fontdata != nil {
		_fontdata = unsafe.Pointer(&fontdata[0])
	}
	_cw := C.Uint32(cw)
	_ch := C.Uint32(ch)
	C.gfxPrimitivesSetFont(_fontdata, _cw, _ch)
}

func SetFontRotation(rotation uint32) {
	_rotation := C.Uint32(rotation)
	C.gfxPrimitivesSetFontRotation(_rotation)
}

func CharacterColor(renderer *SDL_Renderer, x, y int32, c byte, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_c := C.char(c)
	_color := C.Uint32(gfxColor(color))
	return C.characterColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _c, _color) == 0
}

func CharacterRGBA(renderer *SDL_Renderer, x, y int32, c, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_c := C.char(c)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.characterRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _c, _r, _g, _b, _a) == 0
}

func StringColor(renderer *SDL_Renderer, x, y int32, s string, color SDL_Color) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_s := C.CString(s)
	_color := C.Uint32(gfxColor(color))
	return C.stringColor((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _s, _color) == 0
}

func StringRGBA(renderer *SDL_Renderer, x, y int32, s string, r, g, b, a uint8) bool {
	_x := C.Sint16(x)
	_y := C.Sint16(y)
	_s := C.CString(s)
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return C.stringRGBA((*C.SDL_Renderer)(unsafe.Pointer(renderer)), _x, _y, _s, _r, _g, _b, _a) == 0
}

func SDL_imageFilterMMXdetect() bool {
	return C.SDL_imageFilterMMXdetect() > 0
}

func SDL_imageFilterMMXoff() {
	C.SDL_imageFilterMMXoff()
}

func SDL_imageFilterMMXon() {
	C.SDL_imageFilterMMXon()
}

func SDL_imageFilterAdd(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterAdd(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterMean(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMean(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterSub(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterSub(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterAbsDiff(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterAbsDiff(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterMult(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMult(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterMultNor(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMultNor(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterMultDivby2(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMultDivby2(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterMultDivby4(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMultDivby4(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterBitAnd(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterBitAnd(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterBitOr(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterBitOr(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterDiv(src1, src2, dest []byte) bool {
	_len := C.uint(min(len(src1), len(src2), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_src2 := (*C.uchar)(unsafe.Pointer(&src2[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterDiv(_src1, _src2, _dest, _len) == 0
}

func SDL_imageFilterBitNegation(src1, dest []byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterBitNegation(_src1, _dest, _len) == 0
}

func SDL_imageFilterAddByte(src1, dest []byte, c byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_c := C.uchar(c)
	return C.SDL_imageFilterAddByte(_src1, _dest, _len, _c) == 0
}

func SDL_imageFilterAddUint(src1, dest []byte, c uint) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_c := C.uint(c)
	return C.SDL_imageFilterAddUint(_src1, _dest, _len, _c) == 0
}

func SDL_imageFilterSubByte(src1, dest []byte, c byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_c := C.uchar(c)
	return C.SDL_imageFilterSubByte(_src1, _dest, _len, _c) == 0
}

func SDL_imageFilterSubUint(src1, dest []byte, c uint) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_c := C.uint(c)
	return C.SDL_imageFilterSubUint(_src1, _dest, _len, _c) == 0
}

func SDL_imageFilterShiftRight(src1, dest []byte, n byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftRight(_src1, _dest, _len, _n) == 0
}

func SDL_imageFilterShiftRightUint(src1, dest []byte, n byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftRightUint(_src1, _dest, _len, _n) == 0
}

func SDL_imageFilterMultByByte(src1, dest []byte, c byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_c := C.uchar(c)
	return C.SDL_imageFilterMultByByte(_src1, _dest, _len, _c) == 0
}

func SDL_imageFilterShiftRightAndMultByByte(src1, dest []byte, n, c byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	_c := C.uchar(c)
	return C.SDL_imageFilterShiftRightAndMultByByte(_src1, _dest, _len, _n, _c) == 0
}

func SDL_imageFilterShiftLeftByte(src1, dest []byte, n byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftLeftByte(_src1, _dest, _len, _n) == 0
}

func SDL_imageFilterShiftLeftUint(src1, dest []byte, n byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftLeftUint(_src1, _dest, _len, _n) == 0
}

func SDL_imageFilterShiftLeft(src1, dest []byte, n byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_n := C.uchar(n)
	return C.SDL_imageFilterShiftLeft(_src1, _dest, _len, _n) == 0
}

func SDL_imageFilterBinarizeUsingThreshold(src1, dest []byte, t byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_t := C.uchar(t)
	return C.SDL_imageFilterBinarizeUsingThreshold(_src1, _dest, _len, _t) == 0
}

func SDL_imageFilterClipToRange(src1, dest []byte, tmin, tmax byte) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_tmin := C.uchar(tmin)
	_tmax := C.uchar(tmax)
	return C.SDL_imageFilterClipToRange(_src1, _dest, _len, _tmin, _tmax) == 0
}

func SDL_imageFilterNormalizeLinear(src1, dest []byte, cmin, cmax, nmin, nmax int) bool {
	_len := C.uint(min(len(src1), len(dest)))
	if _len == 0 {
		return true
	}
	_src1 := (*C.uchar)(unsafe.Pointer(&src1[0]))
	_dest := (*C.uchar)(unsafe.Pointer(&dest[0]))
	_cmin := C.int(cmin)
	_cmax := C.int(cmax)
	_nmin := C.int(nmin)
	_nmax := C.int(nmax)
	return C.SDL_imageFilterNormalizeLinear(_src1, _dest, _len, _cmin, _cmax, _nmin, _nmax) == 0
}

func RotoZoomSurface(src *SDL_Surface, angle, zoom float64, smooth int) *SDL_Surface {
	_angle := C.double(angle)
	_zoom := C.double(zoom)
	_smooth := C.int(smooth)
	return (*SDL_Surface)(unsafe.Pointer(C.rotozoomSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _angle, _zoom, _smooth)))
}

func RotoZoomSurfaceXY(src *SDL_Surface, angle, zoomx, zoomy float64, smooth int) *SDL_Surface {
	_angle := C.double(angle)
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_smooth := C.int(smooth)
	return (*SDL_Surface)(unsafe.Pointer(C.rotozoomSurfaceXY((*C.SDL_Surface)(unsafe.Pointer(src)), _angle, _zoomx, _zoomy, _smooth)))
}

func ZoomSurface(src *SDL_Surface, zoomx, zoomy float64, smooth int) *SDL_Surface {
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_smooth := C.int(smooth)
	return (*SDL_Surface)(unsafe.Pointer(C.zoomSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _zoomx, _zoomy, _smooth)))
}

func ZoomSurfaceSize(width, height int32, zoomx, zoomy float64) (dstwidth, dstheight int) {
	_width := C.int(width)
	_height := C.int(height)
	_zoomx := C.double(zoomx)
	_zoomy := C.double(zoomy)
	_dstwidth := (*C.int)(unsafe.Pointer(&dstwidth))
	_dstheight := (*C.int)(unsafe.Pointer(&dstheight))
	C.zoomSurfaceSize(_width, _height, _zoomx, _zoomy, _dstwidth, _dstheight)
	return dstwidth, dstheight
}

func ShrinkSurface(src *SDL_Surface, factorx, factory int) *SDL_Surface {
	_factorx := C.int(factorx)
	_factory := C.int(factory)
	return (*SDL_Surface)(unsafe.Pointer(C.shrinkSurface((*C.SDL_Surface)(unsafe.Pointer(src)), _factorx, _factory)))
}

func RotateSurface90Degrees(src *SDL_Surface, numClockwiseTurns int) *SDL_Surface {
	_numClockwiseTurns := C.int(numClockwiseTurns)
	return (*SDL_Surface)(unsafe.Pointer(C.rotateSurface90Degrees((*C.SDL_Surface)(unsafe.Pointer(src)), _numClockwiseTurns)))
}
