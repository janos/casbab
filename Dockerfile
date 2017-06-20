FROM golang
ADD . /go/src/resenje.org/casbab
RUN CGO_ENABLED=0 go build -ldflags -s -w resenje.org/casbab/cmd/casbab

FROM scratch
COPY --from=0 /go/casbab /casbab
ENTRYPOINT ["/casbab"]
