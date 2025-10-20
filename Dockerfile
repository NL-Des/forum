FROM golang:1.24.0

WORKDIR /app

# Importation des dépendances.
COPY go.mod go.sum ./
RUN go mod download

# Copie du code source.
COPY . .

# Docker compile.
RUN go build -o forum ./cmd/server

# Lancement de l'éxécutable.
CMD ["./forum"]