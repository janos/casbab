FROM golang
ADD . /go/src/resenje.org/casbab
RUN cd /go/src/resenje.org/casbab && CGO_ENABLED=0 go build -ldflags "-s -w" -trimpath ./cmd/casbab

FROM scratch
COPY --from=0 /go/src/resenje.org/casbab/casbab /casbab
ENTRYPOINT ["/casbab"]
