GOPATH:=$(CURDIR)
export GOPATH

all: build

fmt:
	gofmt -l -w -s src/

dep:fmt
	#

build:dep
	go build -o bin/goDemo main

clean:
	rm -rfv pkg
	rm -rf bin/demo
	rm -rf status

