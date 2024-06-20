FROM golang:alpine3.19 as builder

WORKDIR /app

COPY src .

WORKDIR /app/presentation/github

RUN go mod tidy
RUN go build -o app.golang

FROM alpine:latest

COPY --from=builder /app/presentation/github/app.golang /app/app.golang

WORKDIR /app

ENTRYPOINT [ "/app/app.golang" ]