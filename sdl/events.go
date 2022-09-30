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

func cSDL_Event(s *SDL_Event) *C.SDL_Event {
	return (*C.SDL_Event)(unsafe.Pointer(s))
}

func SDL_PumpEvents() {
	C.SDL_PumpEvents()
}

func SDL_PeepEvents(events []SDL_Event, numevents int, action SDL_eventaction, minType SDL_EventType, maxType SDL_EventType) int {
	if events == nil {
		return 0
	}

	cRet := C.SDL_PeepEvents(cSDL_Event(&events[0]), cInt(numevents), cSDL_eventaction(action), cUint(minType), cUint(maxType))
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
	cRet := C.SDL_PollEvent(cSDL_Event(event))
	return cRet != 0
}

func SDL_WaitEvent(event *SDL_Event) int {
	cRet := C.SDL_WaitEvent(cSDL_Event(event))
	return int(cRet)
}

func SDL_WaitEventTimeout(event *SDL_Event, timeout int) int {
	cRet := C.SDL_WaitEventTimeout(cSDL_Event(event), cInt(timeout))
	return int(cRet)
}

func SDL_PushEvent(event *SDL_Event) {
	C.SDL_PushEvent(cSDL_Event(event))
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
