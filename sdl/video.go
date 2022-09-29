package sdl

/*
#include "video.h"
*/
import "C"
import "unsafe"

func cSDL_Window(w *SDL_Window) *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(w))
}

func cSDL_DisplayMode(dm *SDL_DisplayMode) *C.SDL_DisplayMode {
	return (*C.SDL_DisplayMode)(unsafe.Pointer(dm))
}

func SDL_GetDisplayName(displayIndex int) string {
	cName := C.SDL_GetDisplayName(cInt(displayIndex))
	if cName == nil {
		return ""
	}
	return SDL_GoString(cName)
}

func SDL_GetNumVideoDisplays() int {
	cRet := C.SDL_GetNumVideoDisplays()
	return int(cRet)
}

func SDL_GetNumVideoDrivers() int {
	cRet := C.SDL_GetNumVideoDrivers()
	return int(cRet)
}

func SDL_GetVideoDriver(index int) string {
	cStr := C.SDL_GetVideoDriver(cInt(index))
	return SDL_GoString(cStr)
}

func SDL_VideoInit(driverName string) int {
	cDriverName := SDL_CreateCString(SDL_GetMemoryPool(), driverName)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cDriverName) // memory free

	cRet := C.SDL_VideoInit(cDriverName.(*cChar))
	return int(cRet)
}

func SDL_VideoQuit() {
	C.SDL_VideoQuit()
}

func SDL_GetCurrentVideoDriver() string {
	cName := C.SDL_GetCurrentVideoDriver()
	if cName == nil {
		return ""
	}
	return SDL_GoString(cName)
}

func SDL_GetNumDisplayModes(displayIndex int) int {
	cRet := C.SDL_GetNumDisplayModes(cInt(displayIndex))
	return int(cRet)
}

func SDL_GetDisplayBounds(displayIndex int, rect *SDL_Rect) int {
	cRet := C.SDL_GetDisplayBounds(cInt(displayIndex), cSDL_Rect(rect))
	return int(cRet)
}

func SDL_GetDisplayUsableBounds(displayIndex int, rect *SDL_Rect) int {
	cRet := C.SDL_GetDisplayUsableBounds(cInt(displayIndex), cSDL_Rect(rect))
	return int(cRet)
}

func SDL_GetDisplayMode(displayIndex int, modeIndex int, mode *SDL_DisplayMode) int {
	cRet := C.SDL_GetDisplayMode(cInt(displayIndex), cInt(modeIndex), cSDL_DisplayMode(mode))
	return int(cRet)
}

func SDL_GetDesktopDisplayMode(displayIndex int, mode *SDL_DisplayMode) int {
	cRet := C.SDL_GetDesktopDisplayMode(cInt(displayIndex), cSDL_DisplayMode(mode))
	return int(cRet)
}

func SDL_GetCurrentDisplayMode(displayIndex int, mode *SDL_DisplayMode) int {
	cRet := C.SDL_GetCurrentDisplayMode(cInt(displayIndex), cSDL_DisplayMode(mode))
	return int(cRet)
}

func SDL_GetClosestDisplayMode(displayIndex int, mode *SDL_DisplayMode, closest *SDL_DisplayMode) *SDL_DisplayMode {
	cDisplayMode := C.SDL_GetClosestDisplayMode(cInt(displayIndex), cSDL_DisplayMode(mode), cSDL_DisplayMode(closest))
	displayMode := (*SDL_DisplayMode)(unsafe.Pointer(cDisplayMode))
	if displayMode == nil {
		return nil
	}
	return displayMode
}

func SDL_GetPointDisplayIndex(p *SDL_Point) int {
	cIndex := C.SDL_GetPointDisplayIndex(cSDL_Point(p))
	index := int(cIndex)
	return index
}

func SDL_GetRectDisplayIndex(r *SDL_Rect) int {
	cIndex := C.SDL_GetRectDisplayIndex(cSDL_Rect(r))
	index := int(cIndex)
	return index
}

func SDL_GetDisplayDPI(displayIndex int, ddpi, hdpi, vdpi *float32) int {
	cDdpi := (*cFloat)(unsafe.Pointer(ddpi))
	cHdpi := (*cFloat)(unsafe.Pointer(hdpi))
	cVdpi := (*cFloat)(unsafe.Pointer(vdpi))
	cRet := C.SDL_GetDisplayDPI(cInt(displayIndex), cDdpi, cHdpi, cVdpi)
	return int(cRet)
}

func SDL_GetWindowDisplayIndex(window *SDL_Window) int {
	cRet := C.SDL_GetWindowDisplayIndex(cSDL_Window(window))
	return int(cRet)
}

func SDL_SetWindowDisplayMode(window *SDL_Window, mode *SDL_DisplayMode) int {
	cRet := C.SDL_SetWindowDisplayMode(cSDL_Window(window), cSDL_DisplayMode(mode))
	return int(cRet)
}

