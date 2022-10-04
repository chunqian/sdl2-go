package sdl

// define
var (
	SDLK_SCANCODE_MASK      = (1 << 30)
	SDL_SCANCODE_TO_KEYCODE = func(x int32) int32 { return x | int32(SDLK_SCANCODE_MASK) }
)

// typedef
type SDL_Scancode = int32
type SDL_KeyCode = int32
type SDL_Keycode = int32
type SDL_Keymod = uint32

// enum
const (
	SDL_SCANCODE_UNKNOWN            SDL_Scancode = 0
	SDL_SCANCODE_A                  SDL_Scancode = 4
	SDL_SCANCODE_B                  SDL_Scancode = 5
	SDL_SCANCODE_C                  SDL_Scancode = 6
	SDL_SCANCODE_D                  SDL_Scancode = 7
	SDL_SCANCODE_E                  SDL_Scancode = 8
	SDL_SCANCODE_F                  SDL_Scancode = 9
	SDL_SCANCODE_G                  SDL_Scancode = 10
	SDL_SCANCODE_H                  SDL_Scancode = 11
	SDL_SCANCODE_I                  SDL_Scancode = 12
	SDL_SCANCODE_J                  SDL_Scancode = 13
	SDL_SCANCODE_K                  SDL_Scancode = 14
	SDL_SCANCODE_L                  SDL_Scancode = 15
	SDL_SCANCODE_M                  SDL_Scancode = 16
	SDL_SCANCODE_N                  SDL_Scancode = 17
	SDL_SCANCODE_O                  SDL_Scancode = 18
	SDL_SCANCODE_P                  SDL_Scancode = 19
	SDL_SCANCODE_Q                  SDL_Scancode = 20
	SDL_SCANCODE_R                  SDL_Scancode = 21
	SDL_SCANCODE_S                  SDL_Scancode = 22
	SDL_SCANCODE_T                  SDL_Scancode = 23
	SDL_SCANCODE_U                  SDL_Scancode = 24
	SDL_SCANCODE_V                  SDL_Scancode = 25
	SDL_SCANCODE_W                  SDL_Scancode = 26
	SDL_SCANCODE_X                  SDL_Scancode = 27
	SDL_SCANCODE_Y                  SDL_Scancode = 28
	SDL_SCANCODE_Z                  SDL_Scancode = 29
	SDL_SCANCODE_1                  SDL_Scancode = 30
	SDL_SCANCODE_2                  SDL_Scancode = 31
	SDL_SCANCODE_3                  SDL_Scancode = 32
	SDL_SCANCODE_4                  SDL_Scancode = 33
	SDL_SCANCODE_5                  SDL_Scancode = 34
	SDL_SCANCODE_6                  SDL_Scancode = 35
	SDL_SCANCODE_7                  SDL_Scancode = 36
	SDL_SCANCODE_8                  SDL_Scancode = 37
	SDL_SCANCODE_9                  SDL_Scancode = 38
	SDL_SCANCODE_0                  SDL_Scancode = 39
	SDL_SCANCODE_RETURN             SDL_Scancode = 40
	SDL_SCANCODE_ESCAPE             SDL_Scancode = 41
	SDL_SCANCODE_BACKSPACE          SDL_Scancode = 42
	SDL_SCANCODE_TAB                SDL_Scancode = 43
	SDL_SCANCODE_SPACE              SDL_Scancode = 44
	SDL_SCANCODE_MINUS              SDL_Scancode = 45
	SDL_SCANCODE_EQUALS             SDL_Scancode = 46
	SDL_SCANCODE_LEFTBRACKET        SDL_Scancode = 47
	SDL_SCANCODE_RIGHTBRACKET       SDL_Scancode = 48
	SDL_SCANCODE_BACKSLASH          SDL_Scancode = 49
	SDL_SCANCODE_NONUSHASH          SDL_Scancode = 50
	SDL_SCANCODE_SEMICOLON          SDL_Scancode = 51
	SDL_SCANCODE_APOSTROPHE         SDL_Scancode = 52
	SDL_SCANCODE_GRAVE              SDL_Scancode = 53
	SDL_SCANCODE_COMMA              SDL_Scancode = 54
	SDL_SCANCODE_PERIOD             SDL_Scancode = 55
	SDL_SCANCODE_SLASH              SDL_Scancode = 56
	SDL_SCANCODE_CAPSLOCK           SDL_Scancode = 57
	SDL_SCANCODE_F1                 SDL_Scancode = 58
	SDL_SCANCODE_F2                 SDL_Scancode = 59
	SDL_SCANCODE_F3                 SDL_Scancode = 60
	SDL_SCANCODE_F4                 SDL_Scancode = 61
	SDL_SCANCODE_F5                 SDL_Scancode = 62
	SDL_SCANCODE_F6                 SDL_Scancode = 63
	SDL_SCANCODE_F7                 SDL_Scancode = 64
	SDL_SCANCODE_F8                 SDL_Scancode = 65
	SDL_SCANCODE_F9                 SDL_Scancode = 66
	SDL_SCANCODE_F10                SDL_Scancode = 67
	SDL_SCANCODE_F11                SDL_Scancode = 68
	SDL_SCANCODE_F12                SDL_Scancode = 69
	SDL_SCANCODE_PRINTSCREEN        SDL_Scancode = 70
	SDL_SCANCODE_SCROLLLOCK         SDL_Scancode = 71
	SDL_SCANCODE_PAUSE              SDL_Scancode = 72
	SDL_SCANCODE_INSERT             SDL_Scancode = 73
	SDL_SCANCODE_HOME               SDL_Scancode = 74
	SDL_SCANCODE_PAGEUP             SDL_Scancode = 75
	SDL_SCANCODE_DELETE             SDL_Scancode = 76
	SDL_SCANCODE_END                SDL_Scancode = 77
	SDL_SCANCODE_PAGEDOWN           SDL_Scancode = 78
	SDL_SCANCODE_RIGHT              SDL_Scancode = 79
	SDL_SCANCODE_LEFT               SDL_Scancode = 80
	SDL_SCANCODE_DOWN               SDL_Scancode = 81
	SDL_SCANCODE_UP                 SDL_Scancode = 82
	SDL_SCANCODE_NUMLOCKCLEAR       SDL_Scancode = 83
	SDL_SCANCODE_KP_DIVIDE          SDL_Scancode = 84
	SDL_SCANCODE_KP_MULTIPLY        SDL_Scancode = 85
	SDL_SCANCODE_KP_MINUS           SDL_Scancode = 86
	SDL_SCANCODE_KP_PLUS            SDL_Scancode = 87
	SDL_SCANCODE_KP_ENTER           SDL_Scancode = 88
	SDL_SCANCODE_KP_1               SDL_Scancode = 89
	SDL_SCANCODE_KP_2               SDL_Scancode = 90
	SDL_SCANCODE_KP_3               SDL_Scancode = 91
	SDL_SCANCODE_KP_4               SDL_Scancode = 92
	SDL_SCANCODE_KP_5               SDL_Scancode = 93
	SDL_SCANCODE_KP_6               SDL_Scancode = 94
	SDL_SCANCODE_KP_7               SDL_Scancode = 95
	SDL_SCANCODE_KP_8               SDL_Scancode = 96
	SDL_SCANCODE_KP_9               SDL_Scancode = 97
	SDL_SCANCODE_KP_0               SDL_Scancode = 98
	SDL_SCANCODE_KP_PERIOD          SDL_Scancode = 99
	SDL_SCANCODE_NONUSBACKSLASH     SDL_Scancode = 100
	SDL_SCANCODE_APPLICATION        SDL_Scancode = 101
	SDL_SCANCODE_POWER              SDL_Scancode = 102
	SDL_SCANCODE_KP_EQUALS          SDL_Scancode = 103
	SDL_SCANCODE_F13                SDL_Scancode = 104
	SDL_SCANCODE_F14                SDL_Scancode = 105
	SDL_SCANCODE_F15                SDL_Scancode = 106
	SDL_SCANCODE_F16                SDL_Scancode = 107
	SDL_SCANCODE_F17                SDL_Scancode = 108
	SDL_SCANCODE_F18                SDL_Scancode = 109
	SDL_SCANCODE_F19                SDL_Scancode = 110
	SDL_SCANCODE_F20                SDL_Scancode = 111
	SDL_SCANCODE_F21                SDL_Scancode = 112
	SDL_SCANCODE_F22                SDL_Scancode = 113
	SDL_SCANCODE_F23                SDL_Scancode = 114
	SDL_SCANCODE_F24                SDL_Scancode = 115
	SDL_SCANCODE_EXECUTE            SDL_Scancode = 116
	SDL_SCANCODE_HELP               SDL_Scancode = 117
	SDL_SCANCODE_MENU               SDL_Scancode = 118
	SDL_SCANCODE_SELECT             SDL_Scancode = 119
	SDL_SCANCODE_STOP               SDL_Scancode = 120
	SDL_SCANCODE_AGAIN              SDL_Scancode = 121
	SDL_SCANCODE_UNDO               SDL_Scancode = 122
	SDL_SCANCODE_CUT                SDL_Scancode = 123
	SDL_SCANCODE_COPY               SDL_Scancode = 124
	SDL_SCANCODE_PASTE              SDL_Scancode = 125
	SDL_SCANCODE_FIND               SDL_Scancode = 126
	SDL_SCANCODE_MUTE               SDL_Scancode = 127
	SDL_SCANCODE_VOLUMEUP           SDL_Scancode = 128
	SDL_SCANCODE_VOLUMEDOWN         SDL_Scancode = 129
	SDL_SCANCODE_KP_COMMA           SDL_Scancode = 133
	SDL_SCANCODE_KP_EQUALSAS400     SDL_Scancode = 134
	SDL_SCANCODE_INTERNATIONAL1     SDL_Scancode = 135
	SDL_SCANCODE_INTERNATIONAL2     SDL_Scancode = 136
	SDL_SCANCODE_INTERNATIONAL3     SDL_Scancode = 137
	SDL_SCANCODE_INTERNATIONAL4     SDL_Scancode = 138
	SDL_SCANCODE_INTERNATIONAL5     SDL_Scancode = 139
	SDL_SCANCODE_INTERNATIONAL6     SDL_Scancode = 140
	SDL_SCANCODE_INTERNATIONAL7     SDL_Scancode = 141
	SDL_SCANCODE_INTERNATIONAL8     SDL_Scancode = 142
	SDL_SCANCODE_INTERNATIONAL9     SDL_Scancode = 143
	SDL_SCANCODE_LANG1              SDL_Scancode = 144
	SDL_SCANCODE_LANG2              SDL_Scancode = 145
	SDL_SCANCODE_LANG3              SDL_Scancode = 146
	SDL_SCANCODE_LANG4              SDL_Scancode = 147
	SDL_SCANCODE_LANG5              SDL_Scancode = 148
	SDL_SCANCODE_LANG6              SDL_Scancode = 149
	SDL_SCANCODE_LANG7              SDL_Scancode = 150
	SDL_SCANCODE_LANG8              SDL_Scancode = 151
	SDL_SCANCODE_LANG9              SDL_Scancode = 152
	SDL_SCANCODE_ALTERASE           SDL_Scancode = 153
	SDL_SCANCODE_SYSREQ             SDL_Scancode = 154
	SDL_SCANCODE_CANCEL             SDL_Scancode = 155
	SDL_SCANCODE_CLEAR              SDL_Scancode = 156
	SDL_SCANCODE_PRIOR              SDL_Scancode = 157
	SDL_SCANCODE_RETURN2            SDL_Scancode = 158
	SDL_SCANCODE_SEPARATOR          SDL_Scancode = 159
	SDL_SCANCODE_OUT                SDL_Scancode = 160
	SDL_SCANCODE_OPER               SDL_Scancode = 161
	SDL_SCANCODE_CLEARAGAIN         SDL_Scancode = 162
	SDL_SCANCODE_CRSEL              SDL_Scancode = 163
	SDL_SCANCODE_EXSEL              SDL_Scancode = 164
	SDL_SCANCODE_KP_00              SDL_Scancode = 176
	SDL_SCANCODE_KP_000             SDL_Scancode = 177
	SDL_SCANCODE_THOUSANDSSEPARATOR SDL_Scancode = 178
	SDL_SCANCODE_DECIMALSEPARATOR   SDL_Scancode = 179
	SDL_SCANCODE_CURRENCYUNIT       SDL_Scancode = 180
	SDL_SCANCODE_CURRENCYSUBUNIT    SDL_Scancode = 181
	SDL_SCANCODE_KP_LEFTPAREN       SDL_Scancode = 182
	SDL_SCANCODE_KP_RIGHTPAREN      SDL_Scancode = 183
	SDL_SCANCODE_KP_LEFTBRACE       SDL_Scancode = 184
	SDL_SCANCODE_KP_RIGHTBRACE      SDL_Scancode = 185
	SDL_SCANCODE_KP_TAB             SDL_Scancode = 186
	SDL_SCANCODE_KP_BACKSPACE       SDL_Scancode = 187
	SDL_SCANCODE_KP_A               SDL_Scancode = 188
	SDL_SCANCODE_KP_B               SDL_Scancode = 189
	SDL_SCANCODE_KP_C               SDL_Scancode = 190
	SDL_SCANCODE_KP_D               SDL_Scancode = 191
	SDL_SCANCODE_KP_E               SDL_Scancode = 192
	SDL_SCANCODE_KP_F               SDL_Scancode = 193
	SDL_SCANCODE_KP_XOR             SDL_Scancode = 194
	SDL_SCANCODE_KP_POWER           SDL_Scancode = 195
	SDL_SCANCODE_KP_PERCENT         SDL_Scancode = 196
	SDL_SCANCODE_KP_LESS            SDL_Scancode = 197
	SDL_SCANCODE_KP_GREATER         SDL_Scancode = 198
	SDL_SCANCODE_KP_AMPERSAND       SDL_Scancode = 199
	SDL_SCANCODE_KP_DBLAMPERSAND    SDL_Scancode = 200
	SDL_SCANCODE_KP_VERTICALBAR     SDL_Scancode = 201
	SDL_SCANCODE_KP_DBLVERTICALBAR  SDL_Scancode = 202
	SDL_SCANCODE_KP_COLON           SDL_Scancode = 203
	SDL_SCANCODE_KP_HASH            SDL_Scancode = 204
	SDL_SCANCODE_KP_SPACE           SDL_Scancode = 205
	SDL_SCANCODE_KP_AT              SDL_Scancode = 206
	SDL_SCANCODE_KP_EXCLAM          SDL_Scancode = 207
	SDL_SCANCODE_KP_MEMSTORE        SDL_Scancode = 208
	SDL_SCANCODE_KP_MEMRECALL       SDL_Scancode = 209
	SDL_SCANCODE_KP_MEMCLEAR        SDL_Scancode = 210
	SDL_SCANCODE_KP_MEMADD          SDL_Scancode = 211
	SDL_SCANCODE_KP_MEMSUBTRACT     SDL_Scancode = 212
	SDL_SCANCODE_KP_MEMMULTIPLY     SDL_Scancode = 213
	SDL_SCANCODE_KP_MEMDIVIDE       SDL_Scancode = 214
	SDL_SCANCODE_KP_PLUSMINUS       SDL_Scancode = 215
	SDL_SCANCODE_KP_CLEAR           SDL_Scancode = 216
	SDL_SCANCODE_KP_CLEARENTRY      SDL_Scancode = 217
	SDL_SCANCODE_KP_BINARY          SDL_Scancode = 218
	SDL_SCANCODE_KP_OCTAL           SDL_Scancode = 219
	SDL_SCANCODE_KP_DECIMAL         SDL_Scancode = 220
	SDL_SCANCODE_KP_HEXADECIMAL     SDL_Scancode = 221
	SDL_SCANCODE_LCTRL              SDL_Scancode = 224
	SDL_SCANCODE_LSHIFT             SDL_Scancode = 225
	SDL_SCANCODE_LALT               SDL_Scancode = 226
	SDL_SCANCODE_LGUI               SDL_Scancode = 227
	SDL_SCANCODE_RCTRL              SDL_Scancode = 228
	SDL_SCANCODE_RSHIFT             SDL_Scancode = 229
	SDL_SCANCODE_RALT               SDL_Scancode = 230
	SDL_SCANCODE_RGUI               SDL_Scancode = 231
	SDL_SCANCODE_MODE               SDL_Scancode = 257
	SDL_SCANCODE_AUDIONEXT          SDL_Scancode = 258
	SDL_SCANCODE_AUDIOPREV          SDL_Scancode = 259
	SDL_SCANCODE_AUDIOSTOP          SDL_Scancode = 260
	SDL_SCANCODE_AUDIOPLAY          SDL_Scancode = 261
	SDL_SCANCODE_AUDIOMUTE          SDL_Scancode = 262
	SDL_SCANCODE_MEDIASELECT        SDL_Scancode = 263
	SDL_SCANCODE_WWW                SDL_Scancode = 264
	SDL_SCANCODE_MAIL               SDL_Scancode = 265
	SDL_SCANCODE_CALCULATOR         SDL_Scancode = 266
	SDL_SCANCODE_COMPUTER           SDL_Scancode = 267
	SDL_SCANCODE_AC_SEARCH          SDL_Scancode = 268
	SDL_SCANCODE_AC_HOME            SDL_Scancode = 269
	SDL_SCANCODE_AC_BACK            SDL_Scancode = 270
	SDL_SCANCODE_AC_FORWARD         SDL_Scancode = 271
	SDL_SCANCODE_AC_STOP            SDL_Scancode = 272
	SDL_SCANCODE_AC_REFRESH         SDL_Scancode = 273
	SDL_SCANCODE_AC_BOOKMARKS       SDL_Scancode = 274
	SDL_SCANCODE_BRIGHTNESSDOWN     SDL_Scancode = 275
	SDL_SCANCODE_BRIGHTNESSUP       SDL_Scancode = 276
	SDL_SCANCODE_DISPLAYSWITCH      SDL_Scancode = 277
	SDL_SCANCODE_KBDILLUMTOGGLE     SDL_Scancode = 278
	SDL_SCANCODE_KBDILLUMDOWN       SDL_Scancode = 279
	SDL_SCANCODE_KBDILLUMUP         SDL_Scancode = 280
	SDL_SCANCODE_EJECT              SDL_Scancode = 281
	SDL_SCANCODE_SLEEP              SDL_Scancode = 282
	SDL_SCANCODE_APP1               SDL_Scancode = 283
	SDL_SCANCODE_APP2               SDL_Scancode = 284
	SDL_SCANCODE_AUDIOREWIND        SDL_Scancode = 285
	SDL_SCANCODE_AUDIOFASTFORWARD   SDL_Scancode = 286
	SDL_SCANCODE_SOFTLEFT           SDL_Scancode = 287
	SDL_SCANCODE_SOFTRIGHT          SDL_Scancode = 288
	SDL_SCANCODE_CALL               SDL_Scancode = 289
	SDL_SCANCODE_ENDCALL            SDL_Scancode = 290
	SDL_NUM_SCANCODES               SDL_Scancode = 512
)

