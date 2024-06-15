FROM golang:1.21-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go-crud-app/go.mod go-crud-app/go.sum ./

RUN go mod download

COPY go-crud-app/ .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
