FROM golang:1.24-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /build

# Copiar arquivos de dependências
COPY go.mod go.sum ./

ENV GOPROXY=https://proxy.golang.org,direct

# Baixar dependências
RUN go mod download

# Copiar o código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /build/app .

# Imagem final
FROM alpine:latest

WORKDIR /app

# Copiar o arquivo .env para o container
COPY .env ./

# Copiar o binário compilado
COPY --from=builder /build/app /usr/local/bin/app

# Expor a porta
EXPOSE 8080

# Executar a aplicação
CMD ["/usr/local/bin/app"]
