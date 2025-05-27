FILES = cesar.go \
		vigenere.go \
		main.go \
		cipher.go 

all:

build:
	go run $(FILES)

test:
	go test -v