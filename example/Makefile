
UNAME ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')
UARCH ?= $(shell uname -m | tr '[:upper:]' '[:lower:]' | sed 's|x86_64|amd64|g')

rebuild: clean update demo

clean:
	rm -f demo-*

update:
	go get -u github.com/eyedeekay/go-i2pcontrol; true

demo:
	GOOS="$(UNAME)" GOARCH="$(UARCH)" go build \
		-a \
		-tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o demo-"$(UNAME)" \
		./
	./demo-linux

demo-linux:
	GOOS=linux GOARCH="$(UARCH)" go build \
		-a \
		-tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o demo-linux \
		./

demo-win:
	GOOS=windows GOARCH="$(UARCH)" go build \
		-a \
		-tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o demo-win \
		./


demo-osx:
	GOOS=darwin GOARCH="$(UARCH)" go build \
		-a \
		-tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o demo-osx \
		./

all: demo-linux demo-win demo-osx
