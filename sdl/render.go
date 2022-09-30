package sdl

/*
#include "render.h"
*/
import "C"
import (
	"unsafe"
)

func cSDL_RendererInfo(info *SDL_RendererInfo) *C.SDL_RendererInfo {
	return (*C.SDL_RendererInfo)(unsafe.Pointer(info))
}

func cSDL_Texture(t *SDL_Texture) *C.SDL_Texture {
	return (*C.SDL_Texture)(unsafe.Pointer(t))
}

func cSDL_Renderer(r *SDL_Renderer) *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
}

func SDL_GetNumRenderDrivers() int {
	i := int(C.SDL_GetNumRenderDrivers())
	return i
}

func SDL_GetRenderDriverInfo(index int, info *SDL_RendererInfo) int {
	ret := int(C.SDL_GetRenderDriverInfo(cInt(index), cSDL_RendererInfo(info)))
	return ret
}

func SDL_CreateWindowAndRenderer(w, h int32, flags SDL_WindowFlags, window *SDL_Window, renderer *SDL_Renderer) int {
	cWindow := cSDL_Window(window)
	cRenderer := cSDL_Renderer(renderer)
	cRet := C.SDL_CreateWindowAndRenderer(
		cInt(w),
		cInt(h),
		cUint32(flags),
		&cWindow,
		&cRenderer)
	return int(cRet)
}

func SDL_CreateRenderer(window *SDL_Window, index int, flags SDL_RendererFlags) *SDL_Renderer {
	cRenderer := C.SDL_CreateRenderer(cSDL_Window(window), cInt(index), cUint32(flags))
	if cRenderer == nil {
		return nil
	}
	return (*SDL_Renderer)(unsafe.Pointer(cRenderer))
}

func SDL_CreateSoftwareRenderer(surface *SDL_Surface) *SDL_Renderer {
	cRenderer := C.SDL_CreateSoftwareRenderer(cSDL_Surface(surface))
	if cRenderer == nil {
		return nil
	}
	return (*SDL_Renderer)(unsafe.Pointer(cRenderer))
}

func SDL_GetRenderer(window *SDL_Window) *SDL_Renderer {
	cRenderer := C.SDL_GetRenderer(cSDL_Window(window))
	if cRenderer == nil {
		return nil
	}
	return (*SDL_Renderer)(unsafe.Pointer(cRenderer))
}

func SDL_GetRendererInfo(renderer *SDL_Renderer, info *SDL_RendererInfo) int {
	cRet := C.SDL_GetRendererInfo(cSDL_Renderer(renderer), cSDL_RendererInfo(info))
	return int(cRet)
}

func SDL_GetRendererOutputSize(renderer *SDL_Renderer, w, h *int32) int {
	cRet := C.SDL_GetRendererOutputSize(
		cSDL_Renderer(renderer),
		(*cInt)(unsafe.Pointer(w)),
		(*cInt)(unsafe.Pointer(h)))
	return int(cRet)
}

func SDL_CreateTexture(renderer *SDL_Renderer, format uint32, access SDL_TextureAccess, w, h int32) *SDL_Texture {
	cTexture := C.SDL_CreateTexture(
		cSDL_Renderer(renderer),
		cUint32(format),
		cInt(access),
		cInt(w),
		cInt(h))
	if cTexture == nil {
		return nil
	}
	return (*SDL_Texture)(unsafe.Pointer(cTexture))
}

func SDL_CreateTextureFromSurface(renderer *SDL_Renderer, surface *SDL_Surface) *SDL_Texture {
	cTexture := C.SDL_CreateTextureFromSurface(cSDL_Renderer(renderer), cSDL_Surface(surface))
	if cTexture == nil {
		return nil
	}
	return (*SDL_Texture)(unsafe.Pointer(cTexture))
}

func SDL_QueryTexture(texture *SDL_Texture, format *uint32, access *int, w *int32, h *int32) int {
	cRet := C.SDL_QueryTexture(
		cSDL_Texture(texture),
		(*cUint32)(unsafe.Pointer(format)),
		(*cInt)(unsafe.Pointer(access)),
		(*cInt)(unsafe.Pointer(w)),
		(*cInt)(unsafe.Pointer(h)))
	return int(cRet)
}

