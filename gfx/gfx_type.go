package sdl

/*
#include "gfx_wrapper.h"
*/
import "C"

// c #define
const (
	FPS_UPPER_LIMIT = 200
	FPS_LOWER_LIMIT = 1
	FPS_DEFAULT = 30

	SMOOTHING_OFF = 0
	SMOOTHING_ON = 1
)

type FPSmanager struct {
	Framecount uint32
	Rateticks  float32
	Baseticks  uint32
	Lastticks  uint32
	Rate       uint32
}
