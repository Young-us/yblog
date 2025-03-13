FROM golang:1.13-alpine AS bulider

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=bulider /app/main .

COPY resources /app/resources

COPY public /app/public

CMD ["./main"]