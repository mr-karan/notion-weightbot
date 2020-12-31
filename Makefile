.PHONY : build run fresh test clean

BIN := weightbot

BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')

build:
	go build -o ${BIN} -ldflags=" -X 'main.buildDate=${BUILD_DATE}'"

build-docker-amd64:
	docker build --build-arg ARCH=amd64 -t weightbot/amd64:latest .

build-docker-arm32v7:
	docker build --build-arg ARCH=arm32v7 -t weightbot/arm32v7:latest .

run:
	./${BIN}

fresh: clean build run

test:
	go test

clean:
	go clean
	- rm -f ${BIN}
