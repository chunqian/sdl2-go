package gfx

/*
#include "gfx_wrapper.h"
*/
import "C"
import (
	"unsafe"

	. "github.com/chunqian/sdl2-go/sdl"
)

func cSDL_Renderer(r *SDL_Renderer) *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
}

func cFPSmanager(manager *FPSmanager) *C.FPSmanager {
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
	C.SDL_initFramerate(cFPSmanager(manager))
}

func SDL_setFramerate(manager *FPSmanager, rate uint) bool {
	cRate := cUint32(rate)
	return int(C.SDL_setFramerate(cFPSmanager(manager), cRate)) == 0
}

func SDL_getFramerate(manager *FPSmanager) int {
	cRet := C.SDL_getFramerate(cFPSmanager(manager))
	return int(cRet)
}

func SDL_getFramecount(manager *FPSmanager) int {
	cRet := C.SDL_getFramecount(cFPSmanager(manager))
	return int(cRet)
}

func SDL_framerateDelay(manager *FPSmanager) uint {
	cRet := C.SDL_framerateDelay(cFPSmanager(manager))
	return uint(cRet)
}

func PixelColor(renderer *SDL_Renderer, x, y int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cColor := cUint32(gfxColor(color))
	return C.pixelColor(cSDL_Renderer(renderer), cX, cY, cColor) == 0
}

func PixelRGBA(renderer *SDL_Renderer, x, y int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.pixelRGBA(cSDL_Renderer(renderer), cX, cY, cR, cG, cB, cA) == 0
}

func HlineColor(renderer *SDL_Renderer, x1, x2, y int32, color SDL_Color) bool {
	cX1 := cInt16(x1)
	cX2 := cInt16(x2)
	cY := cInt16(y)
	cColor := cUint32(gfxColor(color))
	return C.hlineColor(cSDL_Renderer(renderer), cX1, cX2, cY, cColor) == 0
}

func HlineRGBA(renderer *SDL_Renderer, x1, x2, y int32, r, g, b, a uint8) bool {
	cX1 := cInt16(x1)
	cX2 := cInt16(x2)
	cY := cInt16(y)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.hlineRGBA(cSDL_Renderer(renderer), cX1, cX2, cY, cR, cG, cB, cA) == 0
}

func VlineColor(renderer *SDL_Renderer, x, y1, y2 int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY1 := cInt16(y1)
	cY2 := cInt16(y2)
	cColor := cUint32(gfxColor(color))
	return C.vlineColor(cSDL_Renderer(renderer), cX, cY1, cY2, cColor) == 0
}

func VlineRGBA(renderer *SDL_Renderer, x, y1, y2 int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY1 := cInt16(y1)
	cY2 := cInt16(y2)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.vlineRGBA(cSDL_Renderer(renderer), cX, cY1, cY2, cR, cG, cB, cA) == 0
}

func RectangleColor(renderer *SDL_Renderer, x1, y1, x2, y2 int32, color SDL_Color) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cColor := cUint32(gfxColor(color))
	return C.rectangleColor(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cColor) == 0
}

func RectangleRGBA(renderer *SDL_Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.rectangleRGBA(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cR, cG, cB, cA) == 0
}

func RoundedRectangleColor(renderer *SDL_Renderer, x1, y1, x2, y2, rad int32, color SDL_Color) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cRad := cInt16(rad)
	cColor := cUint32(gfxColor(color))
	return C.roundedRectangleColor(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cRad, cColor) == 0
}

func RoundedRectangleRGBA(renderer *SDL_Renderer, x1, y1, x2, y2, rad int32, r, g, b, a uint8) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cRad := cInt16(rad)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.roundedRectangleRGBA(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cRad, cR, cG, cB, cA) == 0
}

func BoxColor(renderer *SDL_Renderer, x1, y1, x2, y2 int32, color SDL_Color) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cColor := cUint32(gfxColor(color))
	return C.boxColor(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cColor) == 0
}

func BoxRGBA(renderer *SDL_Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.boxRGBA(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cR, cG, cB, cA) == 0
}

func RoundedBoxColor(renderer *SDL_Renderer, x1, y1, x2, y2, rad int32, color SDL_Color) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cRad := cInt16(rad)
	cColor := cUint32(gfxColor(color))
	return C.roundedBoxColor(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cRad, cColor) == 0
}