func SDL_SetTextureColorMod(texture *SDL_Texture, r uint8, g uint8, b uint8) int {
	cRet := C.SDL_SetTextureColorMod(
		cSDL_Texture(texture),
		cUint8(r),
		cUint8(g),
		cUint8(b))
	return int(cRet)
}

func SDL_SetTextureAlphaMod(texture *SDL_Texture, alpha uint8) int {
	cRet := C.SDL_SetTextureAlphaMod(cSDL_Texture(texture), cUint8(alpha))
	return int(cRet)
}

func SDL_GetTextureAlphaMod(texture *SDL_Texture, alpha *uint8) int {
	cRet := C.SDL_GetTextureAlphaMod(cSDL_Texture(texture), (*cUint8)(unsafe.Pointer(alpha)))
	return int(cRet)
}

func SDL_SetTextureBlendMode(texture *SDL_Texture, bm SDL_BlendMode) int {
	cRet := C.SDL_SetTextureBlendMode(cSDL_Texture(texture), cSDL_BlendMode(bm))
	return int(cRet)
}

func SDL_GetTextureBlendMode(texture *SDL_Texture, bm *SDL_BlendMode) int {
	cBM := (*cSDL_BlendMode)(unsafe.Pointer(bm))
	cRet := C.SDL_GetTextureBlendMode(cSDL_Texture(texture), cBM)
	return int(cRet)
}

func SDL_SetTextureScaleMode(texture *SDL_Texture, sm SDL_ScaleMode) int {
	cRet := C.SDL_SetTextureScaleMode(cSDL_Texture(texture), cSDL_ScaleMode(sm))
	return int(cRet)
}

func SDL_GetTextureScaleMode(texture *SDL_Texture, sm *SDL_ScaleMode) int {
	cSM := (*cSDL_ScaleMode)(unsafe.Pointer(sm))
	cRet := C.SDL_GetTextureScaleMode(cSDL_Texture(texture), cSM)
	return int(cRet)
}

func SDL_UpdateTexture[T SDL_PixelsArray](texture *SDL_Texture, rect *SDL_Rect, pixels T, pitch int) int {
	if pixels == nil {
		return -1
	}
	emptyInterface := (*emptyInterface)(unsafe.Pointer(&pixels))
	cRet := C.SDL_UpdateTexture(
		cSDL_Texture(texture),
		cSDL_Rect(rect),
		emptyInterface.Word,
		cInt(pitch))

	return int(cRet)
}

func SDL_UpdateYUVTexture(texture *SDL_Texture, rect *SDL_Rect, yPlane []byte, yPitch int, uPlane []byte, uPitch int, vPlane []byte, vPitch int) int {
	var yPlanePtr, uPlanePtr, vPlanePtr *byte
	if yPlane != nil {
		yPlanePtr = &yPlane[0]
	}
	if uPlane != nil {
		uPlanePtr = &uPlane[0]
	}
	if vPlane != nil {
		vPlanePtr = &vPlane[0]
	}

	cRet := C.SDL_UpdateYUVTexture(
		cSDL_Texture(texture),
		cSDL_Rect(rect),
		(*cUint8)(unsafe.Pointer(yPlanePtr)),
		cInt(yPitch),
		(*cUint8)(unsafe.Pointer(uPlanePtr)),
		cInt(uPitch),
		(*cUint8)(unsafe.Pointer(vPlanePtr)),
		cInt(vPitch))
	return int(cRet)
}

func SDL_LockTexture[T SDL_PixelsArray](texture *SDL_Texture, rect *SDL_Rect, pixels *T, pitch *int32) int {
	emptyInterface := (*emptyInterface)(unsafe.Pointer(pixels))
	cPitch := (*cInt)(unsafe.Pointer(pitch))
	cRet := C.SDL_LockTexture(cSDL_Texture(texture), cSDL_Rect(rect), &emptyInterface.Word, cPitch)

	return int(cRet)
}

