FROM golang:1.23

WORKDIR /app

COPY cmd/main.go .
COPY go.mod .
COPY go.sum .
RUN go build -o menuFy .

EXPOSE 8000

CMD ["./menuFy", "-p", "8000:8000"]