func RoundedBoxRGBA(renderer *SDL_Renderer, x1, y1, x2, y2, rad int32, r, g, b, a uint8) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cRad := cInt16(rad)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.roundedBoxRGBA(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cRad, cR, cG, cB, cA) == 0
}

func LineColor(renderer *SDL_Renderer, x1, y1, x2, y2 int32, color SDL_Color) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cColor := cUint32(gfxColor(color))
	return C.lineColor(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cColor) == 0
}

func LineRGBA(renderer *SDL_Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.lineRGBA(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cR, cG, cB, cA) == 0
}

func AALineColor(renderer *SDL_Renderer, x1, y1, x2, y2 int32, color SDL_Color) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cColor := cUint32(gfxColor(color))
	return C.aalineColor(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cColor) == 0
}

func AALineRGBA(renderer *SDL_Renderer, x1, y1, x2, y2 int32, r, g, b, a uint8) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.aalineRGBA(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cR, cG, cB, cA) == 0
}

func ThickLineColor(renderer *SDL_Renderer, x1, y1, x2, y2, width int32, color SDL_Color) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cWidth := cUint8(width)
	cColor := cUint32(gfxColor(color))
	return C.thickLineColor(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cWidth, cColor) == 0
}

func ThickLineRGBA(renderer *SDL_Renderer, x1, y1, x2, y2, width int32, r, g, b, a uint8) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cWidth := cUint8(width)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.thickLineRGBA(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cWidth, cR, cG, cB, cA) == 0
}

func CircleColor(renderer *SDL_Renderer, x, y, rad int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cColor := cUint32(gfxColor(color))
	return C.circleColor(cSDL_Renderer(renderer), cX, cY, cRad, cColor) == 0
}

func CircleRGBA(renderer *SDL_Renderer, x, y, rad int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.circleRGBA(cSDL_Renderer(renderer), cX, cY, cRad, cR, cG, cB, cA) == 0
}

func ArcColor(renderer *SDL_Renderer, x, y, rad, start, end int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cStart := cInt16(start)
	cEnd := cInt16(end)
	cColor := cUint32(gfxColor(color))
	return C.arcColor(cSDL_Renderer(renderer), cX, cY, cRad, cStart, cEnd, cColor) == 0
}

func ArcRGBA(renderer *SDL_Renderer, x, y, rad, start, end int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cStart := cInt16(start)
	cEnd := cInt16(end)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.arcRGBA(cSDL_Renderer(renderer), cX, cY, cRad, cStart, cEnd, cR, cG, cB, cA) == 0
}

func AACircleColor(renderer *SDL_Renderer, x, y, rad int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cColor := cUint32(gfxColor(color))
	return C.aacircleColor(cSDL_Renderer(renderer), cX, cY, cRad, cColor) == 0
}

func AACircleRGBA(renderer *SDL_Renderer, x, y, rad int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.aacircleRGBA(cSDL_Renderer(renderer), cX, cY, cRad, cR, cG, cB, cA) == 0
}

func FilledCircleColor(renderer *SDL_Renderer, x, y, rad int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cColor := cUint32(gfxColor(color))
	return C.filledCircleColor(cSDL_Renderer(renderer), cX, cY, cRad, cColor) == 0
}

func FilledCircleRGBA(renderer *SDL_Renderer, x, y, rad int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.filledCircleRGBA(cSDL_Renderer(renderer), cX, cY, cRad, cR, cG, cB, cA) == 0
}

func EllipseColor(renderer *SDL_Renderer, x, y, rx, ry int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRX := cInt16(rx)
	cRY := cInt16(ry)
	cColor := cUint32(gfxColor(color))
	return C.ellipseColor(cSDL_Renderer(renderer), cX, cY, cRX, cRY, cColor) == 0
}

func EllipseRGBA(renderer *SDL_Renderer, x, y, rx, ry int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRX := cInt16(rx)
	cRY := cInt16(ry)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.ellipseRGBA(cSDL_Renderer(renderer), cX, cY, cRX, cRY, cR, cG, cB, cA) == 0
}

func AAEllipseColor(renderer *SDL_Renderer, x, y, rx, ry int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRX := cInt16(rx)
	cRY := cInt16(ry)
	cColor := cUint32(gfxColor(color))
	return C.aaellipseColor(cSDL_Renderer(renderer), cX, cY, cRX, cRY, cColor) == 0
}

func AAEllipseRGBA(renderer *SDL_Renderer, x, y, rx, ry int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRX := cInt16(rx)
	cRY := cInt16(ry)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.aaellipseRGBA(cSDL_Renderer(renderer), cX, cY, cRX, cRY, cR, cG, cB, cA) == 0
}

