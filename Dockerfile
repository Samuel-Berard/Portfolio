#Etape 1: Build (compilation)
FROM golang:1.25 AS builder
WORKDIR /app

#copier les fichiers go
COPY main.go go.mod ./
RUN go mod download

#compiler l'application
RUN CGO_ENABLED=0 GOOS=linux go build -o portfolio main.go

#Etape 2: Runtime (image finale légère)
FROM alpine:latest
WORKDIR /app

#copier l'executable compilé
COPY --from=builder /app/portfolio .

#copier les fichiers statiques
COPY static/ ./static/
COPY templates/ ./templates/

#exposer le port
EXPOSE 8080

#lancer l'application
CMD ["./portfolio"]