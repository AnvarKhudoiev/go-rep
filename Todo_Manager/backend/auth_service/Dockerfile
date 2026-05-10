FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth-app ./auth_service/main.go

EXPOSE 8081

CMD ["./auth-app"]