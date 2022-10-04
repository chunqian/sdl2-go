package sdl

/*
#include "events.h"
*/
import "C"
import (
	"reflect"
	"sync"
	"unsafe"
)

// define
const (
	SDL_RELEASED = 0
	SDL_PRESSED  = 1
)

// typedef
type SDL_EventType = uint32
type SDL_eventaction = int32

// enum
const (
	SDL_FIRSTEVENT              SDL_EventType = 0
	SDL_QUIT                    SDL_EventType = 0x100
	SDL_APP_TERMINATING         SDL_EventType = 0x101
	SDL_APP_LOWMEMORY           SDL_EventType = 0x102
	SDL_APP_WILLENTERBACKGROUND SDL_EventType = 0x103
	SDL_APP_DIDENTERBACKGROUND  SDL_EventType = 0x104
	SDL_APP_WILLENTERFOREGROUND SDL_EventType = 0x105
	SDL_APP_DIDENTERFOREGROUND  SDL_EventType = 0x106
	SDL_LOCALECHANGED           SDL_EventType = 0x107

	SDL_DISPLAYEVENT SDL_EventType = 0x150

	SDL_WINDOWEVENT SDL_EventType = 0x200
	SDL_SYSWMEVENT  SDL_EventType = 0x201

	SDL_KEYDOWN         SDL_EventType = 0x300
	SDL_KEYUP           SDL_EventType = 0x301
	SDL_TEXTEDITING     SDL_EventType = 0x302
	SDL_TEXTINPUT       SDL_EventType = 0x303
	SDL_KEYMAPCHANGED   SDL_EventType = 0x304
	SDL_TEXTEDITING_EXT SDL_EventType = 0x305

	SDL_MOUSEMOTION     SDL_EventType = 0x400
	SDL_MOUSEBUTTONDOWN SDL_EventType = 0x401
	SDL_MOUSEBUTTONUP   SDL_EventType = 0x402
	SDL_MOUSEWHEEL      SDL_EventType = 0x403

	SDL_JOYAXISMOTION     SDL_EventType = 0x600
	SDL_JOYBALLMOTION     SDL_EventType = 0x601
	SDL_JOYHATMOTION      SDL_EventType = 0x602
	SDL_JOYBUTTONDOWN     SDL_EventType = 0x603
	SDL_JOYBUTTONUP       SDL_EventType = 0x604
	SDL_JOYDEVICEADDED    SDL_EventType = 0x605
	SDL_JOYDEVICEREMOVED  SDL_EventType = 0x606
	SDL_JOYBATTERYUPDATED SDL_EventType = 0x607

	SDL_CONTROLLERAXISMOTION     SDL_EventType = 0x650
	SDL_CONTROLLERBUTTONDOWN     SDL_EventType = 0x651
	SDL_CONTROLLERBUTTONUP       SDL_EventType = 0x652
	SDL_CONTROLLERDEVICEADDED    SDL_EventType = 0x653
	SDL_CONTROLLERDEVICEREMOVED  SDL_EventType = 0x654
	SDL_CONTROLLERDEVICEREMAPPED SDL_EventType = 0x655
	SDL_CONTROLLERTOUCHPADDOWN   SDL_EventType = 0x656
	SDL_CONTROLLERTOUCHPADMOTION SDL_EventType = 0x657
	SDL_CONTROLLERTOUCHPADUP     SDL_EventType = 0x658
	SDL_CONTROLLERSENSORUPDATE   SDL_EventType = 0x659

	SDL_FINGERDOWN   SDL_EventType = 0x700
	SDL_FINGERUP     SDL_EventType = 0x701
	SDL_FINGERMOTION SDL_EventType = 0x702

	SDL_DOLLARGESTURE SDL_EventType = 0x800
	SDL_DOLLARRECORD  SDL_EventType = 0x801
	SDL_MULTIGESTURE  SDL_EventType = 0x802

	SDL_CLIPBOARDUPDATE SDL_EventType = 0x900

	SDL_DROPFILE     SDL_EventType = 0x1000
	SDL_DROPTEXT     SDL_EventType = 0x1001
	SDL_DROPBEGIN    SDL_EventType = 0x1002
	SDL_DROPCOMPLETE SDL_EventType = 0x1003

	SDL_AUDIODEVICEADDED   SDL_EventType = 0x1100
	SDL_AUDIODEVICEREMOVED SDL_EventType = 0x1101

	SDL_SENSORUPDATE SDL_EventType = 0x1200

	SDL_RENDER_TARGETS_RESET SDL_EventType = 0x2000
	SDL_RENDER_DEVICE_RESET  SDL_EventType = 0x2001

	SDL_POLLSENTINEL SDL_EventType = 0x7F00

	SDL_USEREVENT SDL_EventType = 0x8000

	SDL_LASTEVENT SDL_EventType = 0xFFFF
)

