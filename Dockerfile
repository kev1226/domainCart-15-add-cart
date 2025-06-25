# Etapa 1: Compilar el binario
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Compilar el binario para Linux con enlaces estáticos
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o add-cart .

# Etapa 2: Imagen final liviana
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/add-cart .

# Ejecutar el binario
CMD ["./add-cart"]