func FilledEllipseColor(renderer *SDL_Renderer, x, y, rx, ry int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRX := cInt16(rx)
	cRY := cInt16(ry)
	cColor := cUint32(gfxColor(color))
	return C.filledEllipseColor(cSDL_Renderer(renderer), cX, cY, cRX, cRY, cColor) == 0
}

func FilledEllipseRGBA(renderer *SDL_Renderer, x, y, rx, ry int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRX := cInt16(rx)
	cRY := cInt16(ry)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.filledEllipseRGBA(cSDL_Renderer(renderer), cX, cY, cRX, cRY, cR, cG, cB, cA) == 0
}

func PieColor(renderer *SDL_Renderer, x, y, rad, start, end int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cStart := cInt16(start)
	cEnd := cInt16(end)
	cColor := cUint32(gfxColor(color))
	return C.pieColor(cSDL_Renderer(renderer), cX, cY, cRad, cStart, cEnd, cColor) == 0
}

func PieRGBA(renderer *SDL_Renderer, x, y, rad, start, end int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cStart := cInt16(start)
	cEnd := cInt16(end)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.pieRGBA(cSDL_Renderer(renderer), cX, cY, cRad, cStart, cEnd, cR, cG, cB, cA) == 0
}

func FilledPieColor(renderer *SDL_Renderer, x, y, rad, start, end int32, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cStart := cInt16(start)
	cEnd := cInt16(end)
	cColor := cUint32(gfxColor(color))
	return C.filledPieColor(cSDL_Renderer(renderer), cX, cY, cRad, cStart, cEnd, cColor) == 0
}

func FilledPieRGBA(renderer *SDL_Renderer, x, y, rad, start, end int32, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cRad := cInt16(rad)
	cStart := cInt16(start)
	cEnd := cInt16(end)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.filledPieRGBA(cSDL_Renderer(renderer), cX, cY, cRad, cStart, cEnd, cR, cG, cB, cA) == 0
}

func TrigonColor(renderer *SDL_Renderer, x1, y1, x2, y2, x3, y3 int32, color SDL_Color) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cX3 := cInt16(x3)
	cY3 := cInt16(y3)
	cColor := cUint32(gfxColor(color))
	return C.trigonColor(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cX3, cY3, cColor) == 0
}

func TrigonRGBA(renderer *SDL_Renderer, x1, y1, x2, y2, x3, y3 int32, r, g, b, a uint8) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cX3 := cInt16(x3)
	cY3 := cInt16(y3)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.trigonRGBA(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cX3, cY3, cR, cG, cB, cA) == 0
}

func FilledTrigonColor(renderer *SDL_Renderer, x1, y1, x2, y2, x3, y3 int32, color SDL_Color) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cX3 := cInt16(x3)
	cY3 := cInt16(y3)
	cColor := cUint32(gfxColor(color))
	return C.filledTrigonColor(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cX3, cY3, cColor) == 0
}

func FilledTrigonRGBA(renderer *SDL_Renderer, x1, y1, x2, y2, x3, y3 int32, r, g, b, a uint8) bool {
	cX1 := cInt16(x1)
	cY1 := cInt16(y1)
	cX2 := cInt16(x2)
	cY2 := cInt16(y2)
	cX3 := cInt16(x3)
	cY3 := cInt16(y3)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.filledTrigonRGBA(cSDL_Renderer(renderer), cX1, cY1, cX2, cY2, cX3, cY3, cR, cG, cB, cA) == 0
}

func PolygonColor(renderer *SDL_Renderer, vx, vy []int16, color SDL_Color) bool {
	cLen := cInt(min(len(vx), len(vy)))
	if cLen == 0 {
		return true
	}
	cVX := (*cInt16)(unsafe.Pointer(&vx[0]))
	cVY := (*cInt16)(unsafe.Pointer(&vy[0]))
	cColor := cUint32(gfxColor(color))
	return C.polygonColor(cSDL_Renderer(renderer), cVX, cVY, cLen, cColor) == 0
}

func PolygonRGBA(renderer *SDL_Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	cLen := cInt(min(len(vx), len(vy)))
	if cLen == 0 {
		return true
	}
	cVX := (*cInt16)(unsafe.Pointer(&vx[0]))
	cVY := (*cInt16)(unsafe.Pointer(&vy[0]))
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.polygonRGBA(cSDL_Renderer(renderer), cVX, cVY, cLen, cR, cG, cB, cA) == 0
}

