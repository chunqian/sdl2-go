package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import (
	"math"
	"unsafe"
)

func cPoint(p *SDL_Point) *C.SDL_Point {
	return (*C.SDL_Point)(unsafe.Pointer(p))
}

func cRect(r *SDL_Rect) *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(r))
}

func cFPoint(fp *SDL_FPoint) *C.SDL_FPoint {
	return (*C.SDL_FPoint)(unsafe.Pointer(fp))
}

func cFRect(fr *SDL_FRect) *C.SDL_FRect {
	return (*C.SDL_FRect)(unsafe.Pointer(fr))
}

func SDL_PointInRect(p *SDL_Point, r *SDL_Rect) bool {
	if (p.X >= r.X) && (p.X < (r.X + r.W)) &&
		(p.Y >= r.Y) && (p.Y < (r.Y + r.H)) {
		return true
	}
	return false
}

func SDL_RectEmpty(a *SDL_Rect) bool {
	return a == nil || a.W <= 0 || a.H <= 0
}

func SDL_RectEquals(a *SDL_Rect, b *SDL_Rect) bool {
	if (a != nil) && (b != nil) &&
		(a.X == b.X) && (a.Y == b.Y) &&
		(a.W == b.W) && (a.H == b.H) {
		return true
	}
	return false
}

func SDL_HasIntersection(a *SDL_Rect, b *SDL_Rect) bool {
	if a == nil || b == nil {
		return false
	}

	// Special case for empty rects
	if SDL_RectEmpty(a) || SDL_RectEmpty(b) {
		return false
	}

	if a.X >= b.X+b.W || a.X+a.W <= b.X || a.Y >= b.Y+b.H || a.Y+a.H <= b.Y {
		return false
	}

	return true
}

func SDL_IntersectRect(a *SDL_Rect, b *SDL_Rect, result *SDL_Rect) bool {
	result = &SDL_Rect{}

	if a == nil || b == nil {
		return false
	}

	// Special case for empty rects
	if SDL_RectEmpty(a) || SDL_RectEmpty(b) {
		result.W = 0
		result.H = 0
		return false
	}

	aMin := a.X
	aMax := aMin + a.W
	bMin := b.X
	bMax := bMin + b.W
	if bMin > aMin {
		aMin = bMin
	}
	result.X = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.W = aMax - aMin

	aMin = a.Y
	aMax = aMin + a.H
	bMin = b.Y
	bMax = bMin + b.H
	if bMin > aMin {
		aMin = bMin
	}
	result.Y = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.H = aMax - aMin

	return !SDL_RectEmpty(result)
}

func SDL_UnionRect(a *SDL_Rect, b *SDL_Rect, result *SDL_Rect) {
	result = &SDL_Rect{}

	if a == nil || b == nil {
		return
	}

	// Special case for empty rects
	if SDL_RectEmpty(a) {
		*result = *b
		return
	} else if SDL_RectEmpty(b) {
		*result = *a
		return
	} else if SDL_RectEmpty(a) && SDL_RectEmpty(b) {
		return
	}

	aMin := a.X
	aMax := aMin + a.W
	bMin := b.X
	bMax := bMin + b.W
	if bMin < aMin {
		aMin = bMin
	}
	result.X = aMin
	if bMax > aMax {
		aMax = bMax
	}
	result.W = aMax - aMin

	aMin = a.Y
	aMax = aMin + a.H
	bMin = b.Y
	bMax = bMin + b.H
	if bMin < aMin {
		aMin = bMin
	}
	result.Y = aMin
	if bMax > aMax {
		aMax = bMax
	}
	result.H = aMax - aMin

	return
}

