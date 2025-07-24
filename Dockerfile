FROM golang:1.24
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o /go/bin/main ./cmd/api/main.go
EXPOSE 8080
CMD ["app"]


#RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/main ./cmd/api/main.go
#
#FROM alpine:latest
#WORKDIR /root/
#COPY --from=builder /go/bin/main .
#EXPOSE 8080
#CMD ["./main"]