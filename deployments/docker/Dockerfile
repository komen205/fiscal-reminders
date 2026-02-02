FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /fiscal-reminders ./cmd/fiscal-reminders

# Final stage
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /fiscal-reminders /app/fiscal-reminders

# Default timezone for Portugal
ENV TZ=Europe/Lisbon

ENTRYPOINT ["/app/fiscal-reminders"]