const (
	SDL_ADDEVENT  SDL_eventaction = 0
	SDL_PEEKEVENT SDL_eventaction = 1
	SDL_GETEVENT  SDL_eventaction = 2
)

// struct
type SDL_CommonEvent struct {
	Type      uint32
	Timestamp uint32
}

type SDL_DisplayEvent struct {
	Type      uint32
	Timestamp uint32
	Display   uint32
	Event     uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Data1     int32
}

type SDL_WindowEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Event     uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Data1     int32
	Data2     int32
}

type SDL_KeyboardEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	State     uint8
	Repeat    uint8
	Padding2  uint8
	Padding3  uint8
	Keysym    SDL_Keysym
}

type SDL_TextEditingEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Text      [32]int8
	Start     int32
	Length    int32
}

type SDL_TextEditingExtEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Text      *int8
	Start     int32
	Length    int32
}

type SDL_TextInputEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Text      [32]int8
}

type SDL_MouseMotionEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	State     uint32
	X         int32
	Y         int32
	XRel      int32
	YRel      int32
}

type SDL_MouseButtonEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	Button    uint8
	State     uint8
	Clicks    uint8
	Padding1  uint8
	X         int32
	Y         int32
}

type SDL_MouseWheelEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	X         int32
	Y         int32
	Direction uint32
	PreciseX  float32
	PreciseY  float32
}

type SDL_JoyAxisEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Axis      uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Value     int16
	Padding4  uint16
}

type SDL_JoyBallEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Ball      uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	XRel      int16
	YRel      int16
}

type SDL_JoyHatEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Hat       uint8
	Value     uint8
	Padding1  uint8
	Padding2  uint8
}

type SDL_JoyButtonEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Button    uint8
	State     uint8
	Padding1  uint8
	Padding2  uint8
}

type SDL_JoyDeviceEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
}

type SDL_JoyBatteryEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Level     int32
}

type SDL_ControllerAxisEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Axis      uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Value     int16
	Padding4  uint16
}

type SDL_ControllerButtonEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Button    uint8
	State     uint8
	Padding1  uint8
	Padding2  uint8
}

type SDL_ControllerDeviceEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
}

type SDL_ControllerTouchpadEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Touchpad  int32
	Finger    int32
	X         float32
	Y         float32
	Pressure  float32
}

type SDL_ControllerSensorEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Sensor    int32
	Data      [3]float32
}

type SDL_AudioDeviceEvent struct {
	Type      uint32
	Timestamp uint32
	Which     uint32
	Iscapture uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
}

type SDL_TouchFingerEvent struct {
	Type      uint32
	Timestamp uint32
	TouchId   int64
	FingerId  int64
	X         float32
	Y         float32
	Dx        float32
	Dy        float32
	Pressure  float32
	WindowID  uint32
}

type SDL_MultiGestureEvent struct {
	Type       uint32
	Timestamp  uint32
	TouchId    int64
	DTheta     float32
	DDist      float32
	X          float32
	Y          float32
	NumFingers uint16
	Padding    uint16
}

type SDL_DollarGestureEvent struct {
	Type       uint32
	Timestamp  uint32
	TouchId    int64
	GestureId  int64
	NumFingers uint32
	Error      float32
	X          float32
	Y          float32
}

type SDL_DropEvent struct {
	Type      uint32
	Timestamp uint32
	File      *int8
	WindowID  uint32
}

