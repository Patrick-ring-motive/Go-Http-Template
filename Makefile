.PHONY: run
run: main
	go run main

main: *.go go.mod
	bash ./build.sh &
	go build -o main main.go
	chmod +x ./main

.PHONY: all
all: main
