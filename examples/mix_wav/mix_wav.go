/**--------------------------------------------------------
 * name: mix_wav.go
 * author: shenchunqian
 * created: 2022-09-11
 ---------------------------------------------------------*/

package main

import (
	"io/ioutil"

	. "github.com/chunqian/sdl2-go/mix"
	. "github.com/chunqian/sdl2-go/sdl"
	log "github.com/chunqian/tinylog"
)

const windowWidth = 800
const windowHigth = 600

func main() {
	SDL_Init(SDL_INIT_EVERYTHING)

	window := SDL_CreateWindow("sdl2-go mix wav",
		SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, windowWidth, windowHigth, SDL_WINDOW_SHOWN)
	defer SDL_DestroyWindow(window)
	log.Info("window title: {}", SDL_GetWindowTitle(window))

	Mix_OpenAudio(44100, MIX_DEFAULT_FORMAT, 2, 4096)
	defer Mix_CloseAudio()

	data, _ := ioutil.ReadFile("./examples/res/test.wav")

	chunk := Mix_QuickLoad_WAV(data)
	defer Mix_FreeChunk(chunk)

	Mix_PlayChannel(1, chunk, 0)

	renderer := SDL_CreateRenderer(window, -1, 0)
	defer SDL_DestroyRenderer(renderer)

	isquit := false
	var event SDL_Event

	for !isquit {
		for SDL_PollEvent(&event) {
			if event.Type == SDL_QUIT {
				isquit = true
			}
		}
		SDL_RenderPresent(renderer)
		SDL_Delay(16)
	}

	SDL_Quit()
	return
}
