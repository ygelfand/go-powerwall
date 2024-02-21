FROM golang:1.22 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

COPY main.go main.go
COPY internal/ internal/
COPY cmd/ cmd/
ARG BUILD_VERSION
RUN CGO_ENABLED=0 go build -ldflags="-w -s -X github.com/ygelfand/go-powerwall/cmd.debugMode=false -X github.com/ygelfand/go-powerwall/cmd.GoPowerwallVersion=${BUILD_VERSION}"  -o go-powerwall main.go

FROM scratch
COPY --from=builder /workspace/go-powerwall /go-powerwall

ENTRYPOINT ["/go-powerwall"]
CMD ["proxy"]
