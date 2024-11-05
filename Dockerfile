# Stage 1: Construir la aplicación GO
FROM golang:alpine AS builder

# Instalar GIT para descargar módulos de GO
RUN apk update && apk add --no-cache git

# Configurar el directorio de trabajo
WORKDIR /app

# Copiar el source code
COPY . .

# Iniciar Go Modules
RUN go mod init to-do-app

# Añadir las dependencias
RUN go get github.com/gin-gonic/gin
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/postgres

# Descargar las dependencias
RUN go mod tidy

# Construir la aplicación main
RUN go build -o main .

# Stage 2: Ejecutar la aplicación
FROM golang:alpine

# Configurar el directorio de trabajo
WORKDIR /app

# Copiar la aplicación construido en el contenedor
COPY --from=builder /app/main .

# Usar puerto 8080 para conexiones entrantes
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]
