package sdl

/*
#include "render.h"
*/
import "C"
import (
	"unsafe"
)

func (info *SDL_RendererInfo) cSDL_RendererInfo() *C.SDL_RendererInfo {
	return (*C.SDL_RendererInfo)(unsafe.Pointer(info))
}

func (t *SDL_Texture) cSDL_Texture() *C.SDL_Texture {
	return (*C.SDL_Texture)(unsafe.Pointer(t))
}

func (r *SDL_Renderer) cSDL_Renderer() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
}

func SDL_GetNumRenderDrivers() int {
	i := int(C.SDL_GetNumRenderDrivers())
	return i
}

func SDL_GetRenderDriverInfo(index int, info *SDL_RendererInfo) int {
	ret := int(C.SDL_GetRenderDriverInfo(cInt(index), info.cSDL_RendererInfo()))
	return ret
}

func SDL_CreateWindowAndRenderer(w, h int32, flags SDL_WindowFlags, window *SDL_Window, renderer *SDL_Renderer) int {
	cWindow := window.cSDL_Window()
	cRenderer := renderer.cSDL_Renderer()
	cRet := C.SDL_CreateWindowAndRenderer(
		cInt(w),
		cInt(h),
		cUint32(flags),
		&cWindow,
		&cRenderer)
	return int(cRet)
}

func SDL_CreateRenderer(window *SDL_Window, index int, flags SDL_RendererFlags) *SDL_Renderer {
	cRenderer := C.SDL_CreateRenderer(window.cSDL_Window(), cInt(index), cUint32(flags))
	if cRenderer == nil {
		return nil
	}
	return (*SDL_Renderer)(unsafe.Pointer(cRenderer))
}

func SDL_CreateSoftwareRenderer(surface *SDL_Surface) *SDL_Renderer {
	cRenderer := C.SDL_CreateSoftwareRenderer(surface.cSDL_Surface())
	if cRenderer == nil {
		return nil
	}
	return (*SDL_Renderer)(unsafe.Pointer(cRenderer))
}

func SDL_GetRenderer(window *SDL_Window) *SDL_Renderer {
	cRenderer := C.SDL_GetRenderer(window.cSDL_Window())
	if cRenderer == nil {
		return nil
	}
	return (*SDL_Renderer)(unsafe.Pointer(cRenderer))
}

func SDL_GetRendererInfo(renderer *SDL_Renderer, info *SDL_RendererInfo) int {
	cRet := C.SDL_GetRendererInfo(renderer.cSDL_Renderer(), info.cSDL_RendererInfo())
	return int(cRet)
}

func SDL_GetRendererOutputSize(renderer *SDL_Renderer, w, h *int32) int {
	cRet := C.SDL_GetRendererOutputSize(
		renderer.cSDL_Renderer(),
		(*cInt)(unsafe.Pointer(w)),
		(*cInt)(unsafe.Pointer(h)))
	return int(cRet)
}

