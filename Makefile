all: test-examples

clean:
	rm -f ./examples/bin/basic_window

test-sdl:
	go build ./sdl/

test-gfx:
	go build ./gfx/

test-examples:
	go build -o ./examples/bin/basic_window ./examples/basic_window/
