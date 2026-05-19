FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o soma cmd/main.go

ENTRYPOINT ["./soma"]