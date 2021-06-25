.PHONY: recovery clean

os=$(shell uname -s | tr '[:upper:]' '[:lower:]')
arch="amd64"

recovery:
	@GOOS=${os} GOARCH=${arch} CGO_ENABLED=0 go build ./cmd/recovery

clean:
	@rm -f i2i-client

