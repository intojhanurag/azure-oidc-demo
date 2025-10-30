FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app .

FROM alpine:latest

ENV PORT=8080
EXPOSE 8080

COPY --from=builder /app/app /app/app

ENTRYPOINT ["/app/app"]
