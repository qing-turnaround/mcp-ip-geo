FROM golang:1.24.2-alpine AS builder

ENV APP_HOME /app
WORKDIR "$APP_HOME"

COPY . ./

RUN go mod download

RUN go build -o mcp-ip-geo ./cmd/mcp-ip-geo

FROM alpine:latest

COPY --from=builder /app/mcp-ip-geo /app/mcp-ip-geo

WORKDIR /app

EXPOSE 8000

ENTRYPOINT ["./mcp-ip-geo", "--address", "0.0.0.0:8000"]