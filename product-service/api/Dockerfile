# Этап 1: Сборка
FROM golang:1.23.1-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем все остальные файлы
COPY . .

# Собираем приложение
RUN go build -o product-service ./cmd/main.go

# Этап 2: Создаем легковесный образ для выполнения
FROM alpine:latest

WORKDIR /root/

# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /app/product-service .

# Указываем порт, на котором будет работать приложение
EXPOSE 8080

# Запускаем приложение
CMD ["./product-service"]

