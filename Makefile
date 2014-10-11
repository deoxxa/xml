all: c14n_example

c14n_example: c14n/main.go c14n/example/main.go
	go build -a -o c14n_example ./c14n/example

clean:
	rm c14n_example
