package sdl

/*
#include "events.h"
*/
import "C"
import (
	"reflect"
	"sync"
	"unsafe"

	. "github.com/chunqian/memory"
)

// go struct to c struct
func (s *SDL_CommonEvent) cSDL_CommonEvent() *C.SDL_CommonEvent {
	return (*C.SDL_CommonEvent)(unsafe.Pointer(s))
}

func (s *SDL_DisplayEvent) cSDL_DisplayEvent() *C.SDL_DisplayEvent {
	return (*C.SDL_DisplayEvent)(unsafe.Pointer(s))
}

func (s *SDL_WindowEvent) cSDL_WindowEvent() *C.SDL_WindowEvent {
	return (*C.SDL_WindowEvent)(unsafe.Pointer(s))
}

func (s *SDL_KeyboardEvent) cSDL_KeyboardEvent() *C.SDL_KeyboardEvent {
	return (*C.SDL_KeyboardEvent)(unsafe.Pointer(s))
}

func (s *SDL_TextEditingEvent) cSDL_TextEditingEvent() *C.SDL_TextEditingEvent {
	return (*C.SDL_TextEditingEvent)(unsafe.Pointer(s))
}

func (s *SDL_TextInputEvent) cSDL_TextInputEvent() *C.SDL_TextInputEvent {
	return (*C.SDL_TextInputEvent)(unsafe.Pointer(s))
}

func (s *SDL_MouseMotionEvent) cSDL_MouseMotionEvent() *C.SDL_MouseMotionEvent {
	return (*C.SDL_MouseMotionEvent)(unsafe.Pointer(s))
}

func (s *SDL_MouseButtonEvent) cSDL_MouseButtonEvent() *C.SDL_MouseButtonEvent {
	return (*C.SDL_MouseButtonEvent)(unsafe.Pointer(s))
}

func (s *SDL_MouseWheelEvent) cSDL_MouseWheelEvent() *C.SDL_MouseWheelEvent {
	return (*C.SDL_MouseWheelEvent)(unsafe.Pointer(s))
}

func (s *SDL_JoyAxisEvent) cSDL_JoyAxisEvent() *C.SDL_JoyAxisEvent {
	return (*C.SDL_JoyAxisEvent)(unsafe.Pointer(s))
}

func (s *SDL_JoyBallEvent) cSDL_JoyBallEvent() *C.SDL_JoyBallEvent {
	return (*C.SDL_JoyBallEvent)(unsafe.Pointer(s))
}

func (s *SDL_JoyHatEvent) cSDL_JoyHatEvent() *C.SDL_JoyHatEvent {
	return (*C.SDL_JoyHatEvent)(unsafe.Pointer(s))
}

func (s *SDL_JoyButtonEvent) cSDL_JoyButtonEvent() *C.SDL_JoyButtonEvent {
	return (*C.SDL_JoyButtonEvent)(unsafe.Pointer(s))
}

func (s *SDL_JoyDeviceEvent) cSDL_JoyDeviceEvent() *C.SDL_JoyDeviceEvent {
	return (*C.SDL_JoyDeviceEvent)(unsafe.Pointer(s))
}

func (s *SDL_ControllerAxisEvent) cSDL_ControllerAxisEvent() *C.SDL_ControllerAxisEvent {
	return (*C.SDL_ControllerAxisEvent)(unsafe.Pointer(s))
}

func (s *SDL_ControllerButtonEvent) cSDL_ControllerButtonEvent() *C.SDL_ControllerButtonEvent {
	return (*C.SDL_ControllerButtonEvent)(unsafe.Pointer(s))
}

func (s *SDL_ControllerDeviceEvent) cSDL_ControllerDeviceEvent() *C.SDL_ControllerDeviceEvent {
	return (*C.SDL_ControllerDeviceEvent)(unsafe.Pointer(s))
}

func (s *SDL_ControllerTouchpadEvent) cSDL_ControllerTouchpadEvent() *C.SDL_ControllerTouchpadEvent {
	return (*C.SDL_ControllerTouchpadEvent)(unsafe.Pointer(s))
}

func (s *SDL_ControllerSensorEvent) cSDL_ControllerSensorEvent() *C.SDL_ControllerSensorEvent {
	return (*C.SDL_ControllerSensorEvent)(unsafe.Pointer(s))
}