var (
	SDLK_UNKNOWN = 0

	SDLK_RETURN     SDL_KeyCode = '\r'
	SDLK_ESCAPE     SDL_KeyCode = '\x1B'
	SDLK_BACKSPACE  SDL_KeyCode = '\b'
	SDLK_TAB        SDL_KeyCode = '\t'
	SDLK_SPACE      SDL_KeyCode = ' '
	SDLK_EXCLAIM    SDL_KeyCode = '!'
	SDLK_QUOTEDBL   SDL_KeyCode = '"'
	SDLK_HASH       SDL_KeyCode = '#'
	SDLK_PERCENT    SDL_KeyCode = '%'
	SDLK_DOLLAR     SDL_KeyCode = '$'
	SDLK_AMPERSAND  SDL_KeyCode = '&'
	SDLK_QUOTE      SDL_KeyCode = '\''
	SDLK_LEFTPAREN  SDL_KeyCode = '('
	SDLK_RIGHTPAREN SDL_KeyCode = ')'
	SDLK_ASTERISK   SDL_KeyCode = '*'
	SDLK_PLUS       SDL_KeyCode = '+'
	SDLK_COMMA      SDL_KeyCode = ','
	SDLK_MINUS      SDL_KeyCode = '-'
	SDLK_PERIOD     SDL_KeyCode = '.'
	SDLK_SLASH      SDL_KeyCode = '/'
	SDLK_0          SDL_KeyCode = '0'
	SDLK_1          SDL_KeyCode = '1'
	SDLK_2          SDL_KeyCode = '2'
	SDLK_3          SDL_KeyCode = '3'
	SDLK_4          SDL_KeyCode = '4'
	SDLK_5          SDL_KeyCode = '5'
	SDLK_6          SDL_KeyCode = '6'
	SDLK_7          SDL_KeyCode = '7'
	SDLK_8          SDL_KeyCode = '8'
	SDLK_9          SDL_KeyCode = '9'
	SDLK_COLON      SDL_KeyCode = ':'
	SDLK_SEMICOLON  SDL_KeyCode = ';'
	SDLK_LESS       SDL_KeyCode = '<'
	SDLK_EQUALS     SDL_KeyCode = '='
	SDLK_GREATER    SDL_KeyCode = '>'
	SDLK_QUESTION   SDL_KeyCode = '?'
	SDLK_AT         SDL_KeyCode = '@'

	// Skip uppercase letters
	SDLK_LEFTBRACKET  SDL_KeyCode = '['
	SDLK_BACKSLASH    SDL_KeyCode = '\\'
	SDLK_RIGHTBRACKET SDL_KeyCode = ']'
	SDLK_CARET        SDL_KeyCode = '^'
	SDLK_UNDERSCORE   SDL_KeyCode = '_'
	SDLK_BACKQUOTE    SDL_KeyCode = '`'
	SDLK_a            SDL_KeyCode = 'a'
	SDLK_b            SDL_KeyCode = 'b'
	SDLK_c            SDL_KeyCode = 'c'
	SDLK_d            SDL_KeyCode = 'd'
	SDLK_e            SDL_KeyCode = 'e'
	SDLK_f            SDL_KeyCode = 'f'
	SDLK_g            SDL_KeyCode = 'g'
	SDLK_h            SDL_KeyCode = 'h'
	SDLK_i            SDL_KeyCode = 'i'
	SDLK_j            SDL_KeyCode = 'j'
	SDLK_k            SDL_KeyCode = 'k'
	SDLK_l            SDL_KeyCode = 'l'
	SDLK_m            SDL_KeyCode = 'm'
	SDLK_n            SDL_KeyCode = 'n'
	SDLK_o            SDL_KeyCode = 'o'
	SDLK_p            SDL_KeyCode = 'p'
	SDLK_q            SDL_KeyCode = 'q'
	SDLK_r            SDL_KeyCode = 'r'
	SDLK_s            SDL_KeyCode = 's'
	SDLK_t            SDL_KeyCode = 't'
	SDLK_u            SDL_KeyCode = 'u'
	SDLK_v            SDL_KeyCode = 'v'
	SDLK_w            SDL_KeyCode = 'w'
	SDLK_x            SDL_KeyCode = 'x'
	SDLK_y            SDL_KeyCode = 'y'
	SDLK_z            SDL_KeyCode = 'z'

	SDLK_CAPSLOCK SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CAPSLOCK)

	SDLK_F1  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F1)
	SDLK_F2  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F2)
	SDLK_F3  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F3)
	SDLK_F4  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F4)
	SDLK_F5  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F5)
	SDLK_F6  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F6)
	SDLK_F7  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F7)
	SDLK_F8  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F8)
	SDLK_F9  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F9)
	SDLK_F10 SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F10)
	SDLK_F11 SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F11)
	SDLK_F12 SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F12)

	SDLK_PRINTSCREEN SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PRINTSCREEN)
	SDLK_SCROLLLOCK  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SCROLLLOCK)
	SDLK_PAUSE       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAUSE)
	SDLK_INSERT      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_INSERT)
	SDLK_HOME        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_HOME)
	SDLK_PAGEUP      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAGEUP)
	SDLK_DELETE      SDL_KeyCode = '\x7F'
	SDLK_END         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_END)
	SDLK_PAGEDOWN    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAGEDOWN)
	SDLK_RIGHT       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RIGHT)
	SDLK_LEFT        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LEFT)
	SDLK_DOWN        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DOWN)
	SDLK_UP          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_UP)

	SDLK_NUMLOCKCLEAR SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_NUMLOCKCLEAR)
	SDLK_KP_DIVIDE    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DIVIDE)
	SDLK_KP_MULTIPLY  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MULTIPLY)
	SDLK_KP_MINUS     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MINUS)
	SDLK_KP_PLUS      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PLUS)
	SDLK_KP_ENTER     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_ENTER)
	SDLK_KP_1         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_1)
	SDLK_KP_2         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_2)
	SDLK_KP_3         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_3)
	SDLK_KP_4         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_4)
	SDLK_KP_5         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_5)
	SDLK_KP_6         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_6)
	SDLK_KP_7         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_7)
	SDLK_KP_8         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_8)
	SDLK_KP_9         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_9)
	SDLK_KP_0         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_0)
	SDLK_KP_PERIOD    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PERIOD)

	SDLK_APPLICATION    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_APPLICATION)
	SDLK_POWER          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_POWER)
	SDLK_KP_EQUALS      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EQUALS)
	SDLK_F13            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F13)
	SDLK_F14            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F14)
	SDLK_F15            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F15)
	SDLK_F16            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F16)
	SDLK_F17            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F17)
	SDLK_F18            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F18)
	SDLK_F19            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F19)
	SDLK_F20            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F20)
	SDLK_F21            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F21)
	SDLK_F22            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F22)
	SDLK_F23            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F23)
	SDLK_F24            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F24)
	SDLK_EXECUTE        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EXECUTE)
	SDLK_HELP           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_HELP)
	SDLK_MENU           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MENU)
	SDLK_SELECT         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SELECT)
	SDLK_STOP           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_STOP)
	SDLK_AGAIN          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AGAIN)
	SDLK_UNDO           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_UNDO)
	SDLK_CUT            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CUT)
	SDLK_COPY           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_COPY)
	SDLK_PASTE          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PASTE)
	SDLK_FIND           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_FIND)
	SDLK_MUTE           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MUTE)
	SDLK_VOLUMEUP       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_VOLUMEUP)
	SDLK_VOLUMEDOWN     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_VOLUMEDOWN)
	SDLK_KP_COMMA       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_COMMA)
	SDLK_KP_EQUALSAS400 SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EQUALSAS400)

	SDLK_ALTERASE   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_ALTERASE)
	SDLK_SYSREQ     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SYSREQ)
	SDLK_CANCEL     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CANCEL)
	SDLK_CLEAR      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CLEAR)
	SDLK_PRIOR      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PRIOR)
	SDLK_RETURN2    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RETURN2)
	SDLK_SEPARATOR  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SEPARATOR)
	SDLK_OUT        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_OUT)
	SDLK_OPER       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_OPER)
	SDLK_CLEARAGAIN SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CLEARAGAIN)
	SDLK_CRSEL      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CRSEL)
	SDLK_EXSEL      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EXSEL)

	SDLK_KP_00              SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_00)
	SDLK_KP_000             SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_000)
	SDLK_THOUSANDSSEPARATOR SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_THOUSANDSSEPARATOR)
	SDLK_DECIMALSEPARATOR   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DECIMALSEPARATOR)
	SDLK_CURRENCYUNIT       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CURRENCYUNIT)
	SDLK_CURRENCYSUBUNIT    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CURRENCYSUBUNIT)
	SDLK_KP_LEFTPAREN       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LEFTPAREN)
	SDLK_KP_RIGHTPAREN      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_RIGHTPAREN)
	SDLK_KP_LEFTBRACE       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LEFTBRACE)
	SDLK_KP_RIGHTBRACE      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_RIGHTBRACE)
	SDLK_KP_TAB             SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_TAB)
	SDLK_KP_BACKSPACE       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_BACKSPACE)
	SDLK_KP_A               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_A)
	SDLK_KP_B               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_B)
	SDLK_KP_C               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_C)
	SDLK_KP_D               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_D)
	SDLK_KP_E               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_E)
	SDLK_KP_F               SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_F)
	SDLK_KP_XOR             SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_XOR)
	SDLK_KP_POWER           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_POWER)
	SDLK_KP_PERCENT         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PERCENT)
	SDLK_KP_LESS            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LESS)
	SDLK_KP_GREATER         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_GREATER)
	SDLK_KP_AMPERSAND       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_AMPERSAND)
	SDLK_KP_DBLAMPERSAND    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DBLAMPERSAND)
	SDLK_KP_VERTICALBAR     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_VERTICALBAR)
	SDLK_KP_DBLVERTICALBAR  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DBLVERTICALBAR)
	SDLK_KP_COLON           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_COLON)
	SDLK_KP_HASH            SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_HASH)
	SDLK_KP_SPACE           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_SPACE)
	SDLK_KP_AT              SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_AT)
	SDLK_KP_EXCLAM          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EXCLAM)
	SDLK_KP_MEMSTORE        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMSTORE)
	SDLK_KP_MEMRECALL       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMRECALL)
	SDLK_KP_MEMCLEAR        SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMCLEAR)
	SDLK_KP_MEMADD          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMADD)
	SDLK_KP_MEMSUBTRACT     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMSUBTRACT)
	SDLK_KP_MEMMULTIPLY     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMMULTIPLY)
	SDLK_KP_MEMDIVIDE       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMDIVIDE)
	SDLK_KP_PLUSMINUS       SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PLUSMINUS)
	SDLK_KP_CLEAR           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_CLEAR)
	SDLK_KP_CLEARENTRY      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_CLEARENTRY)
	SDLK_KP_BINARY          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_BINARY)
	SDLK_KP_OCTAL           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_OCTAL)
	SDLK_KP_DECIMAL         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DECIMAL)
	SDLK_KP_HEXADECIMAL     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_HEXADECIMAL)

	SDLK_LCTRL  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LCTRL)
	SDLK_LSHIFT SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LSHIFT)
	SDLK_LALT   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LALT)
	SDLK_LGUI   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LGUI)
	SDLK_RCTRL  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RCTRL)
	SDLK_RSHIFT SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RSHIFT)
	SDLK_RALT   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RALT)
	SDLK_RGUI   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RGUI)

	SDLK_MODE SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MODE)

	SDLK_AUDIONEXT    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIONEXT)
	SDLK_AUDIOPREV    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOPREV)
	SDLK_AUDIOSTOP    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOSTOP)
	SDLK_AUDIOPLAY    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOPLAY)
	SDLK_AUDIOMUTE    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOMUTE)
	SDLK_MEDIASELECT  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIASELECT)
	SDLK_WWW          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_WWW)
	SDLK_MAIL         SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MAIL)
	SDLK_CALCULATOR   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CALCULATOR)
	SDLK_COMPUTER     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_COMPUTER)
	SDLK_AC_SEARCH    SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_SEARCH)
	SDLK_AC_HOME      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_HOME)
	SDLK_AC_BACK      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_BACK)
	SDLK_AC_FORWARD   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_FORWARD)
	SDLK_AC_STOP      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_STOP)
	SDLK_AC_REFRESH   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_REFRESH)
	SDLK_AC_BOOKMARKS SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_BOOKMARKS)

	SDLK_BRIGHTNESSDOWN SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_BRIGHTNESSDOWN)
	SDLK_BRIGHTNESSUP   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_BRIGHTNESSUP)
	SDLK_DISPLAYSWITCH  SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DISPLAYSWITCH)
	SDLK_KBDILLUMTOGGLE SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KBDILLUMTOGGLE)
	SDLK_KBDILLUMDOWN   SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KBDILLUMDOWN)
	SDLK_KBDILLUMUP     SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KBDILLUMUP)
	SDLK_EJECT          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EJECT)
	SDLK_SLEEP          SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SLEEP)
	SDLK_APP1           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_APP1)
	SDLK_APP2           SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_APP2)

	SDLK_AUDIOREWIND      SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOREWIND)
	SDLK_AUDIOFASTFORWARD SDL_KeyCode = SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AUDIOFASTFORWARD)
)

