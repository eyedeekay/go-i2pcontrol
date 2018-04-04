
UNAME ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')
UARCH ?= $(shell uname -m | tr '[:upper:]' '[:lower:]' | sed 's|x86_64|amd64|g')

check:
	go build ..

r: rebuild

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
		./demo.go
	./demo-linux

demo-linux:
	GOOS=linux GOARCH="$(UARCH)" go build \
		-a \
		-tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o demo-linux \
		./demo.go

demo-win:
	GOOS=windows GOARCH="$(UARCH)" go build \
		-a \
		-tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o demo-win \
		./demo.go


demo-osx:
	GOOS=darwin GOARCH="$(UARCH)" go build \
		-a \
		-tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o demo-osx \
		./demo.go

all: demo-linux demo-win demo-osx

curl:
	/usr/bin/curl --data-binary '{"jsonrpc":"2.0","id":"curltext","method":"Authenticate","params":[ API: 1, Password: "default"]}' -H 'content-type:text/plain;' http://127.0.0.1:7650/
