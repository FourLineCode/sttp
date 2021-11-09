main = cmd/main.go
bin = bin

all: clean
	go run $(main)

build: clean
	go build -o $(bin) $(main)

run:
	./$(bin)

test: clean build run

clean:
	rm -rf $(bin)