func SDL_LockTextureToSurface(texture *SDL_Texture, rect *SDL_Rect, surface *SDL_Surface) int {
	cSurface := cSDL_Surface(surface)
	cRet := C.SDL_LockTextureToSurface(cSDL_Texture(texture), cSDL_Rect(rect), &cSurface)
	return int(cRet)
}

func SDL_UnlockTexture(texture *SDL_Texture) {
	C.SDL_UnlockTexture(cSDL_Texture(texture))
}

func SDL_RenderTargetSupported(renderer *SDL_Renderer) bool {
	cRet := C.SDL_RenderTargetSupported(cSDL_Renderer(renderer))
	return cRet != 0
}

func SDL_SetRenderTarget(renderer *SDL_Renderer, texture *SDL_Texture) int {
	cRet := C.SDL_SetRenderTarget(cSDL_Renderer(renderer), cSDL_Texture(texture))
	return int(cRet)
}

func SDL_GetRenderTarget(renderer *SDL_Renderer) *SDL_Texture {
	cSDL_Texture := C.SDL_GetRenderTarget(cSDL_Renderer(renderer))
	return (*SDL_Texture)(unsafe.Pointer(cSDL_Texture))
}

func SDL_RenderSetLogicalSize(renderer *SDL_Renderer, w, h int32) int {
	cRet := C.SDL_RenderSetLogicalSize(cSDL_Renderer(renderer), cInt(w), cInt(h))
	return int(cRet)
}

func SDL_RenderGetLogicalSize(renderer *SDL_Renderer, w, h *int32) {
	C.SDL_RenderGetLogicalSize(
		cSDL_Renderer(renderer),
		(*cInt)(unsafe.Pointer(w)),
		(*cInt)(unsafe.Pointer(h)))
}

func SDL_RenderSetViewport(renderer *SDL_Renderer, rect *SDL_Rect) int {
	cRet := C.SDL_RenderSetViewport(cSDL_Renderer(renderer), cSDL_Rect(rect))
	return int(cRet)
}

func SDL_RenderGetViewport(renderer *SDL_Renderer, rect *SDL_Rect) {
	C.SDL_RenderGetViewport(cSDL_Renderer(renderer), cSDL_Rect(rect))
}

func SDL_RenderIsClipEnabled(renderer *SDL_Renderer) bool {
	cRet := C.SDL_RenderIsClipEnabled(cSDL_Renderer(renderer))
	return cRet != 0
}

func SDL_RenderSetClipRect(renderer *SDL_Renderer, rect *SDL_Rect) int {
	cRet := C.SDL_RenderSetClipRect(cSDL_Renderer(renderer), cSDL_Rect(rect))
	return int(cRet)
}

func SDL_RenderGetClipRect(renderer *SDL_Renderer, rect *SDL_Rect) {
	C.SDL_RenderGetClipRect(cSDL_Renderer(renderer), cSDL_Rect(rect))
}

func SDL_RenderSetScale(renderer *SDL_Renderer, scaleX, scaleY float32) int {
	cRet := C.SDL_RenderSetScale(
		cSDL_Renderer(renderer),
		cFloat(scaleX),
		cFloat(scaleY))
	return int(cRet)
}

func SDL_RenderGetScale(renderer *SDL_Renderer, scaleX, scaleY *float32) {
	C.SDL_RenderGetScale(
		cSDL_Renderer(renderer),
		(*cFloat)(unsafe.Pointer(scaleX)),
		(*cFloat)(unsafe.Pointer(scaleY)))
}

func SDL_RenderSetIntegerScale(renderer *SDL_Renderer, v SDL_bool) int {
	cRet := C.SDL_RenderSetIntegerScale(cSDL_Renderer(renderer), cSDL_bool(v))
	return int(cRet)
}

func SDL_RenderGetIntegerScale(renderer *SDL_Renderer) bool {
	SDL_ClearError()
	cRet := C.SDL_RenderGetIntegerScale(cSDL_Renderer(renderer))
	return cRet != 0
}

