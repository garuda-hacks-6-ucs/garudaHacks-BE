FROM golang:1.21-alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/main ./cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/bin/main .
EXPOSE 8080
CMD ["./main"]