const (
	KMOD_NONE     SDL_Keymod = 0x0000
	KMOD_LSHIFT   SDL_Keymod = 0x0001
	KMOD_RSHIFT   SDL_Keymod = 0x0002
	KMOD_LCTRL    SDL_Keymod = 0x0040
	KMOD_RCTRL    SDL_Keymod = 0x0080
	KMOD_LALT     SDL_Keymod = 0x0100
	KMOD_RALT     SDL_Keymod = 0x0200
	KMOD_LGUI     SDL_Keymod = 0x0400
	KMOD_RGUI     SDL_Keymod = 0x0800
	KMOD_NUM      SDL_Keymod = 0x1000
	KMOD_CAPS     SDL_Keymod = 0x2000
	KMOD_MODE     SDL_Keymod = 0x4000
	KMOD_SCROLL   SDL_Keymod = 0x8000
	KMOD_CTRL     SDL_Keymod = KMOD_LCTRL | KMOD_RCTRL
	KMOD_SHIFT    SDL_Keymod = KMOD_LSHIFT | KMOD_RSHIFT
	KMOD_ALT      SDL_Keymod = KMOD_LALT | KMOD_RALT
	KMOD_GUI      SDL_Keymod = KMOD_LGUI | KMOD_RGUI
	KMOD_RESERVED SDL_Keymod = KMOD_SCROLL
)
