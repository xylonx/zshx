PROJECT:=zshx

.PHONY: build

default: ${PROJECT}

${PROJECT}: build

build: main.go
	go build -o ${PROJECT} $<