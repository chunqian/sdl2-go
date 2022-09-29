all: test-sdl test-gfx test-raster test-img test-ttf test-examples

clean:
	rm -f ./examples/bin/basic_window

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

test-examples:
	go build -o ./examples/bin/basic_window ./examples/basic_window/
