# Base image golang:1.16
FROM golang:1.16

# Thiết lập thư mục làm việc
WORKDIR /app

# Sao chép mã nguồn vào thư mục làm việc
COPY . .

# Biên dịch ứng dụng
RUN go build -o main .

# Thiết lập cổng mặc định
EXPOSE 8080

# Chạy ứng dụng khi container được khởi chạy
CMD ["./main"]