func SDL_GetWindowDisplayMode(window *SDL_Window, mode *SDL_DisplayMode) int {
	cRet := C.SDL_GetWindowDisplayMode(cSDL_Window(window), cSDL_DisplayMode(mode))
	return int(cRet)
}

func SDL_GetWindowPixelFormat(window *SDL_Window) uint32 {
	cRet := C.SDL_GetWindowPixelFormat(cSDL_Window(window))
	return uint32(cRet)
}

func SDL_CreateWindow(title string, x, y, w, h int32, flags SDL_WindowFlags) *SDL_Window {
	cTitle := SDL_CreateCString(SDL_GetMemoryPool(), title)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cTitle) // memory free

	var cWindow = C.SDL_CreateWindow(cTitle.(*cChar), cInt(x), cInt(y), cInt(w), cInt(h), cUint(flags))
	if cWindow == nil {
		return nil
	}
	return (*SDL_Window)(unsafe.Pointer(cWindow))
}

func SDL_CreateWindowFrom(data unsafe.Pointer) *SDL_Window {
	cWindow := C.SDL_CreateWindowFrom(data)
	if cWindow == nil {
		return nil
	}
	return (*SDL_Window)(unsafe.Pointer(cWindow))
}

func SDL_DestroyWindow(window *SDL_Window) {
	SDL_ClearError()
	C.SDL_DestroyWindow(cSDL_Window(window))
}

func SDL_GetWindowID(window *SDL_Window) uint32 {
	cId := C.SDL_GetWindowID(cSDL_Window(window))
	return uint32(cId)
}

func SDL_GetWindowFromID(id uint32) *SDL_Window {
	cWindow := C.SDL_GetWindowFromID(cUint(id))
	if cWindow == nil {
		return nil
	}
	return (*SDL_Window)(unsafe.Pointer((cWindow)))
}

func SDL_GetWindowFlags(window *SDL_Window) SDL_WindowFlags {
	cRet := C.SDL_GetWindowFlags(cSDL_Window(window))
	return SDL_WindowFlags(cRet)
}

func SDL_SetWindowTitle(window *SDL_Window, title string) {
	cTitle := SDL_CreateCString(SDL_GetMemoryPool(), title)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cTitle) // memory free

	C.SDL_SetWindowTitle(cSDL_Window(window), cTitle.(*cChar))
}

func SDL_GetWindowTitle(window *SDL_Window) string {
	cTitle := C.SDL_GetWindowTitle(cSDL_Window(window))
	return SDL_GoString(cTitle)
}

func SDL_SetWindowIcon(window *SDL_Window, icon *SDL_Surface) {
	C.SDL_SetWindowIcon(cSDL_Window(window), cSDL_Surface(icon))
}

func SDL_SetWindowData(window *SDL_Window, name string, userdata unsafe.Pointer) unsafe.Pointer {
	cName := SDL_CreateCString(SDL_GetMemoryPool(), name)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cName) // memory free

	return C.SDL_SetWindowData(cSDL_Window(window), cName.(*cChar), userdata)
}

func SDL_GetWindowData(window *SDL_Window, name string) unsafe.Pointer {
	cName := SDL_CreateCString(SDL_GetMemoryPool(), name)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cName) // memory free

	cRet := C.SDL_GetWindowData(cSDL_Window(window), cName.(*cChar))
	return cRet
}

func SDL_SetWindowPosition(window *SDL_Window, x, y int32) {
	C.SDL_SetWindowPosition(cSDL_Window(window), cInt(x), cInt(y))
}

func SDL_GetWindowPosition(window *SDL_Window, x, y *int32) {
	cX := (*cInt)(unsafe.Pointer(x))
	cY := (*cInt)(unsafe.Pointer(y))
	C.SDL_GetWindowPosition(cSDL_Window(window), cX, cY)
}

func SDL_SetWindowResizable(window *SDL_Window, resizable SDL_bool) {
	C.SDL_SetWindowResizable(cSDL_Window(window), cSDL_bool(resizable))
}

func SDL_SetWindowSize(window *SDL_Window, w, h int32) {
	C.SDL_SetWindowSize(cSDL_Window(window), cInt(w), cInt(h))
}

func SDL_GetWindowSize(window *SDL_Window, w, h *int32) {
	cW := (*cInt)(unsafe.Pointer(w))
	cH := (*cInt)(unsafe.Pointer(h))
	C.SDL_GetWindowSize(cSDL_Window(window), cW, cH)
}

func SDL_SetWindowMinimumSize(window *SDL_Window, minW, minH int32) {
	C.SDL_SetWindowMinimumSize(cSDL_Window(window), cInt(minW), cInt(minH))
}

