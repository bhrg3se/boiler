FROM golang:1.16 as builder

WORKDIR /go/src/boiler

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

ENV CGO_ENABLED=0
RUN go build -o boiler

FROM alpine:latest
WORKDIR /boiler

COPY --from=builder /go/src/boiler/boiler .
COPY setup/config.toml /etc/boiler/config.toml

CMD ["./boiler"]
