package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"

// struct
type SDL_Locale struct {
	Language string // A language name, like "en" for English.
	Country  string // A country, like "US" for America. Can be empty.
}

func SDL_GetPreferredLocales() (locale *SDL_Locale) {
	cLocale := C.SDL_GetPreferredLocales()
	if cLocale == nil {
		return nil
	}
	locale.Language = createGoString(cLocale.language)
	locale.Country = createGoString(cLocale.country)
	return
}
