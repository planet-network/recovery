.PHONY: recovery clean

os=$(shell uname -s | tr '[:upper:]' '[:lower:]')
arch="amd64"

i2i-client:
	@GOOS=${os} GOARCH=${arch} CGO_ENABLED=0 go build ./cmd/recovery

clean:
	@rm -f i2i-client
