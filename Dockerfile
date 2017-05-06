FROM golang
ADD . /go/src/resenje.org/casbab
RUN go build -ldflags -w resenje.org/casbab/cmd/casbab

FROM scratch
COPY --from=0 /go/casbab /casbab
ENTRYPOINT ["/casbab"]
