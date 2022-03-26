build: bin
	CGO_ENABLED=0 go build \
		-o bin/audisp \
	.

build-linux:
	make build GOOS=linux
	if [ "upx not found" != "$(shell which upx)" ]; then upx bin/audisp; fi

test:
	go test -cover \
		github.com/athoune/audisp-go/fmt \
		github.com/athoune/audisp-go/audisp

bin:
	mkdir -p bin
