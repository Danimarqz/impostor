# ==========================================
# Stage 1: Build Frontend (Node.js)
# ==========================================
FROM node:18-alpine AS builder-web

WORKDIR /app/web
COPY web/package*.json ./
RUN npm install

COPY web/ ./
RUN npm run build

# ==========================================
# Stage 2: Build Backend (Go)
# ==========================================
FROM golang:1.22-alpine AS builder-go

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Build the binary, explicitly disabling CGO for static linking compatibility/portability
RUN CGO_ENABLED=0 GOOS=linux go build -o impostor-server ./cmd/server/main.go

# ==========================================
# Stage 3: Final Runtime Image
# ==========================================
FROM alpine:latest

WORKDIR /root/

# Copy Backend Binary
COPY --from=builder-go /app/impostor-server .

# Copy Frontend Assets to ./public (as expected by server.go)
COPY --from=builder-web /app/web/dist ./public

# Expose port
EXPOSE 8080

# Run
CMD ["./impostor-server"]
