FROM docker.io/library/golang:1.15 as builder

WORKDIR /workspace
COPY pkg pkg
COPY go.mod go.mod
COPY main.go main.go

RUN CGO_ENABLED=0 go build -o server .

FROM alpine

EXPOSE 9090
COPY --from=builder /workspace/server /server
CMD ["/server"]
