all: test

clean:
	rm -f ./examples/bin/basic_window

test:
	go build -o ./examples/bin/basic_window ./examples/basic_window/
