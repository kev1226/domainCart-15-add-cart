# Build stage
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o add-cart .

# Final image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/add-cart .

EXPOSE 3035

CMD ["./add-cart"]
