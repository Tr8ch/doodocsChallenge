FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o doodocs ./cmd/doodocsApp/main.go
FROM alpine
WORKDIR /app
COPY --from=builder /app/doodocs ./
CMD ["./doodocs"]
