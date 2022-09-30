//go:build !nomidi

package mix

/*
#define SDL_SUPPORTED_MIDI_BACKENDS
#include "mixer_wrapper.h"
*/
import "C"
