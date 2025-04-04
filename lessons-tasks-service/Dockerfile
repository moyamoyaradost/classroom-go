# Используем многоступенчатую сборку
FROM golang:1.24-alpine AS builder

# Устанавливаем зависимости для protoc и компиляции
RUN apk add --no-cache protoc protobuf-dev git make

# Устанавливаем Go плагины для генерации gRPC кода
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /app
COPY . .

# Генерируем gRPC-код
RUN protoc --go_out=. --go-grpc_out=. \
    --proto_path=./api/v1 \
    ./api/v1/lessons_tasks.proto

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o lessons-tasks-service ./cmd/main.go

# Финальный образ
FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache wget

COPY --from=builder /app/lessons-tasks-service .
COPY --from=builder /app/migrations ./migrations

HEALTHCHECK --interval=30s --timeout=10s \
  CMD wget --spider http://localhost:50051 || exit 1

EXPOSE 50051
CMD ["./lessons-tasks-service"]
