FROM golang:1.24.0

WORKDIR /app

# Copier les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source
COPY . .

# COMPILER (Docker fait ça pour vous)
RUN go build -o forum ./cmd/server

# EXÉCUTER le binaire compilé (pas le .go !)
CMD ["./forum"]