func SDL_SetRenderDrawColor(renderer *SDL_Renderer, r, g, b, a uint8) int {
	cRet := C.SDL_SetRenderDrawColor(
		cSDL_Renderer(renderer),
		cUint8(r),
		cUint8(g),
		cUint8(b),
		cUint8(a))
	return int(cRet)
}

func SetDrawColorArray(renderer *SDL_Renderer, bs ...uint8) int {
	cBs := []cUint8{0, 0, 0, 255}
	for i := 0; i < len(cBs) && i < len(bs); i++ {
		cBs[i] = cUint8(bs[i])
	}

	cRet := C.SDL_SetRenderDrawColor(
		cSDL_Renderer(renderer),
		cBs[0],
		cBs[1],
		cBs[2],
		cBs[3])
	return int(cRet)
}

func SDL_GetRenderDrawColor(renderer *SDL_Renderer, r, g, b, a *uint8) int {
	cRet := C.SDL_GetRenderDrawColor(
		cSDL_Renderer(renderer),
		(*cUint8)(unsafe.Pointer(r)),
		(*cUint8)(unsafe.Pointer(g)),
		(*cUint8)(unsafe.Pointer(b)),
		(*cUint8)(unsafe.Pointer(a)))
	return int(cRet)
}

func SDL_SetRenderDrawBlendMode(renderer *SDL_Renderer, bm SDL_BlendMode) int {
	cRet := C.SDL_SetRenderDrawBlendMode(cSDL_Renderer(renderer), cSDL_BlendMode(bm))
	return int(cRet)
}

func SDL_GetRenderDrawBlendMode(renderer *SDL_Renderer, bm *SDL_BlendMode) int {
	cBM := (*cSDL_BlendMode)(unsafe.Pointer(bm))
	cRet := C.SDL_GetRenderDrawBlendMode(cSDL_Renderer(renderer), cBM)
	return int(cRet)
}

func SDL_RenderClear(renderer *SDL_Renderer) int {
	cRet := C.SDL_RenderClear(cSDL_Renderer(renderer))
	return int(cRet)
}

func SDL_RenderDrawPoint(renderer *SDL_Renderer, x, y int32) int {
	cRet := C.SDL_RenderDrawPoint(cSDL_Renderer(renderer), cInt(x), cInt(y))
	return int(cRet)
}

func SDL_RenderDrawPoints(renderer *SDL_Renderer, points []SDL_Point, count int) int {
	cRet := C.SDL_RenderDrawPoints(
		cSDL_Renderer(renderer),
		cSDL_Point(&points[0]),
		cInt(count))
	return int(cRet)
}

func SDL_RenderDrawLine(renderer *SDL_Renderer, x1, y1, x2, y2 int32) int {
	cRet := C.SDL_RenderDrawLine(
		cSDL_Renderer(renderer),
		cInt(x1),
		cInt(y1),
		cInt(x2),
		cInt(y2))
	return int(cRet)
}

func SDL_RenderDrawLines(renderer *SDL_Renderer, points []SDL_Point, count int) int {
	cRet := C.SDL_RenderDrawLines(
		cSDL_Renderer(renderer),
		cSDL_Point(&points[0]),
		cInt(count))
	return int(cRet)
}

func SDL_RenderDrawRect(renderer *SDL_Renderer, rect *SDL_Rect) int {
	cRet := C.SDL_RenderDrawRect(cSDL_Renderer(renderer), cSDL_Rect(rect))
	return int(cRet)
}

func SDL_RenderDrawRects(renderer *SDL_Renderer, rects []SDL_Rect, count int) int {
	cRet := C.SDL_RenderDrawRects(
		cSDL_Renderer(renderer),
		cSDL_Rect(&rects[0]),
		cInt(count))
	return int(cRet)
}

func SDL_RenderFillRect(renderer *SDL_Renderer, rect *SDL_Rect) int {
	cRet := C.SDL_RenderFillRect(cSDL_Renderer(renderer), cSDL_Rect(rect))
	return int(cRet)
}