func SDL_EnclosePoints(points []SDL_Point, count int32, clip *SDL_Rect, result *SDL_Rect) bool {
	result = &SDL_Rect{}

	if count == 0 {
		return false
	}

	var minX, minY, maxX, maxY int32
	if clip != nil {
		added := false
		clipMinX := clip.X
		clipMinY := clip.Y
		clipMaxX := clip.X + clip.W - 1
		clipMaxY := clip.Y + clip.H - 1

		// If the clip has no size, we're done
		if SDL_RectEmpty(clip) {
			return false
		}

		for _, val := range points {
			// Check if the point is inside the clip rect
			if val.X < clipMinX || val.X > clipMaxX || val.Y < clipMinY || val.Y > clipMaxY {
				continue
			}

			if !added {
				// If it's the first point
				minX = val.X
				maxX = val.X
				minY = val.Y
				maxY = val.Y
				added = true
			}

			// Find mins and maxes
			if val.X < minX {
				minX = val.X
			} else if val.X > maxX {
				maxX = val.X
			}
			if val.Y < minY {
				minY = val.Y
			} else if val.Y > maxY {
				maxY = val.Y
			}
		}
	} else {
		for i, val := range points {
			if i == 0 {
				// Populate the first point
				minX = val.X
				maxX = val.X
				minY = val.Y
				maxY = val.Y
				continue
			}

			// Find mins and maxes
			if val.X < minX {
				minX = val.X
			} else if val.X > maxX {
				maxX = val.X
			}
			if val.Y < minY {
				minY = val.Y
			} else if val.Y > maxY {
				maxY = val.Y
			}
		}
	}

	result.X = minX
	result.Y = minY
	result.W = (maxX - minX) + 1
	result.H = (maxY - minY) + 1

	return true
}

func computeOutCode(rect *SDL_Rect, x, y int32) int {
	code := 0
	if y < rect.Y {
		code |= codeTop
	} else if y >= rect.Y+rect.H {
		code |= codeBottom
	}
	if x < rect.X {
		code |= codeLeft
	} else if x >= rect.X+rect.W {
		code |= codeRight
	}

	return code
}

