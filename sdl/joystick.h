#ifndef _GO_SDL_JOYSTICK_H // if {
#define _GO_SDL_JOYSTICK_H

#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2, 24, 0)) // if {

#if defined(WARN_OUTDATED) // if {
#pragma message("SDL_JoystickPathForIndex is not supported before SDL 2.24.0")
#pragma message("SDL_JoystickPath is not supported before SDL 2.24.0")
#pragma message("SDL_JoystickGetFirmwareVersion is not supported before SDL 2.24.0")
#pragma message("SDL_JoystickAttachVirtualEx is not supported before SDL 2.24.0")
#endif // }

typedef struct SDL_VirtualJoystickDesc {
    Uint16 version;     // `SDL_VIRTUAL_JOYSTICK_DESC_VERSION`
    Uint16 type;        // `SDL_JoystickType`
    Uint16 naxes;       // the number of axes on this joystick
    Uint16 nbuttons;    // the number of buttons on this joystick
    Uint16 nhats;       // the number of hats on this joystick
    Uint16 vendor_id;   // the USB vendor ID of this joystick
    Uint16 product_id;  // the USB product ID of this joystick
    Uint16 padding;     // unused
    Uint32 button_mask; // A mask of which buttons are valid for this controller e.g. (1 << SDL_CONTROLLER_BUTTON_A)
    Uint32 axis_mask;   // A mask of which axes are valid for this controller e.g. (1 << SDL_CONTROLLER_AXIS_LEFTX)
    const char *name;   // the name of the joystick

    void *userdata;                                                  // User data pointer passed to callbacks
    void(SDLCALL *Update)(void *userdata);                           // Called when the joystick state should be updated
    void(SDLCALL *SetPlayerIndex)(void *userdata, int player_index); // Called when the player index is set
    int(SDLCALL *Rumble)(void *userdata, Uint16 low_frequency_rumble,
                         Uint16 high_frequency_rumble); // Implements SDL_JoystickRumble()
    int(SDLCALL *RumbleTriggers)(void *userdata, Uint16 left_rumble,
                                 Uint16 right_rumble);                        // Implements SDL_JoystickRumbleTriggers()
    int(SDLCALL *SetLED)(void *userdata, Uint8 red, Uint8 green, Uint8 blue); // Implements SDL_JoystickSetLED()
    int(SDLCALL *SendEffect)(void *userdata, const void *data, int size);     // Implements SDL_JoystickSendEffect()

} SDL_VirtualJoystickDesc;

static const char *SDL_JoystickPathForIndex(int device_index) { return NULL; }

static const char *SDL_JoystickPath(SDL_Joystick *joystick) { return NULL; }

static Uint16 SDL_JoystickGetFirmwareVersion(SDL_Joystick *joystick) { return 0; }

static inline int SDL_JoystickAttachVirtualEx(const SDL_VirtualJoystickDesc *desc) { return -1; }

#endif // }
#endif // }
