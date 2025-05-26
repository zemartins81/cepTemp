# Weather-CEP-API

API em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin).

## Funcionalidades

- Consulta de CEP via ViaCEP
- Consulta de clima via WeatherAPI
- Conversão de temperaturas (Celsius, Fahrenheit e Kelvin)
- Tratamento de erros para CEPs inválidos ou não encontrados
- Documentação Swagger
- Containerização com Docker

## Instalação e Execução

### API NO GCP
https://ceptemp-960436301750.southamerica-east1.run.app

### Execução Local

1. Clone o repositório
2. Configure a chave da API no arquivo `config/config.go` ou via variável de ambiente `WEATHER_API_KEY`
3. Execute a aplicação:

```bash
go run main.go
```

### Execução com Docker Compose

1. Configure a chave da API no arquivo `.env` ou diretamente no comando:

```bash
WEATHER_API_KEY=sua_chave_api 
docker-compose up
```

## Endpoints da API

### GET /weather/:cep

Retorna a temperatura atual para a localidade do CEP informado.

**Parâmetros:**
- `cep`: CEP válido com 8 dígitos numéricos

**Respostas:**
- `200 OK`: Retorna as temperaturas em Celsius, Fahrenheit e Kelvin
- `422 Unprocessable Entity`: CEP com formato inválido
- `404 Not Found`: CEP não encontrado
- `500 Internal Server Error`: Erro interno do servidor

**Exemplo de Resposta de Sucesso:**
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

### GET /health

Endpoint de verificação de saúde da API.

**Resposta:**
- `200 OK`: API funcionando corretamente

## Deploy no Google Cloud Run

Consulte o arquivo [docs/deploy.md](docs/deploy.md) para instruções detalhadas sobre como realizar o deploy da aplicação no Google Cloud Run.

## Testes

Execute os testes automatizados com:

```bash
go test ./tests/...
```

## Licença

Este projeto está licenciado sob a licença MIT.
