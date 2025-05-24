FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o weather-cep-api .

# Imagem final
FROM alpine:latest

WORKDIR /app

# Copiar o binário compilado
COPY --from=builder /app/weather-cep-api .

# Expor a porta
EXPOSE 8080

# Definir variável de ambiente para a chave da API
ENV WEATHER_API_KEY=""

# Executar a aplicação
CMD ["./weather-cep-api"]
