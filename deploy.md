# Instruções para Deploy no Google Cloud Run

Este documento contém as instruções para realizar o deploy da aplicação Weather-CEP-API no Google Cloud Run.

## Pré-requisitos

1. Conta no Google Cloud Platform
2. Google Cloud CLI instalado e configurado
3. Docker instalado localmente
4. Chave válida da WeatherAPI

## Passos para Deploy

### 1. Configurar a Chave da API

Antes de realizar o deploy, substitua a chave fictícia da WeatherAPI por uma chave válida:

```bash
export WEATHER_API_KEY="sua_chave_api_aqui"
```

### 2. Construir a Imagem Docker

```bash
docker build -t gcr.io/[SEU_PROJETO_GCP]/weather-cep-api .
```

### 3. Enviar a Imagem para o Container Registry

```bash
docker push gcr.io/[SEU_PROJETO_GCP]/weather-cep-api
```

### 4. Deploy no Cloud Run

```bash
gcloud run deploy weather-cep-api \
  --image gcr.io/[SEU_PROJETO_GCP]/weather-cep-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars="WEATHER_API_KEY=sua_chave_api_aqui"
```

### 5. Verificar o Serviço

Após o deploy, o Google Cloud Run fornecerá uma URL para acessar o serviço. Você pode testar a API com:

```bash
curl https://[URL_DO_SERVICO]/weather/01001000
```

## Notas Importantes

- Certifique-se de que a chave da WeatherAPI seja válida
- O serviço será escalado automaticamente conforme a demanda
- O Google Cloud Run só cobra pelo tempo em que o serviço está processando requisições
