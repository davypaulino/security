FILES = cesar.go \
		vigenere.go \
		main.go \
		cipher.go 

all: build

build:
	go run $(FILES)

test:
	go test -v