type SDL_SensorEvent struct {
	Type      uint32
	Timestamp uint32
	Which     int32
	Data      [6]float32
}

type SDL_QuitEvent struct {
	Type      uint32
	Timestamp uint32
}

type SDL_OSEvent struct {
	Type      uint32
	Timestamp uint32
}

type SDL_UserEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Code      int32
	Data1     unsafe.Pointer
	Data2     unsafe.Pointer
}

type SDL_SysWMEvent struct {
	Type      uint32
	Timestamp uint32
	Msg       *SDL_SysWMmsg
}

type SDL_Event struct {
	Type    uint32
	Padding [56]uint8 // padding
}

func cEvent(s *SDL_Event) *C.SDL_Event {
	return (*C.SDL_Event)(unsafe.Pointer(s))
}

func (event SDL_TextEditingExtEvent) TextAsString() string {
	cStr := (*cChar)(unsafe.Pointer(event.Text))
	defer C.free(unsafe.Pointer(cStr))

	return SDL_GoString(cStr)
}

func (event SDL_DropEvent) FileAsString() string {
	cStr := (*cChar)(unsafe.Pointer(event.File))
	defer C.free(unsafe.Pointer(cStr))

	return SDL_GoString(cStr)
}

func SDL_PumpEvents() {
	C.SDL_PumpEvents()
}

func SDL_PeepEvents(events []SDL_Event, numevents int, action SDL_eventaction, minType SDL_EventType, maxType SDL_EventType) int {
	if events == nil {
		return 0
	}

	cRet := C.SDL_PeepEvents(cEvent(&events[0]), cInt(numevents), cEventaction(action), cUint(minType), cUint(maxType))
	return int(cRet)
}

func SDL_HasEvent(eventType SDL_EventType) bool {
	cRet := C.SDL_HasEvent(cUint(eventType))
	return cRet != 0
}

func SDL_HasEvents(minType SDL_EventType, maxType SDL_EventType) bool {
	cRet := C.SDL_HasEvents(cUint(minType), cUint(maxType))
	return cRet != 0
}

func SDL_FlushEvent(eventType SDL_EventType) {
	C.SDL_FlushEvent(cUint(eventType))
}

func SDL_FlushEvents(minType SDL_EventType, maxType SDL_EventType) {
	C.SDL_FlushEvents(cUint(minType), cUint(maxType))
}

func SDL_PollEvent(event *SDL_Event) bool {
	cRet := C.SDL_PollEvent(cEvent(event))
	return cRet != 0
}

func SDL_WaitEvent(event *SDL_Event) int {
	cRet := C.SDL_WaitEvent(cEvent(event))
	return int(cRet)
}

func SDL_WaitEventTimeout(event *SDL_Event, timeout int) int {
	cRet := C.SDL_WaitEventTimeout(cEvent(event), cInt(timeout))
	return int(cRet)
}

func SDL_PushEvent(event *SDL_Event) {
	C.SDL_PushEvent(cEvent(event))
}

type SDL_EventFilterCallback = func(userdata any, event *SDL_Event) int32

type SDL_EventWatcher struct {
	callback SDL_EventFilterCallback
	userdata any
	removed  bool
}

var SDL_event_watchers_lock sync.Mutex
var SDL_EventOK SDL_EventWatcher
var SDL_EventSpecific SDL_EventWatcher
var SDL_event_watchers []*SDL_EventWatcher
var SDL_event_watchers_count int = 0

//export callEventFilter
func callEventFilter(_ unsafe.Pointer, cEvent *C.SDL_Event) cInt {
	var ret int32
	var event *SDL_Event

	if SDL_EventOK.callback != nil {
		event = (*SDL_Event)(unsafe.Pointer(cEvent))
		ret = SDL_EventOK.callback(SDL_EventOK.userdata, event)
	}
	return cInt(ret)
}

func SDL_SetEventFilter(filter SDL_EventFilterCallback, userdata any) {
	SDL_event_watchers_lock.Lock()
	defer SDL_event_watchers_lock.Unlock()

	SDL_EventOK.callback = filter
	SDL_EventOK.userdata = userdata

	C.SDL_SetEventFilter(C.SDL_EventFilterCallback(C.callEventFilter), nil)
}

