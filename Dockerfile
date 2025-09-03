FROM golang:1.24-alpine3.22

EXPOSE 8090

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd

CMD ["./main"]