func (s *SDL_AudioDeviceEvent) cSDL_AudioDeviceEvent() *C.SDL_AudioDeviceEvent {
	return (*C.SDL_AudioDeviceEvent)(unsafe.Pointer(s))
}

func (s *SDL_TouchFingerEvent) cSDL_TouchFingerEvent() *C.SDL_TouchFingerEvent {
	return (*C.SDL_TouchFingerEvent)(unsafe.Pointer(s))
}

func (s *SDL_MultiGestureEvent) cSDL_MultiGestureEvent() *C.SDL_MultiGestureEvent {
	return (*C.SDL_MultiGestureEvent)(unsafe.Pointer(s))
}

func (s *SDL_DollarGestureEvent) cSDL_DollarGestureEvent() *C.SDL_DollarGestureEvent {
	return (*C.SDL_DollarGestureEvent)(unsafe.Pointer(s))
}

func (s *SDL_DropEvent) cSDL_DropEvent() *C.SDL_DropEvent {
	return (*C.SDL_DropEvent)(unsafe.Pointer(s))
}

func (s *SDL_SensorEvent) cSDL_SensorEvent() *C.SDL_SensorEvent {
	return (*C.SDL_SensorEvent)(unsafe.Pointer(s))
}

func (s *SDL_QuitEvent) cSDL_QuitEvent() *C.SDL_QuitEvent {
	return (*C.SDL_QuitEvent)(unsafe.Pointer(s))
}

func (s *SDL_OSEvent) cSDL_OSEvent() *C.SDL_OSEvent {
	return (*C.SDL_OSEvent)(unsafe.Pointer(s))
}

func (s *SDL_UserEvent) cSDL_UserEvent() *C.SDL_UserEvent {
	return (*C.SDL_UserEvent)(unsafe.Pointer(s))
}

func (s *SDL_SysWMEvent) cSDL_SysWMEvent() *C.SDL_SysWMEvent {
	return (*C.SDL_SysWMEvent)(unsafe.Pointer(s))
}

func (s *SDL_Event) cSDL_Event() *C.SDL_Event {
	return (*C.SDL_Event)(unsafe.Pointer(s))
}

// go function wapper for c function
// Pump the event loop, gathering events from the input devices.
func SDL_PumpEvents() {
	C.SDL_PumpEvents()
}

// Check the event queue for messages and optionally return them.
func SDL_PeepEvents(events []SDL_Event, numevents int, action SDL_eventaction, minType SDL_EventType, maxType SDL_EventType) int {
	if events == nil {
		return 0
	}

	cRet := C.SDL_PeepEvents((&events[0]).cSDL_Event(), cInt(numevents), cSDL_eventaction(action), cSDL_EventType(minType), cSDL_EventType(maxType))
	return int(cRet)
}

// Check for the existence of a certain event type in the event queue.
func SDL_HasEvent(eventType SDL_EventType) bool {
	cRet := C.SDL_HasEvent(cSDL_EventType(eventType))
	return cRet != 0
}

// Check for the existence of certain event types in the event queue.
func SDL_HasEvents(minType SDL_EventType, maxType SDL_EventType) bool {
	cRet := C.SDL_HasEvents(cSDL_EventType(minType), cSDL_EventType(maxType))
	return cRet != 0
}

// Clear events of a specific type from the event queue.
func SDL_FlushEvent(eventType SDL_EventType) {
	C.SDL_FlushEvent(cSDL_EventType(eventType))
}

// Clear events of a range of types from the event queue.
func SDL_FlushEvents(minType SDL_EventType, maxType SDL_EventType) {
	C.SDL_FlushEvents(cSDL_EventType(minType), cSDL_EventType(maxType))
}

// Poll for currently pending events.
func SDL_PollEvent(event *SDL_Event) bool {
	cRet := C.SDL_PollEvent(event.cSDL_Event())
	return cRet != 0
}

// Wait indefinitely for the next available event.
func SDL_WaitEvent(event *SDL_Event) int {
	cRet := C.SDL_WaitEvent(event.cSDL_Event())
	return int(cRet)
}

// Wait until the specified timeout (in milliseconds) for the next available event.
func SDL_WaitEventTimeout(event *SDL_Event, timeout int) int {
	cRet := C.SDL_WaitEventTimeout(event.cSDL_Event(), cInt(timeout))
	return int(cRet)
}

// Add an event to the event queue.
func SDL_PushEvent(event *SDL_Event) {
	C.SDL_PushEvent(event.cSDL_Event())
}