func SDL_GetWindowMinimumSize(window *SDL_Window, w, h *int32) {
	cW := (*cInt)(unsafe.Pointer(w))
	cH := (*cInt)(unsafe.Pointer(h))
	C.SDL_GetWindowMinimumSize(cSDL_Window(window), cW, cH)
}

func SDL_SetWindowMaximumSize(window *SDL_Window, maxW, maxH int32) {
	C.SDL_SetWindowMaximumSize(cSDL_Window(window), cInt(maxW), cInt(maxH))
}

func SDL_GetWindowMaximumSize(window *SDL_Window, w, h *int32) {
	cW := (*cInt)(unsafe.Pointer(w))
	cH := (*cInt)(unsafe.Pointer(h))
	C.SDL_GetWindowMaximumSize(cSDL_Window(window), cW, cH)
}

func SDL_SetWindowBordered(window *SDL_Window, bordered SDL_bool) {
	C.SDL_SetWindowBordered(cSDL_Window(window), cSDL_bool(bordered))
}

func SDL_ShowWindow(window *SDL_Window) {
	C.SDL_ShowWindow(cSDL_Window(window))
}

func SDL_HideWindow(window *SDL_Window) {
	C.SDL_HideWindow(cSDL_Window(window))
}

func SDL_RaiseWindow(window *SDL_Window) {
	C.SDL_RaiseWindow(cSDL_Window(window))
}

func SDL_MaximizeWindow(window *SDL_Window) {
	C.SDL_MaximizeWindow(cSDL_Window(window))
}

func SDL_MinimizeWindow(window *SDL_Window) {
	C.SDL_MinimizeWindow(cSDL_Window(window))
}

func SDL_RestoreWindow(window *SDL_Window) {
	C.SDL_RestoreWindow(cSDL_Window(window))
}

func SDL_SetWindowFullscreen(window *SDL_Window, flags uint32) int {
	cRet := C.SDL_SetWindowFullscreen(cSDL_Window(window), cUint(flags))
	return int(cRet)
}

func SDL_GetWindowSurface(window *SDL_Window) *SDL_Surface {
	surface := (*SDL_Surface)(unsafe.Pointer(C.SDL_GetWindowSurface(cSDL_Window(window))))
	if surface == nil {
		return nil
	}
	return surface
}

func SDL_UpdateWindowSurface(window *SDL_Window) int {
	cRet := C.SDL_UpdateWindowSurface(cSDL_Window(window))
	return int(cRet)
}

func SDL_UpdateWindowSurfaceRects(window *SDL_Window, rects []SDL_Rect) int {
	cRet := C.SDL_UpdateWindowSurfaceRects(cSDL_Window(window), cSDL_Rect(&rects[0]), cInt(len(rects)))
	return int(cRet)
}

func SDL_SetWindowGrab(window *SDL_Window, grabbed SDL_bool) {
	C.SDL_SetWindowGrab(cSDL_Window(window), cSDL_bool(grabbed))
}

func SDL_GetWindowGrab(window *SDL_Window) bool {
	cRet := C.SDL_GetWindowGrab(cSDL_Window(window))
	return cRet != 0
}

func SDL_SetWindowBrightness(window *SDL_Window, brightness float32) int {
	cRet := C.SDL_SetWindowBrightness(cSDL_Window(window), cFloat(brightness))
	return int(cRet)
}

func SDL_GetWindowBrightness(window *SDL_Window) float32 {
	return float32(C.SDL_GetWindowBrightness(cSDL_Window(window)))
}

func SDL_SetWindowGammaRamp(window *SDL_Window, red, green, blue *[256]uint16) int {
	cRed := (*cUint16)(unsafe.Pointer(red))
	cGreen := (*cUint16)(unsafe.Pointer(green))
	cBlue := (*cUint16)(unsafe.Pointer(blue))
	cRet := C.SDL_SetWindowGammaRamp(cSDL_Window(window), cRed, cGreen, cBlue)
	return int(cRet)
}

func SDL_GetWindowGammaRamp(window *SDL_Window, red, green, blue *[256]uint16) int {
	cRed := (*cUint16)(unsafe.Pointer(red))
	cGreen := (*cUint16)(unsafe.Pointer(green))
	cBlue := (*cUint16)(unsafe.Pointer(blue))
	cRet := C.SDL_GetWindowGammaRamp(cSDL_Window(window), cRed, cGreen, cBlue)
	return int(cRet)
}

func SDL_SetWindowOpacity(window *SDL_Window, opacity float32) int {
	cRet := C.SDL_SetWindowOpacity(cSDL_Window(window), cFloat(opacity))
	return int(cRet)
}

func SDL_GetWindowOpacity(window *SDL_Window, opacity *float32) int {
	cRet := C.SDL_GetWindowOpacity(cSDL_Window(window), (*cFloat)(unsafe.Pointer(opacity)))
	return int(cRet)
}

