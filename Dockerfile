# Sử dụng Golang làm base image
FROM golang:1.24 AS builder

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Copy toàn bộ source code vào container
COPY . .

# Tải các thư viện cần thiết
RUN go mod download

# Biên dịch mã nguồn
RUN go build -o main ./cmd/main.go

# Tạo image chạy thực tế
FROM debian:latest

# Thiết lập thư mục làm việc
WORKDIR /root/

# Tắt proxy Golang
ENV GOPROXY=direct

# Copy binary từ builder
COPY --from=builder /app/main .

# Copy file .env vào container
COPY .env /root/.env

# Expose cổng mà ứng dụng chạy
EXPOSE 8080

# Chạy ứng dụng
CMD ["./main"]