func SDL_RenderFillRects(renderer *SDL_Renderer, rects []SDL_Rect, count int) int {
	cRet := C.SDL_RenderFillRects(
		cSDL_Renderer(renderer),
		cSDL_Rect(&rects[0]),
		cInt(count))
	return int(cRet)
}

func SDL_RenderCopy(renderer *SDL_Renderer, texture *SDL_Texture, src, dst *SDL_Rect) int {
	var cRet cInt

	if dst == nil {
		cRet = C.SDL_RenderCopy(
			cSDL_Renderer(renderer),
			cSDL_Texture(texture),
			cSDL_Rect(src),
			cSDL_Rect(dst))
		return int(cRet)
	}

	cRet = C.RenderCopy(
		cSDL_Renderer(renderer),
		cSDL_Texture(texture),
		cSDL_Rect(src),
		cInt(dst.X), cInt(dst.Y), cInt(dst.W), cInt(dst.H))
	return int(cRet)
}

func SDL_RenderCopyEx(renderer *SDL_Renderer, texture *SDL_Texture, src, dst *SDL_Rect, angle float64, center *SDL_Point, flip SDL_RendererFlip) int {
	cRet := C.SDL_RenderCopyEx(
		cSDL_Renderer(renderer),
		cSDL_Texture(texture),
		cSDL_Rect(src),
		cSDL_Rect(dst),
		cDouble(angle),
		cSDL_Point(center),
		cSDL_RendererFlip(flip))
	return int(cRet)
}

func SDL_RenderDrawPointF(renderer *SDL_Renderer, x, y float32) int {
	cRet := C.SDL_RenderDrawPointF(cSDL_Renderer(renderer), cFloat(x), cFloat(y))
	return int(cRet)
}

func SDL_RenderDrawPointsF(renderer *SDL_Renderer, points []SDL_FPoint, count int) int {
	cRet := C.SDL_RenderDrawPointsF(
		cSDL_Renderer(renderer),
		cSDL_FPoint(&points[0]),
		cInt(count))
	return int(cRet)
}

func SDL_RenderDrawLineF(renderer *SDL_Renderer, x1, y1, x2, y2 float32) int {
	cRet := C.SDL_RenderDrawLineF(
		cSDL_Renderer(renderer),
		cFloat(x1),
		cFloat(y1),
		cFloat(x2),
		cFloat(y2))
	return int(cRet)
}

func SDL_RenderDrawLinesF(renderer *SDL_Renderer, points []SDL_FPoint, count int) int {
	cRet := C.SDL_RenderDrawLinesF(
		cSDL_Renderer(renderer),
		cSDL_FPoint(&points[0]),
		cInt(count))
	return int(cRet)
}

func SDL_RenderDrawRectF(renderer *SDL_Renderer, rect *SDL_FRect) int {
	cRet := C.SDL_RenderDrawRectF(cSDL_Renderer(renderer), cSDL_FRect(rect))
	return int(cRet)
}

func SDL_RenderDrawRectsF(renderer *SDL_Renderer, rects []SDL_FRect, count int) int {
	cRet := C.SDL_RenderDrawRectsF(
		cSDL_Renderer(renderer),
		cSDL_FRect(&rects[0]),
		cInt(count))
	return int(cRet)
}

func SDL_RenderFillRectF(renderer *SDL_Renderer, rect *SDL_FRect) int {
	cRet := C.SDL_RenderFillRectF(cSDL_Renderer(renderer), cSDL_FRect(rect))
	return int(cRet)
}

func SDL_RenderFillRectsF(renderer *SDL_Renderer, rects []SDL_FRect, count int) int {
	cRet := C.SDL_RenderFillRectsF(
		cSDL_Renderer(renderer),
		cSDL_FRect(&rects[0]),
		cInt(count))
	return int(cRet)
}

func SDL_RenderCopyF(renderer *SDL_Renderer, texture *SDL_Texture, src *SDL_Rect, dst *SDL_FRect) int {
	cRet := C.SDL_RenderCopyF(
		cSDL_Renderer(renderer),
		cSDL_Texture(texture),
		cSDL_Rect(src),
		cSDL_FRect(dst))
	return int(cRet)
}

