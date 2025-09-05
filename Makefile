build:
	go build -o ./gb ./main.go

install_linux:
	make build
	mv ./gb ~/go/bin/gb