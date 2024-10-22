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

FROM alpine:3.18

WORKDIR /root/

# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /app/product-service .

# Указываем порт, на котором будет работать приложение
EXPOSE 8080

CMD [ "./product-service" ]


# RUN apk update && apk add --no-cache postgresql-client
# COPY ../db/init-db.sh /usr/local/bin/
# RUN chmod +x /usr/local/bin/

# Команда запуска контейнера
# ENTRYPOINT ["sh", "/usr/local/bin/init-db.sh"]

