FROM golang AS builder

WORKDIR /build

ADD go.mod .
COPY . .

RUN go build -o server cmd/server/server.go

FROM alpine

ARG CONFIG_PATH
ENV CONFIG_PATH=$CONFIG_PATH

EXPOSE 9000

WORKDIR /build

COPY --from=builder /build/server /server

COPY .env /
COPY config config

ENTRYPOINT ["/server"]