func SDL_CreateTexture(renderer *SDL_Renderer, format uint32, access SDL_TextureAccess, w, h int32) *SDL_Texture {
	cTexture := C.SDL_CreateTexture(
		renderer.cSDL_Renderer(),
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
	cTexture := C.SDL_CreateTextureFromSurface(renderer.cSDL_Renderer(), surface.cSDL_Surface())
	if cTexture == nil {
		return nil
	}
	return (*SDL_Texture)(unsafe.Pointer(cTexture))
}

func SDL_QueryTexture(texture *SDL_Texture, format *uint32, access *int, w *int32, h *int32) int {
	cRet := C.SDL_QueryTexture(
		texture.cSDL_Texture(),
		(*cUint32)(unsafe.Pointer(format)),
		(*cInt)(unsafe.Pointer(access)),
		(*cInt)(unsafe.Pointer(w)),
		(*cInt)(unsafe.Pointer(h)))
	return int(cRet)
}

func SDL_SetTextureColorMod(texture *SDL_Texture, r uint8, g uint8, b uint8) int {
	cRet := C.SDL_SetTextureColorMod(
		texture.cSDL_Texture(),
		cUint8(r),
		cUint8(g),
		cUint8(b))
	return int(cRet)
}

func SDL_SetTextureAlphaMod(texture *SDL_Texture, alpha uint8) int {
	cRet := C.SDL_SetTextureAlphaMod(texture.cSDL_Texture(), cUint8(alpha))
	return int(cRet)
}

func SDL_GetTextureAlphaMod(texture *SDL_Texture, alpha *uint8) int {
	cRet := C.SDL_GetTextureAlphaMod(texture.cSDL_Texture(), (*cUint8)(unsafe.Pointer(alpha)))
	return int(cRet)
}

func SDL_SetTextureBlendMode(texture *SDL_Texture, bm SDL_BlendMode) int {
	cRet := C.SDL_SetTextureBlendMode(texture.cSDL_Texture(), cSDL_BlendMode(bm))
	return int(cRet)
}

func SDL_GetTextureBlendMode(texture *SDL_Texture, bm *SDL_BlendMode) int {
	cBM := (*cSDL_BlendMode)(unsafe.Pointer(bm))
	cRet := C.SDL_GetTextureBlendMode(texture.cSDL_Texture(), cBM)
	return int(cRet)
}

func SDL_SetTextureScaleMode(texture *SDL_Texture, sm SDL_ScaleMode) int {
	cRet := C.SDL_SetTextureScaleMode(texture.cSDL_Texture(), cSDL_ScaleMode(sm))
	return int(cRet)
}

func SDL_GetTextureScaleMode(texture *SDL_Texture, sm *SDL_ScaleMode) int {
	cSM := (*cSDL_ScaleMode)(unsafe.Pointer(sm))
	cRet := C.SDL_GetTextureScaleMode(texture.cSDL_Texture(), cSM)
	return int(cRet)
}

func SDL_UpdateTexture[T SDL_PixelsArray](texture *SDL_Texture, rect *SDL_Rect, pixels T, pitch int) int {
	if pixels == nil {
		return -1
	}
	emptyInterface := (*emptyInterface)(unsafe.Pointer(&pixels))
	cRet := C.SDL_UpdateTexture(
		texture.cSDL_Texture(),
		rect.cSDL_Rect(),
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
		texture.cSDL_Texture(),
		rect.cSDL_Rect(),
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
	cRet := C.SDL_LockTexture(texture.cSDL_Texture(), rect.cSDL_Rect(), &emptyInterface.Word, cPitch)

	return int(cRet)
}

func SDL_LockTextureToSurface(texture *SDL_Texture, rect *SDL_Rect, surface *SDL_Surface) int {
	cSurface := surface.cSDL_Surface()
	cRet := C.SDL_LockTextureToSurface(texture.cSDL_Texture(), rect.cSDL_Rect(), &cSurface)
	return int(cRet)
}

func SDL_UnlockTexture(texture *SDL_Texture) {
	C.SDL_UnlockTexture(texture.cSDL_Texture())
}

func SDL_RenderTargetSupported(renderer *SDL_Renderer) bool {
	cRet := C.SDL_RenderTargetSupported(renderer.cSDL_Renderer())
	return cRet != 0
}

func SDL_SetRenderTarget(renderer *SDL_Renderer, texture *SDL_Texture) int {
	cRet := C.SDL_SetRenderTarget(renderer.cSDL_Renderer(), texture.cSDL_Texture())
	return int(cRet)
}

func SDL_GetRenderTarget(renderer *SDL_Renderer) *SDL_Texture {
	cSDL_Texture := C.SDL_GetRenderTarget(renderer.cSDL_Renderer())
	return (*SDL_Texture)(unsafe.Pointer(cSDL_Texture))
}

func SDL_RenderSetLogicalSize(renderer *SDL_Renderer, w, h int32) int {
	cRet := C.SDL_RenderSetLogicalSize(renderer.cSDL_Renderer(), cInt(w), cInt(h))
	return int(cRet)
}

func SDL_RenderGetLogicalSize(renderer *SDL_Renderer, w, h *int32) {
	C.SDL_RenderGetLogicalSize(
		renderer.cSDL_Renderer(),
		(*cInt)(unsafe.Pointer(w)),
		(*cInt)(unsafe.Pointer(h)))
}

func SDL_RenderSetViewport(renderer *SDL_Renderer, rect *SDL_Rect) int {
	cRet := C.SDL_RenderSetViewport(renderer.cSDL_Renderer(), rect.cSDL_Rect())
	return int(cRet)
}

func SDL_RenderGetViewport(renderer *SDL_Renderer, rect *SDL_Rect) {
	C.SDL_RenderGetViewport(renderer.cSDL_Renderer(), rect.cSDL_Rect())
}

func SDL_RenderIsClipEnabled(renderer *SDL_Renderer) bool {
	cRet := C.SDL_RenderIsClipEnabled(renderer.cSDL_Renderer())
	return cRet != 0
}

func SDL_RenderSetClipRect(renderer *SDL_Renderer, rect *SDL_Rect) int {
	cRet := C.SDL_RenderSetClipRect(renderer.cSDL_Renderer(), rect.cSDL_Rect())
	return int(cRet)
}

func SDL_RenderGetClipRect(renderer *SDL_Renderer, rect *SDL_Rect) {
	C.SDL_RenderGetClipRect(renderer.cSDL_Renderer(), rect.cSDL_Rect())
}

func SDL_RenderSetScale(renderer *SDL_Renderer, scaleX, scaleY float32) int {
	cRet := C.SDL_RenderSetScale(
		renderer.cSDL_Renderer(),
		cFloat(scaleX),
		cFloat(scaleY))
	return int(cRet)
}

func SDL_RenderGetScale(renderer *SDL_Renderer, scaleX, scaleY *float32) {
	C.SDL_RenderGetScale(
		renderer.cSDL_Renderer(),
		(*cFloat)(unsafe.Pointer(scaleX)),
		(*cFloat)(unsafe.Pointer(scaleY)))
}

func SDL_RenderSetIntegerScale(renderer *SDL_Renderer, v SDL_bool) int {
	cRet := C.SDL_RenderSetIntegerScale(renderer.cSDL_Renderer(), cSDL_bool(v))
	return int(cRet)
}

func SDL_RenderGetIntegerScale(renderer *SDL_Renderer) bool {
	SDL_ClearError()
	cRet := C.SDL_RenderGetIntegerScale(renderer.cSDL_Renderer())
	return cRet != 0
}

func SDL_SetRenderDrawColor(renderer *SDL_Renderer, r, g, b, a uint8) int {
	cRet := C.SDL_SetRenderDrawColor(
		renderer.cSDL_Renderer(),
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
		renderer.cSDL_Renderer(),
		cBs[0],
		cBs[1],
		cBs[2],
		cBs[3])
	return int(cRet)
}

func SDL_GetRenderDrawColor(renderer *SDL_Renderer, r, g, b, a *uint8) int {
	cRet := C.SDL_GetRenderDrawColor(
		renderer.cSDL_Renderer(),
		(*cUint8)(unsafe.Pointer(r)),
		(*cUint8)(unsafe.Pointer(g)),
		(*cUint8)(unsafe.Pointer(b)),
		(*cUint8)(unsafe.Pointer(a)))
	return int(cRet)
}

func SDL_SetRenderDrawBlendMode(renderer *SDL_Renderer, bm SDL_BlendMode) int {
	cRet := C.SDL_SetRenderDrawBlendMode(renderer.cSDL_Renderer(), cSDL_BlendMode(bm))
	return int(cRet)
}

func SDL_GetRenderDrawBlendMode(renderer *SDL_Renderer, bm *SDL_BlendMode) int {
	cBM := (*cSDL_BlendMode)(unsafe.Pointer(bm))
	cRet := C.SDL_GetRenderDrawBlendMode(renderer.cSDL_Renderer(), cBM)
	return int(cRet)
}

func SDL_RenderClear(renderer *SDL_Renderer) int {
	cRet := C.SDL_RenderClear(renderer.cSDL_Renderer())
	return int(cRet)
}

func SDL_RenderDrawPoint(renderer *SDL_Renderer, x, y int32) int {
	cRet := C.SDL_RenderDrawPoint(renderer.cSDL_Renderer(), cInt(x), cInt(y))
	return int(cRet)
}

func SDL_RenderDrawPoints(renderer *SDL_Renderer, points []SDL_Point, count int) int {
	cRet := C.SDL_RenderDrawPoints(
		renderer.cSDL_Renderer(),
		points[0].cSDL_Point(),
		cInt(count))
	return int(cRet)
}

func SDL_RenderDrawLine(renderer *SDL_Renderer, x1, y1, x2, y2 int32) int {
	cRet := C.SDL_RenderDrawLine(
		renderer.cSDL_Renderer(),
		cInt(x1),
		cInt(y1),
		cInt(x2),
		cInt(y2))
	return int(cRet)
}

func SDL_RenderDrawLines(renderer *SDL_Renderer, points []SDL_Point, count int) int {
	cRet := C.SDL_RenderDrawLines(
		renderer.cSDL_Renderer(),
		points[0].cSDL_Point(),
		cInt(count))
	return int(cRet)
}

func SDL_RenderDrawRect(renderer *SDL_Renderer, rect *SDL_Rect) int {
	cRet := C.SDL_RenderDrawRect(renderer.cSDL_Renderer(), rect.cSDL_Rect())
	return int(cRet)
}

func SDL_RenderDrawRects(renderer *SDL_Renderer, rects []SDL_Rect, count int) int {
	cRet := C.SDL_RenderDrawRects(
		renderer.cSDL_Renderer(),
		rects[0].cSDL_Rect(),
		cInt(count))
	return int(cRet)
}

func SDL_RenderFillRect(renderer *SDL_Renderer, rect *SDL_Rect) int {
	cRet := C.SDL_RenderFillRect(renderer.cSDL_Renderer(), rect.cSDL_Rect())
	return int(cRet)
}

func SDL_RenderFillRects(renderer *SDL_Renderer, rects []SDL_Rect, count int) int {
	cRet := C.SDL_RenderFillRects(
		renderer.cSDL_Renderer(),
		rects[0].cSDL_Rect(),
		cInt(count))
	return int(cRet)
}

func SDL_RenderCopy(renderer *SDL_Renderer, texture *SDL_Texture, src, dst *SDL_Rect) int {
	var cRet cInt

	if dst == nil {
		cRet = C.SDL_RenderCopy(
			renderer.cSDL_Renderer(),
			texture.cSDL_Texture(),
			src.cSDL_Rect(),
			dst.cSDL_Rect())
		return int(cRet)
	}

	cRet = C.RenderCopy(
		renderer.cSDL_Renderer(),
		texture.cSDL_Texture(),
		src.cSDL_Rect(),
		cInt(dst.X), cInt(dst.Y), cInt(dst.W), cInt(dst.H))
	return int(cRet)
}

func SDL_RenderCopyEx(renderer *SDL_Renderer, texture *SDL_Texture, src, dst *SDL_Rect, angle float64, center *SDL_Point, flip SDL_RendererFlip) int {
	cRet := C.SDL_RenderCopyEx(
		renderer.cSDL_Renderer(),
		texture.cSDL_Texture(),
		src.cSDL_Rect(),
		dst.cSDL_Rect(),
		cDouble(angle),
		center.cSDL_Point(),
		cSDL_RendererFlip(flip))
	return int(cRet)
}

func SDL_RenderDrawPointF(renderer *SDL_Renderer, x, y float32) int {
	cRet := C.SDL_RenderDrawPointF(renderer.cSDL_Renderer(), cFloat(x), cFloat(y))
	return int(cRet)
}

func SDL_RenderDrawPointsF(renderer *SDL_Renderer, points []SDL_FPoint, count int) int {
	cRet := C.SDL_RenderDrawPointsF(
		renderer.cSDL_Renderer(),
		points[0].cSDL_FPoint(),
		cInt(count))
	return int(cRet)
}

func SDL_RenderDrawLineF(renderer *SDL_Renderer, x1, y1, x2, y2 float32) int {
	cRet := C.SDL_RenderDrawLineF(
		renderer.cSDL_Renderer(),
		cFloat(x1),
		cFloat(y1),
		cFloat(x2),
		cFloat(y2))
	return int(cRet)
}

func SDL_RenderDrawLinesF(renderer *SDL_Renderer, points []SDL_FPoint, count int) int {
	cRet := C.SDL_RenderDrawLinesF(
		renderer.cSDL_Renderer(),
		points[0].cSDL_FPoint(),
		cInt(count))
	return int(cRet)
}

func SDL_RenderDrawRectF(renderer *SDL_Renderer, rect *SDL_FRect) int {
	cRet := C.SDL_RenderDrawRectF(renderer.cSDL_Renderer(), rect.cSDL_FRect())
	return int(cRet)
}

func SDL_RenderDrawRectsF(renderer *SDL_Renderer, rects []SDL_FRect, count int) int {
	cRet := C.SDL_RenderDrawRectsF(
		renderer.cSDL_Renderer(),
		rects[0].cSDL_FRect(),
		cInt(count))
	return int(cRet)
}

func SDL_RenderFillRectF(renderer *SDL_Renderer, rect *SDL_FRect) int {
	cRet := C.SDL_RenderFillRectF(renderer.cSDL_Renderer(), rect.cSDL_FRect())
	return int(cRet)
}

func SDL_RenderFillRectsF(renderer *SDL_Renderer, rects []SDL_FRect, count int) int {
	cRet := C.SDL_RenderFillRectsF(
		renderer.cSDL_Renderer(),
		rects[0].cSDL_FRect(),
		cInt(count))
	return int(cRet)
}

func SDL_RenderCopyF(renderer *SDL_Renderer, texture *SDL_Texture, src *SDL_Rect, dst *SDL_FRect) int {
	cRet := C.SDL_RenderCopyF(
		renderer.cSDL_Renderer(),
		texture.cSDL_Texture(),
		src.cSDL_Rect(),
		dst.cSDL_FRect())
	return int(cRet)
}

func SDL_RenderCopyExF(renderer *SDL_Renderer, texture *SDL_Texture, src *SDL_Rect, dst *SDL_FRect, angle float64, center *SDL_FPoint, flip SDL_RendererFlip) int {
	cRet := C.SDL_RenderCopyExF(
		renderer.cSDL_Renderer(),
		texture.cSDL_Texture(),
		src.cSDL_Rect(),
		dst.cSDL_FRect(),
		cDouble(angle),
		center.cSDL_FPoint(),
		cSDL_RendererFlip(flip))
	return int(cRet)
}

func SDL_RenderReadPixels(renderer *SDL_Renderer, rect *SDL_Rect, format uint32, pixels unsafe.Pointer, pitch int) int {
	cRet := C.SDL_RenderReadPixels(
		renderer.cSDL_Renderer(),
		rect.cSDL_Rect(),
		cUint32(format),
		pixels,
		cInt(pitch))
	return int(cRet)
}

func SDL_RenderPresent(renderer *SDL_Renderer) {
	C.SDL_RenderPresent(renderer.cSDL_Renderer())
}

func SDL_DestroyTexture(texture *SDL_Texture) {
	SDL_ClearError()
	C.SDL_DestroyTexture(texture.cSDL_Texture())
}

func SDL_DestroyRenderer(renderer *SDL_Renderer) {
	SDL_ClearError()
	C.SDL_DestroyRenderer(renderer.cSDL_Renderer())
}

func SDL_RenderFlush(renderer *SDL_Renderer) int {
	cRet := C.SDL_RenderFlush(renderer.cSDL_Renderer())
	return int(cRet)
}

func SDL_GL_BindTexture(texture *SDL_Texture, texw, texh *float32) int {
	cRet := C.SDL_GL_BindTexture(
		texture.cSDL_Texture(),
		(*cFloat)(unsafe.Pointer(texw)),
		(*cFloat)(unsafe.Pointer(texh)))
	return int(cRet)
}

func SDL_GL_UnbindTexture(texture *SDL_Texture) int {
	cRet := C.SDL_GL_UnbindTexture(texture.cSDL_Texture())
	return int(cRet)
}

func SDL_RenderGetMetalLayer(renderer *SDL_Renderer) (layer unsafe.Pointer) {
	layer = C.SDL_RenderGetMetalLayer(renderer.cSDL_Renderer())
	return
}

func SDL_RenderGetMetalCommandEncoder(renderer *SDL_Renderer) (encoder unsafe.Pointer) {
	encoder = C.SDL_RenderGetMetalCommandEncoder(renderer.cSDL_Renderer())
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
		texture.cSDL_Texture(),
		rect.cSDL_Rect(),
		(*cUint8)(unsafe.Pointer(yPlanePtr)),
		cInt(yPitch),
		(*cUint8)(unsafe.Pointer(uvPlanePtr)),
		cInt(uvPitch))
	return int(cRet)
}

func SDL_RenderGeometry(renderer *SDL_Renderer, texture *SDL_Texture, vertices []SDL_Vertex, indices []int32) int {
	cTexture := texture.cSDL_Texture()
	cVertices := (*C.SDL_Vertex)(unsafe.Pointer(&vertices[0]))
	cNumVertices := cInt(len(vertices))
	var cIndices *cInt
	cNumIndices := cInt(0)
	if indices != nil {
		cIndices = (*cInt)(unsafe.Pointer(&indices[0]))
		cNumIndices = cInt(len(indices))
	}

	cRet := C.SDL_RenderGeometry(renderer.cSDL_Renderer(), cTexture, cVertices, cNumVertices, cIndices, cNumIndices)
	return int(cRet)
}

func SDL_RenderGeometryRaw(renderer *SDL_Renderer, texture *SDL_Texture, xy []float32, xyStride int, color []SDL_Color, colorStride int, uv []float32, uvStride int, numVertices int, indices interface{}) int {
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

	cTexture := texture.cSDL_Texture()
	cXY := (*cFloat)(&xy[0])
	cXYStride := cInt(xyStride)
	cColor := (*C.SDL_Color)(unsafe.Pointer(&color[0]))
	cColorStride := cInt(colorStride)
	cUV := (*cFloat)(&uv[0])
	cUVStride := cInt(uvStride)
	cNumVertices := cInt(len(xy))
	cNumIndices := cInt(numIndices)
	cSizeIndices := cInt(sizeIndices)

	cRet := C.RenderGeometryRaw(renderer.cSDL_Renderer(), cTexture, cXY, cXYStride, cColor, cColorStride, cUV, cUVStride, cNumVertices, cIndices, cNumIndices, cSizeIndices)
	return int(cRet)
}

func SDL_SetTextureUserData(texture *SDL_Texture, userdata unsafe.Pointer) int {
	cRet := C.SDL_SetTextureUserData(texture.cSDL_Texture(), userdata)
	return int(cRet)
}

func SDL_GetTextureUserData(texture *SDL_Texture) (userdata unsafe.Pointer) {
	userdata = C.SDL_GetTextureUserData(texture.cSDL_Texture())
	return
}

func SDL_RenderWindowToLogical(renderer *SDL_Renderer, windowX, windowY int, logicalX, logicalY *float32) {
	cWindowX := cInt(windowX)
	cWindowY := cInt(windowY)
	cLogicalX := (*cFloat)(unsafe.Pointer(logicalX))
	cLogicalY := (*cFloat)(unsafe.Pointer(logicalY))
	C.SDL_RenderWindowToLogical(renderer.cSDL_Renderer(), cWindowX, cWindowY, cLogicalX, cLogicalY)
}

func SDL_RenderLogicalToWindow(renderer *SDL_Renderer, logicalX, logicalY float32, windowX, windowY *int) {
	cLogicalX := cFloat(logicalX)
	cLogicalY := cFloat(logicalY)
	cWindowX := (*cInt)(unsafe.Pointer(windowX))
	cWindowY := (*cInt)(unsafe.Pointer(windowY))
	C.SDL_RenderLogicalToWindow(renderer.cSDL_Renderer(), cLogicalX, cLogicalY, cWindowX, cWindowY)
	return
}

func SDL_RenderSetVSync(renderer *SDL_Renderer, vsync SDL_bool) int {
	cVsync := cInt(vsync)
	cRet := C.SDL_RenderSetVSync(renderer.cSDL_Renderer(), cVsync)
	return int(cRet)
}

func SDL_RenderGetWindow(renderer *SDL_Renderer) (window *SDL_Window) {
	cWindow := C.SDL_RenderGetWindow(renderer.cSDL_Renderer())
	window = (*SDL_Window)(unsafe.Pointer(cWindow))
	return
}
