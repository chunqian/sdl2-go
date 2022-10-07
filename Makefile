all: test-sdl test-gfx test-raster test-img test-ttf test-mix test-examples

clean:
	rm -f ./examples/bin/basic_window
	rm -f ./examples/bin/audio_wav
	rm -f ./examples/bin/mix_wav
	rm -f ./examples/bin/render_bmp

test-sdl:
	go build ./sdl/

test-gfx:
	go build ./gfx/

test-raster:
	go build ./raster/

test-img:
	go build ./img/

test-ttf:
	go build ./ttf/

test-mix:
	go build ./mix/

test-examples:
	go build -o ./examples/bin/basic_window ./examples/basic_window/
	go build -o ./examples/bin/audio_wav ./examples/audio_wav/
	go build -o ./examples/bin/mix_wav ./examples/mix_wav/
	go build -o ./examples/bin/render_bmp ./examples/render_bmp/
