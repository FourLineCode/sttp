main = cmd/main.go
bin = sttp

all: clean
	go run $(main)

build: clean
	go build -o $(bin) $(main)

run:
	./$(bin)

clean:
	rm -f $(bin)