FROM golang AS builder

ENV CGO_ENABLED=0 GO111MODUL=on

WORKDIR /build

ADD go.mod .
COPY . .

RUN go build -o server cmd/server/server.go

FROM alpine

ARG CONFIG_PATH
ENV CONFIG_PATH=$CONFIG_PATH

EXPOSE 8400

WORKDIR /build

COPY --from=builder /build/server /server

COPY .env /
COPY config config

ENTRYPOINT ["/server"]
