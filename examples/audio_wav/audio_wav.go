/**--------------------------------------------------------
 * name: audio_wav.go
 * author: shenchunqian
 * created: 2022-10-07
 ---------------------------------------------------------*/

package main

import (
	"math"

	. "github.com/chunqian/sdl2-go/sdl"
	log "github.com/chunqian/tinylog"
)

const windowWidth = 800
const windowHigth = 600

type Sound struct {
	Data     []uint8
	Position uint32
}

func callback(userdata any, stream []uint8) {
	sound := userdata.(*Sound)
	size := math.Min(float64(len(stream)), float64(len(sound.Data))-float64(sound.Position))
	// stream reset
	for k := range stream {
		stream[k] = 0
	}
	for i := 0; i < int(size); i++ {
		stream[i] = sound.Data[i+int(sound.Position)]
	}
	sound.Position = sound.Position + uint32(size)
}

func main() {
	if SDL_Init(SDL_INIT_EVERYTHING) != 0 {
		log.Error("sdl init failed.")
	}

	var soundSpec SDL_AudioSpec
	var sound Sound

	if SDL_LoadWAV("./examples/res/test.wav", &soundSpec, &sound.Data) == nil {
		log.Error("wav load failed.")
	}
	defer SDL_FreeWAV(sound.Data)

	want := SDL_AudioSpec{
		Freq:     44100,
		Format:   uint16(AUDIO_S16LSB),
		Channels: 2,
		Samples:  4096,
	}
	SDL_SetAudioCallback(&want, callback, &sound)

	SDL_OpenAudio(&want, nil)
	defer SDL_CloseAudio()

	// 播放
	SDL_PauseAudio(0)

	window := SDL_CreateWindow("sdl2-go audio wav",
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