func AAPolygonColor(renderer *SDL_Renderer, vx, vy []int16, color SDL_Color) bool {
	cLen := cInt(min(len(vx), len(vy)))
	if cLen == 0 {
		return true
	}
	cVX := (*cInt16)(unsafe.Pointer(&vx[0]))
	cVY := (*cInt16)(unsafe.Pointer(&vy[0]))
	cColor := cUint32(gfxColor(color))
	return C.aapolygonColor(cSDL_Renderer(renderer), cVX, cVY, cLen, cColor) == 0
}

func AAPolygonRGBA(renderer *SDL_Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	cLen := cInt(min(len(vx), len(vy)))
	if cLen == 0 {
		return true
	}
	cVX := (*cInt16)(unsafe.Pointer(&vx[0]))
	cVY := (*cInt16)(unsafe.Pointer(&vy[0]))
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.aapolygonRGBA(cSDL_Renderer(renderer), cVX, cVY, cLen, cR, cG, cB, cA) == 0
}

func FilledPolygonColor(renderer *SDL_Renderer, vx, vy []int16, color SDL_Color) bool {
	cLen := cInt(min(len(vx), len(vy)))
	if cLen == 0 {
		return true
	}
	cVX := (*cInt16)(unsafe.Pointer(&vx[0]))
	cVY := (*cInt16)(unsafe.Pointer(&vy[0]))
	cColor := cUint32(gfxColor(color))
	return C.filledPolygonColor(cSDL_Renderer(renderer), cVX, cVY, cLen, cColor) == 0
}

func FilledPolygonRGBA(renderer *SDL_Renderer, vx, vy []int16, r, g, b, a uint8) bool {
	cLen := cInt(min(len(vx), len(vy)))
	if cLen == 0 {
		return true
	}
	cVX := (*cInt16)(unsafe.Pointer(&vx[0]))
	cVY := (*cInt16)(unsafe.Pointer(&vy[0]))
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.filledPolygonRGBA(cSDL_Renderer(renderer), cVX, cVY, cLen, cR, cG, cB, cA) == 0
}

func TexturedPolygon(renderer *SDL_Renderer, vx, vy []int16, surface *SDL_Surface, textureDX, textureDY int) bool {
	cLen := cInt(min(len(vx), len(vy)))
	if cLen == 0 {
		return true
	}
	cVX := (*cInt16)(unsafe.Pointer(&vx[0]))
	cVY := (*cInt16)(unsafe.Pointer(&vy[0]))
	cSurface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	cTextureDX := cInt(textureDX)
	cTextureDY := cInt(textureDY)
	return C.texturedPolygon(cSDL_Renderer(renderer), cVX, cVY, cLen, cSurface, cTextureDX, cTextureDY) == 0
}

func BezierColor(renderer *SDL_Renderer, vx, vy []int16, s int, color SDL_Color) bool {
	cLen := cInt(min(len(vx), len(vy)))
	if cLen == 0 {
		return true
	}
	cVX := (*cInt16)(unsafe.Pointer(&vx[0]))
	cVY := (*cInt16)(unsafe.Pointer(&vy[0]))
	cS := cInt(s)
	cColor := cUint32(gfxColor(color))
	return C.bezierColor(cSDL_Renderer(renderer), cVX, cVY, cLen, cS, cColor) == 0
}

func BezierRGBA(renderer *SDL_Renderer, vx, vy []int16, s int, r, g, b, a uint8) bool {
	cLen := cInt(min(len(vx), len(vy)))
	if cLen == 0 {
		return true
	}
	cVX := (*cInt16)(unsafe.Pointer(&vx[0]))
	cVY := (*cInt16)(unsafe.Pointer(&vy[0]))
	cS := cInt(s)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.bezierRGBA(cSDL_Renderer(renderer), cVX, cVY, cLen, cS, cR, cG, cB, cA) == 0
}

func SetFont(fontdata []byte, cw, ch uint32) {
	cFontdata := unsafe.Pointer(nil)
	if fontdata != nil {
		cFontdata = unsafe.Pointer(&fontdata[0])
	}
	cCW := cUint32(cw)
	cCH := cUint32(ch)
	C.gfxPrimitivesSetFont(cFontdata, cCW, cCH)
}

