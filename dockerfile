# ----------------------------
# Etap 1: Budowanie aplikacji
# ----------------------------
FROM golang:1.25 AS builder

# Ustaw katalog roboczy w kontenerze
WORKDIR /app

# Skopiuj pliki zależności i pobierz moduły
COPY go.mod go.sum ./
RUN go mod download

# Skopiuj cały kod źródłowy
COPY . .

# Zbuduj aplikację (entrypoint w cmd/api)
RUN go build -o /app/server ./cmd/api

# ----------------------------
# Etap 2: Minimalny runtime
# ----------------------------
FROM debian:bookworm-slim

# Ustaw katalog roboczy
WORKDIR /app

# Skopiuj binarkę z etapu build
COPY --from=builder /app/server .

# Otwórz port aplikacji
EXPOSE 8080

# Zmienna środowiskowa dla ewentualnej konfiguracji (np. prod/dev)
ENV PORT=8080

# Uruchom aplikację
CMD ["./server"]