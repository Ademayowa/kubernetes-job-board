# ---- Build Stage ----
FROM golang:1.25-alpine AS builder

# If you want faster installs, add build tools:
RUN apk add --no-cache git

WORKDIR /app

# Copy only go.mod and go.sum first for caching
COPY go.mod go.sum ./

# Download modules (cached unless these files change)
RUN go mod download && go mod verify

# Copy the entire project
COPY . .

# Build static binary
# CGO_ENABLED=0 ensures compatibility with distroless/static
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/main.go

# ---- Runtime Stage ----
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .

EXPOSE 8080

USER nonroot
ENTRYPOINT ["./main"]