func SDL_RenderCopyExF(renderer *SDL_Renderer, texture *SDL_Texture, src *SDL_Rect, dst *SDL_FRect, angle float64, center *SDL_FPoint, flip SDL_RendererFlip) int {
	cRet := C.SDL_RenderCopyExF(
		cSDL_Renderer(renderer),
		cSDL_Texture(texture),
		cSDL_Rect(src),
		cSDL_FRect(dst),
		cDouble(angle),
		cSDL_FPoint(center),
		cSDL_RendererFlip(flip))
	return int(cRet)
}

func SDL_RenderReadPixels(renderer *SDL_Renderer, rect *SDL_Rect, format uint32, pixels unsafe.Pointer, pitch int) int {
	cRet := C.SDL_RenderReadPixels(
		cSDL_Renderer(renderer),
		cSDL_Rect(rect),
		cUint32(format),
		pixels,
		cInt(pitch))
	return int(cRet)
}

func SDL_RenderPresent(renderer *SDL_Renderer) {
	C.SDL_RenderPresent(cSDL_Renderer(renderer))
}

func SDL_DestroyTexture(texture *SDL_Texture) {
	SDL_ClearError()
	C.SDL_DestroyTexture(cSDL_Texture(texture))
}

func SDL_DestroyRenderer(renderer *SDL_Renderer) {
	SDL_ClearError()
	C.SDL_DestroyRenderer(cSDL_Renderer(renderer))
}

func SDL_RenderFlush(renderer *SDL_Renderer) int {
	cRet := C.SDL_RenderFlush(cSDL_Renderer(renderer))
	return int(cRet)
}

func SDL_GL_BindTexture(texture *SDL_Texture, texw, texh *float32) int {
	cRet := C.SDL_GL_BindTexture(
		cSDL_Texture(texture),
		(*cFloat)(unsafe.Pointer(texw)),
		(*cFloat)(unsafe.Pointer(texh)))
	return int(cRet)
}

func SDL_GL_UnbindTexture(texture *SDL_Texture) int {
	cRet := C.SDL_GL_UnbindTexture(cSDL_Texture(texture))
	return int(cRet)
}

func SDL_RenderGetMetalLayer(renderer *SDL_Renderer) (layer unsafe.Pointer) {
	layer = C.SDL_RenderGetMetalLayer(cSDL_Renderer(renderer))
	return
}

func SDL_RenderGetMetalCommandEncoder(renderer *SDL_Renderer) (encoder unsafe.Pointer) {
	encoder = C.SDL_RenderGetMetalCommandEncoder(cSDL_Renderer(renderer))
	return
}

func SDL_UpdateNVTexture(texture *SDL_Texture, rect *SDL_Rect, yPlane []byte, yPitch int, uvPlane []byte, uvPitch int) int {
	var yPlanePtr, uvPlanePtr *byte
	if yPlane != nil {
		yPlanePtr = &yPlane[0]
	}
	if uvPlane != nil {
		uvPlanePtr = &uvPlane[0]
	}

	cRet := C.SDL_UpdateNVTexture(
		cSDL_Texture(texture),
		cSDL_Rect(rect),
		(*cUint8)(unsafe.Pointer(yPlanePtr)),
		cInt(yPitch),
		(*cUint8)(unsafe.Pointer(uvPlanePtr)),
		cInt(uvPitch))
	return int(cRet)
}

func SDL_RenderGeometry(renderer *SDL_Renderer, texture *SDL_Texture, vertices []SDL_Vertex, indices []int32) int {
	cTexture := cSDL_Texture(texture)
	cVertices := (*C.SDL_Vertex)(unsafe.Pointer(&vertices[0]))
	cNumVertices := cInt(len(vertices))
	var cIndices *cInt
	cNumIndices := cInt(0)
	if indices != nil {
		cIndices = (*cInt)(unsafe.Pointer(&indices[0]))
		cNumIndices = cInt(len(indices))
	}

	cRet := C.SDL_RenderGeometry(cSDL_Renderer(renderer), cTexture, cVertices, cNumVertices, cIndices, cNumIndices)
	return int(cRet)
}

