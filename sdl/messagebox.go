package sdl

/*
#include "messagebox.h"
*/
import "C"
import "unsafe"
import "reflect"

// typedef
type SDL_MessageBoxFlags = int32
type SDL_MessageBoxButtonFlags = int32
type SDL_MessageBoxColorType = int32

// enum
const (
	SDL_MESSAGEBOX_ERROR                 SDL_MessageBoxFlags = 16
	SDL_MESSAGEBOX_WARNING               SDL_MessageBoxFlags = 32
	SDL_MESSAGEBOX_INFORMATION           SDL_MessageBoxFlags = 64
	SDL_MESSAGEBOX_BUTTONS_LEFT_TO_RIGHT SDL_MessageBoxFlags = 128
	SDL_MESSAGEBOX_BUTTONS_RIGHT_TO_LEFT SDL_MessageBoxFlags = 256
)

const (
	SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT SDL_MessageBoxButtonFlags = 1
	SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT SDL_MessageBoxButtonFlags = 2
)

const (
	SDL_MESSAGEBOX_COLOR_BACKGROUND        SDL_MessageBoxColorType = 0
	SDL_MESSAGEBOX_COLOR_TEXT              SDL_MessageBoxColorType = 1
	SDL_MESSAGEBOX_COLOR_BUTTON_BORDER     SDL_MessageBoxColorType = 2
	SDL_MESSAGEBOX_COLOR_BUTTON_BACKGROUND SDL_MessageBoxColorType = 3
	SDL_MESSAGEBOX_COLOR_BUTTON_SELECTED   SDL_MessageBoxColorType = 4
	SDL_MESSAGEBOX_COLOR_MAX               SDL_MessageBoxColorType = 5
)

// struct
type SDL_MessageBoxButtonData struct {
	Flags    uint32
	Buttonid int32
	Text     *int8
}

type SDL_MessageBoxColor struct {
	R uint8
	G uint8
	B uint8
}

type SDL_MessageBoxColorScheme struct {
	Colors [5]SDL_MessageBoxColor
}

type SDL_MessageBoxData struct {
	Flags       uint32
	Window      *SDL_Window
	Title       *int8
	Message     *int8
	Numbuttons  int32
	Buttons     *SDL_MessageBoxButtonData
	ColorScheme *SDL_MessageBoxColorScheme
}

func cMessageBoxColor(mc *SDL_MessageBoxColor) *C.SDL_MessageBoxColor {
	return (*C.SDL_MessageBoxColor)(unsafe.Pointer(mc))
}

func cMessageBoxColorScheme(mcs *SDL_MessageBoxColorScheme) *C.SDL_MessageBoxColorScheme {
	return (*C.SDL_MessageBoxColorScheme)(unsafe.Pointer(mcs))
}

func cMessageBoxButtonData(mbd *SDL_MessageBoxButtonData) *C.SDL_MessageBoxButtonData {
	return (*C.SDL_MessageBoxButtonData)(unsafe.Pointer(mbd))
}

func cMessageBoxData(md *SDL_MessageBoxData) *C.SDL_MessageBoxData {
	return (*C.SDL_MessageBoxData)(unsafe.Pointer(md))
}

func (data *SDL_MessageBoxData) ButtonsAsSlice() []SDL_MessageBoxButtonData {
	var b []SDL_MessageBoxButtonData
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh.Len = int(data.Numbuttons)
	sh.Cap = int(data.Numbuttons)
	sh.Data = uintptr(unsafe.Pointer(data.Buttons))
	return b
}

func (data *SDL_MessageBoxData) TitleAsString() string {
	cStr := (*cChar)(unsafe.Pointer(data.Title))
	defer C.free(unsafe.Pointer(cStr))

	return SDL_GoString(cStr)
}

func (data *SDL_MessageBoxData) MessageAsString() string {
	cStr := (*cChar)(unsafe.Pointer(data.Message))
	defer C.free(unsafe.Pointer(cStr))

	return SDL_GoString(cStr)
}

func (button *SDL_MessageBoxButtonData) TextAsString() string {
	cStr := (*cChar)(unsafe.Pointer(button.Text))
	defer C.free(unsafe.Pointer(cStr))

	return SDL_GoString(cStr)
}

type MessageBoxButtonData struct {
	Flags    uint32
	Buttonid int32
	Text     string
}

type MessageBoxData struct {
	Flags       uint32
	Window      *SDL_Window
	Title       string
	Message     string
	Numbuttons  int32
	Buttons     []MessageBoxButtonData
	ColorScheme *SDL_MessageBoxColorScheme
}

func SDL_ShowSimpleMessageBox(flags SDL_MessageBoxFlags, title, message string, window *SDL_Window) int {
	cTitle := SDL_CreateCString(SDL_GetMemoryPool(), title)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cTitle)

	cMessage := SDL_CreateCString(SDL_GetMemoryPool(), message)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cMessage)

	cRet := C.SDL_ShowSimpleMessageBox(cUint32(flags), cTitle.(*cChar), cMessage.(*cChar), cWindow(window))
	return int(cRet)
}

func SDL_ShowMessageBox(data *MessageBoxData) (buttonid int32) {
	cTitle := SDL_CreateCString(SDL_GetMemoryPool(), data.Title)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cTitle)

	cMessage := SDL_CreateCString(SDL_GetMemoryPool(), data.Message)
	defer SDL_DestroyCString(SDL_GetMemoryPool(), cMessage)

	var cbuttons []SDL_MessageBoxButtonData
	var cbtntexts []*cChar
	defer func(texts []*cChar) {
		for _, t := range texts {
			defer SDL_DestroyCString(SDL_GetMemoryPool(), t)
		}
	}(cbtntexts)

	for _, btn := range data.Buttons {
		ctext := SDL_CreateCString(SDL_GetMemoryPool(), btn.Text)
		cbtn := SDL_MessageBoxButtonData{
			Flags:    btn.Flags,
			Buttonid: btn.Buttonid,
			Text:     (*int8)(unsafe.Pointer(ctext.(*cChar))),
		}

		cbuttons = append(cbuttons, cbtn)
		cbtntexts = append(cbtntexts, ctext.(*cChar))
	}

	var buttonPtr *SDL_MessageBoxButtonData
	if len(cbuttons) > 0 {
		buttonPtr = &cbuttons[0]
	}
	data2 := &SDL_MessageBoxData{
		Flags:       data.Flags,
		Window:      data.Window,
		Title:       (*int8)(unsafe.Pointer(cTitle.(*cChar))),
		Message:     (*int8)(unsafe.Pointer(cMessage.(*cChar))),
		Numbuttons:  data.Numbuttons,
		Buttons:     buttonPtr,
		ColorScheme: data.ColorScheme,
	}

	cData2 := cMessageBoxData(data2)
	cButtonid := C.ShowMessageBox(*cData2)
	buttonid = int32(cButtonid)
	return buttonid
}
