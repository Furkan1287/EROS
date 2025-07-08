# User Service Dockerfile
FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o user-service .
CMD ["./user-service"]
