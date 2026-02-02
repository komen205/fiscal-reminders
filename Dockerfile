# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY *.go ./

RUN go build -o fiscal-reminders .

# Run stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/fiscal-reminders .
COPY config.json .

# Set timezone to Lisbon
ENV TZ=Europe/Lisbon

CMD ["./fiscal-reminders"]

