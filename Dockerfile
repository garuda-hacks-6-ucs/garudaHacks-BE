FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod edit -replace govtech-api=.
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]