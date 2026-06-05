# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server cmd/server/main.go

# Stage 2: Minimal runtime image
FROM alpine:3.19

WORKDIR /app

# Install certificates for secure connections (HTTPS)
RUN apk --no-cache add ca-certificates

# Create a non-root user (Blueprint section 11.3)
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/server /app/server

# Adjust ownership
RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/server"]