func SDL_RenderGeometryRaw(renderer *SDL_Renderer, texture *SDL_Texture, xy []float32, xyStride int, color []SDL_Color, colorStride int, uv []float32, uvStride int, numVertices int, indices any) int {
	sizeIndices := 0
	numIndices := 0
	cIndices := unsafe.Pointer(nil)

	switch t := indices.(type) {
	case []byte:
		cIndices = unsafe.Pointer(&t[0])
		sizeIndices = 1
		numIndices = len(t)
	case []int8:
		cIndices = unsafe.Pointer(&t[0])
		sizeIndices = 1
		numIndices = len(t)
	case []int16:
		cIndices = unsafe.Pointer(&t[0])
		sizeIndices = 2
		numIndices = len(t)
	case []uint16:
		cIndices = unsafe.Pointer(&t[0])
		sizeIndices = 2
		numIndices = len(t)
	case []int32:
		cIndices = unsafe.Pointer(&t[0])
		sizeIndices = 4
		numIndices = len(t)
	case []uint32:
		cIndices = unsafe.Pointer(&t[0])
		sizeIndices = 4
		numIndices = len(t)
	}

	cTexture := cSDL_Texture(texture)
	cXY := (*cFloat)(&xy[0])
	cXYStride := cInt(xyStride)
	cColor := cSDL_Color(&color[0])
	cColorStride := cInt(colorStride)
	cUV := (*cFloat)(&uv[0])
	cUVStride := cInt(uvStride)
	cNumVertices := cInt(len(xy))
	cNumIndices := cInt(numIndices)
	cSizeIndices := cInt(sizeIndices)

	cRet := C.RenderGeometryRaw(cSDL_Renderer(renderer), cTexture, cXY, cXYStride, cColor, cColorStride, cUV, cUVStride, cNumVertices, cIndices, cNumIndices, cSizeIndices)
	return int(cRet)
}

func SDL_SetTextureUserData(texture *SDL_Texture, userdata unsafe.Pointer) int {
	cRet := C.SDL_SetTextureUserData(cSDL_Texture(texture), userdata)
	return int(cRet)
}

func SDL_GetTextureUserData(texture *SDL_Texture) (userdata unsafe.Pointer) {
	userdata = C.SDL_GetTextureUserData(cSDL_Texture(texture))
	return
}

func SDL_RenderWindowToLogical(renderer *SDL_Renderer, windowX, windowY int, logicalX, logicalY *float32) {
	cWindowX := cInt(windowX)
	cWindowY := cInt(windowY)
	cLogicalX := (*cFloat)(unsafe.Pointer(logicalX))
	cLogicalY := (*cFloat)(unsafe.Pointer(logicalY))
	C.SDL_RenderWindowToLogical(cSDL_Renderer(renderer), cWindowX, cWindowY, cLogicalX, cLogicalY)
}

func SDL_RenderLogicalToWindow(renderer *SDL_Renderer, logicalX, logicalY float32, windowX, windowY *int) {
	cLogicalX := cFloat(logicalX)
	cLogicalY := cFloat(logicalY)
	cWindowX := (*cInt)(unsafe.Pointer(windowX))
	cWindowY := (*cInt)(unsafe.Pointer(windowY))
	C.SDL_RenderLogicalToWindow(cSDL_Renderer(renderer), cLogicalX, cLogicalY, cWindowX, cWindowY)
	return
}

func SDL_RenderSetVSync(renderer *SDL_Renderer, vsync SDL_bool) int {
	cVsync := cInt(vsync)
	cRet := C.SDL_RenderSetVSync(cSDL_Renderer(renderer), cVsync)
	return int(cRet)
}

func SDL_RenderGetWindow(renderer *SDL_Renderer) (window *SDL_Window) {
	cWindow := C.SDL_RenderGetWindow(cSDL_Renderer(renderer))
	window = (*SDL_Window)(unsafe.Pointer(cWindow))
	return
}
