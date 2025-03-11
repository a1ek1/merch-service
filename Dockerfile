FROM golang:1.23 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o merch-service ./cmd/main.go

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/merch-service .
COPY --from=builder /app/migrations ./migrations
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["/app/merch-service"]
