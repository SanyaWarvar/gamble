# Используем официальный образ Go
FROM golang:latest AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o user_service ./cmd/

# Используем образ Debian для запуска
FROM debian:bullseye-slim

# Устанавливаем необходимые библиотеки
RUN apt-get update && apt-get install -y --no-install-recommends \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

# Копируем собранное приложение из предыдущего этапа
COPY --from=builder /app/user_service /user_service

COPY .env ./

# Указываем команду для запуска
ENTRYPOINT ["/user_service"]