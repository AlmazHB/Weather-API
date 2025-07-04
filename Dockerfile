FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/server


RUN go build -o /app/app

WORKDIR /app

CMD ["./app"]



