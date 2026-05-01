FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем конкретно todo_service
RUN go build -o todo-app ./todo_service/main.go

EXPOSE 8082

CMD ["./todo-app"]