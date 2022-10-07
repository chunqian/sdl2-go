/**--------------------------------------------------------
 * name: basic_window.go
 * author: shenchunqian
 * created: 2022-09-11
 ---------------------------------------------------------*/

package main

import (
	. "github.com/chunqian/sdl2-go/sdl"
	log "github.com/chunqian/tinylog"
)

const windowWidth = 800
const windowHigth = 600

func main() {
	SDL_Init(SDL_INIT_EVERYTHING)

	window := SDL_CreateWindow("SDL2-Go Basic Window",
		SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, windowWidth, windowHigth, SDL_WINDOW_SHOWN)
	defer SDL_DestroyWindow(window)
	log.Info("window title: {}", SDL_GetWindowTitle(window))

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
