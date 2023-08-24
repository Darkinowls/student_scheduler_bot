# Stage 1: Build the Go application
FROM golang:1.20 AS builder
WORKDIR /code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Stage 2: Create the final image using Distroless base
FROM gcr.io/distroless/base
ENV TZ=Europe/Kyiv
COPY --from=builder /code/main /
COPY --from=builder /code/.env /
CMD ["/main"]
