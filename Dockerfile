# Sử dụng Go Image được build sẵn
FROM golang:1.16-alpine

# Thiết lập biến môi trường GOPATH và PATH
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Thiết lập thư mục làm việc trong container
WORKDIR /app

# Sao chép file go.mod và go.sum vào thư mục làm việc
COPY go.mod go.sum ./

# Tải về các phụ thuộc của ứng dụng
RUN go mod download

# Sao chép toàn bộ source code vào thư mục làm việc
COPY . .

# Build ứng dụng
RUN go build -o main .

# Thiết lập lệnh mặc định để chạy ứng dụng khi container được bật
CMD ["/app/main"]
