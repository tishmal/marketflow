# Stage 1: билдим Go-приложение
FROM golang:1.24-alpine AS builder

# Рабочая директория
WORKDIR /app

# Копируем go.mod и загружаем зависимости
COPY go.mod ./
RUN go mod download

# Копируем остальной проект
COPY . .

# Сборка бинарника (для Linux)
RUN GOOS=linux GOARCH=amd64 go build -o marketflow ./cmd/marketflow

# Stage 2: минимальный рантайм-образ
FROM alpine:latest

# Создаём рабочую директорию
WORKDIR /app

# Копируем бинарник из builder stage
COPY --from=builder /app/marketflow .

# Установка сертификатов, если используешь HTTPS-запросы
RUN apk add --no-cache ca-certificates

# Открываем порт, если используешь web API
EXPOSE 8080

# Команда по умолчанию при старте контейнера
CMD ["./marketflow"]
