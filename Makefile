all: c14n_example dom_example

c14n_example: c14n/main.go c14n/example/main.go
	go build -o c14n_example ./c14n/example

dom_example: dom/*.go dom/example/main.go serialiser/*.go
	go build -o dom_example ./dom/example

clean:
	rm -f c14n_example dom_example
