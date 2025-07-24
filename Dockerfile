FROM golang:1.21-alpine AS builder
WORKDIR /app
ENV GOPROXY=https://proxy.golang.org,direct
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -o /app/main ./cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]