/**--------------------------------------------------------
 * name: render_bmp.go
 * author: shenchunqian
 * created: 2022-10-07
 ---------------------------------------------------------*/

package main

import (
	. "github.com/chunqian/sdl2-go/sdl"
	log "github.com/chunqian/tinylog"
)

const windowWidth = 800
const windowHeight = 600

func main() {
	SDL_Init(SDL_INIT_EVERYTHING)

	window := SDL_CreateWindow("sdl2-go render bmp",
		SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, windowWidth, windowHeight, SDL_WINDOW_SHOWN)
	defer SDL_DestroyWindow(window)
	log.Info("window title: {}", SDL_GetWindowTitle(window))

	renderer := SDL_CreateRenderer(window, -1, SDL_RENDERER_ACCELERATED|SDL_RENDERER_PRESENTVSYNC)
	defer SDL_DestroyRenderer(renderer)

	surface := SDL_LoadBMP("./examples/res/test.bmp")
	defer SDL_FreeSurface(surface)

	texture := SDL_CreateTextureFromSurface(renderer, surface)

	isquit := false
	var event SDL_Event

	for !isquit {
		for SDL_PollEvent(&event) {
			if event.Type == SDL_QUIT {
				isquit = true
			}
		}

		src := SDL_Rect{X: 0, Y: 0, W: 512, H: 512}
		dst := SDL_Rect{X: 100, Y: 50, W: 512, H: 512}

		SDL_SetRenderDrawColor(renderer, 0, 0, 0, 255)
		SDL_RenderClear(renderer)

		SDL_RenderCopy(renderer, texture, &src, &dst)

		SDL_RenderPresent(renderer)
		SDL_Delay(16)
	}

	SDL_Quit()
	return
}
