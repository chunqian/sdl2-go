#ifndef _GO_SDL_HINTS_H // if {
#define _GO_SDL_HINTS_H

#include "sdl_wrapper.h"

void goHintCallback(char *name, char *oldValue, char *newValue);

static void hintCallback(void *userdata, const char *name, const char *oldValue, const char *newValue) {
    goHintCallback((char *)name, (char *)oldValue, (char *)newValue);
}

static void addHintCallback(const char *name) { SDL_AddHintCallback(name, hintCallback, NULL); }

static void delHintCallback(const char *name) { SDL_DelHintCallback(name, hintCallback, NULL); }

#if !SDL_VERSION_ATLEAST(2, 0, 22) // if {
#define SDL_HINT_IME_SUPPORT_EXTENDED_TEXT ""
#define SDL_HINT_MOUSE_RELATIVE_MODE_CENTER ""
#define SDL_HINT_MOUSE_AUTO_CAPTURE ""
#define SDL_HINT_VIDEO_FOREIGN_WINDOW_OPENGL ""
#define SDL_HINT_VIDEO_FOREIGN_WINDOW_VULKAN ""
#define SDL_HINT_QUIT_ON_LAST_WINDOW_CLOSE ""
#define SDL_HINT_JOYSTICK_ROG_CHAKRAM ""
#define SDL_HINT_X11_WINDOW_TYPE ""
#define SDL_HINT_VIDEO_WAYLAND_PREFER_LIBDECOR ""

#define SDL_HINT_JOYSTICK_HIDAPI_JOY_CONS ""
#define SDL_HINT_JOYSTICK_HIDAPI_SWITCH_HOME_LED ""
#endif // }

#if !SDL_VERSION_ATLEAST(2, 24, 0) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_ResetHint is not supported before SDL 2.24.0")
#endif // }

static inline SDL_bool SDL_ResetHint(const char *name) { return SDL_FALSE; }

#define SDL_HINT_MOUSE_RELATIVE_WARP_MOTION ""
#define SDL_HINT_TRACKPAD_IS_TOUCH_ONLY ""
// #define SDL_HINT_JOYSTICK_HIDAPI_JOY_CONS ""
#define SDL_HINT_JOYSTICK_HIDAPI_COMBINE_JOY_CONS ""
// #define SDL_HINT_JOYSTICK_HIDAPI_SWITCH_HOME_LED ""
#define SDL_HINT_JOYSTICK_HIDAPI_JOYCON_HOME_LED ""
#define SDL_HINT_JOYSTICK_HIDAPI_SWITCH_PLAYER_LED ""
#define SDL_HINT_JOYSTICK_HIDAPI_NINTENDO_CLASSIC ""
#define SDL_HINT_JOYSTICK_HIDAPI_SHIELD ""
#define SDL_HINT_WINDOWS_DPI_AWARENESS ""
#define SDL_HINT_WINDOWS_DPI_SCALING ""
#define SDL_HINT_DIRECTINPUT_ENABLED ""
#define SDL_HINT_VIDEO_WAYLAND_MODE_EMULATION ""
#define SDL_HINT_KMSDRM_DEVICE_INDEX ""
#define SDL_HINT_LINUX_DIGITAL_HATS ""
#define SDL_HINT_LINUX_HAT_DEADZONES ""
#define SDL_HINT_MAC_OPENGL_ASYNC_DISPATCH ""

#endif // }
#endif // }
