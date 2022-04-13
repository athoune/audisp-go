build: bin
	CGO_ENABLED=0 go build \
		-o bin/audisp \
	.

build-sshd: bin
	CGO_ENABLED=0 go build \
		-o bin/audisp-sshd \
	./cli/sshd/main.go

build-sons: bin
	CGO_ENABLED=0 go build \
		-o bin/audisp-sons \
	./cli/sons/main.go

build-linux:
	make build GOOS=linux
	if [ "upx not found" != "$(shell which upx)" ]; then upx bin/audisp; fi

build-sshd-linux:
	make build-sshd GOOS=linux
	if [ "upx not found" != "$(shell which upx)" ]; then upx bin/audisp-sshd; fi

build-sons-linux:
	make build-sons GOOS=linux
	if [ "upx not found" != "$(shell which upx)" ]; then upx bin/audisp-sons; fi

test:
	go test -cover \
		github.com/athoune/audisp-go/fmt \
		github.com/athoune/audisp-go/audisp

bin:
	mkdir -p bin
