# Build stage
FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN apk add curl
RUN go build -o main main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.deb -o migrate.deb

# Final stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.deb .
RUN apk add dpkg
RUN dpkg -i ./migrate.deb
RUN rm ./migrate.deb
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration
EXPOSE 5000
CMD ["./main"]
