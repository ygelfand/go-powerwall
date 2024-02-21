VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
            echo v0)
all: 
		go build \
        -tags release \
        -ldflags '-X github.com/ygelfand/go-powerwall/cmd.GoPowerwallVersion=$(VERSION)' \
        -o bin/go-powerwall main.go


