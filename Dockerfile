# Estágio de Compilação
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Copia os arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila o binário
RUN go build -o main cmd/api/main.go

# Estágio de Execução (Imagem leve)
FROM alpine:latest

WORKDIR /app

# Copia o binário e a pasta static do estágio anterior
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
# Se você tiver um arquivo .env, pode copiar também, 
# mas o Compose já está injetando as variáveis.

EXPOSE 8080

CMD ["./main"]