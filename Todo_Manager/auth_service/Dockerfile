FROM golang:1.25-alpine

WORKDIR /app

# Устанавливаем зависимости для сборки (если нужны)
RUN apk add --no-cache git

# Копируем go.mod и go.sum из корня (так как проект один)
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем конкретно auth_service
RUN go build -o auth-app ./auth_service/main.go

EXPOSE 8081

CMD ["./auth-app"]