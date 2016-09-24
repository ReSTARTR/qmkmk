VERSION=0.0.1
BIN_NAME=qmkmk

build:
	go build -ldflags "-X main.version=${VERSION}" -o bin/${BIN_NAME}


