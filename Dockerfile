# Step 1: Build the Go binary
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o expense-api ./cmd/server/main.go

# Step 2: Create a tiny production image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/expense-api .
EXPOSE 8080
CMD ["./expense-api"]
