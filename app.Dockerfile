FROM golang:1.23-alpine AS builder

ARG APP_DIR=/auth

COPY . ${APP_DIR}/
WORKDIR ${APP_DIR}/

RUN go mod download
RUN go build -ldflags="-s -w" -o ./bin/auth_server cmd/*.go

FROM alpine:3.13

ARG APP_DIR=/auth
WORKDIR /root/

COPY --from=builder ${APP_DIR}/bin/auth_server .

# CMD ["./auth_server"]
ENTRYPOINT ["tail", "-f", "/dev/null"]
