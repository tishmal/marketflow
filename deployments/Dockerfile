FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o marketflow ./cmd/marketflow

# Stage 2:
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/marketflow .

# сертификаты, если использую HTTPS/REST
RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["./marketflow"]
