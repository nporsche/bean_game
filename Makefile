export GOPATH:=$(GOPATH):$(CURDIR)

all: dep
	go install main

dep:
	go get github.com/nporsche/goyaml
	go get github.com/nporsche/np-golang-logger

clean:
	-rm -rf ./bin
