# Stage 1: Application Build
FROM golang:1.22.1  

ARG SERVICE_BINARY=cron_recibo
WORKDIR /app

# Copiar os arquivos necessários
COPY . .

# Instalar as dependências
RUN go mod tidy

# Construir a aplicação
RUN go build -o /app/${SERVICE_BINARY} ./cmd/agendamento/*.go

# Verificar se o executável foi realmente criado
# RUN ls -la /app

# # Instalar certificados CA e criar um usuário não root
# RUN apk add --no-cache ca-certificates \
#     && addgroup -S appgroup && adduser -S appuser -G appgroup

# # Tornar o executável executável e ajustar permissões
# RUN chmod +x /app/${SERVICE_BINARY} \
#     && chown -R appuser:appgroup /app

# USER appuser
 
  CMD ["./cron_recibo"] 
