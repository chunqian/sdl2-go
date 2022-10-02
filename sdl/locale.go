package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"

type Locale struct {
	Language string // A language name, like "en" for English.
	Country  string // A country, like "US" for America. Can be empty.
}

func SDL_GetPreferredLocales() (locale *Locale) {
	cLocale := C.SDL_GetPreferredLocales()
	if cLocale == nil {
		return nil
	}
	locale.Language = SDL_GoString(cLocale.language)
	locale.Country = SDL_GoString(cLocale.country)
	return
}