func SDL_GetEventFilter(filter *SDL_EventFilterCallback, userdata *any) bool {
	var event_ok SDL_EventWatcher

	SDL_event_watchers_lock.Lock()
	defer SDL_event_watchers_lock.Unlock()

	event_ok = SDL_EventOK
	*filter = event_ok.callback
	*userdata = event_ok.userdata
	if event_ok.callback != nil {
		return true
	}
	return false
}

func SDL_event_watchers_delete(watchers []*SDL_EventWatcher, watcher *SDL_EventWatcher) []*SDL_EventWatcher {
	retWatchers := watchers[:0]
	for _, v := range watchers {
		if v != watcher {
			retWatchers = append(retWatchers, v)
		} else {
			SDL_DestroyDataStructures(SDL_GetMemoryPool(), v)
		}
	}
	return retWatchers
}

//export callEventWatch
func callEventWatch(_ unsafe.Pointer, cEvent *C.SDL_Event) cInt {
	var ret int32
	var event *SDL_Event

	for i := range SDL_event_watchers {
		if !SDL_event_watchers[i].removed {
			event = (*SDL_Event)(unsafe.Pointer(cEvent))
			ret = SDL_event_watchers[i].callback(SDL_event_watchers[i].userdata, event)
		}
	}
	return cInt(ret)
}

func SDL_AddEventWatch(filter SDL_EventFilterCallback, userdata any) {
	SDL_event_watchers_lock.Lock()
	defer SDL_event_watchers_lock.Unlock()

	var event_watchers *SDL_EventWatcher
	event_watchers = SDL_CreateDataStructures(SDL_GetMemoryPool(), SDL_EventWatcher{})
	event_watchers.callback = filter
	event_watchers.userdata = userdata
	event_watchers.removed = false

	if len(SDL_event_watchers) == 0 {
		C.SDL_AddEventWatch(C.SDL_EventFilterCallback(C.callEventWatch), nil)
	}
	SDL_event_watchers = append(SDL_event_watchers, event_watchers)
}

func SDL_DelEventWatch(filter SDL_EventFilterCallback, userdata any) {
	SDL_event_watchers_lock.Lock()
	defer SDL_event_watchers_lock.Unlock()

	for i := range SDL_event_watchers {
		callback := SDL_event_watchers[i].callback
		p1 := reflect.ValueOf(callback).Pointer()
		p2 := reflect.ValueOf(filter).Pointer()

		if p1 == p2 && SDL_event_watchers[i].userdata == userdata {
			SDL_event_watchers[i].removed = true
			SDL_event_watchers = SDL_event_watchers_delete(SDL_event_watchers, SDL_event_watchers[i])
			break
		}
	}

	if len(SDL_event_watchers) == 0 {
		C.SDL_DelEventWatch(C.SDL_EventFilterCallback(C.callEventWatch), nil)
	}
}

//export callFilterEvents
func callFilterEvents(_ unsafe.Pointer, cEvent *C.SDL_Event) cInt {
	var ret int32
	var event *SDL_Event

	if SDL_EventSpecific.callback != nil {
		event = (*SDL_Event)(unsafe.Pointer(cEvent))
		ret = SDL_EventSpecific.callback(SDL_EventSpecific.userdata, event)
	}
	return cInt(ret)
}

func SDL_FilterEvents(filter SDL_EventFilterCallback, userdata any) {
	SDL_event_watchers_lock.Lock()
	defer SDL_event_watchers_lock.Unlock()

	SDL_EventSpecific.callback = filter
	SDL_EventSpecific.userdata = userdata

	C.SDL_FilterEvents(C.SDL_EventFilterCallback(C.callFilterEvents), nil)
	// clear
	SDL_EventSpecific = SDL_EventWatcher{}
}

func SDL_EventState(eventType SDL_EventType, state int) uint8 {
	cRet := C.SDL_EventState(cUint(eventType), cInt(state))
	return uint8(cRet)
}

func SDL_RegisterEvents(numevents int) uint32 {
	cRet := C.SDL_RegisterEvents(cInt(numevents))
	return uint32(cRet)
}