func SDL_IntersectRectAndLine(a *SDL_Rect, X1, Y1, X2, Y2 *int32) bool {
	if SDL_RectEmpty(a) {
		return false
	}

	x1 := *X1
	y1 := *Y1
	x2 := *X2
	y2 := *Y2
	rectX1 := a.X
	rectY1 := a.Y
	rectX2 := a.X + a.W - 1
	rectY2 := a.Y + a.H - 1

	// Check if the line is entirely inside the rect
	if x1 >= rectX1 && x1 <= rectX2 && x2 >= rectX1 && x2 <= rectX2 &&
		y1 >= rectY1 && y1 <= rectY2 && y2 >= rectY1 && y2 <= rectY2 {
		return true
	}

	// Check if the line is entirely outside the rect
	if (x1 < rectX1 && x2 < rectX1) || (x1 > rectX2 && x2 > rectX2) ||
		(y1 < rectY1 && y2 < rectY1) || (y1 > rectY2 && y2 > rectY2) {
		return false
	}

	// Check if the line is horizontal
	if y1 == y2 {
		if x1 < rectX1 {
			*X1 = rectX1
		} else if x1 > rectX2 {
			*X1 = rectX2
		}
		if x2 < rectX1 {
			*X2 = rectX1
		} else if x2 > rectX2 {
			*X2 = rectX2
		}

		return true
	}

	// Check if the line is vertical
	if x1 == x2 {
		if y1 < rectY1 {
			*Y1 = rectY1
		} else if y1 > rectY2 {
			*Y1 = rectY2
		}
		if y2 < rectY1 {
			*Y2 = rectY1
		} else if y2 > rectY2 {
			*Y2 = rectY2
		}

		return true
	}

	// Use Cohen-Sutherland algorithm when all shortcuts fail
	outCode1 := computeOutCode(a, x1, y1)
	outCode2 := computeOutCode(a, x2, y2)
	for outCode1 != 0 || outCode2 != 0 {
		if outCode1&outCode2 != 0 {
			return false
		}

		if outCode1 != 0 {
			var x, y int32
			if outCode1&codeTop != 0 {
				y = rectY1
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode1&codeBottom != 0 {
				y = rectY2
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode1&codeLeft != 0 {
				x = rectX1
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			} else if outCode1&codeRight != 0 {
				x = rectX2
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			}

			x1 = x
			y1 = y
			outCode1 = computeOutCode(a, x, y)
		} else {
			var x, y int32
			if outCode2&codeTop != 0 {
				y = rectY1
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode2&codeBottom != 0 {
				y = rectY2
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode2&codeLeft != 0 {
				x = rectX1
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			} else if outCode2&codeRight != 0 {
				x = rectX2
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			}

			x2 = x
			y2 = y
			outCode2 = computeOutCode(a, x, y)
		}
	}

	*X1 = x1
	*Y1 = y1
	*X2 = x2
	*Y2 = y2

	return true
}

func SDL_PointInFRect(p *SDL_FPoint, r *SDL_FRect) bool {
	if (p.X >= r.X) && (p.X < (r.X + r.W)) &&
		(p.Y >= r.Y) && (p.Y < (r.Y + r.H)) {
		return true
	}
	return false
}

func SDL_FRectEmpty(a *SDL_FRect) bool {
	return a == nil || a.W <= 0 || a.H <= 0
}

func SDL_FRectEqualsEpsilon(a *SDL_FRect, b *SDL_FRect, epsilon float32) bool {
	if (a != nil) && (b != nil) && (a == b ||
		(float32(math.Abs(float64(a.X-b.X))) <= epsilon) &&
			(float32(math.Abs(float64(a.Y-b.Y))) <= epsilon) &&
			(float32(math.Abs(float64(a.W-b.W))) <= epsilon) &&
			(float32(math.Abs(float64(a.H-b.H))) <= epsilon)) {
		return true
	}
	return false
}

func SDL_FRectEquals(a *SDL_FRect, b *SDL_FRect) bool {
	if (a != nil) && (b != nil) &&
		(a.X == b.X) && (a.Y == b.Y) &&
		(a.W == b.W) && (a.H == b.H) {
		return true
	}
	return false
}

func SDL_HasIntersectionF(a *SDL_FRect, b *SDL_FRect) bool {
	if a == nil || b == nil {
		return false
	}

	// Special case for empty rects
	if SDL_FRectEmpty(a) || SDL_FRectEmpty(b) {
		return false
	}

	if a.X >= b.X+b.W || a.X+a.W <= b.X || a.Y >= b.Y+b.H || a.Y+a.H <= b.Y {
		return false
	}

	return true
}

func SDL_IntersectFRect(a *SDL_FRect, b *SDL_FRect, result *SDL_FRect) bool {
	result = &SDL_FRect{}

	if a == nil || b == nil {
		return false
	}

	// Special case for empty rects
	if SDL_FRectEmpty(a) || SDL_FRectEmpty(b) {
		result.W = 0
		result.H = 0
		return false
	}

	aMin := a.X
	aMax := aMin + a.W
	bMin := b.X
	bMax := bMin + b.W
	if bMin > aMin {
		aMin = bMin
	}
	result.X = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.W = aMax - aMin

	aMin = a.Y
	aMax = aMin + a.H
	bMin = b.Y
	bMax = bMin + b.H
	if bMin > aMin {
		aMin = bMin
	}
	result.Y = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.H = aMax - aMin

	return !SDL_FRectEmpty(result)
}

func SDL_UnionFRect(a *SDL_FRect, b *SDL_FRect, result *SDL_FRect) {
	result = &SDL_FRect{}

	if a == nil || b == nil {
		return
	}

	// Special case for empty rects
	if SDL_FRectEmpty(a) {
		*result = *b
		return
	} else if SDL_FRectEmpty(b) {
		*result = *a
		return
	} else if SDL_FRectEmpty(a) && SDL_FRectEmpty(b) {
		return
	}

	aMin := a.X
	aMax := aMin + a.W
	bMin := b.X
	bMax := bMin + b.W
	if bMin < aMin {
		aMin = bMin
	}
	result.X = aMin
	if bMax > aMax {
		aMax = bMax
	}
	result.W = aMax - aMin

	aMin = a.Y
	aMax = aMin + a.H
	bMin = b.Y
	bMax = bMin + b.H
	if bMin < aMin {
		aMin = bMin
	}
	result.Y = aMin
	if bMax > aMax {
		aMax = bMax
	}
	result.H = aMax - aMin

	return
}

func SDL_EncloseFPoints(points []SDL_FPoint, count int32, clip *SDL_FRect, result *SDL_FRect) bool {
	result = &SDL_FRect{}

	if count == 0 {
		return false
	}

	var minX, minY, maxX, maxY float32
	if clip != nil {
		added := false
		clipMinX := clip.X
		clipMinY := clip.Y
		clipMaxX := clip.X + clip.W - 1
		clipMaxY := clip.Y + clip.H - 1

		// If the clip has no size, we're done
		if SDL_FRectEmpty(clip) {
			return false
		}

		for _, val := range points {
			// Check if the point is inside the clip rect
			if val.X < clipMinX || val.X > clipMaxX || val.Y < clipMinY || val.Y > clipMaxY {
				continue
			}

			if !added {
				// If it's the first point
				minX = val.X
				maxX = val.X
				minY = val.Y
				maxY = val.Y
				added = true
			}

			// Find mins and maxes
			if val.X < minX {
				minX = val.X
			} else if val.X > maxX {
				maxX = val.X
			}
			if val.Y < minY {
				minY = val.Y
			} else if val.Y > maxY {
				maxY = val.Y
			}
		}
	} else {
		for i, val := range points {
			if i == 0 {
				// Populate the first point
				minX = val.X
				maxX = val.X
				minY = val.Y
				maxY = val.Y
				continue
			}

			// Find mins and maxes
			if val.X < minX {
				minX = val.X
			} else if val.X > maxX {
				maxX = val.X
			}
			if val.Y < minY {
				minY = val.Y
			} else if val.Y > maxY {
				maxY = val.Y
			}
		}
	}

	result.X = minX
	result.Y = minY
	result.W = (maxX - minX) + 1
	result.H = (maxY - minY) + 1

	return true
}

const (
	codeBottom = 1
	codeTop    = 2
	codeLeft   = 4
	codeRight  = 8
)

func computeFOutCode(rect *SDL_FRect, x, y float32) int {
	code := 0
	if y < rect.Y {
		code |= codeTop
	} else if y >= rect.Y+rect.H {
		code |= codeBottom
	}
	if x < rect.X {
		code |= codeLeft
	} else if x >= rect.X+rect.W {
		code |= codeRight
	}

	return code
}

func SDL_IntersectFRectAndLine(a *SDL_FRect, X1, Y1, X2, Y2 *float32) bool {
	if SDL_FRectEmpty(a) {
		return false
	}

	x1 := *X1
	y1 := *Y1
	x2 := *X2
	y2 := *Y2
	rectX1 := a.X
	rectY1 := a.Y
	rectX2 := a.X + a.W - 1
	rectY2 := a.Y + a.H - 1

	// Check if the line is entirely inside the rect
	if x1 >= rectX1 && x1 <= rectX2 && x2 >= rectX1 && x2 <= rectX2 &&
		y1 >= rectY1 && y1 <= rectY2 && y2 >= rectY1 && y2 <= rectY2 {
		return true
	}

	// Check if the line is entirely outside the rect
	if (x1 < rectX1 && x2 < rectX1) || (x1 > rectX2 && x2 > rectX2) ||
		(y1 < rectY1 && y2 < rectY1) || (y1 > rectY2 && y2 > rectY2) {
		return false
	}

	// Check if the line is horizontal
	if y1 == y2 {
		if x1 < rectX1 {
			*X1 = rectX1
		} else if x1 > rectX2 {
			*X1 = rectX2
		}
		if x2 < rectX1 {
			*X2 = rectX1
		} else if x2 > rectX2 {
			*X2 = rectX2
		}

		return true
	}

	// Check if the line is vertical
	if x1 == x2 {
		if y1 < rectY1 {
			*Y1 = rectY1
		} else if y1 > rectY2 {
			*Y1 = rectY2
		}
		if y2 < rectY1 {
			*Y2 = rectY1
		} else if y2 > rectY2 {
			*Y2 = rectY2
		}

		return true
	}

	// Use Cohen-Sutherland algorithm when all shortcuts fail
	outCode1 := computeFOutCode(a, x1, y1)
	outCode2 := computeFOutCode(a, x2, y2)
	for outCode1 != 0 || outCode2 != 0 {
		if outCode1&outCode2 != 0 {
			return false
		}

		if outCode1 != 0 {
			var x, y float32
			if outCode1&codeTop != 0 {
				y = rectY1
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode1&codeBottom != 0 {
				y = rectY2
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode1&codeLeft != 0 {
				x = rectX1
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			} else if outCode1&codeRight != 0 {
				x = rectX2
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			}

			x1 = x
			y1 = y
			outCode1 = computeFOutCode(a, x, y)
		} else {
			var x, y float32
			if outCode2&codeTop != 0 {
				y = rectY1
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode2&codeBottom != 0 {
				y = rectY2
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode2&codeLeft != 0 {
				x = rectX1
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			} else if outCode2&codeRight != 0 {
				x = rectX2
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			}

			x2 = x
			y2 = y
			outCode2 = computeFOutCode(a, x, y)
		}
	}

	*X1 = x1
	*Y1 = y1
	*X2 = x2
	*Y2 = y2

	return true
}