type SDL_EventFilter = func(interface{}, *SDL_Event) int32

type SDL_EventWatcher struct {
	callback SDL_EventFilter
	userdata interface{}
	removed  bool
}

var SDL_event_watchers_lock sync.Mutex
var SDL_EventOK SDL_EventWatcher
var SDL_EventSpecific SDL_EventWatcher
var SDL_event_watchers []*SDL_EventWatcher
var SDL_event_watchers_count int = 0

//export SDL_EventFilterWrapper
func SDL_EventFilterWrapper(_ unsafe.Pointer, cEvent *C.SDL_Event) cInt {
	var ret int32
	var event *SDL_Event

	if SDL_EventOK.callback != nil {
		event = (*SDL_Event)(unsafe.Pointer(cEvent))
		ret = SDL_EventOK.callback(SDL_EventOK.userdata, event)
	}
	return cInt(ret)
}

// Set up a filter to process all events before they change internal state and are posted to the internal event queue.
func SDL_SetEventFilter(filter SDL_EventFilter, userdata interface{}) {
	SDL_event_watchers_lock.Lock()
	defer SDL_event_watchers_lock.Unlock()

	SDL_EventOK.callback = filter
	SDL_EventOK.userdata = userdata

	C.SDL_SetEventFilter(C.SDL_EventFilter(C.SDL_EventFilterWrapper), nil)
}

// Query the current event filter.
func SDL_GetEventFilter(filter *SDL_EventFilter, userdata *interface{}) bool {
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
			MP_Free(SDL_GetMemoryPool(), unsafe.Pointer(v))
		}
	}
	return retWatchers
}

//export SDL_EventWatchWrapper
func SDL_EventWatchWrapper(_ unsafe.Pointer, cEvent *C.SDL_Event) cInt {
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

// Add a callback to be triggered when an event is added to the event queue.
func SDL_AddEventWatch(filter SDL_EventFilter, userdata interface{}) {
	SDL_event_watchers_lock.Lock()
	defer SDL_event_watchers_lock.Unlock()

	var event_watchers *SDL_EventWatcher
	event_watchers = (*SDL_EventWatcher)(MP_Malloc(SDL_GetMemoryPool(), uint32(unsafe.Sizeof(*event_watchers))))
	event_watchers.callback = filter
	event_watchers.userdata = userdata
	event_watchers.removed = false

	if len(SDL_event_watchers) == 0 {
		C.SDL_AddEventWatch(C.SDL_EventFilter(C.SDL_EventWatchWrapper), nil)
	}
	SDL_event_watchers = append(SDL_event_watchers, event_watchers)
}

// Remove an event watch callback added with SDL_AddEventWatch().
func SDL_DelEventWatch(filter SDL_EventFilter, userdata interface{}) {
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
		C.SDL_DelEventWatch(C.SDL_EventFilter(C.SDL_EventWatchWrapper), nil)
	}
}

//export SDL_FilterEventsWrapper
func SDL_FilterEventsWrapper(_ unsafe.Pointer, cEvent *C.SDL_Event) cInt {
	var ret int32
	var event *SDL_Event

	if SDL_EventSpecific.callback != nil {
		event = (*SDL_Event)(unsafe.Pointer(cEvent))
		ret = SDL_EventSpecific.callback(SDL_EventSpecific.userdata, event)
	}
	return cInt(ret)
}

// Run a specific filter function on the current event queue, removing any events for which the filter returns 0.
func SDL_FilterEvents(filter SDL_EventFilter, userdata interface{}) {
	SDL_event_watchers_lock.Lock()
	defer SDL_event_watchers_lock.Unlock()

	SDL_EventSpecific.callback = filter
	SDL_EventSpecific.userdata = userdata

	C.SDL_FilterEvents(C.SDL_EventFilter(C.SDL_FilterEventsWrapper), nil)
	// clear
	SDL_EventSpecific = SDL_EventWatcher{}
}

// Set the state of processing events by type.
func SDL_EventState(eventType SDL_EventType, state int) uint8 {
	cRet := C.SDL_EventState(cSDL_EventType(eventType), cInt(state))
	return uint8(cRet)
}

// Allocate a set of user-defined events, and return the beginning event number for that set of events.
func SDL_RegisterEvents(numevents int) uint32 {
	cRet := C.SDL_RegisterEvents(cInt(numevents))
	return uint32(cRet)
}
