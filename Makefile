BUILD_DIR := bin

.DEFAULT_GOAL=help

build:
	@CGO_ENABLED=0 go build -o ${BUILD_DIR}/tcp ./cmd/server

run: build
	@./${BUILD_DIR}/tcp