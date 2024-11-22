FROM golang:alpine AS builder
WORKDIR /root
ADD . .
RUN go mod tidy
RUN env CGO_ENABLED=0 go build -v -trimpath -ldflags '-w -s'
FROM scratch
COPY --from=builder /root/webd /webd
ENTRYPOINT ["/webd"]
