# Bước 1: Build ứng dụng Go
FROM golang:1.25.1-alpine AS builder

WORKDIR /app


COPY go.* ./

RUN go mod download

# Copy toàn bộ source code
COPY . .

# Build binary
RUN go build -o main .

# (nếu có file .env hoặc cấu hình, copy luôn)
COPY .env .

# Mở port (ví dụ app chạy ở 8080)
EXPOSE 8080

# Chạy app với tham số "api"
CMD ["./main", "api"]
