FROM --platform=linux/arm64/v8 golang:1.21-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go-crud-app/go.mod go-crud-app/go.sum ./

RUN go mod download

COPY go-crud-app/ .

RUN go build -o main .

FROM --platform=linux/arm64/v8 alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env . 

EXPOSE 8080

CMD ["./main"]
