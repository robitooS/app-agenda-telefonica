# Build da aplicação
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/app/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .
COPY migrations ./migrations
RUN mkdir logs

EXPOSE 8080

CMD ["/app/main"]