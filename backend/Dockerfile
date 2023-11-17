FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o server main.go

FROM alpine:3.18
RUN apk update && apk add --no-cache git
COPY --from=builder /app/server /bin/server
WORKDIR /app
EXPOSE 8080

CMD ["/bin/server"]
