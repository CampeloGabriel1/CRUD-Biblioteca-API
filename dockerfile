# ESTÁGIO 1: Compilação (Imagem pesada com Go instalado)
FROM golang:1.25.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# Gera o binário chamado "main"
RUN go build -o main main.go

# ESTÁGIO 2: Execução (Imagem ultra leve)
FROM alpine:latest
WORKDIR /root/
# Copiamos apenas o binário gerado no estágio anterior
COPY --from=builder /app/main .
# Executa a aplicação
CMD ["./main"]