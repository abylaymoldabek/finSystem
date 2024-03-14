FROM golang:latest as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o app .

FROM alpine:latest  

COPY --from=builder /app/app /usr/local/bin/app

COPY config.json .

EXPOSE 8088

CMD ["app"]
