# Menggunakan base image golang
FROM golang:alpine

# Setel direktori kerja di dalam container
WORKDIR /app

RUN apk add --update --no-cache gcc musl-dev

# sqlite need this -> remove if you not used it or set 0
ENV CGO_ENABLED=1
ENV GIN_MODE="release"

# Menyalin go.mod dan go.sum
COPY go.mod go.sum ./

# Unduh dependencies
RUN go mod download

# Menyalin seluruh kode sumber aplikasi
COPY . .

# Membuat aplikasi
RUN go build -o main .

# Menentukan port aplikasi yang akan dijalankan
EXPOSE 8080

# Menjalankan aplikasi
CMD ["./main"]
