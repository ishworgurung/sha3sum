all: build install clean

build:
	go build src/sha3sum.go

install: build
	cp ./sha3sum /usr/bin/

clean: 
	rm -f ./sha3sum