func SetFontRotation(rotation uint32) {
	cRotation := cUint32(rotation)
	C.gfxPrimitivesSetFontRotation(cRotation)
}

func CharacterColor(renderer *SDL_Renderer, x, y int32, c byte, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cC := cChar(c)
	cColor := cUint32(gfxColor(color))
	return C.characterColor(cSDL_Renderer(renderer), cX, cY, cC, cColor) == 0
}

func CharacterRGBA(renderer *SDL_Renderer, x, y int32, c, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)
	cC := cChar(c)
	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.characterRGBA(cSDL_Renderer(renderer), cX, cY, cC, cR, cG, cB, cA) == 0
}

func StringColor(renderer *SDL_Renderer, x, y int32, s string, color SDL_Color) bool {
	cX := cInt16(x)
	cY := cInt16(y)

	cS := SDL_CreateCString(SDL_GetMemoryPool(), s)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cS) // memory free

	cColor := cUint32(gfxColor(color))
	return C.stringColor(cSDL_Renderer(renderer), cX, cY, (*cChar)(unsafe.Pointer(cS)), cColor) == 0
}

func StringRGBA(renderer *SDL_Renderer, x, y int32, s string, r, g, b, a uint8) bool {
	cX := cInt16(x)
	cY := cInt16(y)

	cS := SDL_CreateCString(SDL_GetMemoryPool(), s)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cS) // memory free

	cR := cUint8(r)
	cG := cUint8(g)
	cB := cUint8(b)
	cA := cUint8(a)
	return C.stringRGBA(cSDL_Renderer(renderer), cX, cY, (*cChar)(unsafe.Pointer(cS)), cR, cG, cB, cA) == 0
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
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterAdd(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterMean(src1, src2, dest []byte) bool {
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMean(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterSub(src1, src2, dest []byte) bool {
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterSub(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterAbsDiff(src1, src2, dest []byte) bool {
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterAbsDiff(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterMult(src1, src2, dest []byte) bool {
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMult(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterMultNor(src1, src2, dest []byte) bool {
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMultNor(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterMultDivby2(src1, src2, dest []byte) bool {
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMultDivby2(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterMultDivby4(src1, src2, dest []byte) bool {
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterMultDivby4(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterBitAnd(src1, src2, dest []byte) bool {
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterBitAnd(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterBitOr(src1, src2, dest []byte) bool {
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterBitOr(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterDiv(src1, src2, dest []byte) bool {
	cLen := cUint(min(len(src1), len(src2), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cSrc2 := (*cUchar)(unsafe.Pointer(&src2[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterDiv(cSrc1, cSrc2, cDest, cLen) == 0
}

func SDL_imageFilterBitNegation(src1, dest []byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	return C.SDL_imageFilterBitNegation(cSrc1, cDest, cLen) == 0
}

func SDL_imageFilterAddByte(src1, dest []byte, c byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cC := cUchar(c)
	return C.SDL_imageFilterAddByte(cSrc1, cDest, cLen, cC) == 0
}

func SDL_imageFilterAddUint(src1, dest []byte, c uint) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cC := cUint(c)
	return C.SDL_imageFilterAddUint(cSrc1, cDest, cLen, cC) == 0
}

func SDL_imageFilterSubByte(src1, dest []byte, c byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cC := cUchar(c)
	return C.SDL_imageFilterSubByte(cSrc1, cDest, cLen, cC) == 0
}

func SDL_imageFilterSubUint(src1, dest []byte, c uint) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cC := cUint(c)
	return C.SDL_imageFilterSubUint(cSrc1, cDest, cLen, cC) == 0
}

func SDL_imageFilterShiftRight(src1, dest []byte, n byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cN := cUchar(n)
	return C.SDL_imageFilterShiftRight(cSrc1, cDest, cLen, cN) == 0
}

func SDL_imageFilterShiftRightUint(src1, dest []byte, n byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cN := cUchar(n)
	return C.SDL_imageFilterShiftRightUint(cSrc1, cDest, cLen, cN) == 0
}

func SDL_imageFilterMultByByte(src1, dest []byte, c byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cC := cUchar(c)
	return C.SDL_imageFilterMultByByte(cSrc1, cDest, cLen, cC) == 0
}

func SDL_imageFilterShiftRightAndMultByByte(src1, dest []byte, n, c byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cN := cUchar(n)
	cC := cUchar(c)
	return C.SDL_imageFilterShiftRightAndMultByByte(cSrc1, cDest, cLen, cN, cC) == 0
}

func SDL_imageFilterShiftLeftByte(src1, dest []byte, n byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cN := cUchar(n)
	return C.SDL_imageFilterShiftLeftByte(cSrc1, cDest, cLen, cN) == 0
}

func SDL_imageFilterShiftLeftUint(src1, dest []byte, n byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cN := cUchar(n)
	return C.SDL_imageFilterShiftLeftUint(cSrc1, cDest, cLen, cN) == 0
}

func SDL_imageFilterShiftLeft(src1, dest []byte, n byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cN := cUchar(n)
	return C.SDL_imageFilterShiftLeft(cSrc1, cDest, cLen, cN) == 0
}

func SDL_imageFilterBinarizeUsingThreshold(src1, dest []byte, t byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cT := cUchar(t)
	return C.SDL_imageFilterBinarizeUsingThreshold(cSrc1, cDest, cLen, cT) == 0
}

func SDL_imageFilterClipToRange(src1, dest []byte, tmin, tmax byte) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cTmin := cUchar(tmin)
	cTmax := cUchar(tmax)
	return C.SDL_imageFilterClipToRange(cSrc1, cDest, cLen, cTmin, cTmax) == 0
}

func SDL_imageFilterNormalizeLinear(src1, dest []byte, cmin, cmax, nmin, nmax int) bool {
	cLen := cUint(min(len(src1), len(dest)))
	if cLen == 0 {
		return true
	}
	cSrc1 := (*cUchar)(unsafe.Pointer(&src1[0]))
	cDest := (*cUchar)(unsafe.Pointer(&dest[0]))
	cCmin := cInt(cmin)
	cCmax := cInt(cmax)
	cNmin := cInt(nmin)
	cNmax := cInt(nmax)
	return C.SDL_imageFilterNormalizeLinear(cSrc1, cDest, cLen, cCmin, cCmax, cNmin, cNmax) == 0
}

func RotoZoomSurface(src *SDL_Surface, angle, zoom float64, smooth int) *SDL_Surface {
	cAngle := cDouble(angle)
	cZoom := cDouble(zoom)
	cSmooth := cInt(smooth)
	cSurface := C.rotozoomSurface((*C.SDL_Surface)(unsafe.Pointer(src)), cAngle, cZoom, cSmooth)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func RotoZoomSurfaceXY(src *SDL_Surface, angle, zoomx, zoomy float64, smooth int) *SDL_Surface {
	cAngle := cDouble(angle)
	cZoomx := cDouble(zoomx)
	cZoomy := cDouble(zoomy)
	cSmooth := cInt(smooth)
	cSurface := C.rotozoomSurfaceXY((*C.SDL_Surface)(unsafe.Pointer(src)), cAngle, cZoomx, cZoomy, cSmooth)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func ZoomSurface(src *SDL_Surface, zoomx, zoomy float64, smooth int) *SDL_Surface {
	cZoomx := cDouble(zoomx)
	cZoomy := cDouble(zoomy)
	cSmooth := cInt(smooth)
	cSurface := C.zoomSurface((*C.SDL_Surface)(unsafe.Pointer(src)), cZoomx, cZoomy, cSmooth)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func ZoomSurfaceSize(width, height int32, zoomx, zoomy float64, dstwidth, dstheight *int) {
	cWidth := cInt(width)
	cHeight := cInt(height)
	cZoomx := cDouble(zoomx)
	cZoomy := cDouble(zoomy)
	cDstwidth := (*cInt)(unsafe.Pointer(dstwidth))
	cDstheight := (*cInt)(unsafe.Pointer(dstheight))
	C.zoomSurfaceSize(cWidth, cHeight, cZoomx, cZoomy, cDstwidth, cDstheight)
}

func ShrinkSurface(src *SDL_Surface, factorx, factory int) *SDL_Surface {
	cFactorx := cInt(factorx)
	cFactory := cInt(factory)
	cSurface := C.shrinkSurface((*C.SDL_Surface)(unsafe.Pointer(src)), cFactorx, cFactory)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}

func RotateSurface90Degrees(src *SDL_Surface, numClockwiseTurns int) *SDL_Surface {
	cNumClockwiseTurns := cInt(numClockwiseTurns)
	cSurface := C.rotateSurface90Degrees((*C.SDL_Surface)(unsafe.Pointer(src)), cNumClockwiseTurns)
	return (*SDL_Surface)(unsafe.Pointer(cSurface))
}