func SDL_IsScreenSaverEnabled() bool {
	cRet := C.SDL_IsScreenSaverEnabled()
	return cRet != 0
}

func SDL_EnableScreenSaver() {
	C.SDL_EnableScreenSaver()
}

func SDL_DisableScreenSaver() {
	C.SDL_DisableScreenSaver()
}

func SDL_GL_LoadLibrary(path string) int {
	cPath := SDL_CreateCString(SDL_GetMemoryPool(), path)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cPath) // memory free

	cRet := C.SDL_GL_LoadLibrary(cPath.(*cChar))
	return int(cRet)
}

func SDL_GL_GetProcAddress(proc string) unsafe.Pointer {
	cProc := SDL_CreateCString(SDL_GetMemoryPool(), proc)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cProc) // memory free

	return C.SDL_GL_GetProcAddress(cProc.(*cChar))
}

func SDL_GL_UnloadLibrary() {
	C.SDL_GL_UnloadLibrary()
}

func SDL_GL_ExtensionSupported(extension string) bool {
	cExtension := SDL_CreateCString(SDL_GetMemoryPool(), extension)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cExtension) // memory free

	cRet := C.SDL_GL_ExtensionSupported(cExtension.(*cChar))
	return cRet != 0
}

func SDL_GL_SetAttribute(attr SDL_GLattr, value int) int {
	cRet := C.SDL_GL_SetAttribute(C.SDL_GLattr(attr), cInt(value))
	return int(cRet)
}

func SDL_GL_GetAttribute(attr SDL_GLattr, value *int) int {
	cValue := (*cInt)(unsafe.Pointer(value))
	cRet := C.SDL_GL_GetAttribute(C.SDL_GLattr(attr), cValue)
	return int(cRet)
}

func SDL_GL_CreateContext(window *SDL_Window) SDL_GLContext {
	context := SDL_GLContext(C.SDL_GL_CreateContext(cSDL_Window(window)))
	if context == nil {
		return nil
	}
	return context
}

func SDL_GL_MakeCurrent(window *SDL_Window, glcontext SDL_GLContext) int {
	cRet := C.SDL_GL_MakeCurrent(cSDL_Window(window), C.SDL_GLContext(glcontext))
	return int(cRet)
}

func SDL_GL_SetSwapInterval(interval int) int {
	cRet := C.SDL_GL_SetSwapInterval(cInt(interval))
	return int(cRet)
}

func SDL_GL_GetSwapInterval() int {
	cRet := C.SDL_GL_GetSwapInterval()
	return int(cRet)
}

func SDL_GL_GetDrawableSize(window *SDL_Window, w, h *int32) {
	cW := (*cInt)(unsafe.Pointer(w))
	cH := (*cInt)(unsafe.Pointer(h))
	C.SDL_GL_GetDrawableSize(cSDL_Window(window), cW, cH)
}

func SDL_GL_SwapWindow(window *SDL_Window) {
	C.SDL_GL_SwapWindow(cSDL_Window(window))
}

func SDL_GL_DeleteContext(context SDL_GLContext) {
	C.SDL_GL_DeleteContext(C.SDL_GLContext(context))
}

func SDL_FlashWindow(window *SDL_Window, operation SDL_FlashOperation) int {
	cRet := C.SDL_FlashWindow(cSDL_Window(window), C.SDL_FlashOperation(operation))
	return int(cRet)
}

func SDL_SetWindowAlwaysOnTop(window *SDL_Window, onTop SDL_bool) {
	C.SDL_SetWindowAlwaysOnTop(cSDL_Window(window), cSDL_bool(onTop))
}

func SDL_GetWindowKeyboardGrab(window *SDL_Window, grabbed SDL_bool) {
	C.SDL_SetWindowKeyboardGrab(cSDL_Window(window), cSDL_bool(grabbed))
}

func SDL_GetWindowICCProfile(window *SDL_Window, size *uint) unsafe.Pointer {
	cSize := (*cSize_t)(unsafe.Pointer(size))
	cIccProfile := C.SDL_GetWindowICCProfile(cSDL_Window(window), cSize)
	return cIccProfile
}

func SDL_SetWindowMouseRect(window *SDL_Window, rect *SDL_Rect) int {
	cRect := (*C.SDL_Rect)(unsafe.Pointer(rect))
	cRet := C.SDL_SetWindowMouseRect(cSDL_Window(window), cRect)
	return int(cRet)
}

func SDL_GetWindowMouseRect(window *SDL_Window) (rect *SDL_Rect) {
	cRect := C.SDL_GetWindowMouseRect(cSDL_Window(window))
	rect = (*SDL_Rect)(unsafe.Pointer(cRect))
	return
}
