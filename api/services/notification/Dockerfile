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
RUN go build -o order-service ./cmd/main.go

FROM alpine:3.18

WORKDIR /root/

# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /app/order-service .

# Указываем порт, на котором будет работать приложение
EXPOSE 8080

CMD [ "./order-service" ]



