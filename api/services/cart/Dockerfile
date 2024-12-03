# Указываем базовый образ для сборки
FROM golang:1.23.1-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем Go модули и загружаем зависимости
COPY ./go.mod ./go.sum ./
RUN go mod download

# Копируем остальные файлы исходного кода
COPY . .

# Сборка бинарного файла
RUN go build -o cart-service ./cmd/main.go

FROM alpine:3.18


# Копируем бинарник из этапа сборки
COPY --from=builder /app/cart-service .

# Открываем порт
EXPOSE 8080

CMD [ "./cart-service" ]

# RUN apk update && apk add --no-cache postgresql-client

# COPY ../db/init-db.sh /usr/local/bin/
# RUN chmod +x /usr/local/bin/

# Команда для запуска приложения
# ENTRYPOINT ["sh", "/usr/local/bin/create-db.sh"]