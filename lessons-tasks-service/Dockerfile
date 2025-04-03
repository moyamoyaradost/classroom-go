# Используем многоступенчатую сборку
FROM golang:1.24-alpine AS builder


WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o lessons-tasks-service ./cmd/main.go

# Финальный образ
FROM alpine:latest
WORKDIR /app

# Устанавливаем зависимости для healthcheck
RUN apk add --no-cache wget

# Копируем бинарник и миграции
COPY --from=builder /app/lessons-tasks-service .
COPY --from=builder /app/migrations ./migrations

# Настраиваем здоровье приложения
HEALTHCHECK --interval=30s --timeout=10s \
  CMD wget --spider http://localhost:50051 || exit 1

EXPOSE 50051
CMD ["./lessons-tasks-service"]
